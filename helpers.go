package main

import (
	"encoding/json"
	"fmt"
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

func sendMessage(text string) {
	http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%v&text=%s", "846079612:AAEAH62nzgQWY51hEoRsupx9OgdDIPAsMl8", 253637452, url.QueryEscape(text)))
}
