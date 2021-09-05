package schema

type DatabaseQueryParam struct {
	DatabaseId int `form:"databaseId" binding:"required"`
}
