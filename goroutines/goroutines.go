package main

import (
	"fmt"
)

var ch = make(chan string, 10)
var chStart = make(chan struct{})
var chError = make(chan error, 10)

func Print() error {
	<-chStart
	ch <- fmt.Sprintf("Print")
	return nil
}

func PrintError() error {
	<-chStart
	return fmt.Errorf("PrintError")
}

func gorontines(tasks []func() error) {
	for i := 0; i < len(tasks); i++ {
		err := tasks[i]()
		if err != nil {
			chError <- err
		}
	}
}

func main() {
	var tasks []func() error
	tasks = append(tasks, Print)
	tasks = append(tasks, PrintError)
	for n := 0; n < 10; n++ {
		go gorontines(tasks)
	}
	close(chStart)
	n := 0
	for {
		select {
		case err := <-chError:
			fmt.Printf("%s\n", err)
			n++
			if n == 5 {
				mem := len(ch)
				//fmt.Printf("len = %v\n", mem)
				for i := 0; i < mem; i++ {
					fmt.Println(<-ch)
				}
				return
			}
		case printPrint := <-ch:
			fmt.Println(printPrint)
		}
	}
}
