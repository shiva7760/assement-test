package src

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type BlogPost struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p BlogPost) String() string {
	return fmt.Sprintf("%s %s (ID: %d) ", p.Title, p.Content, p.ID)
}
