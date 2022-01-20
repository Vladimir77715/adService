package database

type Ad struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Price       int      `json:"price"`
	ImageLinks  []string `json:"image_links,omitempty"`
}
