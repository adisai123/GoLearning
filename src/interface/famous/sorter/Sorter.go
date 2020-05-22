package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

type book struct {
	Title string
	Price int
}
func (a list) Len() int{ 
	return len(a)
}
func (a list) Swap(i,j int)  {
	a[i], a[j] = a[j], a[i]
}

func (a list) Less(i,j int)bool{
	return a[i].Price < a[j].Price
}

func (a list) String()string{
	// var buff  strings.Builder
	// buff.WriteString("{")
	// for _, s := range a {
	// 	buff.WriteRune('\n')
	// 	buff.WriteString(s.title )
	// 	buff.WriteString(":")
	// 	buff.WriteString(strconv.Itoa(s.price))
	
	// }
	// buff.WriteRune('\n')
	// buff.WriteString("}")
	data , err := json.MarshalIndent(a,""," ")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

type list []*book
func main(){
	b := list{
		{
			Title:"my3",
			Price: 500,

		},
		{
			Title: "my2",
			Price: 200,
		},

	}
	sort.Sort(b)
	fmt.Println(b)
	sort.Sort(sort.Reverse(b))
	fmt.Println(b)

}