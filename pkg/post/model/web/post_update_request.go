package web

type PostUpdateRequest struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
