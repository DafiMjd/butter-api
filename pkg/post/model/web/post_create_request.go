package web

type PostCreateRequest struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
}
