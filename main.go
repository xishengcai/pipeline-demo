package main

import (
	"fmt"
	"time"
)

func main(){
	for true  {
		fmt.Println(time.Now())
		time.Sleep(100000000)
	}
}