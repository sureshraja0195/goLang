package main

import "fmt"

func printHello() string {
	return "Hello World!"
}

type cards []string

func (d cards) print() {

	for i, v := range d {
		fmt.Println(i, v)
	}

}
