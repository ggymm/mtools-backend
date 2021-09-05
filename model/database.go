package model

import (
	"errors"
	"github.com/google/wire"
	"xorm.io/xorm"
)

var DatabaseModelSet = wire.NewSet(wire.Struct(new(DatabaseModel), "*"))

type DatabaseModel struct {
	DB *xorm.Engine
}

type Database struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INTEGER"`
	ShowName string `json:"showName" xorm:"TEXT"`
	Driver   string `json:"driver" xorm:"TEXT"`
	Host     string `json:"host" xorm:"TEXT"`
	Port     string `json:"port" xorm:"TEXT"`
	Name     string `json:"name" xorm:"TEXT"`
	Username string `json:"username" xorm:"TEXT"`
	Password string `json:"password" xorm:"TEXT"`
}

func (m *Database) TableName() string {
	return "database"
}

func (m *DatabaseModel) GetList() ([]*Database, error) {
	dbList := make([]*Database, 0)
	err := m.DB.Find(&dbList)
	if err != nil {
		return nil, err
	}
	return dbList, nil
}

func (m *DatabaseModel) Get(id int) (*Database, error) {
	var db = new(Database)
	db.Id = id
	if has, err := m.DB.Get(db); err == nil {
		if has {
			return db, nil
		} else {
			return nil, errors.New("数据不存在")
		}
	} else {
		return nil, err
	}
}
