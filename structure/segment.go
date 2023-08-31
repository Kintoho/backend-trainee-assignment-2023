package structure

type Segment struct {
	Id   int    `json:"id" db:"id" swaggerignore:"true"`
	Slug string `json:"slug" binding:"required"`
}
