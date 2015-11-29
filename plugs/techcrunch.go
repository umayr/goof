package plugs

import (
	"time"
	"goof/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/tj/go-debug"
	"github.com/briandowns/spinner"
)

const (
	baseUrl = "http://techcrunch.com/"
	name = "Tech Crunch"
	layout = "2006-01-02 15:04:05"
)

type TechCrunch struct {
	debug   debug.DebugFunction
	spinner *spinner.Spinner
	page    int

}

func NewTechCrunch() (*TechCrunch) {
	return &TechCrunch{page: 1, debug:debug.Debug(name), spinner: spinner.New(spinner.CharSets[14], 100 * time.Millisecond)}
}

func (t *TechCrunch) Load(url string) (doc *goquery.Document) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	return doc
}

func (t *TechCrunch) Meta(url string, post *models.Post) {
	doc := t.Load(url)

	t.debug("Loading: %s", url);

	m := new(models.Meta)

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		t, e := s.Attr("name")
		if e {
			val, _ := s.Attr("content")
			switch t {
			case "tag":
				m.Tags = append(m.Tags, val)
				break
			case "category":
				m.Categories = append(m.Categories, val)
				break
			case "timestamp":
				t, _ := time.Parse(layout, val)
				m.Time = t
				break
			case "author":
				m.Author = val
				break
			case "description":
				m.Description = val
				break
			case "content":
				m.Content = val
				break
			}

		}
	})

	post.Meta = *m
}

func (t *TechCrunch) Posts(url string) (posts []models.Post) {
	doc := t.Load(url)

	channel := make(chan models.Post)
	s := doc.Find("ul.river li .block-content")

	for i := range s.Nodes {
		go func(s *goquery.Selection) {
			title := s.Find("h2.post-title").Text()
			url, exists := s.Find("h2.post-title > a").Attr("href")

			post := models.Post{Title: title, Url: url, Origin: name}
			if exists {
				t.Meta(url, &post)
			}
			t.debug("%s", post.Json())

			channel <- post
		}(s.Eq(i))
	}

	for {
		select {
		case p := <-channel:
			posts = append(posts, p)
			t.debug("Length types.Posts: %d", len(posts));
			t.debug("Length Nodes: %d", len(s.Nodes));
			if len(s.Nodes) == len(posts) {
				t.spinner.Stop()
				return posts
			}
		case <-time.After(50 * time.Millisecond):
			t.spinner.Start()
		}
	}
}

func (t *TechCrunch) Next() (posts []models.Post) {
	url := baseUrl + "page/" + string(t.page)
	return t.Posts(url)
}