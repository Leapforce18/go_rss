package rss

import (
	"encoding/xml"
	"net/http"
)

/*type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}*/

type Entry struct {
	Title string `xml:"title"`
	/*Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Guid        string    `xml:"guid"`
	Enclosure   Enclosure `xml:"enclosure"`
	PubDate     string    `xml:"pubDate"`*/
}

type Atom struct {
	Title string `xml:"title"`
	/*Link           string `xml:"link"`
	Description    string `xml:"description"`
	Language       string `xml:"language"`
	Copyright      string `xml:"copyright"`
	ManagingEditor string `xml:"managingEditor"`
	WebMaster      string `xml:"webMaster"`
	Ttl            int    `xml:"ttl"`
	PubDate        string `xml:"pubDate"`
	Generator      string `xml:"generator"`*/
	Entries []Entry `xml:"entry"`
}

// ReadAtom reads rss2.0 feed
//
func ReadAtom(url string) (*Atom, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	atom := Atom{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&atom)
	if err != nil {

		return nil, err
	}

	return &atom, nil
}
