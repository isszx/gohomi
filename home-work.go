package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

func average(arr []int) float64 { // Task 1
	if len(arr) == 0 {
		return 0
	}
	sum := 0
	for _, a := range arr {
		sum += a
	}
	return float64(sum) / float64(len(arr))
}

func max(words []string) string { // Task 2.1
	word := ""
	for _, a := range words {
		if utf8.RuneCountInString(a) > utf8.RuneCountInString(word) {
			word = a
		}
	}
	return word
}

func reverse(arr []int64) []int64 { // Task 2.2
	newSlice := make([]int64, len(arr))
	for i, a := range arr {
		newSlice[i] = a
	}
	for i, j := 0, len(newSlice) - 1; i < len(newSlice) / 2; i, j = i + 1, j - 1 {
		newSlice[i], newSlice[j] = newSlice[j], newSlice[i]
	}
	return newSlice
}

func printSorted(m map[int]string) { // Task 3
	var tmp []int
	for k := range m {
		tmp = append(tmp, k)
	}
	sort.Ints(tmp)
	for _, v := range tmp {
		fmt.Printf("%s ", m[v])
	}
	fmt.Println()
}

func main() {
	fmt.Println("\t\t\tTask #1\t\t\t")
	fmt.Println(average([]int{1,2,3,4,5,6}))
	fmt.Println(average([]int{}))

	aw := []int{1,1,1,1,1,3,3,3,3,3}
	fmt.Println(average(aw))

	fmt.Println("\n\t\t\tTask #2.1\t\t\t")
	fmt.Println("Max length word:\t", max([]string{"one", "two", "three"}))
	fmt.Println("Max length word:\t", max([]string{"one", "two"}))
	fmt.Println("Max length word:\t", max([]string{"su", "sudo", "super user"}))

	fmt.Println("\n\t\t\tTask #2.2\t\t\t")
	fmt.Println("Reverse\t", reverse([]int64{1,2,3,4,20}))

	t2p2 := []int64{1,21}
	fmt.Println("\nBefore\t", t2p2)
	fmt.Println("Reverse\t", reverse(t2p2))
	fmt.Println("After\t", t2p2)

	fmt.Println("\n\t\t\tTask #3\t\t\t")
	printSorted(map[int]string{10: "aa", 0: "bb", 500: "cc"})
	printSorted(map[int]string{2: "a", 0: "b", 1: "c"})
	printSorted(map[int]string{99: "a", 0: "bb", -99: "ccc"})
}
