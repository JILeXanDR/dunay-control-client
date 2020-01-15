package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func genKey(length int) []byte {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return []byte(b.String())
}

func beautyJSON(v interface{}) []byte {
	j, _ := json.MarshalIndent(v, "", "    ")
	return j
}

func sendSkypeMessage(text string, recipients []string) {
	send := func(conversationID string) {
		println("sendSkypeMessage", text)

		body := map[string]interface{}{
			"text":            text,
			"conversation_id": conversationID,
		}

		b, err := json.Marshal(body)
		if err != nil {
			println("err", err.Error())
			return
		}

		req, err := http.NewRequest(http.MethodPost, config.BotAPI.Endpoint, bytes.NewReader(b))
		if err != nil {
			println("err", err.Error())
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Api-Key", config.BotAPI.Token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			println("err", err.Error())
			return
		}
		defer resp.Body.Close()

		parseResult(resp, nil)
	}

	for _, conversationID := range recipients {
		send(conversationID)
	}
}

func parseResult(resp *http.Response, err error) {
	if err != nil {
		println("err", err.Error())
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("err", err.Error())
		return
	}

	println("content", string(content))
}

func filterByFunc(slice []interface{}, filterFunc func(val interface{}) bool) []interface{} {
	var res []interface{}
	for _, val := range slice {
		if filterFunc(val) {
			res = append(res, val)
		}
	}
	return res
}

func castToSliceOfStrings(slice []interface{}) []string {
	var res []string
	for _, val := range slice {
		res = append(res, val.(string))
	}
	return res
}

func castToSliceOfInterfaces(slice []string) []interface{} {
	var res []interface{}
	for _, val := range slice {
		res = append(res, val)
	}
	return res
}

func privateConversations() []string {
	return castToSliceOfStrings(filterByFunc(castToSliceOfInterfaces(config.BotAPI.Recipients), func(val interface{}) bool {
		return strings.HasPrefix(val.(string), "8:")
	}))
}

func formatTime(t time.Time) string {
	return t.Format("15:04:05")
}

func getOpenEmotion(t time.Time) string {
	if t.Hour() >= 9 && t.Minute() >= 0 {
		return "(snail) "
	}
	if t.Hour() >= 8 && t.Minute() >= 0 {
		return "(hedgehog)"
	}
	if t.Hour() >= 7 && t.Minute() >= 0 {
		return "(hendance)"
	}
	if t.Hour() >= 6 && t.Minute() >= 0 {
		return "(werewolfhowl)"
	}
	return "(monkey) "
}

func getCloseEmotion(t time.Time) string {
	if (t.Hour() >= 0 && t.Minute() >= 0) && t.Hour() < 18 {
		return "(drunk)"
	}
	if t.Hour() >= 22 && t.Minute() >= 0 {
		return "(sleepy)"
	}
	if t.Hour() >= 20 && t.Minute() >= 0 {
		return "(waiting)"
	}
	if t.Hour() >= 19 && t.Minute() >= 0 {
		return "(nerdy)"
	}
	if t.Hour() >= 18 && t.Minute() >= 0 {
		return "(wasntme)"
	}
	return "(cool)"
}
