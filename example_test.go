package main

import (
	"github.com/JILeXanDR/dunay-control-client/testdata"
	"github.com/JILeXanDR/dunay-control-client/venbest"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"testing"
)

func TestExample(t *testing.T) {
	server := testdata.NewFakeServer()
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	require.NoError(t, err, "parse URL of fake server")

	port, err := strconv.Atoi(serverURL.Port())
	require.NoError(t, err, "convert string port to int")

	f, err := os.Open("./testdata/rsa_key.pub")
	b, err := ioutil.ReadAll(f)

	logger := logrus.New()
	logger.WithField("module", "client")
	logger.SetLevel(logrus.TraceLevel)

	client := venbest.NewClient(venbest.ClientOptions{
		HardCodedLoginData: []byte(``),
		RSAPublicKey:       string(b),
		ServerHost:         serverURL.Hostname(),
		ServerPort:         port,
		Username:           "test",
		PPKNum:             0,
		Pwd:                "",
		LicenseKey:         nil,
		Logger:             logger,
	})

	go func() {
		if err := client.Start(); err != nil {
			logger.WithError(err).Error("starting venbest client")
		}
	}()

	println("tests...")

	t.Run("test1", func(t *testing.T) {
		<-client.States
		println("got state")
	})

	t.Run("test2", func(t *testing.T) {
		<-client.Events
		println("got event")
	})

	t.Run("test3", func(t *testing.T) {
		<-client.Errors
		println("got error")
	})

	println("end.")
}
