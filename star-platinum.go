package main

import (
	"flag"
	"fmt"
	"github.com/DanielHilton/star-platinum/lib"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	endpoint := flag.String("endpoint", "http://localhost", "The endpoint you want to attack")
	standQuantity := flag.Int("quantity", 1, "The amount of stands you want to attack your endpoint with")

	flag.Parse()

	notice := fmt.Sprintf("Attacking %s endpoint with %d stands", *endpoint, *standQuantity)
	fmt.Println(notice)

	stands := assembleStands(*endpoint, *standQuantity)

	for _, stand := range stands {
		go stand.StartAttacking()
	}

	go waitForSignal(signals, done)
	<-done

	for _, stand := range stands {
		go stand.StopAttacking()
	}
	fmt.Println("Exiting")
}

func waitForSignal(sigChan <-chan os.Signal, done chan<- bool) {
	sig := <-sigChan
	fmt.Println(sig)
	done <- true
}

func assembleStands(endpoint string, quantity int) []*star_platinum.Stand {
	var stands []*star_platinum.Stand

	for i := 0; i < quantity; i++ {
		stands = append(stands, star_platinum.NewStand(endpoint, 1))
	}

	return stands
}
