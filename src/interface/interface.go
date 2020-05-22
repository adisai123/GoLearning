package main

import "fmt"

type printer interface{
	print()
}

type teacher struct{
	id int
}
func(t teacher) print(){
	fmt.Println("teacher",t.id)
}
type person struct{
	name string
}
func (p person) print(){
	fmt.Println("person",p.name)
}
type list []printer
func (ls list) print(){
	for _, l  := range ls {
		l.print()
	}
}
func main(){
	p := []person {
		{"aditya"},
		{"nupur"},
		{"saish"},
	}
	t := teacher{1}
	var l list
	for _, p := range p {
		l = append(l,p)		
	}
	l = append(l,t)

	l.print()
}
