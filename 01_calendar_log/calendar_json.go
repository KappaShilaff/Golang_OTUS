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

func (t EventJson) JsonInsert(d InserIventInterface) error{
	err := d.EventInsert(EventDate{t.Age, t.Month, t.Day}, EventTime{t.Start, t.End, t.Name})
	if err != nil {
		return fmt.Errorf("[Insert ERROR] %s. %v", t.Name, err)
	}
	return nil
}

func ReadJsonDir(dir string) ([]EventJson, error) {
	var errSum error
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
			continue
		}
		if err != nil {
			errSum = fmt.Errorf("%v\n[ERROR fileText] %v", errSum, err)
		} else {
			if err = json.Unmarshal(fileText, &mp[i]); err != nil {
				errSum = fmt.Errorf("%v\n[ERROR Unmarshal] %v", errSum, err)
			}
		}
	}
	return mp, errSum
}
