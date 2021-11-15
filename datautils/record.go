package datautils

type Record struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	ThumbURL  string   `json:"thumb_url"`
	Tags      []string `json:"tags"`
	UpdatedAt int64    `json:"updated_at"`
	ImageURLs []string `json:"image_urls"`
}

func (r *Record) getItem(item string) string {
	var res string
	switch item {
	case "title":
		res = r.Title
	case "content":
		res = r.Content
	case "thubmURL":
		res = r.ThumbURL
	}
	return res
}
