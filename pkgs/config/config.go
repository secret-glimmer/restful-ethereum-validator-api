package config

import (
	"os"

	"github.com/joho/godotenv"
)

type QuickNode struct {
	Http      string
	WebSocket string
}

type Config struct {
	ServerPort string
	QuckNode   QuickNode
}

func NewConfig() *Config {
	return &Config{
		ServerPort: "8000", // Default Server Port
		QuckNode: QuickNode{
			Http:      "",
			WebSocket: "",
		},
	}
}

func (config *Config) LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	config.ServerPort = os.Getenv("SERVER_PORT")
	if config.ServerPort == "" {
		config.ServerPort = "8000"
	}

	config.QuckNode.Http = os.Getenv("QUICK_NODE_HTTP_URL")
	config.QuckNode.WebSocket = os.Getenv("QUICK_NODE_WEBSOCKET_URL")
	return nil
}
