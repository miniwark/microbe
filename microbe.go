/*

Main program for the Microbe robots experiment.

*/

package main

import (
	"fmt"

	"github.com/miniwark/microbe/organs"
)

const (
	VERSION = "0.0.1"
)

func main() {
	// main loop of the bot
	for {
		alive := communicate.Alive()
		if alive == "status alive" {
			fmt.Println(alive)
		}
	}
}
