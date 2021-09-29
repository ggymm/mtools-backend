package schema

type GenCode struct {
	DatabaseId int      `json:"databaseId" binding:"required"`
	Tables     []string `json:"tables" binding:"required"`

	OutputCover bool `json:"outputCover"`
	OnlyEntity  bool `json:"onlyEntity"`

	UseLombok bool `json:"useLombok"`
	UseParent bool `json:"useParent"`
	AutoFill  bool `json:"autoFill"`

	UseOriginTable   bool `json:"useOriginTable"`
	UseOriginColumn  bool `json:"useOriginColumn"`
	FormatDateColumn bool `json:"formatDateColumn"`

	AutoClassComment bool `json:"autoClassComment"`

	GenFrontEnd       bool `json:"genFrontEnd"`
	UseCommentAsLabel bool `json:"useCommentAsLabel"`

	Package        string `json:"package" binding:"required"`
	ParentPackage  string `json:"parentPackage"`
	ExcludeColumn  string `json:"excludeColumn"`
	AutoFillColumn string `json:"autoFillColumn"`
	Output         string `json:"output" binding:"required"`
}
