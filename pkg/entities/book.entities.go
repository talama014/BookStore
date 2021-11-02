package entities

import "time"

type Book struct {
	ID             uint      `json:"-"`
	Title          string    `json:"title"`
	Total_pages    int       `json:"total_pages"`
	Isbn           string    `json:"isbn"`
	Quarantine     uint      `json:"quarantine"`
	Publisher_date time.Time `json:"publisher_date"`
	Description    string    `json:"description"`
	AuthorID       uint      `json:"-"`
	PublisherID    uint      `json:"-"`

	PublishersRefer uint      `json:"-"`
	Publishers      Publisher `json:"publisher,omitempty" gorm:"foreignKey:PublishersID;references:ID"`

	AuthorRefer uint   `json:"-"`
	Authors     Author `json:",omitempty" gorm:"foreignKey:AuthorID;references:ID"`
}
