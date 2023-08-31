package structure

type User struct {
	Id int `json:"id"`
}

type UserSegment struct {
	User_id int    `json:"user_id"`
	Slug    string `json:"segment_slug"`
}
