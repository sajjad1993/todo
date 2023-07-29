package main

import "fmt"

type SomeInterface interface {
	SomeCommonFunction(message string) error
}

type Parent struct {
}

func (p *Parent) SomeCommonFunction(message string) error {
	fmt.Printf("this is message Prant : %s \n ", message)
	return nil
}

type Child struct {
	Parent
}

func (p *Child) SomeCommonFunction(message string) error {
	fmt.Printf("this is message from Child :  %s \n ", message)
	return nil
}

func (p *Child) SomeDiffrentFunctionChild(message string) error {
	///////////////////
	p.SomeCommonFunction(message)
	///////////////////
	return nil
}

func Call(input SomeInterface) error {
	return input.SomeCommonFunction("message from new object ")
}
func main() {
	new(Child).SomeDiffrentFunctionChild("hello ")
}
