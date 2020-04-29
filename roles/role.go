package roles

// Role model
type Role struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}


