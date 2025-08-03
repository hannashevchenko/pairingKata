package categorizer

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type Category struct {
	value string
}

func (c *Category) UnmarshalText(text []byte) error {
	v, err := Parse(string(text))
	if err != nil {
		return err
	}
	*c = v
	return nil
}

func (c Category) MarshalText() (text []byte, err error) {
	return []byte(c.value), nil
}

func (c Category) String() string {
	return c.value
}

var Rest = Category{value: "Rest"}
var Fizz = Category{value: "Fizz"}
var Buzz = Category{value: "Buzz"}
var FizzBuzz = Category{value: "FizzBuzz"}

func Parse(s string) (Category, error) {
	var m = map[string]Category{
		"Rest":     Rest,
		"Fizz":     Fizz,
		"Buzz":     Buzz,
		"FizzBuzz": FizzBuzz,
	}
	c, ok := m[s]
	if !ok {
		return Category{}, fmt.Errorf("category %s not found", s)
	}
	return c, nil
}

func Categorizer(number int) Category {
	if number%3 == 0 {
		if number%5 == 0 {
			return FizzBuzz
		}
		return Fizz
	}
	if number%5 == 0 {
		return Buzz
	}
	return Rest
}
