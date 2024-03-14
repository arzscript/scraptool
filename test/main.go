package main

import (
	"fmt"

	"github.com/arzscript/scraptool/tweets"
)

func main() {
    // Get a greeting message and print it.
    message := tweets.Hello("Gladys")
    fmt.Println(message)
    tweets.ViewURL1("https://accuweather.com")
}