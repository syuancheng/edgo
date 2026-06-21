package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Source struct {
	Name string
	Age  int
}

type Destination struct {
	Name string
	Age  int
}

func copyCommonFields(destination, source any) error {
	destinationValue := reflect.ValueOf(destination)
	sourceValue := reflect.ValueOf(source)
	if destinationValue.Kind() != reflect.Ptr || destinationValue.IsNil() {
		return errors.New("destination must be a non-nil pointer to a struct")
	}
	if sourceValue.Kind() != reflect.Ptr || sourceValue.IsNil() {
		return errors.New("source must be a non-nil pointer to a struct")
	}

	destinationValue = destinationValue.Elem()
	sourceValue = sourceValue.Elem()
	if destinationValue.Kind() != reflect.Struct || sourceValue.Kind() != reflect.Struct {
		return errors.New("destination and source must point to structs")
	}

	for i := 0; i < destinationValue.NumField(); i++ {
		destinationField := destinationValue.Field(i)
		fieldInfo := destinationValue.Type().Field(i)
		sourceField := sourceValue.FieldByName(fieldInfo.Name)
		if destinationField.CanSet() && sourceField.IsValid() && sourceField.Type().AssignableTo(fieldInfo.Type) {
			destinationField.Set(sourceField)
		}
	}
	return nil
}

func main() {
	source := Source{Name: "Gopher", Age: 15}
	var destination Destination
	if err := copyCommonFields(&destination, &source); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", destination)
}
