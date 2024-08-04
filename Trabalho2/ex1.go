package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func avaliaExp(operandos []int, operadores []string) int {
	result := operandos[0]
	for i := 1; i < len(operandos); i++ {
		if operadores[i-1] == "+" {
			result += operandos[i]
		} else if operadores[i-1] == "-" {
			result -= operandos[i]
		}
	}
	return result
}

func main() {
	input, err := os.Open("calcula.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input.Close()

	dados, err := ioutil.ReadAll(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	entradas := strings.Split(string(dados), "/n")
	saidas := ""

	for i := 0; i < len(entradas); i += 2 {
		m, _ := strconv.Atoi(entradas[i])
		if m == 0 {
			break
		}
		expr := entradas[i+1]
		operandos := make([]int, m)
		operadores := make([]string, m-1)
		opIndex := 0
		for j := 0; j < len(expr); j++ {
			if expr[j] == '+' || expr[j] == '-' {
				operadores[opIndex] = string(expr[j])
				opIndex++
			} else {
				operand, _ := strconv.Atoi(string(expr[j]))
				operandos[opIndex] = operandos[opIndex]*10 + operand
			}
		}
		result := avaliaExp(operandos, operadores)
		saidas += fmt.Sprintf("Teste %d/n%d/n/n", (i/2)+1, result)
	}

	err = ioutil.WriteFile("calcula.out", []byte(saidas), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
