//go:generate go build calendar_log_main.go calendar_json.go calendar_struct.go
//go:generate echo build success

package main

import (
	"flag"
	"fmt"
)

var conf string
var jsonDir string

func init() {
	flag.StringVar(&conf, "from", "", "file to read from")
	flag.StringVar(&jsonDir, "json", "", "files to read from")
}

func main() {
	flag.Parse()
	if jsonDir == "" {
		fmt.Printf("-json dir PLZ\n")
		return
	}
	JsSl, err := ReadJsonDir(jsonDir)
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
