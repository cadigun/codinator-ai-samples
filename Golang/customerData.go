package main

import "fmt"

type CustomerData struct { 
	Id           int    
	CustomerName string 
}

func (data *CustomerData) setdata(id int, name string) { 
	data.Id = id
	data.CustomerName = name
}

func main() {
	customer := CustomerData{} 
	customer.setdata(1, "John Doe") 
	fmt.Println(customer.CustomerName)
}