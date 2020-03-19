// slice project main.go
package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100 //数组和切片都是下标从0开始的
}

func printSlice(s []int) {
	fmt.Printf("%v,len=%d,cap=%d\n", s, len(s), cap(s))
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])

	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println("After updateSlice")
	updateSlice(s1)
	updateSlice(s2)
	fmt.Println("arr=", arr) //因为slice是数组的一个view，
	//slice改变了，数组展示时也跟着变了
	fmt.Println("reslice:")

	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
	fmt.Println("Extending slice")
	//slice不能向前扩展，可以向后扩展，但是不能超过底层数字的容量
	arr[0], arr[2] = 0, 2
	s3 := arr[2:6]
	s4 := s3[3:5] //[s1[3],s1[4]]
	fmt.Println("arr:", arr)
	fmt.Println("s3=", s3)
	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d\n",
		s3, len(s3), cap(s3))
	fmt.Println("s4=", s4)
	fmt.Printf("s4=%v,len(s4)=%d,cap(s4)=%d\n",
		s4, len(s4), cap(s4))
	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	fmt.Println("s5,s6,s7=", s5, s6, s7)
	//s6,s7不是对arr的view，arr的长度不够s6，s7往里面加值
	fmt.Println("arr=", arr)
	fmt.Println("Creating slice")
	var s []int //zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s8 := make([]int, 16)
	s9 := make([]int, 10, 32)
	printSlice(s8)
	printSlice(s9)

	fmt.Println("Copying slice")
	copy(s8, s3)
	printSlice(s8)

	fmt.Println("Deleting elements from slice")
	s8 = append(s8[:3], s8[4:]...)
	fmt.Println("s8=", s8)
	fmt.Println("Popping from front ")
	front := s8[0]
	s8 = s8[1:]
	fmt.Println("Popping from back ")
	tail := s8[len(s8)-1]
	s8 = s8[:len(s8)-1]
	fmt.Println(front, tail)
	printSlice(s8)
}
