package src

type Calculator struct {
	Operations chan operation
	Result     float64
	Logs       chan callog
}

type operation struct {
	operation string
	num1      float64
	num2      float64
}

type callog struct {
	logprefix string
	logmsg    string
}

func NewCalculator() *Calculator {
	return &Calculator{
		Operations: make(chan operation),
		Logs:       make(chan callog),
		Result:     0,
	}
}

func (cal *Calculator) ExecuteOperation(operator string, num1 float64, num2 float64) {
	switch operator {
	case "add":
		cal.Result = num1 + num2
		cal.Operations <- operation{operation: "+", num1: num1, num2: num2}

	case "divide":
		cal.Result = num1 / num2
		cal.Operations <- operation{operation: "/", num1: num1, num2: num2}

	case "substract":
		cal.Result = num1 - num2
		cal.Operations <- operation{operation: "-", num1: num1, num2: num2}

	case "multiply":
		cal.Result = num1 * num2
		cal.Operations <- operation{operation: "*", num1: num1, num2: num2}
	}
}
