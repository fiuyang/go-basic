package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)

	// penurunan nilai konversi
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//konversion type data string
	var name = "Ama"
	var a byte = name[0]
    var aString = string(a)

	fmt.Println(name)
	fmt.Println(aString)
}