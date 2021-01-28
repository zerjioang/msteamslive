package main

import (
	"fmt"
	"github.com/zerjioang/msteamslive/aliver"
	"log"
)

func main(){
	fmt.Println("Starting Microsoft Teams keep-alive faker")
	if err := aliver.Start(); err != nil {
		log.Fatal(err)
	}
}
