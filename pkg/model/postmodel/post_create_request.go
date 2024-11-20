package postmodel

type PostCreateRequest struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
}
