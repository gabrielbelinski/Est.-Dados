package main

import (
	"fmt"
	"os"
)

func calcMd(temp []int, intervalo int) (int, int) {
	n := len(temp)
	sum := 0

	for i := 0; i < intervalo; i++ {
		sum += temp[i]
	}

	min := sum / intervalo
	max := sum / intervalo

	for i := intervalo; i < n; i++ {
		sum = sum - temp[i-intervalo] + temp[i]
		md := sum / intervalo

		if md < min {
			min = md
		}

		if md > max {
			max = md
		}
	}

	return min, max
}

func main() {
	input, err := os.Open("lua.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input.Close()

	output, err := os.Create("lua.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer output.Close()

	var n, m int
	testes := 1

	for {
		fmt.Fscanf(input, "%d %d", &n, &m)
		if n == 0 && m == 0 {
			break
		}

		temp := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscanf(input, "%d", &temp[i])
		}

		fmt.Fprintf(output, "Teste %d/n", testes)
		testes++

		min, max := calcMd(temp, m)
		fmt.Fprintf(output, "%d %d/n/n", min, max)
	}
}
