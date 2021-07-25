package calc

type calculator struct {
	NumberOne float64
	NumberTwo float64
	MathSign  string
}

func CreateCalc() calculator {
	return calculator{}
}

func (c *calculator) Calculate(num1, num2 float64, mathSign string) float64 {
	var result float64

	switch mathSign {
	case "+":
		result = c.add(num1, num2)
	case "-":
		result = c.sub(num1, num2)
	case "*":
		result = c.mult(num1, num2)
	case "/":
		// изначально проверку деления на 0 написал здесь, но приходится возвращать 0 и
		// он вылезвет в принте вместе с "Деление на 0 запрещено", мне показалось это логически не верным
		// переместил в main
		result = c.div(num1, num2)
	}

	return result
}

func (c *calculator) add(num1, num2 float64) float64 {
	return num1 + num2
}

func (c *calculator) sub(num1, num2 float64) float64 {
	return num1 - num2
}

func (c *calculator) mult(num1, num2 float64) float64 {
	return num1 * num2
}

func (c *calculator) div(num1, num2 float64) float64 {
	return num1 / num2
}
