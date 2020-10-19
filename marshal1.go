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
	p1 := person{
		First: "Aswin",
		Last:  "Lakshmanan",
	}

	p2 := person{
		First: "Rama",
		Last:  "Trivetia",
	}

	xp := []person{p1, p2}
	fmt.Println("go data: %+v\n", xp)

	bs, err := json.Marshal(xp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json:", string(bs))

}
