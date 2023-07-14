package calcu

import "fmt"

type Calc struct {
	Num1 int
	Num2 int
	Operation string
}

func (c *Calc) Calculate() int {
	switch c.Operation {
	case "+":
		return c.Num1 + c.Num2
	case "-":
		return c.Num1 - c.Num2
	case "*":
		return c.Num1 * c.Num2
	case "/":
		return c.Num1 / c.Num2
	default:
		return 0
	}
}

func (c *Calc) Show(num int) string {
	return fmt.Sprintf("The result is %d", num)
} 

