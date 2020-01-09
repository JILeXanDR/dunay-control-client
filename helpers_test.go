package main

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	println(formatTime(time.Now()))
}

func TestFormatTime1(t *testing.T) {
	println(getCloseEmotion(time.Now().Add(6 * time.Hour)))
}
