package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Addr        string
	DSN         string
	TOKEN       string
	MaxIdleConn int
}

var config Config

func Load() error {
	config.Addr = os.Getenv("ADDR")
	config.DSN = os.Getenv("DSN")
	config.TOKEN = os.Getenv("TOKEN")
	max_idle, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONN"))
	if err != nil {
		fmt.Println("err max_idle_conn")
		return err
	}
	config.MaxIdleConn = max_idle
	return nil
}

func Get() Config {
	return config
}
