package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
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
