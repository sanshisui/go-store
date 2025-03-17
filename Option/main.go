package main

import "fmt"

func main() {
	//Functional Options（函数选项模式）
	serverFunctional := NewServerFunctional(WithHost("example.com"), WithTLS())
	fmt.Printf("Server Config: %+v\n", serverFunctional)

	
	//Struct Options（结构体选项模式）
	config := ServerConfig{
		Host: "example.com",
		TLS:  true,
	}
	serverStruct := NewServerStruct(config)
	fmt.Printf("Server Config: %+v\n", serverStruct.config)
}