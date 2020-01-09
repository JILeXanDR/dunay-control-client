package main

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	println(formatTime(time.Now()))
}
