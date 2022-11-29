package main

import (
	"fmt"
	"reflect"
)

// No go posso inferir as variaveis

func main() {
	nome := "Rafael"
	var versao float32 = 1.1
	var idade = 28
	fmt.Println("Olá, sr.", nome, "sua idade é", idade, "anos.")
	fmt.Println("Este programa está na versão", versao)

	fmt.Println(reflect.TypeOf(nome), reflect.TypeOf(idade), reflect.TypeOf(versao))
}
