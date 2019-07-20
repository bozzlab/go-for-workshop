package main

// func main() {
// 	fmt.Println("idiot")
// }
var (
	grade string
	num   string
	b     []byte
)

//Users is Users
// type Users struct {
// 	ID       int    `json:"id"`
// 	Name     string `json:"name"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Address  struct {
// 		Street  string `json:"street"`
// 		Suite   string `json:"suite"`
// 		City    string `json:"city"`
// 		Zipcode string `json:"zipcode"`
// 		Geo     struct {
// 			Lat string `json:"lat"`
// 			Lng string `json:"lng"`
// 		} `json:"geo"`
// 	} `json:"address"`
// 	Phone   string `json:"phone"`
// 	Website string `json:"website"`
// 	Company struct {
// 		Name        string `json:"name"`
// 		CatchPhrase string `json:"catchPhrase"`
// 		Bs          string `json:"bs"`
// 	} `json:"company"`
// }

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

	//NumThai("2019-07-18 23:55:43")
	//fmt.Println(NumThai("2019-07-18 23:55:43"))

	// type text string
	// var t text = "gopher"
	// fmt.Printf("Type: %T \n", t)
	//$ main.text

	// type text string
	// var t string = "gopher"
	// fmt.Printf("Type: %T \n", t)
	//$ string

	// type text string
	// var t text = "gopher"
	// var s string = "gopher"
	// fmt.Println(t == s)
	//$ invalid operation: t == s (mismatched types text and string)

	// type text = string
	// var t text = "gopher"
	// var s string = "gopher"
	// fmt.Println(t == s)
	//$ true

	// b, err := ioutil.ReadFile("users.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var users Users

	// b, err = json.Unmarshal(b, &Users)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// b, err = xml.MarshalIndent(&users, " ", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//anonymous function
	// fmt.Println(func() string {
	// 	return "Anonymous func"
	// }())

	//printfib(10)
	//fibo(10)

	//defer stack
	// i := 10
	// defer func(n int) {
	// 	fmt.Println(n)
	// }(i)

	// // display last value of program
	// defer func() {
	// 	fmt.Println(i)
	// }()

	// i = 20

	// fmt.Println("hehe")

	// 	var v interface{}
	// 	v = 1
	// 	fmt.Println(v)
	// 	v = "m"
	// 	fmt.Println(v)

	// 	varidic()
	// 	varidic("a1", "a2")

	//switch
	// var i interface{} = "hello"
	// switch v := i.(type) {
	// case int:
	// 	fmt.Printf("%T\n", v)
	// case string:
	// 	fmt.Printf("%T\n", v)
	// default:
	// 	fmt.Printf("Undefined")
	// }

}

// func varidic(a ...string) {
// 	for _, v := range a {
// 		fmt.Println(v)
// 	}
// }

// func fibonacci() func() int {
// 	x, y := -1, 1
// 	return func() int {
// 		x, y = y, x+y
// 		return y
// 	}
// }

// func printfib(n int) {
// 	f := fibonacci()
// 	for i := 0; i < n; i++ {
// 		fmt.Println(f())
// 	}
// }

// func fibo(n int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		fmt.Println(x)
// 		x, y = y, x+y
// 	}
// }

// // Clouser function wtf
// func fact() (func(), func() int) {
// 	i := 0
// 	return func() {
// 			i++
// 		}, func() int {
// 			return i
// 		}
// }

// // Score is funcion
// func score(score int) string {
// 	if score < 60 {
// 		grade = "F"
// 	} else if score < 70 {
// 		grade = "D"
// 	} else if score < 80 {
// 		grade = "C"
// 	} else if score < 90 {
// 		grade = "B"
// 	} else {
// 		grade = "A"
// 	}
// 	return grade
// }

// func init() {
// 	fmt.Println("Program Start")
// 	fmt.Println("Init must be do, before everything")
// 	fmt.Println("------")

// }

// // NumThai is function
// func NumThai(s string) string {
// 	m := map[string]string{
// 		"0": "๐",
// 		"1": "๑",
// 		"2": "๒",
// 		"3": "๓",
// 		"4": "๔",
// 		"5": "๕",
// 		"6": "๖",
// 		"7": "๗",
// 		"8": "๘",
// 		"9": "๙",
// 	}
// 	for k, v := range m {
// 		s = strings.Replace(s, k, v, -1)
// 		fmt.Println(s)
// 	}
// 	return s
// }
