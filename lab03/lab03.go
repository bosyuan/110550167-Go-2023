package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	// 解析URL中的參數
	query := r.URL.Path
	params := strings.Split(query, "/")
	if len(params) != 4 {
		fmt.Fprint(w, "Error!")
		return
	}
	op := params[1]
	a, err := strconv.Atoi(params[2])
	b, err := strconv.Atoi(params[3])
	for _, char := range params[2] {
		if char < '0' || char > '9' {
			fmt.Fprint(w, "Error!")
			return
		}
	}
	for _, char := range params[3] {
		if char < '0' || char > '9' {
			fmt.Fprint(w, "Error!")
			return
		}
	}
	if err != nil {
		fmt.Fprint(w, "Error!")
		return
	}

	if op == "add" {
		result := a + b
		fmt.Fprintf(w, "%d + %d = %d", a, b, result)
	} else if op == "sub" {
		result := a - b
		fmt.Fprintf(w, "%d - %d = %d", a, b, result)
	} else if op == "mul" {
		result := a * b
		fmt.Fprintf(w, "%d * %d = %d", a, b, result)
	} else if op == "div" {
		if b == 0 {
			fmt.Fprint(w, "Error!")
			return
		}
		result := a / b
		remain := a % b
		fmt.Fprintf(w, "%d / %d = %d, reminder = %d", a, b, result, remain)
	} else {
		fmt.Fprint(w, "Error!")
	}

	// 回傳計算結果
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
