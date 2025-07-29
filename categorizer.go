package categorizer

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type Category struct {
	value string
}

var Rest = Category{value: "Rest"}
var Fizz = Category{value: "Fizz"}

func Categorizer(number int) Category {
	return Rest
}

func main() {
	Categorizer(10)
}
