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

	fmt.Println("current time: ", t.UTC())

	if err == nil {
		fmt.Println("exact time: ", et.UTC().Round(0))
	} else {
		log.Fatalf(err.Error())
	}
}
