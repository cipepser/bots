package facebook

import (
	"net/http"
	"net/url"

	"github.com/cipepser/bots/util"
	"github.com/cipepser/httpclient/sdk"
)

type Feed struct {
	Data []struct {
		CreatedTime string `json:"created_time"`
		Message     string `json:"message"`
		ID          string `json:"id"`
		Story       string `json:"story,omitempty"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

func GetFeed(URL string) (*Feed, error) {
	t := &Token{}
	err := util.GetToken("./token/facebook_app.json", t)
	if err != nil {
		return nil, err
	}

	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return nil, err
	}

	c := &http.Client{}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	vals := url.Values{}
	vals.Add("access_token", t.AccessToken)
	req.URL.RawQuery = vals.Encode()

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	feed := Feed{}
	err = sdk.DecodeBody(resp, &feed)
	if err != nil {
		return nil, err
	}

	return &feed, nil

}
