package main

import (
	"fmt"
	. "github.com/onsi/gomega"
	"math/rand"
)

var parens = []string{"()", "[]", "{}"}

var rng = rand.New(rand.NewSource(rand.Int63()))

func SingleTest(str string, res bool) {

	Expect(ValidBraces(str)).To(Equal(res), fmt.Sprintf("should return %v for %v", res, str))
}

func ValidBraces(braces string) bool {
	var stack []string

	for i := 0; i < len(braces); i++ {
		//fmt.Println(string(braces[i]))
		if braces[i] == '(' || braces[i] == '[' || braces[i] == '{' {

			stack = append(stack, string(braces[i]))
			continue
		}

		if len(stack) < 1 {
			return false
		}

		switch braces[i] {
		case ')':

			x := len(stack) - 1

			if stack[x] == "{" || stack[x] == "[" {
				return false
			}
			stack = stack[:x]

			break

		case '}':

			x := len(stack) - 1

			if stack[x] == "(" || stack[x] == "[" {
				return false
			}
			stack = stack[:x]
			break

		case ']':

			x := len(stack) - 1

			if stack[x] == "(" || stack[x] == "{" {
				return false
			}
			stack = stack[:x]
			break
		}
	}
	if len(stack) < 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(ValidBraces("()"))

}
