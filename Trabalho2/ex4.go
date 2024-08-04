package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("dobra.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input.Close()

	output, err := os.Create("dobra.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	testes := 0

	for {
		scanner.Scan()
		str := scanner.Text()
		n, err := strconv.Atoi(str)

		if err != nil || n == -1 {
			break
		}

		testes++

		nelevado := math.Pow(2, float64(n))
		pedacos := math.Pow(nelevado+1, 2)

		fmt.Fprintf(output, "Teste %d/n", testes)
		fmt.Fprintf(output, "%d/n/n", int64(pedacos))
	}
}
