- open source from google
- good for scalability and effectiveness
- compiled and static-typed
- using packages for management of dependecies

## Packages

- organize code
- Capitalization:
  Uppercase: exported, visible outside the package
  Lower case: unexported, not visible outside the package

## data types

var types determins the how much memory space occupies in storage and bit pattern stored is interpreted.

- boolean true / false
- numeric types : integer / floating
- string
- derived types: pointer, array, structure, union types, function types, slice types, interface, map

## quotes in golang

double quotes: define a string
back quotes: define a raw literal string
single quotew: byte or rune.


## variables

shorthand `:=` only inside of fuc
var: set things to zero values

Here no value is initialized and later assigned to variables. Go automatically assigned default values to it

- For Booleans→ false for integer→ 0 for float→ 0.0 for strings→ ””
- For pointers, maps, slices, functions, interfaces, channels→ nil

## levels of scope:

universe -> package -> file -> block{}

Package Level Scope→ for variables ; not for imports

In below examples, capitalized variables having package scope can be imported to another package

## memory addresses

For every variable value we stored has memory address. To know the stored memory address of variable we use ‘&’ (ampersand) operator.

## pointer

Like C Language,Go also supports pointer. Pointer is used to hold the memory address of a value. It is represented by using “\*” (asterisk) just like the C language.

We can pass memory address i.e. reference instead of bunch of values and we still can change the values. So we don’t need to pass values, instead we can pass addresses.

```
func main() { \
s := 5
call(s)
fmt.Println(s)
}
func call(n int) {
// make a copy of the
 n++
}
//Here o/p is still 5

```

## control flow

control flow stmt : to break the flow of execution by branching, looping, decision making by enabling the program to execute the code based on the conditions.

- if / else if / else
- for loop:

  ```
  for init; condition; increment/decrement{

  }
  ```

-

## switch case

- compare the values of the same type
- set optional default stmt if the condition fails
- use expression if need to compare condition based on the vaue to use

- also switch on types also:

```
func Type(x interface{}) {
 switch x.(type) { //this one is assert
 case int:
 fmt.Println("int")
 case string:
 fmt.Println("string")
 case wish:
 fmt.Println("wish")
 default:
 fmt.Println("Unknown Type")
 }
}
```

## Functions:

The execution of Go program starts from the main function itself func main(){ } . A function takes one or more parameters and return one or more outputs. Like other programming languages, go explicitly needs to return values.

```
func <function name>(parameters)(return types){
 …………….//Executable code
 return <value> //If function is returning
}
```

## variadic functions

pass in any number of parameters of the same type

```
func sum(args ..int)int{ //Variadic function
    total:=0
    for value:=range args{
        total+=value
    }
    return sum
 }
 func main(){
    fmt.Println(sum(54,76,43,23,78)) //Passing multiple values
    data:=[]int{33,87,43,21} //Passing Slice of ints
    fmt.Println(sum(data...))
```

## Function as expression

We can pass function to an expression and call that function by using expression/variable name. This is the only way to have one function inside another.

```
func main(){
message:=func(){
fmt.Println(“Good Morning User”)
}
message() //calling function by expression name/var name as anonymous fun
fmt.Printf(“%T”,message) //Its type will be a func()
```

## Closure

```
func main(){
 sum:=func (x,y int)(int){
 return x+y
 }
 fmt.Println(“Sum is”,sum(1, 2))
}
```

## callbacks

Passing a function as an argument is called callback in golang. Here, we can pass function as an argument.
**call**

```
func main(){
 show([]int{43, 76, 34}, func(n int) {
 fmt.Println(n)
 })
}

 func show(numbers [ ]int, call func(int)) {
 for _, n := range numbers {
 call(n)
 }
}
```

## recursion

```
func factorial(n int) int {
 if n==1{
 fmt.Println(n)
 }
 return n*factorial(n-1) //factorial() will call itself till it reaches to 0
```

## Defer

Go has a special statement defer which schedules a function call to execute before function completes.

```
f,_:=os.Open(“filename.txt”)
defer f.Close()
```

Advantages:

- It keeps our Close() near to the Open() and it’s easy to understand
- deferred func will always run even if runtime panic err occurs

## Pass By value

Strictly speaking, there is only one way to pass parameters in Go is Pass By Value. Whenever a variable is passed as a parameter, a copy of that variable is created and passed to the function and this copy will be at a different memory location.

In case, variable is passed as a pointer argument then again new copy of parameter is created and that will point to the same address.

## Anonymous Function

# Data Structure:

## Array:

var var_name [size]var_type

```
var arr [10]int
```

- once it is created, not allowed to chagne
- not more insert than the array size
- by default array size is 0
- len()

## slice

Slice is same like an array but it has variable length so we don’t need to specify the length to it. It will grow whenever it exceeds its size. Like an array, slice also has index and length but its length can be changed.

- Slice also has continuous segments of memory locations
- The default value of uninitialized slice is nil
- Slices does not store the data. It just provides reference to an array
- As we change the elements of slice, it will modify corresponding elements of that array

**The only difference between slice and array is that there is no need to specify length while creating slice.**

```
var x[ ] float64
```

Here, slice of type float64 elements will be created with length 0 & capacity 0

We can also create slice by using make() also which is available in builtin package of golang

```
 x:=make([ ]string,5) → Here the capacity of slice equal to length of it
```

The length and capacity of slice can be obtained by using len(slice), cap(slice)

```
array:=[ ]float64 {1,2,3,4,5} //slice of float64 elements
 x:=array[0:5]
 fmt.Println(x)
```

You can do slicing on string because string is slice of bytes.

- append
- copy

```
func main() {
 elements := []int{3, 5, 2, 6, 2}
 arrays := make([ ]int, 0, 3) //Creating slice with make( )
 for n := range elements {
 fmt.Println("Elements", n)
 }
 for i := 0; i < 80; i++ {
 arrays = append(arrays, i) // to append i elements to slice
fmt.Println("Len:", len(arrays), "Capacity:", cap(arrays),
 "Value: ", arrays[i])
 }}
```

## Map

An unique key value pair like dictionary is used to lookup values based on the key.

```
map [keyType] valueType

```

This includes unordered pairs. Maps are like literals but unique keys are required.

```
var dictionary map[string]int //map using var
 dictionary[“Zero”]=1
 fmt.Println(dictionary[“Zero”]) //Accessing value using key
```
`dictionary:=make(map[string]int) //map using make()`

The functions provided are len(map) and delete(mapName,key)

### range:
Range : Range is used to iterate over the slice and map along with for loop. When range is used with slice it returns two values- 1st one is index and 2nd one is copy of the element at that index.

## struct:
- Encapsulation → state [“fields”] behaviour [“methods”] export / unexported
- Reusability → Inheritance [“Embedded Types”]
- Polymorphism → Interfaces
- Overriding → Promotion

- user defined type
- we declare the type
- the type has fields
- the type can also have “tags”
- the type has an underlying type
- in this case, underlying type is struct
- we declare variables of the type
- we initialize those variables
- initialize with specific value or with default values.

But in Golang, they are ‘associated’ with struct.

```
type Animal struct{ }
func (e Animal ) eat () { } //Methods defined outside but works on same struct

```

```
mob:=new(Mobile)
```

```
mob:=Mobile{brand:“Samsung”,model:”Galaxy”,price:24500}
```

## Methods
Go supports methods to define struct.

func (receiverName receiverType) functionName(arguments) returnType

Methods having pointer receivers can modify the value to which the receiver points.

## Type Aliasing –

## Embedded Types –

## Promotion
Embedding is the composition not inheritance, but Go supports something called “promotion” where the fields and methods of the embedded type becomes available to the embedding type.

Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying method calls.

## Struct Pointer
We can use struct pointer to access struct fields –
   
```
mob:=&Mobile {“Lenovo”,”A6”,”6700”}
 fmt.Println(mob.brand,mob.model,mob.price)
```
## JSON
JSON(JavaScript Object Notation) is simple text based data interchanging format. Basically we use structs for JSON for putting and receiving the data. It is lightweight and gives faster performance to access the data

## Encoding –
- Marshal → to encode GO values to JSON in string format
- Unmarshal → to decode JSON data to GO values
- Encoding → same as Marshal but in streams
- Decoding → same as unmarshal but from streams

Go has built-in support for JSON encoding by providing functions in “json/encoding” package.

## marshalling
This function will convert the given data structure into JSON in [ ]byte format which we need to convert into string.


## Unmarshalling
```
func Unmarshal(data [ ]byte,v interface{ }
var per Person
err:=json.Unmarshal(b,&per)
)

```
If the JSON fits and matches with struct fields, after the call err will be nil and data from b will have been stored in the struct m, but non-matched fields will not be decoded


## Tags
The tags are given to the fields to attach meta information and acquired by using reflection. It is used to provide info on how struct field is encoded to or decoded from another format (or stored /retrieved from database), but you can use it to store whatever meta-info you want either for another package or for your use.

```
type User struct{
 Name string ` json:”name” xml:”name” ` 
 Age int `json:”omitempty”`
}
```

## encoding

Encoding works with writer and writing streams of JSON data. The GO has provided below function for encoding JSON through writer interface –

```
func (enc *Encoder) Encode (v interface{ }) error
```
```
json.NewEncoder(os.Stdout).Encode(p1)
```

```
If we see the source code, the NewEncoder(w *io.Writer) *Encoder takes writer and returns pointer to an Encoder for encoding and os.Stdout are open files pointing to standard input. Stdout is a pointer to file and pointer to a file implements  func(f *file ) Write(b [ ]byte) (n int,err error) and that means it is implementing this method from Writer interface. (Polymorphism)
```


## decoding
Decoding works with reader and reading streams of JSON data. The GO has provided below function for encoding JSON through writer interface –

```
func (dec *Decoder) Decode (v interface{ }) error
```

NewReader( ) will return a reader and give it to NewDecoder( ) so that it will decode the JSON.

```
json.NewDecoder(reader).Decode(&p1)
```

## Strings:




## Interface:
An interface is an abstract type. It doesn’t expose the representation or internal structure of its values, or set of basic operation they support; it reveals only some of their methods. When you have value of an interface type, you don’t know only what it is; you only know what it can do.


Interfaces are named collection of method signatures only (like java). To implement an interfaces in Go, we need to implement all the methods of that interface.

```
type interface_name interface{
 /* method signature (one or more method sets) */
}
```

This is what we call Polymorphism. Polymorphism is ability to write code that take on different behaviour through its implementation of types.

## Empty Interface
Empty interface has zero methods. It is like Object class (which is superclass of all classes in java and accept any object). Similar to this, empty interface accepts any interface to itself.

## Conversion vs Assertion
Conversion process deals with two concepts like –
- Widening → Converting from Lower to Higher data type
- Narrowing → Converting from Higher to Lower data type 

```
var x=15 //int type
var y=15.45
fmt.Println(float(x)) → 15.00
fmt.Println(int(y)) → 15
```

assertion: `interface_variable.(type)`

# Concurrency:
For this, Go has great support and enhancement for concurrency using goroutines and channels.

## Goroutines:
Goroutine is like thread concept in java which is capable of running with the multiple independent functions. To make function as a goroutine, you just need to add go keyword before the function.

## wait group
To use sync.WaitGroup :
Create instance of sync.WaitGroup      → var wg sync.WaitGroup
Call Add(n) where n is no of goroutines to wait for  → wg.Add(1)
Execute defer Wg.Done( ) in each goroutine to indicate that goroutine is finished
executing to the WaitGroup.
Call wg.Wait( ) where we want to block

## Concurrency vs Parallelism
Concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations. Concurrency is dealing with lots of things at once. Parallelism is about doing things at once.






