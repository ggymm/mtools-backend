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
	"time"

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
	conf          *schema.GenCode
	Fields        []*Field
}

func (h *CoderHandler) GenJavaCode(c *gin.Context) {
	h.conf = new(schema.GenCode)
	if err := ParseJSON(c, &h.conf); err != nil {
		returnFailed(c, validatorErrorData(err))
		return
	}

	// 获取数据库信息
	dbConf, err := h.DatabaseModel.Get(h.conf.DatabaseId)
	if err != nil {
		returnFailed(c, err.Error())
		return
	}

	// 遍历tables
	for _, table := range h.conf.Tables {
		// 查询字段列表
		if h.Fields, err = h.getFields(dbConf, table); err != nil {
			h.Logger.Errorf("获取%s表, 字段出现错误: %v", table, err)
			continue
		}

		// 获取需要生成的文件列表
		files := h.javaFiles()
		for _, file := range files {

			smallCamelTable := tpl.SmallCamel(table)
			bigCamelTable := tpl.BigCamel(table)

			filePath := filepath.Join(h.conf.Output, utils.IfStr(h.conf.UseOriginTable, table, smallCamelTable),
				file.Path, utils.IfStr(h.conf.UseOriginTable, table, bigCamelTable)+file.Suffix)
			fileContent := bytes.NewBufferString("")

			// 构造模板
			tmpl, err := template.New(file.Key).Funcs(tpl.Funcs()).Delims("${", "}").Parse(file.Template)
			if err != nil {
				h.Logger.Errorf("表 [%s], 模板初始化失败 %v", table, err)
				continue
			}

			if h.conf.UseParent {
				columns := h.conf.ExcludeColumn
				excludes := strings.Split(columns, ",")

				// 如果使用父类，那么会有字段需要排除
				for i := 0; i < len(h.Fields); i++ {
					for j := 0; j < len(excludes); j++ {
						if h.Fields[i].ColumnName == excludes[j] {
							h.Fields[i].Exclude = true
						}
					}
				}
			} else {
				// 如果不使用父类
				// 需要判断是否有字段需要自动填充
				if h.conf.AutoFill {
					columns := h.conf.AutoFillColumn
					autoFills := strings.Split(columns, ",")
					for i := 0; i < len(h.Fields); i++ {
						for j := 0; j < len(autoFills); j++ {
							if h.Fields[i].ColumnName == autoFills[j] {
								h.Fields[i].AutoFill = true
								if strings.Contains(autoFills[j], "update") {
									h.Fields[i].AutoFillType = "INSERT_UPDATE"
								} else {
									h.Fields[i].AutoFillType = "INSERT"
								}
							}
						}
					}
				}
			}

			if h.conf.FormatDateColumn {
				for i := 0; i < len(h.Fields); i++ {
					if h.Fields[i].DataType == "datetime" {
						h.Fields[i].FormatDate = true
					}
				}
			}

			// 生成代码
			if err = tmpl.Execute(fileContent, map[string]interface{}{
				"Author":          utils.GetUsername(),
				"Now":             time.Now().Format("2006-01-02 15:04:05"),
				"BasePackageName": h.conf.Package,
				"PackageName":     utils.IfStr(h.conf.UseOriginTable, table, smallCamelTable),
				"TableName":       utils.IfStr(h.conf.UseOriginTable, table, bigCamelTable),
				"OriginTableName": table,
				"Fields":          h.Fields,
				"HasID":           h.hasID(),
				"IDType":          h.IDType(),
				"UseParent":       h.conf.UseParent,
				"ParentPackage":   h.conf.ParentPackage,
				"UseOriginColumn": h.conf.UseOriginColumn,
			}); err != nil {
				h.Logger.Errorf("表 [%s], 生成代码失败 %v", table, err)
				continue
			}

			if h.conf.OutputCover {
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

func (h *CoderHandler) GenGoCode(c *gin.Context) {
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
	defer func() {
		_ = db.Close()
	}()
	if err := db.Select(&fieldList, config.QueryTableFieldList, tableName, conf.Name); err != nil {
		return fieldList, err
	}
	return fieldList, nil
}

func (h *CoderHandler) hasID() bool {
	for _, field := range h.Fields {
		if field.IsKey == "TRUE" {
			return true
		}
	}
	return false
}

func (h *CoderHandler) IDType() string {
	for _, field := range h.Fields {
		if field.IsKey == "TRUE" {
			return tpl.FormatJavaDataType(field.DataType)
		}
	}
	return ""
}

type GenFile struct {
	Key      string
	Path     string
	Suffix   string
	Template string
}

func (h *CoderHandler) javaFiles() (files []*GenFile) {
	files = make([]*GenFile, 0)
	files = append(files, &GenFile{Key: "Controller", Path: "controller", Suffix: "Controller.java", Template: tpl.ControllerTemplate})
	files = append(files, &GenFile{Key: "Service", Path: "service", Suffix: "Service.java", Template: tpl.ServiceTemplate})
	files = append(files, &GenFile{Key: "Mapper", Path: "mapper", Suffix: "Mapper.java", Template: tpl.MapperTemplate})
	files = append(files, &GenFile{Key: "MapperXml", Path: "mapper/xml", Suffix: "Mapper.xml", Template: tpl.MapperXmlTemplate})
	files = append(files, &GenFile{Key: "Entity", Path: "entity", Suffix: ".java", Template: tpl.EntityTemplate})
	if h.conf.GenFrontEnd {
		// files = append(files, &GenFile{Key: "Vue", Path: "vue/views", Suffix: ".vue", Template: tpl.VueTemplate})
	}
	return
}

func (h *CoderHandler) goFiles() (files []*GenFile) {
	files = make([]*GenFile, 0)
	files = append(files, &GenFile{Key: "Handler", Path: "handler", Suffix: ".go", Template: tpl.HandlerTemplate})
	files = append(files, &GenFile{Key: "Model", Path: "model", Suffix: ".go", Template: tpl.ModelTemplate})
	if h.conf.GenFrontEnd {
		// files = append(files, &GenFile{Key: "Vue", Path: "vue/views", Suffix: ".vue", Template: tpl.VueTemplate})
	}
	return
}
