package main

import (
	"fmt"
	"os"
)

func main() {
	var a, v int

	input, err := os.Open("aero.in")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input.Close()

	output, err := os.Create("aero.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer output.Close()

	for i := 1; ; i++ {
		fmt.Fscanf(input, "%d %d", &a, &v)
		if a == 0 && v == 0 {
			break
		}

		trafego := make([]int, a+1)

		for j := 0; j < v; j++ {
			var x, y int
			fmt.Fscanf(input, "%d %d", &x, &y)
			trafego[x]++
			trafego[y]++
		}

		var maior int
		for _, t := range trafego {
			if t > maior {
				maior = t
			}
		}

		fmt.Fprintf(output, "Teste %d/n", i)
		for j := 1; j <= a; j++ {
			if trafego[j] == maior {
				fmt.Fprintf(output, "%d ", j)
			}
		}
		fmt.Fprintf(output, "/n/n")
	}
}
