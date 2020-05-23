# GoLearning

It is the mandatory have a package as a first statement and for main method it will be a "main" package only.
 
 

Command to run go :

    aditya@aditya-nupur:~/GOProjects/Aditya/src/mygo/goCamp/src/learning1$ go run main.go 
    hello
    aditya@aditya-nupur:~/GOProjects/Aditya/src/mygo/goCamp/src/learning1$ go build main.go 
    aditya@aditya-nupur:~/GOProjects/Aditya/src/mygo/goCamp/src/learning1$ ls
    main  main.go
    aditya@aditya-nupur:~/GOProjects/Aditya/src/mygo/goCamp/src/learning1$ ./main 
    hello
    aditya@aditya-nupur:~/GOProjects/Aditya/src/mygo/goCamp/src/learning1$ 

    Every variable has a type and that type cannot be changes

    int - 1, 3
    float - -.5, 0.1 , .1
    bool - true , false
    
    default values : 
    numberic (int or float)  - 0 , string  - "" , bool - false , pointer - nil

    unused variable error occurs only at block level
            var speed int 
            func main(){
               var speed  int  //unused variables 
            }
    _ identifier - skipping blank return values (like a black-hole)

   multiple declarations:
        var firstName, lastName string = "", ""
        var a,b int 
        var complex complex64 = 0
        var (
            speed int 
            heat float64
            off boolean
            n = 3
        )

    Type inference : figure out the type automatically
        var safe = true 
    
    short declaration cant be declared outside the block or at package level (At the package scope you can only use use declaration starts with a keywoard): 
        safe := true

        multiple short declaration at a time
        safe, speed := true, 50
    
    Redeclaration : 
    Short declaration can initialize new variables and assign to existing variables at the same time
        var safe bool
        safe, measure := true, 10
        
        OR 
        safe := true
        safe, speed := false, 300 
        same as 
        safe = false
        speed := 300
    var x = 10
	var y int
	y = x
	x = 221
	fmt.Println("previous vlu", y) // 10
	fmt.Println("current vlu", x)  // 221

    how to get time :  var t time.Time = Time.Now()
    swap variable:
    x , y = y , x


Given the following statement, which of the following is a correct assignment statement?

    var (
    a int
    c bool
    )
    
    a = _   
    i, y = _, true
    _ , _ = a , b
    cannot use _ as value


-----------------------------------------------
Type conversion: (converting a value to another type)
    speed := 100
    force := 1.9
	//speed = speed * int(force)
	speed = int(float64(speed) * force)
    var orange int32
    var blue int 
    orange = blue //int32 can not convert to int
    (so you can not perform any operation with two different type)
    speed * force or speed / force is not possible

------Get input from terminal --------
    var args = os.Args
	fmt.Println(args[0])
    how to get length of argument : 
    fmt.Println(len(arg))

--------------Printf-------------------
 %T - print type for ex : int
 %[2]v - print any type (second argument) (index starts at 1)
 %s - string
 %d - int
 %f - float (%.1f - precision (0,1,2,3....))
 %t - bool
 %b : print bit

var x float64 = 3 / 2  = it will return 1 as both are int

statements can not be used as like an expression, there should be one statement in the single line of go
    var y = 6 + n--

-------------------------------------------------------------------------

    ` symbol used when you need not to escap and display it as it is like \n will be displayed as it is, also called raw string.
strconv :
   var mystring string = "11.1s"
	myconvertedNumber, err := strconv.ParseFloat(mystring, 64)
    strconv.Itoa() int to string 
    strconv.FormatBool( true ) boolean to string
    + - concanatenate operator is used only between string , not between habrid operand

find length:
    currentLen := len("français")
	fmt.Println("Current Length ", currentLen)
	currentLen = utf8.RuneCountInString("français")
	fmt.Println("current length ", currentLen)

strings package:

strings.ToUpper(s)
strings.Repeat("!", l)
strings.Split(content,"\n")
strings.Fields(content) - split by spaces

var bit byte

predeclared type
    bool, string, int, float, int8...64 or rune is same, uint, uint8(byte), uint16, uint32, uint64
    float32, float64, complex64, complex128

you can't check with two different type
    if <int variable> < <int64> //will through an error

    to get max value use math packages math.MaxInt8 math.MinInt8


Creating defined type (named type) : To create define type you need existing type.
    Ex = type Duration int64, but you can not use it in the expression with int64 variable type, (ex if statement comparasion can not be possible between Duration type variable and int64 type variable)
    solution := you need to convert it to one of 
    var ns int64 
    var ms Duration 
    ms = Duration(ns)

Reason of creating defined type := you can attach methods to type, you can attach funtionalities to type.
    Example of method := func (d Duration) Hours() float64

There is no type hierarchy in go 
    type Duration int64
    type myDuration Duration //inderline type will be int64 only
    type (
        Gram int
        Kilogram Gram
        Ton Kilogram 
    )
    var(
        salt Gram = 10 
        apple Kilogram = 1
        truck Ton = 1 
    )


constants can't change once declared.
    const one int =  1
    var two = 2 
    const constTwo = two // not allowed , becuase variable belong to run time and const belong to compile time

    const (
        one int = 1
        tw0
    )
    it will initialise two with 1 ,bacuse it repeats expression of previous expression
    typeless const
    const min = 10
	const yy = 3.14 * min  //only possible for constants and not for varibles
    An untyped numberic constant can be used with all numberic values together
    const min = 13
    var myfloat float64 = min // this is possible as min is typeless constant

    i := 1 //int
    i := 3.34 //float64
    i = 'A' //rune

    i := 1 * 32.5 //possible with numeric values
    i := true + 23.4 // not possible with bool and numberic const

iota: to generate increasing values
    const (
	monday    = iota + 2
    _ 
	tuesday   = iota + 2
	wednesday = iota
)
	fmt.Println("wed", wednesday)  // wed 3

Logical not operator ;
    var on bool
    on = !!on // on = false

nil : (set when yet to initialise)
    default value for pointers, slices, maps, interfaces, channels

short it :
    if n ,err :=strconv.Atoi("123"); err == nil{
     //note n and err not available outside this block or in the chain of else if or in else block   
    }

switch:  Even though more than one condition matches it will execute only one based on sequence.
    city := "Pune"
    switch city {         //switch type
        case "Pune" :    // case type shoule match with switch type
        fmt.Println("city") // varible created in this case block will not be visible to 
                            //outside not even in other block
        default:             //you can move it in any sequence , only one default allowed 
        fmt.Println("where?") // executed if no condition matches
        case "mumbai" , "delhi" : // you can add as many condition as you can in one 
                                    // with quoma seperated         
    }

    fallthrough - used when you want to exeute condition without even if no match (it wont check condition)
    switch {  // if no type , defalut is bool
        case i > 0 :
        fmt.Println("positive number")
        fallthrough    //it should be last statement
        case i > 100:
        fmt.Println("big positive number")
        fallthrough
        default: 
        fmt.Println("number")
    }

    short switch :
    switch i:=0 ; true {

    }

There is only one loop in go i.e. for loop
        for i := 0 ; i < 10 ; i++ {

        }
    queries:
        for{
            //infinite loop
            break  // to break out of loop
            //or you can use labled break
            break queries
        }
    query:
        for <condition> {
            continue // to use to continue for next iteration ; will skip lines below continue
            // or you can use labled continue
            continue query
        }

        strings.Fields("string") : create string slice from string

        range used on string , slice , array , map or a channel
        for i,v := range os.Args { /i - index , v - value

        }
        for _,v := range os.Args[1:]{ // starts from  second element

        }
        for i := range os.Args{

        }

generate random number : 
        rand.Intn(10) // will generate number between 0 ... 9
        rand.seed()   //Seed uses the provided seed value to initialize the default Source to a deterministic state. // each time it will produce different set of random numbers , use combination wity rand.Intn() or other methods
        	rand.Seed(time.Now().UnixNano())
	        for i := 0; i < 10; i++ {
		fmt.Printf("%3d",
			rand.Intn(10))
	} 


goto :
    func someFunc(){
        goto label //not possible
          var i  int // you can not jump to label using goto before declaration of variables
        label:
        goto label
    }

Arrays are composite littral :
    [4] string {"aditya","b"} //rest of element will get default value 
    [4] string {
        "aditya",
        "b",               // this quama is needed as it multiline else not required
    }
 books := [100] years{} // you can use constant variables 
  books :=  [...]string{"aditya","nupur"} // now it will be automatically 2 sized array.
  mutiDimen := [3][4]int{
		{1, 2, 3, 4},
	}
  
  if element type is not same then you can not compare array
  if length type are not same then they are not compareable 

  KEYED array:(you can provide index while assigning value; in below example 2 is the index)

   rate := [...]floate64{
       2:30,
        67,
        5:10      // element size will be 7 : 0,0,30,0,0,10,67  
   }

   const (
     	ETH  = 9 - iota //9
       WAN             //8
       INCX            //7
   )

   type (
		bookCase     [3]int
		cabinateCase [3]int
	)
	b := bookCase{2, 3, 4}
	c := cabinateCase{2, 3, 4}
	n := [3]int{2, 3, 4}
	//c==b not possible
	//b==n or c==n possile
	fmt.Println(b, c, n)
	fmt.Println(b == n)
	fmt.Println(c == n)
    	fmt.Println(c == cabinateCase(b))

    difference between array and slices:
        Array can not shrink or grow
        Array element can not new element or delete element
        array length can not be changes
        slices zero value is nil
        arryas zero value is elemets zero value
    slice uses: slice header (pointer , length , capacity)
        capacity is - length of baking array when slice created, cap() 
        backing array is the array used to create slice and len() and cap() of the slice will be set with the length of backing array
        new backing array created if no enough capacity.
    Append one slice to another
        You can concatenate two slices using the three dots notation:
        a := []int{1, 2}
        b := []int{11, 22}
        a = append(a, b...) // 
    
    copy : both dest and source type should be same 
        copies element based on the length of the smallest slice
        even := []int{2,3,4} odd:= []int{5,4}
        N :=copy(even,odd)// result := 5,4,4 , N -how many element copied
    full slice expression : 
        newSlice := slicable[start:stop:capacity]   
    
    If two slices or array and slice pointing to same backing array , then modification will reflect both


Error :
    error := errors.New(" issues occured")
    var err error
    err = fmt.Errorf("my error")

Map : (disordered , nofix length , len method returns how many keys that map has)
 var myDictionary map[int]string
 myDictionary[1] ="aditya"
delete(mydictionary, 1) // how to delete it 

    another way to use shortcut to creat map 
    m:= make(map[int]string)
    m["aditya"] = 1
    m["nupur"] = 2

make :
    does memory allocation for built-in modules like , map , slice , and channel , while new is for types memory allocation
    make(type, length,capacity)
    append is expensive operation, to set memory at the time of creation use make. 
go supports variable length arguments:
    func myfunc(arg ...int) {}

defer:
    you can use many defer statements in one function ; they will execute in reverse order when the program executes

function as a value type : 
    type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
    advantage you can pass function as value

Panic and Recover: (no try catch block in go)
    panic is builtin function to break the normal flow
    
init() , main() : will be called automatically

files ----
    files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println("err", err)
		return
	}
	for _, files := range files {
		fmt.Println(files.Name())
	}

A string value is actually a read-only byte slice. The ellipsis passes the byte elemets of the name string as individual elements to the append function
    var names []byte
    name := "aditya"
    names = append(names, name...)

   	stringSlice := []string{"asd", "asd", stringValue}
	mysorted := sort.StringSlice(stringSlice)

time : command measures the execution time of a program
    time go run main.go
    o/p:
        real    0m1.322s
        user    0m0.473s
        sys     0m0.166s

The backing array is shared amoung the same string values. That is why they are read only.
unless necessory do not convert to string , comsume memory.
Be efficient: Do not use string concat (+ operator).
Instead, create a new byte slice as a buffer 

var buff = make([]byte, 0 , size)

for i:=0 ; i<size ;i++{
    buff = append(buff, text[i])
}

MAP : 
 var myMap map[string]int 
 var myMap map[string]int {
     "aditya" : 99,
 }
    value,bool := myMap["aditya"]

    If element doesn't exist in a map, the map returns a zero value depending on the element type of the map.
    You cannot add element inside the map without initialising it.
    if you copy from one map to other and modify oter map , it will reflect to other map as well.
    myCopiedMap := myMap
    Map header only has pointer.
    String header has pointer and length
    slice and array header has pointer , length and capacity
    if you add same key & value into map , existing value will be overriden

    delete(myMap,key) //no need to check if key exist or not.
    myMap = nil  //to delete all element but element still exist in the memory
     for key:= range myMap{
         delete(myMap,key)  //this for loop is just a single operation behind the scenes go changes this whole loop to single clear command mapclear().
     } 



scanning user input :
    input stream line by line into a buffer:
    Command line Input , Network Input or file 
    default stop character is new line

go run . < main.go //file will be treated as standard input to the program 
in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		fmt.Println(in.Text())   // print file line by line
}
	if err := in.Err(); err != nil {
		fmt.Println(err)          // no need to add it in the for loop
	}

To create set you need to use map in go (unique values)
    //	in.Split(bufio.ScanWords)  -- use can use it to scan for words instead of new line



// you can search content of the websites as well : where website content will be act as a file to a go input (curl -s - only show contents)

 curl -s https://goaditya.blogspot.com/2020/05/linux-commands-for-beginners.html | go run .


//regular expression
var regx *regexp.Regexp = regexp.MustCompile(`[^ a-z | A-Z | 0-9]+`)  //costly operation bust be run once , must declare outside method or outside loop
    str := regx.ReplaceAllString(in.Text(), "")

Struct : allow to create be blueprint for set of related type to a single type. (fix at compile type , you can not add fields at run type)

with slice , array or map , you can create of same type , in struct you can add different type as well because they may have different types of fields.

type Person struct{
    name string,
    age int
}

aditya := Person{
name : "aditya",
age : 10,
}
or
aditya := Person{ /// when you dont declare field name , order is important
"adtiay",
10,
}

var aditya Person
aditya.name = "vasanti"
aditya.age=10


when you create new struct from exist it creates clone , modification of one struct will not reflect  others
type person struct{
name string
}

func main() {
var person1 person
person1.name = "aditya"
person2 := person1
person2.name = "saish"
fmt.Println("",person1) //output aditya
}


go does not have inheritance or is-a relation
go use composition has-a relationship  (go called it embedding )

anonymous has a relation :

type Person struct {
name string
mobile int
}
type Struent struct {
Person // anynomous
cardId string
}

-------------------

type Person struct {
name string
mobile int
}
type Student struct {
Person // anynomous
cardId string
name string
}

func main() {
var person1 Person
person1.name = "aditya"
person2 := person1
person2.name = "saish"
fmt.Println("Hello, playground",person1)

 stud := Student{
Person : Person{
"aditya",
56788888,
},

}
fmt.Println(stud.mobile) // 56788888
fmt.Println(stud.name) // blank


---------json marshaling ------------
type person struct {
name string
FirstName string     `json:"Name,omitempty"`   this is field tag
password string      `json:"_"`   
}
func main() {
var student person
student.name = "aditya"
out, _ := json.Marshal(student)
fmt.Println("Hello, playground",string(out))   // name is not exported as n is lower case will not be marshaled, only exported field will be marshaled
}

you can redirect your output into the file:
    go run main.go > users.json


Go does not support function overloading , constructor

Declaring package level variable is a bad practic.
Go is a pass by value language: variable pass the function is a copy (clone) and not actual value 

func myFunc(m int) (n int){
    n = m
    return  // it is same as return n l; also called as naked retrun
}

pointer :
    * - finds value , & - find address

    A map value is a pointer. So, do not use pointers to map values. (if map is passed as a parameter function can modify it and will affect orignal map as well)

Methods :
    struct Person{
        name string
    }

    func (p Person) myPerson(){
        fmt.Println(p)
    }
    func (p *Person)changeName(n  string){
        p.name = n
    }
    var p Person
    p.name="sad"
    (&p).changeName("happy")
    another benifit , you can create fun with same name for  dfferent diffrent struct type
    method belongs to type , function belongs to a package
    methods can be attached to any types : int , string , float64, array , struct, slice , map , chan , func
    A type and its methods should be in same package
    you can call methods with nil values variable as well

interface ://
    interface is a protocol - a contract
    We can describe comman behaviour with the help of interface.
    interface is an abstract type.
    interface only describe expected behaviour.
    any type wants to implement interface has to implement all the methods , no partitial implementation allowed
    The bigger the interface the weaker the abstraction
    you can assigne any type of value to printer var as long as that type implements interface all methods.
    to convert interface variable to implementation type
      g , isGame := interfacevar.(*game) 
      or if any type satisfise given interface you can check :
      g , isGame := interfaceVar.(interface{ discount(float64)})
    package main

import "fmt"

type printer interface{
	print()  // if any type implements this print() will be implementing this printer interface
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


empty interfaces:
    type empty interface {

    } 
    or inferface {}

    var any interface{}
    any = "aditya"
    any = 10
    any = any.(int) * 2  // you need to convert inface type to required type

    type switch :
    Detects and extracts the dynamic value from an inteface value
    
    switch e := v.(type){

    }

    famous interfaces:
    fmt.Stringer, io.Writer, io.Reader, sort.Interface, error, json.Marshaler
    fmt.Stringer - need to implement String() which will written string
    type book struct{
	name string
}
func (b book) String() string{
	return fmt.Sprintf("book.name%s",b.name)
}
func main()  {
	b := book { "Ramayan"}
	fmt.Println("book",b)  // b it will indirectly call String()
} 
strings.Builder can efficiently combine multiple string values.

//sort need to implemnt below interface   	sort.Sort(b) // or reverse : sort.Sort(sort.Reverse(b))
func (a list) Len() int{ 
	return len(a)
}
func (a list) Swap(i,j int)  {
	a[i], a[j] = a[j], a[i]
}

func (a list) Less(i,j int)bool{
	return a[i].price < a[j].price
}

if err != nil {
		log.Fatal(err)   // it is same as using fmt.print and return
	}

how to joing string:
s1 := []string{"ad", "das"}
	totalString := strings.Join(s1, "-")

both slice , map , channels , pointers, function (all this is reference type) will be modifed if passed to any function as a argument

int, bool, float, struct are not reference type (they are value type), if you want to pass its reference ,you need to use it using & symbol

channels :
    when we have main go routine and child go routine (by using go before calling method/function) child routine does not get priority to run 
    if you don't use channel :
    main go routine starts creating child go rountine and checks any other pending code to run inside main routine then program exist. 

    Channels are used to communicate between go routines(only way to communicate between go routines)
    import (
	"fmt"
	"net/http"
	"time"
)

var (
	site = []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
		"http://faceboosk2.com",
	}
)

func main() {
	c := make(chan bool)
	for _, link := range site {
		go checkLink(link, c)
		bools := <-c   // wating for the message coming from channel
		if !bools {
			time.Sleep(30 * time.Second)
			go checkLink(link, c)
			<-c  
		}
	}

}

func checkLink(l string, c chan bool) {
	_, err := http.Get(l)
	fmt.Println("checking for ling", l)
	for _, i := range []int{1, 2, 3, 4, 5, 6} {
		fmt.Print(i)
	}
	if err != nil {
		fmt.Println(l, "might be down")
		c <- false
		return
	}
	fmt.Println(l, "is up")
	c <- true //sending message to channel

}


function literal:(anonymous function):

    

//packages list
    math, strings, fmt, strconv, utf8, os, path, runtime, rand, errors, "io/ioutil", sort
    BUFIO.SCANNER, regx, encoding/json ,log,net/http

Problems:

1. masking characters


2. 
Wrap the given text for 40 characters per line. For example, for the following input, the program should print the following output.

INPUT:

Hello world, how is it going? It is ok The weather is beautiful.
OUTPUT:

Hello world, how is it going? It is ok.
The weather is beautiful.


https://github.com/inancgumus/learngo.git

https://github.com/StephenGrider/GoCasts.git