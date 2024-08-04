package main

import (
	"fmt"
	"os"
)

// par ou impar
func main() {
	in, err := os.Open("par.in")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create("par.out")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	var N, teste, A, B, result int
	var J1, J2 string
	teste = 1
	for {
		fmt.Fscanf(in, "%d", &N)
		if N == 0 {
			break
		}
		fmt.Fprintf(out, "Teste %d/n", teste)
		fmt.Fscanf(in, "%s/n%s", &J1, &J2)
		for i := 0; i < N; i++ {
			fmt.Fscanf(in, "%d %d", &A, &B)
			result = A + B

			if result%2 == 0 {
				fmt.Fprintf(out, "%s/n", J1)

			} else {
				fmt.Fprintf(out, "%s/n", J2)
			}

		}
		teste++

	}
}
