package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer client.Close()

	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Print("Enter second number: ")
	fmt.Scanf("%d", &num2)

	args := Args{A: num1, B: num2}
	var reply int
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		fmt.Println("Error calling RPC:", err)
		return
	}
	fmt.Println("Add Result:", reply)

	var subtractReply int
	err = client.Call("Arith.Subtract", args, &subtractReply)
	if err != nil {
		fmt.Println("Error calling RPC:", err)
		return
	}
	fmt.Println("Subtract Result:", subtractReply)
	fmt.Println("\nRPC calls completed successfully")
}
