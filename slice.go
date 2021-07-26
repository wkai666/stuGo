package main

import "fmt"

func main()  {
	data := [...]int{0, 1, 2, 3, 4, 5, 10:9}

	s := data[0:4]
	fmt.Println(&data[0], &s[0])
	s[0] += 100
	s[1] += 200

	s2 := data[:2]
	fmt.Println(&data[0], &s[0], &s2[0])
	s2 = append(s2, 100, 200)
	fmt.Println(&data[0], &s[0], &s2[0])

	fmt.Println(s)
	fmt.Printf("the lens of s is: %d, caps of s is %d\n", len(s), cap(s))
	fmt.Println(data)
	fmt.Printf("the pointer s and data is: %p, %p\n", &s[0], &data[0])
}
