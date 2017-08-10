// The zero sequence
// http://oeis.org/A000004
package main

import (
	"fmt"
)

const N int = 100

func main() {
	for i := 0; i < N; i++ {
		fmt.Print("0, ")
	}
}
