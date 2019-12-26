package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	AESPassword       string   `json:"aes_password"`
	RSAPublicKey      string   `json:"rsa_public_key"`
	BotAPIEndpoint    string   `json:"bot_api_endpoint"`
	BotAPIToken       string   `json:"bot_api_token"`
	BotAPIRecipients  []string `json:"bot_api_recipients"`
	VenbestServer     string   `json:"venbest_server"`
	VenbestPort       int      `json:"venbest_port"`
	VenbestUsername   string   `json:"venbest_username"`
	VenbestPPKNum     uint     `json:"venbest_ppk_num"`
	VenbestPwd        string   `json:"venbest_pwd"`
	VenbestLicenseKey []uint   `json:"venbest_license_key"`
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
