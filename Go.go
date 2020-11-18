package main

import (
	"fmt"
)

func main() {
	var i, j, s, n int
	var a [10000]int
	var b [10000]int
	var c [10000]int
	var d [10000]int
	s = 0
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&a[i])
		fmt.Scan(&b[i])
	}
	for i = 0; i < n; i++ {
		if a[i] >= b[i] {
			d[i] = a[i]
			c[i] = b[i]
		} else if b[i] > a[i] {
			c[i] = a[i]
			d[i] = b[i]
		}
	}
	for i = 0; i < n; i++ {
		if c[i]%2 == 0 {
			for j = c[i] + 1; j < d[i]; j = j + 2 {
				s = s + j
			}
		}
		if c[i]%2 == 1 {
			for j = c[i] + 2; j < d[i]; j = j + 2 {
				s = s + j
			}
		}
		fmt.Printf("%d\n", s)
		s = 0
	}

}
