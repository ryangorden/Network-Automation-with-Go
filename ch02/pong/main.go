package main

import (
	"fmt"

	"github.com/ryangorden/Network-Automation-with-Go/ch02/ping"
)

func main() {
	s := ping.Send()
	fmt.Println(s)
}
