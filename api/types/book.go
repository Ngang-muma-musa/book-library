package types

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Qauntity int    `json:"quantity"`
}