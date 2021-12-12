package ctrlflow

func condition() {
	ifelse()
	switchCond()
	switchValue()
	swtichType()
}

func ifelse() {
	// if else if else
	a := 1
	if a < 0 {
		p("a less than zero")
	} else if a > 0 {
		p("a greater than zero")
	} else {
		p("a is zero")
	}

	// if else assign
	if b := 1; b < 0 {
		p("b less than zero")
	} else {
		p("else -> b= ", b)
	}
}

func switchCond() {
	a := 1
	switch {
	case a < 0:
		p("a less than zero")
	case a > 0:
		p("a greater than zero")
	default:
		p("a is zero")
	}
}

func switchValue() {
	switch a := 1; a {
	case 0:
		p("a is zero")
	case 1, 2:
		p("a is 1 or 2")
	default:
		p("else, a=", a)
	}
}

func swtichType() {
	var a interface{} = 1
	switch a.(type) {
	case int:
		p("int a=", a)
	case string:
		p("string a=", a)
	default:
		p("else a=", a)
	}
}
