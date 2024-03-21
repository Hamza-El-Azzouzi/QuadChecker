package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			for i := 1; i <= x; i++ {
				if j == 1 {
					if i == 1 {
						z01.PrintRune('/')
					} else if i == x {
						z01.PrintRune('\\')
					} else {
						z01.PrintRune('*')
					}
				} else if j == y {
					if i == 1 {
						z01.PrintRune('\\')
					} else if i == x {
						z01.PrintRune('/')
					} else {
						z01.PrintRune('*')
					}
				} else {
					if i == 1 || i == x {
						z01.PrintRune('*')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func main() {
	QuadA(5, 3)
	QuadA(5, 1)
	QuadA(1, 1)
	QuadA(1, 5)
}
