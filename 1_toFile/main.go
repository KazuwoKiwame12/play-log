package main

import (
	"log"
	"os"
	"time"
)

func main() {
	output, err := os.Create("output.log")
	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(output, "result_", log.Lmicroseconds)
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		logger.Printf("%d回目の出力です\n", i)
	}
}
