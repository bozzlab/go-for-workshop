package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// func main() {
// 	fmt.Println("idiot")
// }
var grade string

func main() {

	dat, err := ioutil.ReadFile("grade.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(strings.NewReader(string(dat)))

	rec, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	rec = rec[1:]
	for _, row := range rec {
		var g float64
		if s, err := strconv.ParseFloat(row[7], 64); err == nil {
			g = s
		}
		fmt.Println(row[0], score(int(g)))
	}
}

// Score is funcion
func score(score int) string {
	if score < 60 {
		grade = "F"
	} else if score < 70 {
		grade = "D"
	} else if score < 80 {
		grade = "C"
	} else if score < 90 {
		grade = "B"
	} else {
		grade = "A"
	}
	return grade
}
