package models

import (
	"testing"
	"time"
)

func TestBook(t *testing.T) {
	book := Book{
		ID:        1,
		Title:     "Clean Code",
		Author:    "Robert C. Martin",
		Price:     200000,
		CreatedAt: time.Now(),
	}

	if book.ID != 1 {
		t.Errorf("ID: kutilgan 1, olindi %d", book.ID)
	}

	if book.Title != "Clean Code" {
		t.Errorf("Title: kutilgan Clean Code, olindi %s", book.Title)
	}

	if book.Author != "Robert C. Martin" {
		t.Errorf("Author: kutilgan Robert C. Martin, olindi %s", book.Author)
	}

	if book.Price != 200000 {
		t.Errorf("Price: kutilgan 200000, olindi %f", book.Price)
	}
}