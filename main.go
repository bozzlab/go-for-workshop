package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	fmt.Println("idiot")
// }
var (
	grade string
	num   string
)

func main() {

	// dat, err := ioutil.ReadFile("grade.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// r := csv.NewReader(strings.NewReader(string(dat)))
	// rec, err := r.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// rec = rec[1:]
	// for _, row := range rec {
	// 	var g float64
	// 	if s, err := strconv.ParseFloat(row[7], 64); err == nil {
	// 		g = s
	// 	}
	// 	fmt.Println(row[0], score(int(g)))
	// }

	// m := map[string]string{
	// 	"A": "apple",
	// 	"B": "bat",
	// 	"C": "cat",
	// }
	// for _, k := range m {
	// 	fmt.Println(k)
	// }
	// //	v, ok := m["Z"]
	// if v, ok := m["X"]; ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println(ok)
	// }

	// if v, ok := m["C"]; !ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println(ok)
	// }

	NumThai("2019-07-18 23:55:43")
	//fmt.Println(NumThai("2019-07-18 23:55:43"))

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

func init() {
	fmt.Println("Program Start")
	fmt.Println("Init must be do, before everything")
	fmt.Println("------")

}

// NumThai is function
func NumThai(s string) string {
	m := map[string]string{
		"0": "๐",
		"1": "๑",
		"2": "๒",
		"3": "๓",
		"4": "๔",
		"5": "๕",
		"6": "๖",
		"7": "๗",
		"8": "๘",
		"9": "๙",
	}
	for k, v := range m {
		s = strings.Replace(s, k, v, -1)
		fmt.Println(s)
	}
	return s
}
