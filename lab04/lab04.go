package main

import (
	"log"
	"net/http"
	"strconv"
	"html/template"
	"strings"
)

type CalculatorData struct {
	Expression string
	Result string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "error.html")
}

func gcd(a, b int)int{
	if a%b == 0{
		return b
	}
	return gcd(b, a%b)
}

func Lab4(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	params := strings.Split(query, "&")
	op:= ""
	num1 := 0
	num2 := 0
	str1 := ""
	str2 := ""
	data := CalculatorData {
		Expression: "",
		Result: "",
	}
	if len(params) != 3 {
		ErrorHandler(w, r)
		return
	}

	for i, param := range params {
		keyValue := strings.Split(param, "=")
		if len(keyValue) == 2 {

			a := keyValue[1]
			if i == 0 {
				if (keyValue[0] != "op") {
					ErrorHandler(w, r)
					return
				}
				op = a
				continue
			}
			// Attempt to convert the valueStr to an integer
			value, err := strconv.Atoi(a)
			for _, char := range a {
				if char < '0' || char > '9' {
					ErrorHandler(w, r)
					return
				}
			}
			if err != nil {
				ErrorHandler(w, r)
				return
			}

			if i == 1 {
				if (keyValue[0] != "num1") {
					ErrorHandler(w, r)
					return
				}
				num1 = value
				str1 = a
			} else if i == 2 {
				if (keyValue[0] != "num2") {
					ErrorHandler(w, r)
					return
				}
				num2 = value
				str2 = a
			}
		} else {
			print("bad argument")
			ErrorHandler(w, r)
			return
		}
	}

	if (op == "add") {
		data.Expression = str1 + " + " + str2
		data.Result = strconv.Itoa(num1 + num2)
	} else if (op == "sub") {
		data.Expression = str1 + " - " + str2
		data.Result = strconv.Itoa(num1 - num2)
	} else if (op == "mul") {
		data.Expression = str1 + " * " + str2
		data.Result = strconv.Itoa(num1 * num2)
	} else if (op == "div") {
		if (num2 == 0) {
			ErrorHandler(w, r)
			return
		}
		data.Expression = str1 + " / " + str2
		data.Result = strconv.Itoa(num1 / num2)
	} else if (op == "gcd") {
		data.Expression = "GCD(" + str1 + ", " + str2 + ")"
		data.Result = strconv.Itoa(gcd(num1, num2))
	} else if (op == "lcm") {
		data.Expression = "LCM(" + str1 + ", " + str2 + ")"
		data.Result = strconv.Itoa((num1 * num2) / gcd(num1, num2))
	} else {
		ErrorHandler(w, r)
		return
	}

	err := template.Must(template.ParseFiles("index.html")).Execute(w, data)
	if err != nil {
		ErrorHandler(w, r)
		return
	}

}

func main() {
	http.HandleFunc("/", Lab4)
	log.Fatal(http.ListenAndServe(":8084", nil))
}
