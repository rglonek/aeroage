package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// process the logic
	answer, _ := base64.StdEncoding.DecodeString("SXQncyBhIG5ldyBjYXIK")

	// print the answer
	fmt.Println(string(answer))
}
