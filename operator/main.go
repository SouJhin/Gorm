package main

import "fmt"

var i = 10

func findIndex(nums []int, target int) [2]int {
	for i, temp, length, cur, res := 0, target-nums[i], len(nums), i, [2]int{}; i < length; i++ {
		if nums[i] == temp {
			res[0] = i
			res[1] = cur
			return res
		}
	}
	return [2]int{}
}

func main() {
	a := 1 // 001
	b := 5 // 101
	fmt.Printf("a =====> 🚀🚀🚀 %v\n", a)
	fmt.Printf("a & b =====> 🚀🚀🚀 %v\n", a&b)
	fmt.Printf("a & b =====> 🚀🚀🚀 %v\n", a|b)
	for ; i < 15; i++ {
		fmt.Printf("i =====> 🚀🚀🚀 %v\n", i)
	}
	fmt.Printf("i =====> 🚀🚀🚀 %v\n", i)
	var aa [3]int
	fmt.Printf("aa =====> 🚀🚀🚀 %v\n", aa)
	var cityArr = [4]string{"b", "s", "g", "ss"}
	fmt.Printf("cityArr =====> 🚀🚀🚀 %v\n", cityArr)
	var newArr = [...]int{0: 1, 1: 20, 5: 111}
	fmt.Printf("newArr =====> 🚀🚀🚀 %v\n", newArr)
	for i, s := 0, len(newArr); i < s; i++ {
		fmt.Printf("newArr[i] =====> 🚀🚀🚀 %v\n", newArr[i])
	}
	var mulArr = [2][2]string{{"11", "sss"}, {"nana", "cccc"}}
	fmt.Printf("mulArr =====> 🚀🚀🚀 %v\n", mulArr)
	nums := []int{1, 7, 3, 2}
	fmt.Printf("findIndex(nums, 8) =====> 🚀🚀🚀 %v\n", findIndex(nums, 8))
}
