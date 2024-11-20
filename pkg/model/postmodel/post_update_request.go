package postmodel

type PostUpdateRequest struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
