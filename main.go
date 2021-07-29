package main

import "fmt"

func main()  {
	var num int

	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	if (num & (num-1) == 0) {
		fmt.Print(num, " Is pow of two")
	}else {
		fmt.Print(num, " Isn't pow of two")
	}
}