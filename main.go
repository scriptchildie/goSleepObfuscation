package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {

	for {
		toBEncrypted()
		err := encryptFunc(reflect.ValueOf(toBEncrypted).Pointer(), 3) //encrypt function toBEncrypted(), sleep for 3 seconds , decrypt function
		if err != nil {
			log.Fatalf("%v", err)

		}
		toBEncrypted()
		fmt.Println("-----------------")
	}

}

// This function is encrypted during sleeping
func toBEncrypted() (uint32, error) {
	fmt.Println("is it working?")
	return 0, nil
}
