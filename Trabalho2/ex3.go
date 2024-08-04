package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Tesouro struct {
	valor []int
}

func (t *Tesouro) addValor(valor int) {
	t.valor = append(t.valor, valor)
}

func (t Tesouro) checkDiv(x, y int) bool {
	valorTotal := x + y + t.valorTotal()
	if valorTotal%2 != 0 {
		return false
	}
	tmp := valorTotal / 2
	return t.findSubSum(tmp)
}

func (t Tesouro) findSubSum(val int) bool {
	n := len(t.valor)
	dp := make([][]bool, n+1)

	for i := range dp {
		dp[i] = make([]bool, val+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= val; j++ {
			if j < t.valor[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-t.valor[i-1]]
			}
		}
	}
	return dp[n][val]
}

func (t Tesouro) valorTotal() int {
	total := 0
	for _, v := range t.valor {
		total += v
	}
	return total
}

func main() {
	input, _ := os.Open("tesouro.in")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	output, _ := os.Create("tesouro.out")
	defer output.Close()
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	var x, y, n int
	test := 1
	for {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d %d", &x, &y, &n)
		if x == 0 && y == 0 && n == 0 {
			break
		}
		tesouro := Tesouro{}
		for i := 0; i < n; i++ {
			scanner.Scan()
			var valor int
			fmt.Sscanf(scanner.Text(), "%d", &valor)
			tesouro.addValor(valor)
		}
		sort.Ints(tesouro.valor)
		fmt.Fprintf(writer, "Teste %d/n", test)
		if tesouro.checkDiv(x, y) {
			fmt.Fprintln(writer, "S")
		} else {
			fmt.Fprintln(writer, "N")
		}
		fmt.Fprintln(writer)
		test++
	}
}
