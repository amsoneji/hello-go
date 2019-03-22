package main

import (
	"fmt"
	"time"
)

func main() {
	hello()
	fmt.Println("\n\nAbout to Sleep!")
	time.Sleep(10 * time.Second)
	fmt.Println("Done with Hello examples!")
}