package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVal(t *testing.T) {
	t.Run("A=positiveTest", func(t *testing.T) {
		if ValidBraces("()") != true {
			assert.Equal(t, ValidBraces("()"), true, "это правильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "()")
		}

		if ValidBraces("[]") != true {
			assert.Equal(t, ValidBraces("[]"), true, "это правильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "[]")
		}
		if ValidBraces("{}") != true {
			assert.Equal(t, ValidBraces("{}"), true, "это правильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "{}")
		}
		if ValidBraces("{}()[]") != true {
			assert.Equal(t, ValidBraces("{}()[]"), true, "это правильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "{}()[]")
		}
		if ValidBraces("([{}])") != true {
			assert.Equal(t, ValidBraces("([{}])"), true, "это правильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "([{}])")
		}
		if ValidBraces("{}({})[]") != true {
			assert.Equal(t, ValidBraces("{}({})[]"), true, "это правильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "{}({})[]")
		}
		if ValidBraces("(({{[[]]}}))") != true {
			assert.Equal(t, ValidBraces("(({{[[]]}}))"), true, "это правильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "(({{[[]]}}))")
		}
		if ValidBraces("(({[]{()}}[{(())}[]]))") != true {
			assert.Equal(t, ValidBraces("(({[]{()}}[{(())}[]]))"), true, "это правильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "(({[]{()}}[{(())}[]]))")
		}
		if ValidBraces("{[]}(){()}") != true {
			assert.Equal(t, ValidBraces("{[]}(){()}"), true, "это правильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "{[]}(){()}")
		}
		if ValidBraces("[{()()()}]") != true {
			assert.Equal(t, ValidBraces("[{()()()}]"), true, "это правильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "[{()()()}]")
		}
	})
	t.Run("A=positiveTest", func(t *testing.T) {
		if ValidBraces("(){") != false {
			assert.Equal(t, ValidBraces("(){"), true, "это неправильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "(){")
		}
		if ValidBraces("}{") != false {
			assert.Equal(t, ValidBraces("}{"), true, "это неправильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "}{")
		}
		if ValidBraces("[}") != false {
			assert.Equal(t, ValidBraces("[}"), true, "это неправильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "(){")
		}
		if ValidBraces("(}") != false {
			assert.Equal(t, ValidBraces("(}"), true, "это неправильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "(}")
		}
		if ValidBraces("[)") != false {
			assert.Equal(t, ValidBraces("[)"), true, "это неправильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "[)")
		}
		if ValidBraces("([{)]}") != false {
			assert.Equal(t, ValidBraces("([{)]}"), true, "это неправильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "([{)]}")
		}
		if ValidBraces("(((({{}})[]))") != false {
			assert.Equal(t, ValidBraces("(((({{}})[]))"), true, "это неправильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "(((({{}})[]))")
		}
		if ValidBraces("()(){}[(])") != false {
			assert.Equal(t, ValidBraces("()(){}[(])"), true, "это неправильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "()(){}[(])")
		}
		if ValidBraces("{(})") != false {
			assert.Equal(t, ValidBraces("{(})"), true, "это неправильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", "{(})")
		}
		if ValidBraces(")(}{[]") != false {
			assert.Equal(t, ValidBraces(")(}{[]"), true, "это неправильная скобочная последовательность ")
			t.Errorf("incorrect sequence processing %v", ")(}{[]")
		}
	})

}
