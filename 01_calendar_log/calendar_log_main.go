//go:generate go build calendar_log_main.go calendar_json.go calendar_struct.go
//go:generate echo build success

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
)

type Config struct {
	JsonDir string `config:"jsonDir"`
}

var conf string
var JsonDir string

func init() {
	flag.StringVar(&conf, "conf", "", "file to read from")
}

func main() {
	flag.Parse()
	if conf == "" {
		conf = "./config"
	}
	cfg := Config{JsonDir}
	loader := confita.NewLoader(
		file.NewBackend(conf + "/config.json"),
		flags.NewBackend(),
	)
	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		fmt.Printf("[CONFIG ERROR]%s\n", err)
		return
	}
	JsSl, err := ReadJsonDir(cfg.JsonDir)
	if err != nil {
		fmt.Printf("[Invalid JsonDir]%v\n", err)
		return
	}
	kek := CrStruct{}
	for _, Js := range JsSl {
		Js.JsonInsert(&kek)
	}
	_ = kek.PrintAllEvents()
}
