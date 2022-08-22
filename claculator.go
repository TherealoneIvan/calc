package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tmp := make(chan []string)
	var sample string = "21 + 2, 32 + 5222, 7123 - 10, 9 / 3, 10 * 20, 2 * 4"
	output := make(chan string)
	go scan(sample, tmp)
	go calculate(tmp, output)
	for i := range output {
		fmt.Println(i)
	}
}

func calculate(input chan []string, output chan string) {
	for {
		select {
		case a := <-input:
			if len(a) == 0 {
				close(output)
				return
			}
			switch a[1] {
			case "-":
				first, _ := strconv.Atoi(string(a[0]))
				second, _ := strconv.Atoi(string(a[2]))
				res := first - second
				output <- fmt.Sprint(a, " = ", res)
			case "+":
				first, _ := strconv.Atoi(string(a[0]))
				second, _ := strconv.Atoi(string(a[2]))
				res := first + second
				output <- fmt.Sprint(a, " = ", res)
			case "/":
				first, _ := strconv.Atoi(string(a[0]))
				second, _ := strconv.Atoi(string(a[2]))
				res := first / second
				output <- fmt.Sprint(a, " = ", res)
			case "*":
				first, _ := strconv.Atoi(string(a[0]))
				second, _ := strconv.Atoi(string(a[2]))
				res := first * second
				output <- fmt.Sprint(a, " = ", res)
			}

		}
	}
}

func scan(string2 string, input chan []string) {
	sliceString := strings.Split(string2, ",")
	for _, i := range sliceString {
		slices := strings.ReplaceAll(i, " ", "")
		s := make([]string, 3)
		var b = false
		for _, j := range slices {
			switch j {
			case '-':
				s[1] = string(j)
				b = true
			case '+':
				s[1] = string(j)
				b = true
			case '*':
				s[1] = string(j)
				b = true
			case '/':
				s[1] = string(j)
				b = true
			default:
				if b {
					s[2] += string(j)
				} else {
					s[0] += string(j)
				}
			}
		}
		input <- s
	}
	close(input)
}
