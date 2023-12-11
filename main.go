package main

import (
	"fmt"

	"github.com/SoulKyu/go-website/http"
)

func main() {
	r := http.SetupRouter()
	fmt.Println("Hello World !")
	r.Run(":80")
}
