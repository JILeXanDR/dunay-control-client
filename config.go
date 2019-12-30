package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type botAPIConfig struct {
	Endpoint   string   `json:"endpoint"`
	Token      string   `json:"token"`
	Recipients []string `json:"recipients"`
}

type venbestConfig struct {
	Server     string `json:"server"`
	Port       int    `json:"port"`
	Username   string `json:"username"`
	PPKNumber  uint   `json:"ppk_number"`
	Password   string `json:"password"`
	LicenseKey []uint `json:"license_key"`
}

type Config struct {
	DebugServerPort string        `json:"debug_server_port"`
	AESPassword     string        `json:"aes_password"`
	RSAPublicKey    string        `json:"rsa_public_key"`
	BotAPI          botAPIConfig  `json:"bot_api"`
	Venbest         venbestConfig `json:"venbest"`
}

func readConfig(path string, config *Config) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &config); err != nil {
		return err
	}

	return nil
}
