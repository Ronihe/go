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



