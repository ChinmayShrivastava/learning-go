package main

import (
	"errors"
	"fmt"
)

func PhoneWorking(pwr bool) (string, error) {
	if pwr {
		return "Phone is working", nil
	}
	return "", errors.New("Phone is not working")
}

func main() {
	phone, err := PhoneWorking(true)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(phone)
	}
}
