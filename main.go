package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

type Request struct {
	ID     int    `db:"id"`
	Value  string `db:"value"`
	Answer string `db:"answer"`
}

var schema string = "CREATE TABLE `requests` (	" +
	"  	`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY," +
	"		  	`value` varchar(255) NOT NULL," +
	"           `answer` varchar(255) NOT NULL	)"

func main() {
	//fmt.Println(ValidBraces("()"))
	insertToBD("()", ValidBraces("()"))
}

func insertToBD(val string, ans bool) {
	conn, err := sqlx.Connect("mysql", "Actor_For_Hits_Back:iu3nYfCE27SKAET@/db_for_test")
	if err != nil {
		panic(err)
	}
	//conn.MustExec(schema)
	stmt := `INSERT INTO requests (value, answer) VALUES(?,?)`
	res, err := conn.Exec(stmt, val, ans)
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created request with id:%d", id)

	var request Request
	err = conn.Get(&request, "select * from requests where id=?", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(request)
	//_, err = conn.Exec("UPDATE requests set value=\"John\" where id=?", id)
	//if err != nil {
	//	panic(err)
	//}
	//_, err = conn.Exec("DELETE FROM requests where id=?", id)
	//if err != nil {
	//	panic(err)
	//}
}
