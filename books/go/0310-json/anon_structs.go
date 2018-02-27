package main

import (
	"encoding/json"
	"fmt"
)

// :show start
var jsonBlob = []byte(`
{
  "_total": 1,
  "_links": {
	"self": "https://api.twitch.tv/kraken/channels/foo/subscriptions?direction=ASC&limit=25&offset=0",
	"next": "https://api.twitch.tv/kraken/channels/foo/subscriptions?direction=ASC&limit=25&offset=25"
  },
  "subscriptions": [
	{
	  "created_at": "2011-11-23T02:53:17Z",
	  "_id": "abcdef0000000000000000000000000000000000",
	  "_links": {
		"self": "https://api.twitch.tv/kraken/channels/foo/subscriptions/bar"
	  },
	  "user": {
		"display_name": "bar",
		"_id": 123456,
		"name": "bar",
		"created_at": "2011-06-16T18:23:11Z",
		"updated_at": "2014-10-23T02:20:51Z",
		"_links": {
		  "self": "https://api.twitch.tv/kraken/users/bar"
		}
	  }
	}
  ]
}
`)

// :show end

func main() {
	// :show start
	var js struct {
		Total int `json:"_total"`
		Links struct {
			Next string `json:"next"`
		} `json:"_links"`
		Subs []struct {
			Created string `json:"created_at"`
			User    struct {
				Name string `json:"name"`
				ID   int    `json:"_id"`
			} `json:"user"`
		} `json:"subscriptions"`
	}

	err := json.Unmarshal(jsonBlob, &js)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", js)
	// :show end
}
