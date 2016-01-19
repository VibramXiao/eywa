package models

import (
	"fmt"
	. "github.com/vivowares/octopus/Godeps/_workspace/src/gopkg.in/olivere/elastic.v3"
	. "github.com/vivowares/octopus/configs"
	. "github.com/vivowares/octopus/utils"
	"log"
	"strings"
)

var IndexClient *Client

func CloseIndexClient() error {
	return nil
}

func InitializeIndexClient() error {
	url := fmt.Sprintf("http://%s:%d", Config().Indices.Host, Config().Indices.Port)
	client, err := NewClient(
		SetURL(url),
		setLogger(ESLogger),
	)
	if err != nil {
		return err
	}
	_, _, err = client.Ping(url).Do()
	if err != nil {
		return err
	}
	IndexClient = client
	return nil
}

func setLogger(logger *log.Logger) func(*Client) error {
	switch strings.ToUpper(Config().Logging.Indices.Level) {
	case "INFO":
		return SetInfoLog(logger)
	case "WARN", "ERROR", "CRITICAL":
		return SetErrorLog(logger)
	case "DEBUG":
		return SetTraceLog(logger)
	default:
		return SetErrorLog(logger)
	}
}
