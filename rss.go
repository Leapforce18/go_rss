package rss

import (
	"encoding/xml"
	"net/http"

	rss2_0 "github.com/leapforce-libraries/go_rss/2.0"
)

// ReadRss2_0 reads rss2.0 feed
//
func ReadRss2_0(url string) (*rss2_0.Rss, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rss := rss2_0.Rss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {

		return nil, err
	}

	return &rss, nil
}
