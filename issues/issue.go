package issues

// Issue model
type Issue struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	FinishedAt  string `json:"-"`
}

// IssueDTO to bind requests
type IssueDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	AuthorID    int    `json:"author" validate:"required"`
}
