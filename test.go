package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		err := eval(text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print("> ")
	}
}

func parse(text string) (aFloat []float64, aOp []string, err error) {
	expr := strings.Split(text, " ")

	//Check valid number of calculation
	if len(expr)%2 != 1 {
		err = errors.New("The number of elements does not valid")
		return
	}

	//add first number to array number
	var tempFloat float64
	tempFloat, err = strconv.ParseFloat(expr[0], 64)
	if err != nil {
		return
	} else {
		aFloat = append(aFloat, tempFloat)
	}

	for i := 1; i < len(expr); i = i + 2 {
		//add operator to array
		if expr[i] == "+" || expr[i] == "-" || expr[i] == "*" || expr[i] == "/" {
			aOp = append(aOp, expr[i])
		} else {
			err = errors.New("Invalid operator")
			return
		}

		//add number to array
		tempFloat, err = strconv.ParseFloat(expr[i+1], 64)
		if err != nil {
			return
		} else {
			aFloat = append(aFloat, tempFloat)
		}

		//check division zero
		if expr[i] == "/" && tempFloat == 0 {
			err = errors.New("Divide by zero")
			return
		}
	}
	return
}

func eval(text string) (err error) {
	var aFloat []float64
	var aOp []string

	var total float64
	aFloat, aOp, err = parse(text)
	if err != nil {
		return
	}

	total = aFloat[0]

	for i := 0; i < len(aOp); i++ {
		switch aOp[i] {
		case "+":
			total = total + aFloat[i+1]
		case "-":
			total = total - aFloat[i+1]
		case "*":
			total = total * aFloat[i+1]
		case "/":
			total = total / aFloat[i+1]
		}
	}
	fmt.Println(text + " = " + strconv.FormatFloat(total, 'f', -1, 64))
	return
}
