package main

import "fmt"

func change1(result []int) {
	fmt.Printf("外部1 %p\n%v\n len:%d cap:%d\n", result, result, len(result), cap(result))
	result[0] = 9
}
func change2(result []int) {
	result = append(result, 92)
	fmt.Printf("外部2 %p\n%v\n len:%d cap:%d\n", result, result, len(result), cap(result))
}
func change3(result *[]int) {
	fmt.Printf("外部3-1 %p\n%v\n len:%d cap:%d\n", result, *result, len(*result), cap(*result))
	*result = append(*result, 92)
	fmt.Printf("外部3-2 %p\n%v\n len:%d cap:%d\n", *result, *result, len(*result), cap(*result))
}

func main() {
	result := make([]int, 1)
	fmt.Printf("内部 %p\n%v\n len:%d cap:%d\n", result, result, len(result), cap(result))
	change1(result)
	fmt.Printf("内部 %p\n%v\n len:%d cap:%d\n", result, result, len(result), cap(result))
	change2(result)
	fmt.Printf("内部 %p\n%v\n len:%d cap:%d\n", result, result, len(result), cap(result))
	change3(&result)
	fmt.Printf("内部3 %p\n%v\n len:%d cap:%d\n", result, result, len(result), cap(result))
}
