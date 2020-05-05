package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	t := time.Now()
	et, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	fmt.Printf("current time: %s\n", t.UTC())

	if err == nil {
		fmt.Printf("exact time: %s\n", et.UTC().Round(0))
	} else {
		log.Fatalf(err.Error())
	}
}
