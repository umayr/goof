package types
import "time"

type Meta struct {
	Author      string      `json:"author"`
	Content     string      `json:"content"`
	Description string      `json:"description"`
	Time        time.Time   `json:"time"`
	Categories  []string    `json:"categories"`
	Tags        []string    `json:"tags"`
}