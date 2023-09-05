package main

import (
	"fmt"
	"math"
	"strconv"
)

func narcissistic(num int) bool {
	numStr := strconv.Itoa(num)
	numDigit := len(numStr)

	sum := 0

	for _, digitStr := range numStr {
		digit, _ := strconv.Atoi(string(digitStr))
		sum += int(math.Pow(float64(digit), float64(numDigit)))
	}

	return sum == num
}

func findOutlier(arr []int) int {
	evenCount, oddCount := 0, 0
	var even, odd int

	for _, num := range arr {
		if num%2 == 0 {
			evenCount++
			even = num
		} else {
			oddCount++
			odd = num
		}

		if evenCount > 1 && oddCount == 1 {
			return odd
		}

		if oddCount > 1 && evenCount == 1 {
			return even
		}
	}

	return 0
}

func findNeedle(haystack []string, needle string) int {
	for i, value := range haystack {
		if value == needle {
			return i
		}
	}
	return -1
}

func blueOcean(arr1 []int, arr2 []int) []int {
	result := []int{}

	elementsInArr := make(map[int]bool)

	for _, num := range arr1 {
		if !elementsInArr[num] {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	//NARCISSISTIC NUMBER
	resultNarcissistic1 := narcissistic(153)
	resultNarcissistic2 := narcissistic(111)

	println("narcissistic(153) =", resultNarcissistic1)
	println("narcissistic(111) =", resultNarcissistic2)

	//PARITY OUTLINER
	arr1 := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	arr2 := []int{160, 3, 1719, 19, 11, 13, -21}
	arr3 := []int{11, 13, 15, 19, 9, 13, -21}

	resOutlier1 := findOutlier(arr1)
	resOutlier2 := findOutlier(arr2)
	resOutlier3 := findOutlier(arr3)

	fmt.Println("Result 1:", resOutlier1) // Should return 11
	fmt.Println("Result 2:", resOutlier2) // Should return 160
	fmt.Println("Result 3:", resOutlier3) // Should return 0 (false)

	//NEEDLE IN THE HAYSTACK
	haystack := []string{"red", "blue", "yellow", "black", "grey"}
	needle := "red"

	index := findNeedle(haystack, needle)

	fmt.Printf("Index of '%s' in haystack: %d\n", needle, index)

	// BLUE OCEAN REVERSE
	arr4 := []int{1, 2, 3, 4, 6, 10}
	arr5 := []int{1}
	resultBlueOcean1 := blueOcean(arr4, arr5)
	fmt.Println("Result 1:", resultBlueOcean1)

	arr6 := []int{1, 5, 5, 5, 5, 3}
	arr7 := []int{5}
	resultBlueOcean2 := blueOcean(arr6, arr7)
	fmt.Println("Result 2:", resultBlueOcean2)
}
