package handler

import (
	"database/sql"
	"go.uber.org/zap"
	"path/filepath"
	"strings"

	"mtools-backend/config"
	"mtools-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var DatabaseHandlerSet = wire.NewSet(wire.Struct(new(DatabaseHandler), "*"))

type DatabaseHandler struct {
	Logger *zap.SugaredLogger
	Config *config.GlobalConfig
}

type DatabaseConfig struct {
	Id       string `json:"id"`
	Host     string `json:"host"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type table struct {
	TableName    string         `db:"TABLE_NAME"`
	TableComment sql.NullString `db:"TABLE_COMMENT"`
}

func (h *DatabaseHandler) GetTableList(c *gin.Context) {
	databaseId := c.GetString("databaseId")
	// 获取全部配置列表
	configList := make([]DatabaseConfig, 0)
	utils.ReadJSON(filepath.Join(h.Config.Coder.ConfigFolder, "database_config.json"), &configList)
	if len(databaseId) == 0 {
		list := make([]map[string]interface{}, 0)
		for _, databaseConfig := range configList {
			if err, tables := h.getTableList(databaseConfig); err != nil {
				h.Logger.Errorf("获取数据库 %s 表出现错误: %s", databaseConfig.Name, err)
				continue
			} else {
				list = append(list, map[string]interface{}{
					"DatabaseName": databaseConfig.Name,
					"Tables":       tables,
				})
			}
		}
		returnSuccess(c, list)
		return
	} else {
		var selectedConfig DatabaseConfig
		for _, databaseConfig := range configList {
			if databaseConfig.Id == databaseId {
				selectedConfig = databaseConfig
				break
			}
		}
		if err, tables := h.getTableList(selectedConfig); err != nil {
			h.Logger.Errorf("获取数据库 %s 表出现错误: %s", selectedConfig.Name, err)
			returnFailed(c, err)
			return
		} else {
			returnSuccess(c, map[string]interface{}{
				"DatabaseName": selectedConfig.Name,
				"Tables":       tables,
			})
			return
		}
	}
}

// 获取待生成代码的表信息
func (h *DatabaseHandler) getTableList(c DatabaseConfig) (error, []table) {
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
