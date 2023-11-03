package main

import "fmt"

func pattern1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%s", "*")
		}
		fmt.Println()
	}
}

func pattern2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%s", "*")
		}
		fmt.Println()
	}
}

func pattern3(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d", j+1)
		}
		fmt.Println()
	}
}

func pattern4(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d", i+1)
		}
		fmt.Println()
	}
}

func pattern5(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= n-i-1; j++ {
			fmt.Printf("%s", "*")
		}
		fmt.Println()
	}
}

func pattern6(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= n-i-1; j++ {
			fmt.Printf("%d", j+1)
		}
		fmt.Println()
	}
}

func pattern7(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < 2*n; j++ {
			if j >= n-1-i && j <= n-1+i {
				fmt.Printf("%s", "*")
			} else {
				fmt.Printf("%s", " ")
			}
		}
		fmt.Println()
	}
}

func pattern8(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			fmt.Printf("%s", " ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("%s", "*")
		}
		for j := 0; j < n-1-i; j++ {
			fmt.Printf("%s", " ")
		}
		fmt.Println()
	}

}

func pattern9(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%s", " ")
		}
		for j := 0; j < 2*n-(2*i+1); j++ {
			fmt.Printf("%s", "*")
		}
		for j := 0; j < i; j++ {
			fmt.Printf("%s", " ")
		}
		fmt.Println()
	}

}

func pattern10(n int) {
	for i := 0; i < 2*n-1; i++ {
		stars := i
		if i > n-1 {
			stars = 2*n - i - 2
		}
		for j := 0; j < stars+1; j++ {
			fmt.Printf("%s", "*")
		}
		fmt.Println()

	}
}

func pattern11(n int) {

	/*
	   solution @VamsikrishnaSiddu
	*/
	// for i := 0; i < n; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		if (i%2 == 0 && j%2 == 0) || (i%2 != 0 && j%2 != 0) {
	// 			fmt.Printf("%d", 1)
	// 		} else {
	// 			fmt.Printf("%d", 0)
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	start := 1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			start = 1
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%d", start)
			start = 1 - start
		}
		fmt.Println()
	}

}

func pattern12(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d", j+1)
		}

		for j := 0; j < 2*n-2-2*i; j++ {
			fmt.Printf("%s", "*")
		}

		for j := i; j >= 0; j-- {
			fmt.Printf("%d", j+1)
		}

		fmt.Println()
	}
}

func pattern13(n int) {
	count := 1

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", count)
			count++
		}
		fmt.Println()
	}
}

func pattern14(n int) {

	for i := 0; i < n; i++ {
		startChar := 'A'
		for j := 0; j <= i; j++ {
			fmt.Printf("%c ", startChar)
			startChar++
		}
		fmt.Println()
	}
}

func pattern15(n int) {

	for i := 0; i < n; i++ {
		startChar := 'A'
		for j := n - 1; j >= i; j-- {
			fmt.Printf("%c ", startChar)
			startChar++
		}
		fmt.Println()
	}
}

func pattern16(n int) {
	startChar := 'A'
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c ", startChar)
		}
		startChar++
		fmt.Println()
	}

}

func pattern17(n int) {
	for i := 0; i < n; i++ {
		startChar := 'A'
		breakpoint := (2*i + 1) / 2
		for j := 0; j < n-1-i; j++ {
			fmt.Printf("%s", " ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("%c", startChar)
			if j < breakpoint {
				startChar++
			} else {
				startChar--
			}

		}
		fmt.Println()
	}
}

func pattern18(n int) {
	for i := 0; i < n; i++ {
		for j := 'A' + n - 1 - i; j <= 'A'+n-1; j++ {
			fmt.Printf("%c", j)
		}
		fmt.Println()
	}
}

func pattern19(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf("*")
		}

		for j := 0; j < 2*i; j++ {
			fmt.Printf(" ")
		}

		for j := 0; j < n-i; j++ {
			fmt.Printf("*")
		}

		fmt.Println()
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}

		for j := 0; j < 2*n-2*i-2; j++ {
			fmt.Printf(" ")
		}

		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}

		fmt.Println()
	}
}

func pattern20(n int) {
	for i := 0; i < 2*n; i++ {
		spaces := 2*n - 2*i - 2
		stars := i
		if i >= n {
			spaces = 2*i - 2*n + 2
			stars = 2*n - i - 2
		}
		for j := 0; j <= stars; j++ {
			fmt.Printf("%s", "*")
		}

		for j := 0; j < spaces; j++ {
			fmt.Printf("%s", " ")
		}

		for j := 0; j <= stars; j++ {
			fmt.Printf("%s", "*")
		}

		fmt.Println()
	}
}

func pattern21(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Println()
	}
}

func pattern22(n int) {
	for i := 0; i < 2*n-1; i++ {
		for j := 0; j < 2*n-1; j++ {
			top := i
			left := j
			bottom := 2*n - 2 - i
			right := 2*n - 2 - j
			result := n - min(min(top, bottom), min(left, right))
			fmt.Printf("%d", result)
		}
		fmt.Println()
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	//pattern1(n)
	//pattern2(n)
	//pattern3(n)
	//pattern4(n)
	//pattern5(n)
	//pattern6(n)
	//pattern7(n)
	//pattern8(n)
	//pattern9(n)
	//pattern10(n)
	//pattern11(n)
	//pattern12(n)
	//pattern13(n)
	//pattern14(n)
	//pattern15(n)
	//pattern16(n)
	//pattern17(n)
	//pattern18(n)
	//pattern19(n)
	//pattern20(n)
	//pattern21(n)
	pattern22(n)
}
