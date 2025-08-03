package main

import (
	"errors"
	"fmt"
	"log"
	"pairingKata/categorizer"
	"pairingKata/cli"
	"reflect"
	"strconv"
)

func main() {
	input, err := cli.Cli()
	if err != nil {
		log.Fatal(errors.New("failed to read input: " + err.Error()))
	}
	inputInt, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(errors.New(fmt.Sprintf("input should be int, %v is provided", reflect.TypeOf(input).Kind().String())))
	} else {
		fmt.Println(categorizer.Categorizer(inputInt))
	}

}
