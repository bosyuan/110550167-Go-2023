package main

import ( "fmt"
		"strconv"
)
func main() {
	var n int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	result := Sum(n)
	fmt.Println(result)
}

func Sum(n int64) string {
	// TODO: Finish this function
	var s string
	sum := 0
	for i := 1; i<= int(n); i++ {
		if i%7 == 0 {
			continue
		}
		s += strconv.Itoa(i) + "+"
		sum += i
	}
	strAsBytes := []byte(s)
	strAsBytes[len(strAsBytes)-1] = '='
	s = string(strAsBytes)
	s += strconv.Itoa(sum)
	return s
}
