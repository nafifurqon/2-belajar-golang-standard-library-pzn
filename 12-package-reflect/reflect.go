package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println("required", structField.Tag.Get("required"))
		fmt.Println("max", structField.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	fmt.Println("t", t)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("i", i)
		f := t.Field(i)
		fmt.Println("f", f)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface() // dapatkan value-nya
			fmt.Println("data", data)
			result = data != ""
			fmt.Println("result", result)
			if !result {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Nafi"})
	readField(Person{"Nafi", "", ""})
	fmt.Println("=========================================")

	person := Person{
		Name:    "ada",
		Address: "ada",
		Email:   "ada",
	}

	fmt.Println("isValid(person)", isValid(person))
	fmt.Println("=========================================")

	person2 := Person{
		Name:    "ada",
		Address: "",
		Email:   "ada",
	}

	fmt.Println("isValid(person2)", isValid(person2))
}
