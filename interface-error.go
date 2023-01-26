package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai/pembagi
		return result, nil
	}
}

func main() {
    hasil, err := pembagi(10,0) 
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
