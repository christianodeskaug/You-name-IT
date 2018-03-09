package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)


func main() {

	c := make(chan int)

	go readInput(c)

	time.Sleep(5 * 1e9)

	go addUp(c)

	time.Sleep(5 * 1e9)
}

func readInput(c chan int) {

	var tall1 int
	var tall2 int

	fmt.Println("Sett inn ett tall: ")
	fmt.Scan(&tall1)
	fmt.Println("Sett inn ett annet tall: ")
	fmt.Scan(&tall2)

	c <- tall1 //Sender data via channelen

	c <- tall2

	sum := <-c // Sum og resultat fra channel

	fmt.Println("Resultat: ", sum)

}



func addUp(c chan int) {

	tall1, tall2 := <-c, <-c // Får tak i data fra readInput()

	sum := tall1 + tall2

	c <- sum // Sender resultat tilbake til readInput()

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGINT,)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				fmt.Println("Program avsluttet.")
			}
		}
	}()

	code := <-exit_chan
	os.Exit(code)


}
