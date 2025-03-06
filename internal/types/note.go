package types

type Note struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author string `json:"author"`
}
