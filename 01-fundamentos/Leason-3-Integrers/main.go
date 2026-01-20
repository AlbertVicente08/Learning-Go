package main

import "fmt"

func main() {
	//variable of type uint8
	var Valor_positivo_pequenyo uint8
	Valor_positivo_pequenyo = 10
	fmt.Println(Valor_positivo_pequenyo)

	//variable of type int8
	var Int_positivo_negativo_pequenyo int8
	Int_positivo_negativo_pequenyo = -10
	fmt.Println(Int_positivo_negativo_pequenyo)
	// uint16, uint32, uint64
	// int16, int32, int64
	var myint int = 2400898
	fmt.Println(myint)

	myint = int(Int_positivo_negativo_pequenyo)
	myint = int(Valor_positivo_pequenyo)
	// to typecase a value int()

	//myint = int("STRING") ERROR 

	var mybyte byte = 'A'
	fmt.Println(mybyte)
	// byte se usa principalmente para representar datos sin procesar
	// y también para distinguir entre uint8 y byte
	// ya que byte es un alias para uint8.

	// Rune is a an alias for int32
	var myrune rune = 'A'
	fmt.Println(myrune)
	myrune = '😊'
	fmt.Println(myrune)
}

