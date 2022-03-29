package main

import "testing"

func TestVal(t *testing.T) {
	if ValidBraces("()") != true {
		t.Errorf("incorrect sequence processing %v", "()")
	}
	if ValidBraces("[]") != true {
		t.Errorf("incorrect sequence processing %v", "[]")
	}
	if ValidBraces("{}") != true {
		t.Errorf("incorrect sequence processing %v", "{}")
	}
	if ValidBraces("{}()[]") != true {
		t.Errorf("incorrect sequence processing %v", "{}()[]")
	}
	if ValidBraces("([{}])") != true {
		t.Errorf("incorrect sequence processing %v", "([{}])")
	}
	if ValidBraces("{}({})[]") != true {
		t.Errorf("incorrect sequence processing %v", "{}({})[]")
	}
	if ValidBraces("(({{[[]]}}))") != true {
		t.Errorf("incorrect sequence processing %v", "(({{[[]]}}))")
	}
	if ValidBraces("(({[]{()}}[{(())}[]]))") != true {
		t.Errorf("incorrect sequence processing %v", "(({[]{()}}[{(())}[]]))")
	}
	if ValidBraces("{[]}(){()}") != true {
		t.Errorf("incorrect sequence processing %v", "{[]}(){()}")
	}
	if ValidBraces("[{()()()}]") != true {
		t.Errorf("incorrect sequence processing %v", "[{()()()}]")
	}
	if ValidBraces("(){") != false {
		t.Errorf("incorrect sequence processing %v", "(){")
	}
	if ValidBraces("}{") != false {
		t.Errorf("incorrect sequence processing %v", "}{")
	}
	if ValidBraces("[}") != false {
		t.Errorf("incorrect sequence processing %v", "(){")
	}
	if ValidBraces("(}") != false {
		t.Errorf("incorrect sequence processing %v", "(}")
	}
	if ValidBraces("[)") != false {
		t.Errorf("incorrect sequence processing %v", "[)")
	}
	if ValidBraces("([{)]}") != false {
		t.Errorf("incorrect sequence processing %v", "([{)]}")
	}
	if ValidBraces("(((({{}})[]))") != false {
		t.Errorf("incorrect sequence processing %v", "(((({{}})[]))")
	}
	if ValidBraces("()(){}[(])") != false {
		t.Errorf("incorrect sequence processing %v", "()(){}[(])")
	}
	if ValidBraces("{(})") != false {
		t.Errorf("incorrect sequence processing %v", "{(})")
	}
	if ValidBraces(")(}{[]") != false {
		t.Errorf("incorrect sequence processing %v", ")(}{[]")
	}

}
