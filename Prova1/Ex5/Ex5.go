package main

import "fmt"

var (
	m, n int
	tv   [1000][1000]int
)

func main() {
	var i, j, ic, jc, dh, dv, a, b, teste int
	teste = 1

	for {
		_, err := fmt.Scan(&m, &n)
		if err != nil || m == 0 {
			break
		}

		for i = 0; i < m; i++ {
			for j = 0; j < n; j++ {
				fmt.Scan(&tv[i][j])
			}
		}

		dh, dv = 0, 0
		for {
			_, err := fmt.Scan(&a, &b)
			if err != nil || (a == 0 && b == 0) {
				break
			}
			dh += a
			dv += b
		}
		dv = -dv

		fmt.Printf("Teste %d/n", teste)
		teste++
		for i = 0; i < m; i++ {
			for j = 0; j < n; j++ {
				if j != 0 {
					fmt.Print(" ")
				}
				ic = (m + (i-dv)%m) % m
				jc = (n + (j-dh)%n) % n
				fmt.Print(tv[ic][jc])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
