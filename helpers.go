package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
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

func sendTelegramMessage(text string) {
	println("sendTelegramMessage", text)
	http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/sendTelegramMessage?chat_id=%v&text=%s", "846079612:AAEAH62nzgQWY51hEoRsupx9OgdDIPAsMl8", 253637452, url.QueryEscape(text)))
}

func sendSkypeMessage(text string) {
	println("sendSkypeMessage", text)

	body := map[string]interface{}{
		"text":            text,
		"conversation_id": "8:jilexandr",
	}

	b, err := json.Marshal(body)
	if err != nil {
		println("err", err.Error())
		return
	}

	resp, err := http.Post("https://api.telegram.org/web/message", "", bytes.NewReader(b))
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
