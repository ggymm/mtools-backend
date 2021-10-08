package database

import (
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
	"xorm.io/core"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func InitXormDB(appPath string) (*xorm.Engine, func(), error) {
	db, cleanFunc, err := NewXormDB(appPath)
	if err != nil {
		return nil, cleanFunc, err
	}
	return db, cleanFunc, nil
}

func NewXormDB(appPath string) (*xorm.Engine, func(), error) {
	if engine, err := xorm.NewEngine("sqlite", appPath+"mtools-backend.db"); err != nil {
		return nil, nil, err
	} else {
		engine.SetTableMapper(core.SnakeMapper{})
		engine.SetColumnMapper(core.GonicMapper{})
		engine.SetMaxIdleConns(20)
		engine.SetMaxOpenConns(100)

		// 日志输出
		fWriter := &lumberjack.Logger{
			Filename: appPath + "log/mtools-backend.database.log",
			MaxAge:   30,
			MaxSize:  256,
			Compress: true,
		}
		engine.SetLogger(log.NewSimpleLogger(io.MultiWriter(fWriter, os.Stdout)))
		engine.ShowSQL(true)
		engine.Logger().SetLevel(log.LOG_DEBUG)

		// 初始化数据库
		//if _, err := engine.ImportFile("_script/db/init_db.sql"); err != nil {
		//	return nil, nil, err
		//}
		cleanFunc := func() {
			_ = engine.Close()
		}
		return engine, cleanFunc, nil
	}
}
