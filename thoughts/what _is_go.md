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

Package Level Scope→   for variables ; not for imports

In below examples, capitalized variables having package scope can be imported to another package


## memory addresses
For every variable value we stored has memory address. To know the stored memory address of variable we use ‘&’ (ampersand) operator.

## pointer
Like C Language,Go also supports pointer. Pointer is used to hold the memory address of a value. It is represented by using “*” (asterisk) just like the C language.

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



