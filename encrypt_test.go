package main

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	println(Encrypt("xxx"))
}

func TestDecrypt(t *testing.T) {
	println(Decrypt("c3aaa29f002ca75870806e44086700f62ce4d43e902b3888e23ceff797a7a471"))
}
