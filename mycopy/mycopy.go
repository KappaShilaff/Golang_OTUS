package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var f string
var t string
var o int
var l int

type Argv struct {
	from   string
	to     string
	offset int
	limit  int
}

func (a Argv) FromTo() (string, string, error) {
	if a.from == "" || a.to == "" {
		return "", "", fmt.Errorf("[ERROR] -from or -to are empty")
	}
	return a.from, a.to, nil
}

func (a Argv) OffsetLimit() (int, int) {
	return a.offset, a.limit
}

func (a Argv) Offset() int {
	return a.offset
}

func (a Argv) Limit() int {
	return a.limit
}

type ArgvInterface interface {
	FromTo() (string, string, error)
	OffsetLimit() (int, int)
	Limit() int
	Offset() int
}

func mycopyWithOffsetLimit(st ArgvInterface, rFile *os.File, wFile *os.File) error {
	_, err := rFile.Seek((int64)(st.Offset()), io.SeekCurrent)
	if err != nil {
		return err
	}
	if st.Limit() == 0 {
		_, err = io.Copy(wFile, rFile)
		if err != nil {
			return fmt.Errorf("[ERROR] io.Copy error: %s", err)
		}
	} else {
		_, err = io.CopyN(wFile, rFile, (int64)(st.Limit()))
		if err != nil {
			return fmt.Errorf("[ERROR] io.CopyN error: %s", err)
		}
		return nil
	}
	return nil
}

func mycopy(st ArgvInterface) error {
	from, to, err := st.FromTo()
	if err != nil {
		return err
	}
	rFile, err := os.Open(from)
	if err != nil {
		return fmt.Errorf("[ERROR] os.Open(%s) error: %s", from, err)
	}
	wFile, err := os.Create(to)
	if err != nil {
		return fmt.Errorf("[ERROR] os.Create(%s) error: %s", to, err)
	}
	offset, limit := st.OffsetLimit()
	if offset != 0 || limit != 0 {
		err = mycopyWithOffsetLimit(st, rFile, wFile)
		if err != nil {
			return err
		}
		return nil
	}
	_, err = io.Copy(wFile, rFile)
	if err != nil {
		return fmt.Errorf("[ERROR] io.Copy error: %s", err)
	}
	return nil
}

func init() {
	flag.StringVar(&f, "from", "", "file to read from")
	flag.StringVar(&t, "to", "", "file to write to")
	flag.IntVar(&o, "offset", 0, "offset in input file")
	flag.IntVar(&l, "limit", 0, "limit in input file")
}

func main() {
	flag.Parse()
	st := Argv{f, t, o, l}
	err := mycopy(st)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
