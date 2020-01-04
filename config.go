package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"strconv"
	"strings"
)

type botAPIConfig struct {
	Endpoint   string   `mapstructure:"endpoint"`
	Token      string   `mapstructure:"token"`
	Recipients []string `mapstructure:"recipients"`
}

type venbestConfig struct {
	Server     string `mapstructure:"server"`
	Port       int    `mapstructure:"port"`
	Username   string `mapstructure:"username"`
	PPKNumber  uint   `mapstructure:"ppk_number"`
	Password   string `mapstructure:"password"`
	LicenseKey []uint `mapstructure:"license_key"`
}

type Config struct {
	Port         uint          `mapstructure:"port"`
	AESPassword  string        `mapstructure:"aes_password"`
	RSAPublicKey string        `mapstructure:"rsa_public_key"`
	BotAPI       botAPIConfig  `mapstructure:"bot_api"`
	Venbest      venbestConfig `mapstructure:"venbest"`
}

var defaults = map[string]interface{}{
	"port":           8000,
	"aes_password":   "",
	"rsa_public_key": "",

	"bot_api.endpoint":   "",
	"bot_api.token":      "",
	"bot_api.recipients": []string{},

	"venbest.server":      "",
	"venbest.port":        0,
	"venbest.username":    "",
	"venbest.ppk_number":  0,
	"venbest.password":    "",
	"venbest.license_key": []uint{},
}

func readConfig(dst string, config *Config) error {
	dir, basename := path.Split(dst)
	name := strings.TrimSuffix(basename, path.Ext(basename))

	v := viper.New()

	v.AddConfigPath(dir)
	v.SetConfigName(name)

	v.SetEnvPrefix("app")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Config file not found:", dst)
	} else {
		fmt.Println("Using config file:", v.ConfigFileUsed())
	}

	if err := v.Unmarshal(&config); err != nil {
		return err
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		config.Port = uint(defaults["port"].(int))
	} else {
		config.Port = uint(port)
	}

	return nil
}
