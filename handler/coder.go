package handler

import (
	"bytes"
	"database/sql"
	"io/ioutil"
	"mtools-backend/utils"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"mtools-backend/config"
	"mtools-backend/model"
	"mtools-backend/schema"
	"mtools-backend/tpl"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var CoderHandlerSet = wire.NewSet(wire.Struct(new(CoderHandler), "*"))

type CoderHandler struct {
	Logger        *zap.SugaredLogger
	Config        *config.GlobalConfig
	DatabaseModel *model.DatabaseModel
}

func (h *CoderHandler) GenCode(c *gin.Context) {
	conf := new(schema.GenCode)
	if err := ParseJSON(c, &conf); err != nil {
		returnFailed(c, validatorErrorData(err))
		return
	}

	// 获取数据库信息
	dbConf, err := h.DatabaseModel.Get(conf.DatabaseId)
	if err != nil {
		returnFailed(c, err.Error())
		return
	}

	// 遍历tables
	for _, table := range conf.Tables {
		// 查询字段列表
		fields, err := h.getFields(dbConf, table)
		if err != nil {
			h.Logger.Errorf("获取%s表, 字段出现错误: %v", table, err)
			continue
		}

		// 获取需要生成的文件列表
		files := h.genFiles(conf.GenFrontEnd)
		for _, file := range files {

			smallCamelTable := tpl.SmallCamel(table)
			bigCamelTable := tpl.BigCamel(table)

			filePath := filepath.Join(conf.Output, utils.IfStr(conf.UseOriginTable, table, smallCamelTable),
				file.Path, utils.IfStr(conf.UseOriginTable, table, bigCamelTable)+file.Suffix)
			fileContent := bytes.NewBufferString("")

			// 构造模板
			tmpl, err := template.New(file.Key).Funcs(tpl.Funcs()).Delims("${", "}").Parse(file.Template)
			if err != nil {
				h.Logger.Errorf("表 [%s], 模板初始化失败 %v", table, err)
				continue
			}

			if conf.UseParent {
				columns := conf.ExcludeColumn
				excludes := strings.Split(columns, ",")

				// 如果使用父类，那么会有字段需要排除
				for i := 0; i < len(fields); i++ {
					for j := 0; j < len(excludes); j++ {
						if fields[i].ColumnName == excludes[j] {
							fields[i].Exclude = true
						}
					}
				}
			} else {
				// 如果不使用父类
				// 需要判断是否有字段需要自动填充
				if conf.AutoFill {
					columns := conf.AutoFillColumn
					autoFills := strings.Split(columns, ",")
					for i := 0; i < len(fields); i++ {
						for j := 0; j < len(autoFills); j++ {
							if fields[i].ColumnName == autoFills[j] {
								fields[i].AutoFill = true
								if strings.Contains(autoFills[j], "update") {
									fields[i].AutoFillType = "INSERT_UPDATE"
								} else {
									fields[i].AutoFillType = "INSERT"
								}
							}
						}
					}
				}
			}

			if conf.FormatDateColumn {
				for i := 0; i < len(fields); i++ {
					if fields[i].DataType == "datetime" {
						fields[i].FormatDate = true
					}
				}
			}

			// 生成代码
			if err = tmpl.Execute(fileContent, map[string]interface{}{
				"PackageName":     utils.IfStr(conf.UseOriginTable, table, smallCamelTable),
				"TableName":       utils.IfStr(conf.UseOriginTable, table, bigCamelTable),
				"OriginTableName": table,
				"Fields":          fields,
				"UseParent":       conf.UseParent,
				"ParentPackage":   conf.ParentPackage,
				"UseOriginColumn": conf.UseOriginColumn,
			}); err != nil {
				h.Logger.Errorf("表 [%s], 生成代码失败 %v", table, err)
				continue
			}

			if conf.OutputCover {
				_ = os.Remove(filePath)
			}

			// 写入代码到文件
			_ = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
			if err := ioutil.WriteFile(filePath, fileContent.Bytes(), os.ModePerm); err != nil {
				h.Logger.Errorf("表[%s] , 写入文件失败 %v", table, err)
				continue
			}
		}
	}
	returnSuccess(c, nil)
	return
}

type Field struct {
	ColumnName             string         `db:"COLUMN_NAME"`
	ColumnDefault          sql.NullString `db:"COLUMN_DEFAULT"`
	DataType               string         `db:"DATA_TYPE"`
	ColumnComment          string         `db:"COLUMN_COMMENT"`
	IsKey                  string         `db:"IS_KEY"`
	NumericPrecision       sql.NullInt64  `db:"NUMERIC_PRECISION"`
	NumericScale           sql.NullInt64  `db:"NUMERIC_SCALE"`
	CharacterMaximumLength sql.NullInt64  `db:"CHARACTER_MAXIMUM_LENGTH"`
	IsNullable             string         `db:"IS_NULLABLE"`
	ColumnType             string         `db:"COLUMN_TYPE"`
	IsAuto                 string         `db:"IS_AUTO"`
	FormatDate             bool
	Exclude                bool
	AutoFill               bool
	AutoFillType           string
}

type GenFile struct {
	Key      string
	Path     string
	Suffix   string
	Template string
}

func (h *CoderHandler) genFiles(hasWeb bool) (files []*GenFile) {
	files = make([]*GenFile, 0)
	// files = append(files, &GenFile{Key: "Controller", Path: "controller", Suffix: "Controller.java", Template: tpl.ControllerTemplate})
	// files = append(files, &GenFile{Key: "Service", Path: "service", Suffix: "Service.java", Template: tpl.ServiceTemplate})
	// files = append(files, &GenFile{Key: "Mapper", Path: "mapper", Suffix: "Mapper.java", Template: tpl.MapperTemplate})
	// files = append(files, &GenFile{Key: "MapperXml", Path: "mapper/xml", Suffix: "Mapper.xml", Template: tpl.MapperXmlTemplate})
	files = append(files, &GenFile{Key: "Entity", Path: "entity", Suffix: ".java", Template: tpl.EntityTemplate})
	if hasWeb {
		// files = append(files, &GenFile{Key: "Vue", Path: "vue/views", Suffix: ".vue", Template: tpl.VueTemplate})
	}
	return
}

func (h *CoderHandler) getFields(conf *model.Database, tableName string) ([]*Field, error) {
	var (
		fieldList []*Field
		dbUrl     = conf.Username + ":" + conf.Password + "@tcp("
	)
	if strings.Contains(conf.Host, ":") {
		dbUrl += conf.Host + ")/"
	} else {
		dbUrl += conf.Host + ":3306)/"
	}
	dbUrl += conf.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, _ := sqlx.Open("mysql", dbUrl)
	if err := db.Select(&fieldList, config.QueryTableFieldList, tableName, conf.Name); err != nil {
		return fieldList, err
	}
	return fieldList, nil
}
