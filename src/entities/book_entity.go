package entities

import (
	"time"
)

// BookEntity ...
type BookEntity struct {
	ID                uint64    `json:"id"`
	Title             string    `json:"title"`
	Author            string    `json:"author"`
	PublishingCompany string    `json:"publishingCompany"`
	Edition           uint64    `json:"edition"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         time.Time `json:"deletedAt"`
}
