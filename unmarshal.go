package main

import (
	"encoding/json"
	"fmt"
)

type SenosrReading struct {
	Name     string `json : "name"`
	Capacity int    `json : "capacity"`
	Time     string `json : "time"`
}

func main() {
	jsonString := `{"name" : "battery sensor", "capacity" : 40, "time" : "2020-01-21T19:07:28Z"}`

	var reading SenosrReading

	err := json.Unmarshal([]byte(jsonString), &reading)

	if err != nil {
		fmt.Printf(err)
	}

	fmt.Printf("%+v\n", reading)
}
