package schema

type CreateCollectionReq struct {
	ParentId int    `form:"label" binding:"required"`
	Label    string `form:"label" binding:"required"`
	SortNo   int    `form:"sortNo"`
}
