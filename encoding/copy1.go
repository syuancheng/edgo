package main

import (
	"fmt"
	"reflect"
)

/**
Explanation of the code:
The function takes two parameters destp and srcp which are pointers to the destination and source structs respectively.
destv and srcv are reflect.Value objects created from the pointers.
destt is the type of the destination struct.
The function iterates over the fields of the destination struct using a for loop.
For each field, it checks if the field exists in the source struct (srcv.FieldByName (sf.Name)). If the field doesn't exist or the types are not assignable, it continues to the next field.
If the field exists and the types are assignable, it sets the value of the destination field to the value of the source field using destv.Field (i).Set (v).

Note: This function uses reflection (reflect package) to work with the fields of the structs at runtime.
*/

type Bob struct {
	Name string
	Age  int
}

type Mark struct {
	Name string
	Age  int
}

// copyCommonFields copies the common fields from the struct
// pointed to srcp to the struct pointed to by destp.
func copyCommonFields(destp, srcp interface{}) {
	destv := reflect.ValueOf(destp).Elem()
	srcv := reflect.ValueOf(srcp).Elem()

	destt := destv.Type()
	for i := 0; i < destt.NumField(); i++ {
		sf := destt.Field(i)
		v := srcv.FieldByName(sf.Name)
		if !v.IsValid() || !v.Type().AssignableTo(sf.Type) {
			continue
		}
		destv.Field(i).Set(v)
	}
}

func main() {

	bob := Bob{
		Name: "bob",
		Age:  12,
	}

	mark := Mark{}

	fmt.Printf("Bob=%v\n", bob)

	copyCommonFields(mark, bob)

	fmt.Printf("Mark=%v\n", mark)

}
