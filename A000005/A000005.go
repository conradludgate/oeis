// The number of divisors of n
// http://oeis.org/A000005
package main

import "fmt"

const N int = 100

// Simple prime factorisation function
// 60 = 2^2 * 3 * 5 -> map[int]int{2: 2, 3: 1, 5: 1}
func pfactor(n int) map[int]int {
	p := make(map[int]int)

	for i := 2; i < n; i++ {
		if n%i == 0 {
			p[i] += 1
			n /= i
			i = 1
		}
	}

	if n > 1 {
		p[n] += 1
	}

	return p
}

func d(n int) int {
	p := pfactor(n)

	d := 1
	for _, v := range p {
		d *= v + 1
	}

	return d
}

func main() {
	for i := 1; i <= N; i++ {
		fmt.Printf("%d, ", d(i))
	}
}
