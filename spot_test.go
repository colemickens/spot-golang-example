package spot

import (
	"bytes"
	"encoding/json"
	"testing"
)

/*
	TODO:

	extract this into some sort of common
	assistant json testing package

	I've written this pattern too many times.
	Too much repetition.
*/

func testStr(t *testing.T, expected, actual, msg string) {
	if expected != actual {
		t.Error(msg)
		t.Log("Expected:", expected)
		t.Log("Got:", actual)
	}
	return
}

func testInt(t *testing.T, expected, actual int, msg string) {
	if expected != actual {
		t.Error(msg)
		t.Log("Expected:", expected)
		t.Log("Got:", actual)
	}
	return
}

func testFloat64(t *testing.T, expected, actual float64, msg string) {
	if expected != actual {
		t.Error(msg)
		t.Log("Expected:", expected)
		t.Log("Got:", actual)
	}
	return
}

func testBool(t *testing.T, expected, actual bool, msg string) {
	if expected != actual {
		t.Error(msg)
		t.Log("Expected:", expected)
		t.Log("Got:", actual)
	}
	return
}

func TestGetFeed(t *testing.T) {
	_, err := GetFeed("0oCHzmaKo1zRkSHQglD2qqXkT2yJPvzpK")
	if err != nil {
		t.Error(err)
	}
	return
}

func TestDecode(t *testing.T) {
	jsonDec := json.NewDecoder(bytes.NewBufferString(TEST_JSON))
	resp := &SpotApiResponse{}
	if err := jsonDec.Decode(resp); err != nil {
		t.Error("Failed to decode in testdecode:", err)
	}

	testInt(t, 1, resp.Resp.FeedMsgResp.Count, "Count is bad")

	testStr(t, "0oCHzmaKo1zRkSHQglD2qqXkT2yJPvzpK", resp.Resp.FeedMsgResp.Feed.Id, "Feed id is bad")
	testStr(t, "MarkPleskac", resp.Resp.FeedMsgResp.Feed.Name, "Feed name is bad")
	testStr(t, "MarkPleskac", resp.Resp.FeedMsgResp.Feed.Description, "Feed description is bad")
	testStr(t, "ACTIVE", resp.Resp.FeedMsgResp.Feed.Status, "Feed status is bad")
	testInt(t, 1, resp.Resp.FeedMsgResp.Feed.Usage, "Feed usage is bad")
	testInt(t, 7, resp.Resp.FeedMsgResp.Feed.DaysRange, "Feed days range is bad")
	testBool(t, false, resp.Resp.FeedMsgResp.Feed.DetailedMessageShown, "Feed detailed message shown is bad")

	testInt(t, 0, resp.Resp.FeedMsgResp.TotalCount, "Total count is bad")
	testInt(t, 0, resp.Resp.FeedMsgResp.ActivityCount, "Activity count is bad")

	testStr(t, "0", resp.Resp.FeedMsgResp.Messages.Message.AtClientUnixTime, "Message @clientUnixTime is wrong")
	testInt(t, 173211229, resp.Resp.FeedMsgResp.Messages.Message.Id, "Message Id is wrong")
	testStr(t, "0-8247915", resp.Resp.FeedMsgResp.Messages.Message.MessengerId, "Message MessengerId is wrong")
	testInt(t, 1352345702, resp.Resp.FeedMsgResp.Messages.Message.UnixTime, "Message UnixTime is wrong")
	testStr(t, "OK", resp.Resp.FeedMsgResp.Messages.Message.MessageType, "Message MessageType is wrong")
	testFloat64(t, 40.80791, resp.Resp.FeedMsgResp.Messages.Message.Latitude, "Message Latitude is wrong")
	testFloat64(t, -96.7037, resp.Resp.FeedMsgResp.Messages.Message.Longitude, "Message Longitude is wrong")
	testStr(t, "N", resp.Resp.FeedMsgResp.Messages.Message.ShowCustomMsg, "Message ShowCustomMsg is wrong")
	testStr(t, "2012-11-08T03:35:02+0000", resp.Resp.FeedMsgResp.Messages.Message.DateTime, "Message DateTime is wrong")
	testStr(t, "", resp.Resp.FeedMsgResp.Messages.Message.MessageDetail, "Message MessageDetail is wrong")
	testBool(t, false, resp.Resp.FeedMsgResp.Messages.Message.Selected, "Message Selected is wrong")
	testInt(t, 0, resp.Resp.FeedMsgResp.Messages.Message.Altitude, "Message Altitude is wrong")
	testInt(t, 0, resp.Resp.FeedMsgResp.Messages.Message.Hidden, "Message Hidden is wrong")
	testStr(t, "Standard update. All is well.", resp.Resp.FeedMsgResp.Messages.Message.MessageContent, "Message MessageContent is wrong")

}

func TestEncode(t *testing.T) {
	sar := SpotApiResponse{
		Response{
			FeedMessageResponse{
				Count: 1,
				Feed: Feed{
					Id:                   "0oCHzmaKo1zRkSHQglD2qqXkT2yJPvzpK",
					Name:                 "MarkPleskac",
					Description:          "MarkPleskac",
					Status:               "ACTIVE",
					Usage:                1,
					DaysRange:            7,
					DetailedMessageShown: false,
				},
				TotalCount:    0,
				ActivityCount: 0,
				Messages: Messages{
					Message: Message{
						AtClientUnixTime: "0",
						Id:               173211229,
						MessengerId:      "0-8247915",
						MessengerName:    "Mark Pleskac",
						UnixTime:         1352345702,
						MessageType:      "OK",
						Latitude:         40.80791,
						Longitude:        -96.7037,
						ShowCustomMsg:    "N",
						DateTime:         "2012-11-08T03:35:02+0000",
						MessageDetail:    "",
						Selected:         false,
						Altitude:         0,
						Hidden:           0,
						MessageContent:   "Standard update. All is well.",
					},
				},
			},
		},
	}

	b, err := json.MarshalIndent(sar, "", "	")
	if err != nil {
		t.Error("failed to enocde properly")
	}

	testStr(t, TEST_JSON, string(b), "failed to encode properly")
}

const TEST_JSON string = `{
	"response": {
		"feedMessageResponse": {
			"count": 1,
			"feed": {
				"id": "0oCHzmaKo1zRkSHQglD2qqXkT2yJPvzpK",
				"name": "MarkPleskac",
				"description": "MarkPleskac",
				"status": "ACTIVE",
				"usage": 1,
				"daysRange": 7,
				"detailedMessageShown": false
			},
			"totalCount": 0,
			"activityCount": 0,
			"messages": {
				"message": {
					"@clientUnixTime": "0",
					"id": 173211229,
					"messengerId": "0-8247915",
					"messengerName": "Mark Pleskac",
					"unixTime": 1352345702,
					"messageType": "OK",
					"latitude": 40.80791,
					"longitude": -96.7037,
					"showCustomMsg": "N",
					"dateTime": "2012-11-08T03:35:02+0000",
					"messageDetail": "",
					"selected": false,
					"altitude": 0,
					"hidden": 0,
					"messageContent": "Standard update. All is well."
				}
			}
		}
	}
}`
