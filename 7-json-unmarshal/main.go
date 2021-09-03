package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	HelloWorld string `json:"HelloWorld"`
}

func main() {
	foo := `{"HelloWorld": "1", "helloWorld": "2"}`
	var dat T
	_ = json.Unmarshal([]byte(foo), &dat)
	fmt.Println(dat)
}
