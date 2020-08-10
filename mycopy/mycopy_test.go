package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type CreateFileInterface interface {
	Text() string
	Name() string
}

type CreateFileStruct struct {
	path string
	text string
}

func (st CreateFileStruct)Text() string {
	return st.text
}

func (st CreateFileStruct)Name() string {
	return st.path
}

func createTestFile(st CreateFileInterface) {
	file, _ := os.Create(st.Name())
	_, err := file.Write([]byte(st.Text()))
	if err != nil {
		log.Panicf("failed to write: %v", err)
	}
	err = file.Close()
	if err != nil {
		log.Panicf("failed to close: %v", err)
	}
}

func TestMyCopy(t *testing.T) {
	st := CreateFileStruct{"test_file", "text file"}
	createTestFile(st)
	_ = mycopy(Argv{st.path, st.path + "_copy", 0, 0})
	_ = mycopy(Argv{st.path, st.path + "_copyO", 5, 0})
	_ = mycopy(Argv{st.path, st.path + "_copyL", 0, 4})
	_ = mycopy(Argv{st.path, st.path + "_copyOL", 5, 3})
	f1, _ := ioutil.ReadFile(st.path + "_copy")
	assert.Equal(t, fmt.Sprintf("%s", f1), "text file")
	f1, _ = ioutil.ReadFile(st.path + "_copyO")
	assert.Equal(t, fmt.Sprintf("%s", f1), "file")
	f1, _ = ioutil.ReadFile(st.path + "_copyL")
	assert.Equal(t, fmt.Sprintf("%s", f1), "text")
	f1, _ = ioutil.ReadFile(st.path + "_copyOL")
	assert.Equal(t, fmt.Sprintf("%s", f1), "fil")
	_ = os.Remove(st.path)
	_ = os.Remove(st.path +"_copy")
	_ = os.Remove(st.path +"_copyO")
	_ = os.Remove(st.path +"_copyL")
	_ = os.Remove(st.path +"_copyOL")
}
