package models

// Article ...
type Article struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author Author `json:"author"`
}

// Author ...
type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
