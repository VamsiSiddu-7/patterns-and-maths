package main

import (
	"fmt"
	"math"
)

func countDigits(n int) int {
	var count int
	var temp int = n
	for n != 0 {
		rem := n % 10
		if rem != 0 {
			if temp%rem == 0 {
				count++
			}
		}
		n = n / 10
	}
	return count
}

func isArmStrong(n int) bool {
	var count int
	var sum int
	var digits []int = make([]int, 0)
	var temp int = n
	for n != 0 {
		rem := n % 10
		digits = append(digits, rem)
		n = n / 10
		count++
	}

	for i := 0; i < len(digits); i++ {
		sum = sum + int(math.Pow(float64(digits[i]), float64(count)))
	}

	if temp == sum {
		return true
	}
	return false
}

func sumOfDivisors(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				sum = sum + j
			}
		}

	}

	return sum
}

func isPrime(n int) bool {
	var count int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	if count == 2 {
		return true
	}
	return false
}

func reverse(n int) int {
	var minVal = math.MinInt32
	var maxVal = math.MaxInt32

	sum := 0
	for n != 0 {
		rem := n % 10
		sum = sum*10 + rem
		n = n / 10
	}
	if sum < minVal || sum > maxVal {
		return 0
	}
	return sum
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	sum := 0
	temp := x
	for x != 0 {
		rem := x % 10
		sum = sum*10 + rem
		x = x / 10
	}

	if temp == sum {
		return true
	}

	return false

}

func gcd(a, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			temp := b
			b = a % b
			a = temp
		} else {
			temp := a
			a = b % a
			b = temp
		}
	}

	if a == 0 {
		return b
	}
	return a
}

func main() {
	var n int
	fmt.Scan(&n)
	//fmt.Println(countDigits(n))
	//fmt.Println(reverse(n))
	//fmt.Println(gcd(a, b))
	//fmt.Println(isArmStrong(n))
	//fmt.Println(sumOfDivisors(n))
	fmt.Println(isPrime(n))
}
