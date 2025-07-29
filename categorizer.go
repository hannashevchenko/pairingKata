package categorizer

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type Category struct {
	value string
}

var Rest = Category{value: "Rest"}
var Fizz = Category{value: "Fizz"}
var Buzz = Category{value: "Buzz"}
var FizzBuzz = Category{value: "FizzBuzz"}

func Categorizer(number int) Category {
	if number%3 == 0 && number%5 == 0 {
		return FizzBuzz
	}
	if number%3 == 0 {
		return Fizz
	}
	if number%5 == 0 {
		return Buzz
	}
	return Rest
}

func main() {
	Categorizer(10)
}
