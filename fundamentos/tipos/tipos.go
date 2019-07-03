package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numero inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	fmt.Println("********************************************")

	// sem sinal (só positvos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	fmt.Println("********************************************")

	// com sinal... uint8 uint16 uint32 uint64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo de int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	fmt.Println("********************************************")

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O  rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	fmt.Println("********************************************")

	// numeros reais (float32, float64)
	var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo literal 49.99 é", reflect.TypeOf(49.99)) // float64

	fmt.Println("********************************************")

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	fmt.Println("********************************************")

	//string
	s1 := "Olá meu nome é Adaumir"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string s1 é", len(s1))

	fmt.Println("********************************************")

	// stringcom multiplas linhas
	s2 := `Ola
	meu
	nome
	é
	Adaumir`
	fmt.Println("O tamanho da string s2  é", len(s2))

	fmt.Println("********************************************")

	// char???
	// var x char = 'b'
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
