package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	//"math/rand"
)

func main() {
	storage := make([]byte, 1024)

	var (
		envMegabytes       string
		envIntervalSeconds string
		megabytes          int
		intervalSeconds    int
	)
	envMegabytes, envMegabytesDefined := os.LookupEnv("MEGABYTES")
	envIntervalSeconds, envIntervalSecondsDefined := os.LookupEnv("INTERVAL_SECONDS")
	//envTerminate, envTerminateDefined := os.LookupEnv("TERMINATE_ON_COMPLETION")

	if envMegabytesDefined == true {
		megabytes, _ = strconv.Atoi(envMegabytes)
	} else {
		megabytes = 500
	}

	if envIntervalSecondsDefined == true {
		intervalSeconds, _ = strconv.Atoi(envIntervalSeconds)
	} else {
		intervalSeconds = 1
	}

	fmt.Println("MEGABYTES=", megabytes)
	fmt.Println("INTERVAL_SECONDS=", intervalSeconds)

	fmt.Println("Starting loop...")
	for i := 1; i < megabytes; i++ {
		storage = append(storage, make([]byte, 1024*512)...)
		fmt.Println(i)
		time.Sleep(time.Duration(intervalSeconds) * time.Second)
	}
	fmt.Println("Loop done")

	for {
		fmt.Println("Just waiting...")
		time.Sleep(time.Duration(1) * time.Second)
	}

}
