package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    // Capitalized to export
	Content   string    // Capitalized to export
	CreatedAt time.Time // Capitalized to export
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled '%v' has the following content:\n\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	noteData, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, noteData, 0644)
}
