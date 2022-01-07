package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var (
	name, town  = "a", "4"
	age, salary = 1, 2
)

func main() {
	fmt.Println("hello world")
	fmt.Println("Default string values are name", name, "town", town)
	fmt.Println("Default int values are age", age, "salary", salary)

	fmt.Println("data type if name is", reflect.TypeOf(name))

	townValue, err := strconv.Atoi(town)
	if err == nil {
		fmt.Println("total value is", townValue+age)
	}

	fmt.Println("salary address is", &name)

	var ptr *string = &name

	fmt.Println("ptr value is", ptr, " and value stred in that is ", *ptr)

	*ptr = "changed value"

	fmt.Println("name value is", name)

	var ptr1 *int = &age

	fmt.Println("ptr value is", ptr1, " and value stred in that is ", *ptr1)

	fmt.Println("new value is", changeValueByPass(name))

	fmt.Println("name value is", name)

	fmt.Println("new value after using pass by reference is", changeValueByReference(&age))

	fmt.Println("name value after pointer is", age)

	fmt.Println("os env value", os.Environ())

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}

func changeValueByPass(a string) string {
	a = os.Getenv("USER")
	return a
}

func changeValueByReference(name *int) int {
	*name = 34
	return *name
}
