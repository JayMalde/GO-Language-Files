package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study of Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//standard format by go documentation
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-03-2006 Monday"))

	//GOOS="windows" go build for creating an exe executable
	//GOOS="linux" go build for creating an binary executable

}
