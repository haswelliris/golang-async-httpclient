package main

import (
	"fmt"

	"github.com/haswelliris/golang-async-httpclient/httpclient"
)

func main() {
	fmt.Println("The time of synchronize is")
	httpclient.Synchronous()
	fmt.Println("The time of asynchronism way")
	httpclient.Asynchronous()
}
