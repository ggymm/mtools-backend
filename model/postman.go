package model

import (
	"fmt"
	"github.com/google/wire"
	"mtools-backend/schema"
	"mtools-backend/utils"
	"xorm.io/xorm"
)

var PostmanModelSet = wire.NewSet(wire.Struct(new(PostmanModel), "*"))

type PostmanModel struct {
	DB *xorm.Engine
}

type Collection struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INTEGER"`
	ParentId   int    `json:"parentId" xorm:"INTEGER"`
	ParentPath string `json:"parentPath" xorm:"TEXT"`
	Label      string `json:"label" xorm:"TEXT"`
	SortNo     int    `json:"sortNo" xorm:"INTEGER"`
	Favorite   int    `json:"favorite" xorm:"INTEGER"`
	DelFlag    int    `json:"delFlag" xorm:"INTEGER"`
}

type CollectionTree struct {
	Collection `xorm:"extends"`
	Type       string            `json:"type" xorm:"-"`
	Children   []*CollectionTree `json:"children" xorm:"-"`
}

func (m *CollectionTree) TableName() string {
	return "collection"
}

func (m *PostmanModel) GetTree() ([]*CollectionTree, error) {
	list := make([]*CollectionTree, 0)
	err := m.DB.OrderBy("sort_no").Find(&list)
	if err != nil {
		return nil, err
	}
	tree := m.buildTree(list)
	return tree, nil
}

func (m *PostmanModel) buildTree(list []*CollectionTree) []*CollectionTree {
	tree := make([]*CollectionTree, 0)
	treeMap := make(map[int]*CollectionTree)
	for _, item := range list {
		treeItem := new(CollectionTree)
		treeItem.Id = item.Id
		treeItem.ParentId = item.ParentId
		treeItem.Label = item.Label
		treeItem.Type = "collection"
		if item.ParentId == -1 {
			tree = append(tree, treeItem)
		} else {
			treeMap[item.ParentId].Children = append(treeMap[item.ParentId].Children, treeItem)
		}
		treeMap[item.Id] = treeItem
	}
	return tree
}

func (m *PostmanModel) Create(collectionReq *schema.CreateCollectionReq) error {
	collection := new(Collection)
	if err := utils.Copy(collection, collectionReq); err != nil {
		return err
	}
	collection.DelFlag = 1
	if insert, err := m.DB.Insert(collection); err != nil {
		return err
	} else {
		fmt.Println(insert)
		return nil
	}
}
