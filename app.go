package main

import (
  // "fmt"
  "sync"
  "time"
  "os"
  "runtime/trace"
  log "github.com/sirupsen/logrus"
)

func printerMethod(app string, function string, payload int){
	log.Info(app,function," call for ",payload," at ",time.Now())
}

func blockingMethod(app string, instruction int){
	time.Sleep(5 * time.Second)
	printerMethod(app, "Finished", instruction)
}

func blockingMethodWrapper(app string, instruction_feed <-chan int){
	for instruction := range instruction_feed {
		blockingMethod(app, instruction)
	}
	wg.Done()
}

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {
	// hello()
	// fmt.Println("\n\nAbout to Sleep!\n")
	// time.Sleep(10 * time.Second)
	// fmt.Println("Done with Hello examples!")

	trace.Start(os.Stdout)
	defer trace.Stop()

	// Direct example
	// for i := 0; i <= 2; i++ {
	// 	printerMethod("Direct - ","Started",i)
 //  	blockingMethod("Direct - ",i)
 //  }
	
	// unbuffered chan example with one go routine
	// instruction_feed := make(chan int)
	// go blockingMethodWrapper("Unbuffered - ", instruction_feed)
	// wg.Add(1)
	// for i := 0; i <= 3; i++ {
	// 	printerMethod("Unbuffered - ","Started",i)
 //  	instruction_feed <- i
 //  }
 //  close(instruction_feed)
 //  wg.Wait()

  // buffered chan example with one go routine
 //  buffered_instruction_feed := make(chan int, 5)
 //  go blockingMethodWrapper("Buffered - ", buffered_instruction_feed)
	// wg.Add(1)
 //  for i := 0; i <= 7; i++ {
	// 	printerMethod("Buffered - ","Started",i)
 //  	buffered_instruction_feed <- i
 //  }
 //  close(buffered_instruction_feed)
 //  wg.Wait()

  // Unbuffered chan example with multiple go routines
  instruction_feed_2 := make(chan int)
  for i := 0; i <= 6; i++ {
  	printerMethod("Multiple consumers - ","Started",i)
  	wg.Add(1)
  	go blockingMethodWrapper("Multiple consumers - ", instruction_feed_2)
  	instruction_feed_2 <- i
  }
  close(instruction_feed_2)
  wg.Wait()
}