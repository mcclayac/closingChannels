package main

import (
	"fmt"
	"time"
)
/*
func printer(msg string, goCh chan bool) {
	<-goCh
	fmt.Printf("msg : %s\n", msg)
}
*/

func printer(msg string, stopCh chan bool) {
	for {
		select {
		case <-stopCh:
			return

			default:
				fmt.Printf("msg : %s\n", msg)

		}
	}
}

func main() {


/*	goCh := make(chan bool)*/
	goCh := make(chan bool)

	for i:= 0; i < 10; i++ {
		go printer(fmt.Sprintf("printer: %d", i), goCh)
	}

	time.Sleep(5 * time.Second)
	close(goCh)
	time.Sleep(5 * time.Second)


}
