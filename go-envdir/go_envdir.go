package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func ReadDir(dir string) (map[string]string, error) {
	out, err := exec.Command("ls", dir).Output()
	if err != nil {
		return nil, fmt.Errorf("[ERROR ls] %v", err)
	}
	if len(out) == 0 {
		return nil, nil
	}
	outSl := strings.Split(string(out), "\n")

	mp := make(map[string]string, len(outSl) - 1)
	for _, fileName := range outSl{
		fileText, err := ioutil.ReadFile(dir + "/" + fileName)
		if fileName == "" {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("[ERROR fileText] %v", err)
		}
		if fileText[len(fileText) - 1] == '\n' {
			mp[fileName] = string(fileText[:len(fileText)-1])
		} else {
			mp[fileName] = string(fileText)
		}
	}
	return mp, nil
}

func RunCmd(cmd []string, env map[string]string) int {
	for envName, envString := range env {
		_ = os.Setenv(envName, envString)
	}
	cmdCommand := exec.Command(cmd[0], cmd[1:]...)
	cmdCommand.Stdout = os.Stdout
	cmdCommand.Stderr = os.Stderr
	cmdCommand.Stdin = os.Stdin
	if err := cmdCommand.Run() ; err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
	}
	return 0
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("[ERROR] invalid number of arguments")
		return
	}
	mapEnv, err := ReadDir(args[1])
	if err != nil {
		fmt.Printf("[ERROR mapEnv] %v", err)
		return
	}
	os.Exit(RunCmd(args[2:], mapEnv))
}