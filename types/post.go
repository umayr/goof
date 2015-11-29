package types
import "encoding/json"

type Post struct {
	Title string    `json:"title"`
	Meta  Meta      `json:"meta"`
	Url   string    `json:"url"`
}

func (p *Post) Json() (str string) {
	b, _ := json.Marshal(p)
	return string(b)
}