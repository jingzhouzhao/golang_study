package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println(os.Getenv("PATH"))
	fmt.Printf("Go PATH:%s",os.Getenv("GOPATH"))
}