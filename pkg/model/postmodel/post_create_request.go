package postmodel

type PostCreateRequest struct {
	UserId  string
	Content string `json:"content"`
}
