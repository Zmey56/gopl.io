//Exercise 7.16:
//Write a web-based calculator program.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type PageData struct {
	Expression string
	Result     string
}

func evaluateExpression(expression string) (string, error) {
	// Split the expression into operands and operators
	tokens := strings.Fields(expression)

	// Check if the expression contains at least three tokens
	if len(tokens) < 3 {
		return "", fmt.Errorf("Invalid expression")
	}

	// Convert the first token to a float64 value
	result, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return "", fmt.Errorf("Invalid expression")
	}

	// Iterate over the remaining tokens in pairs
	for i := 1; i < len(tokens)-1; i += 2 {
		operator := tokens[i]
		operand := tokens[i+1]

		// Convert the operand to a float64 value
		value, err := strconv.ParseFloat(operand, 64)
		if err != nil {
			return "", fmt.Errorf("Invalid expression")
		}

		// Perform the corresponding arithmetic operation
		switch operator {
		case "+":
			result += value
		case "-":
			result -= value
		case "*":
			result *= value
		case "/":
			result /= value
		default:
			return "", fmt.Errorf("Invalid expression")
		}
	}

	return fmt.Sprintf("%.2f", result), nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("ch7/Exercise_7_16/index.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		expression := r.FormValue("expression")
		result, err := evaluateExpression(expression)
		if err != nil {
			log.Printf("Error evaluating expression: %v\n", err)
			http.Error(w, "Invalid expression", http.StatusBadRequest)
			return
		}

		data := PageData{
			Expression: expression,
			Result:     result,
		}

		tmpl := template.Must(template.ParseFiles("ch7/Exercise_7_16/result.html"))
		tmpl.Execute(w, data)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
