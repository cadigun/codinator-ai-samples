package main

import "testing"

func Testsetdata(t *testing.T) { // Function name should be MixedCaps
	customer := CustomerData{} // Type name is okay
	customer.setdata(1, "John Doe") // Method name should be MixedCaps
	if customer.CustomerName != "John Doe" {
		t.Errorf("Expected 'John Doe', got '%s'", customer.CustomerName)
	}
}