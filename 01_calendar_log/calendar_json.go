package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)


type EventJson struct {
	Name  string `json:""`
	Start int    `json:""`
	End   int    `json:""`
	Age   int    `json:""`
	Month int    `json:""`
	Day   int    `json:""`
}

func (t EventJson) JsonInsert(d InserIventInterface) {
	_ = d.EventInsert(EventDate{t.Age, t.Month, t.Day}, EventTime{t.Start, t.End, t.Name})
}

func ReadJsonDir(dir string) ([]EventJson, error) {
	out, err := exec.Command("ls", dir).Output()
	if err != nil {
		return nil, fmt.Errorf("[ERROR ls] %v", err)
	}
	if len(out) == 0 {
		return nil, nil
	}
	outSl := strings.Split(string(out), "\n")

	mp := make([]EventJson, len(outSl)-1)
	for i, fileName := range outSl {
		fileText, err := ioutil.ReadFile(dir + "/" + fileName)
		if fileName == "" {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("[ERROR fileText] %v", err)
		}
		_ = json.Unmarshal(fileText, &mp[i])
	}
	return mp, nil
}
