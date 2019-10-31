package mdls

type NewNote struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Note struct {
	ID              int64  `json:"id"`
	AuthorAccountID int64  `json:"author_account_id"`
	CreatedAt       int64  `json:"created_at"`
	Title           string `json:"title"`
	Body            string `json:"body"`
}
