package rss

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Guid        string    `xml:"guid"`
	Enclosure   Enclosure `xml:"enclosure"`
	PubDate     string    `xml:"pubDate"`
}

type Channel struct {
	Title          string `xml:"title"`
	Link           string `xml:"link"`
	Description    string `xml:"description"`
	Language       string `xml:"language"`
	Copyright      string `xml:"copyright"`
	ManagingEditor string `xml:"managingEditor"`
	WebMaster      string `xml:"webMaster"`
	Ttl            int    `xml:"ttl"`
	PubDate        string `xml:"pubDate"`
	Generator      string `xml:"generator"`
	Items          []Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}
