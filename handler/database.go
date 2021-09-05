package handler

import (
	"database/sql"
	"go.uber.org/zap"
	"mtools-backend/model"
	"mtools-backend/schema"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"mtools-backend/config"
)

var DatabaseHandlerSet = wire.NewSet(wire.Struct(new(DatabaseHandler), "*"))

type DatabaseHandler struct {
	Logger        *zap.SugaredLogger
	Config        *config.GlobalConfig
	DatabaseModel *model.DatabaseModel
}

type table struct {
	TableName    string         `json:"tableName" db:"TABLE_NAME"`
	TableComment sql.NullString `json:"tableComment" db:"TABLE_COMMENT"`
	Checked      bool           `json:"checked"`
}

func (h *DatabaseHandler) GetDatabaseList(c *gin.Context) {
	list, err := h.DatabaseModel.GetList()
	if err != nil {
		returnFailed(c, err.Error())
		return
	}
	returnSuccess(c, list)
	return
}

func (h *DatabaseHandler) GetTableList(c *gin.Context) {
	query := new(schema.DatabaseQueryParam)
	if err := c.ShouldBindQuery(&query); err != nil {
		returnFailed(c, validatorErrorData(err))
		return
	}

	dbConfig, err := h.DatabaseModel.Get(query.DatabaseId)
	if err != nil {
		returnFailed(c, err.Error())
		return
	}

	if err, tables := h.getTableList(dbConfig); err != nil {
		h.Logger.Errorf("获取数据库 %s 表出现错误: %s", dbConfig.Name, err)
		returnFailed(c, err.Error())
		return
	} else {
		returnSuccess(c, tables)
		return
	}
}

// 获取待生成代码的表信息
func (h *DatabaseHandler) getTableList(c *model.Database) (error, []table) {
	var (
		tableMapList []table
		dbUrl        = c.Username + ":" + c.Password + "@tcp("
	)
	if strings.Contains(c.Host, ":") {
		dbUrl += c.Host + ")/"
	} else {
		dbUrl += c.Host + ":3306)/"
	}
	dbUrl += c.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, _ := sqlx.Open("mysql", dbUrl)
	defer func() {
		_ = db.Close()
	}()
	if err := db.Select(&tableMapList, config.QueryTableListSQL, c.Name); err != nil {
		return err, tableMapList
	}
	return nil, tableMapList
}
