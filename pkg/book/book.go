package book

import "time"

type Book struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Total_pages    int       `json:"total_pages"`
	Isbn           string    `json:"isbn"`
	Publisher_date time.Time `json:""`
}
type Publishers struct {
	Publishers_ID   int    `json:"publishers_id"`
	Publishers_name string `json:"name"`
}

type Authors struct {
	Authors_ID   int    `json:"author_id"`
	Authors_name string `json:"author_name"`
}

type Genres struct {
	Genres_ID int    `json:"genres_id"`
	Genres    string `json:"genres"`
	//Genres_Parent_ID int
}
