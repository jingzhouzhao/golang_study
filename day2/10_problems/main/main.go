package main

import (
	"fmt"
	"math"
)

func main() {

	primeCount := 0
	fmt.Println("101-200中的素数：")
	for i := 101; i <= 200; i++ {
		if isPrime(i) {
			fmt.Printf("%d\t", i)
			primeCount++
		}
	}
	fmt.Println("素数个数:", primeCount)

	//100-999 水仙花数
	fmt.Println("100-999中水仙花数")
	for i := 100; i < 999; i++ {
		a := i / 100 % 10
		b := i / 10 % 10
		c := i / 1 % 10
		if math.Pow(float64(a), float64(3))+math.Pow(float64(b), float64(3))+math.Pow(float64(c), float64(3)) == float64(i) {
			fmt.Printf("%d \t", i)
		}
	}
	//求n的阶乘和
	n := 5
	sum := 0
	for i := 1; i <= n; i++ {
		sum+=factorial(i)
	}
	fmt.Printf("\n %d的阶乘和是%d", n,sum)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func isPrime(n int) bool {
	//偶数一定不是素数
	if n > 2 && n%2 == 0 {
		return false
	}

	//从2遍历到n的方根，看看是否有因子
	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			//发现一个因子
			return false
		}
	}
	return true
}
