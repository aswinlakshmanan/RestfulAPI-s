package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
}

func main() {
	j := `[{"First":"Aswin","Last":"Lakshmanan"},{"First":"Rama","Last":"Trivetia"}]`
	fmt.Println("json: ", j)

	xp := []person{}

	err := json.Unmarshal([]byte(j), &xp)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("go data: %+v\n", xp)

	for i, v := range xp {
		fmt.Println(i, v)
	}

}
