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

	args := Args{A: 5, B: 10}
	var reply int
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		fmt.Println("Error calling RPC:", err)
		return
	}
	fmt.Println("Add Result:", reply)

	errors := client.Call("Arith.Subtract", args, &reply)
	if errors != nil {
		fmt.Println("Error calling RPC:", errors)
		return
	}
	fmt.Println("Subtract Result:", reply)
	fmt.Println("\nRPC calls completed successfully")
}
