package spot

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
)

type SpotApiResponse struct {
	Resp Response `json:"response"`
}

type Response struct {
	FeedMsgResp FeedMessageResponse `json:"feedMessageResponse"`
}

type FeedMessageResponse struct {
	Count         int      `json:"count"`
	Feed          Feed     `json:"feed"`
	TotalCount    int      `json:"totalCount"`
	ActivityCount int      `json:"activityCount"`
	Messages      Messages `json:"messages"` // this ought to be an array but it's not from Spot? this might break...
}

type Feed struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Status               string `json:"status"`
	Usage                int    `json:"usage"`
	DaysRange            int    `json:"daysRange"`
	DetailedMessageShown bool   `json:"detailedMessageShown"`
}

type Messages struct {
	Message Message `json:"message"`
}

type Message struct {
	AtClientUnixTime string  `json:"@clientUnixTime"`
	Id               int     `json:"id"`
	MessengerId      string  `json:"messengerId"`
	MessengerName    string  `json:"messengerName"`
	UnixTime         int     `json:"unixTime"`
	MessageType      string  `json:"messageType"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	ShowCustomMsg    string  `json:"showCustomMsg"`
	DateTime         string  `json:"dateTime"`
	MessageDetail    string  `json:"messageDetail"`
	Selected         bool    `json:"selected"`
	Altitude         int     `json:"altitude"`
	Hidden           int     `json:"hidden"`
	MessageContent   string  `json:"messageContent"`
}

func GetFeed(feedId string) (*SpotApiResponse, error) {
	// you sure you want to skip verification?
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://api.findmespot.com/spot-main-web/consumer/rest-api/2.0/public/feed/" + feedId + "/message.json")
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	spotResp := &SpotApiResponse{}
	if err := dec.Decode(spotResp); err != nil {
		return nil, err
	}

	return spotResp, nil
}
