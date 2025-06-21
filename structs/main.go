package main

import (
	"fmt"
	"structs/ll"
)

func main() {
	var arr []int
	fmt.Println("Initial val of arr - ", arr)
	arr = []int{1, 2, 3, 4, 5}

	var head *structs.Node
	fmt.Println("Initial val of head -", head) // nil
	head = structs.CreateLinkedList(arr)
	// head := structs.CreateLinkedList(arr)

	// fmt.Println(head.val)		// NOTE:can't use val as it is lowercase so its not public

	structs.PrintLinkedList(head)

	// Embedded struct
	user := structs.CreateUser()
	fmt.Println("\nUser Details")
	fmt.Println("User object - ", user) // {1234 abc 10 0xc0000a6090}
	fmt.Println("User name - ", user.Name)
	// fmt.Println("User id - ", user.id)  // as it is private
	fmt.Println("User Level - ", user.Level)
	fmt.Println("User age - ", user.Age, " (Embedded struct example)") // Age is in the PrivateFields but it can be accessed via top Level
	// usage when sending to frontend or something like that
	user.PrivateFields = nil            // bad practice but for learning
	fmt.Println("user object - ", user) // {1234 abc 10 <nil>}

	// struct methods in go
	structs.UseStructMethod()
}
