package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name     string
		Relation string
		Phone    string
		Age      uint8
		Family   []Person
		Friends  []string
	}
	t := Person{
		Name:     "Rick Culpepper",
		Relation: "Self",
		Phone:    "615-521-5763",
		Age:      64,
		Family: []Person{
			{
				Name:     "Josie Culpepper",
				Relation: "Spouse",
				Phone:    "615-428-0991",
				Age:      50,
			},
			{
				Name:     "Marshall Culpepper",
				Relation: "Son",
				Age:      41,
				Family: []Person{
					{
						Name:     "Paige Culpepper",
						Relation: "Spouse",
						Age:      41,
					},
					{
						Name:     "Kayla Culpepper",
						Relation: "Daughter",
						Age:      14,
					},
					{
						Name:     "Wyatt Culpepper",
						Relation: "Son",
						Age:      10,
					},
				},
			},
		},
		Friends: []string{
			"Tommy Gautier",
			"Ray \"Buzzard\" Hudson",
			"David",
			"John Flores",
		},
	}
	if ba, e := json.MarshalIndent(t, "", "  "); e != nil {
		panic(e)
	} else {
		fmt.Println(string(ba))
	}
}
