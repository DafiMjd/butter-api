package web

type UserCreateRequest struct {
	UserId  string `json:"userId"`
	Content string `json:"content"`
}
