//go:generate go build calendar_log_main.go calendar_json.go calendar_struct.go
//go:generate echo build success

package main

import (
	"context"
	"flag"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	"go.uber.org/zap"
	"log"
)

type Config struct {
	JsonDir string `config:"jsonDir"`
}

var conf string
var JsonDir string

func init() {
	flag.StringVar(&conf, "conf", "", "file to read from")
}

//func Logging() *os.File{
//	f, err := os.OpenFile("info.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
//	if err != nil {
//		log.Fatalf("error opening file: %v", err)
//	}
//	log.SetOutput(f)
//	return f
//}

func main() {
	//f := Logging()
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	slogger := logger.Sugar()
	defer slogger.Sync()
	flag.Parse()
	if conf == "" {
		conf = "./config"
	}
	cfg := Config{JsonDir}
	loader := confita.NewLoader(
		file.NewBackend(conf + "/config.json"),
		flags.NewBackend(),
	)
	err = loader.Load(context.Background(), &cfg)
	if err != nil {
		slogger.Fatalf("[CONFIG]%s\n", err)
		return
	}
	JsSl, err := ReadJsonDir(cfg.JsonDir)
	if err != nil {
		slogger.Fatalf("[Invalid JsonDir]%v\n", err)
		return
	}
	kek := CrStruct{}
	for _, Js := range JsSl {
		if err = Js.JsonInsert(&kek); err != nil {
			slogger.Warn(err)
		}
	}
	if err = kek.PrintAllEvents() ; err != nil {
		slogger.Warn(err)
	}
}
