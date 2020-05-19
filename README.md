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
//packages list
    math, strings, fmt, strconv, utf8, os, path, runtime , rand , errors , "io/ioutil",sort
