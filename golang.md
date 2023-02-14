# 1 Install & Initilize App
Install go & use vs code install extensions. Then initilize using:<br> 
>`mkdir <app-name>`; cd <app-name> // mkdir  and cd into app <br> 
> `go mod init <app-name>` // initilize & a go.mod file gets created (module name and version of go) <br>  
_____
# 2 Package Name, main Entry Point & function

>`package main` // naming the package as main  <br><br> 
> `import (` <br> 
> `"fmt"` // Importing packages for functions Ex: Print & Println <br> 
>`)`<br> <br> 
> `func main() {` // Set a main function aka entry point, only one main function per app. (Without this we get decleration error) <br> 
> `fmt.Print("Let's learn go and write here the code/logic")` <br> 
> `}` <br> 
_____
##  Test Run
Running go files <br> 
> `go run main.go` <br> 

_____
# 3 Variables & Consrtants

## 3.1 Initilizing

### Constants : 
> `const tickets = 50` <br>

### Variables :
#### Implicit Sytnax : 
> `message := "Go lang" // its a var, we cant use it for const and we cant defne a type` <br> 
#### Explicit Syntax : <br> 
> `var name = "Go lang"` <br> 
> `var remainingtickets = 50` <br>
> `var name string = "Go lang"` \\ with datatype <br>
> Note : Until the variable is use it will be underlined by the linter

#### Using Variables : 
> `fmt.Println("Let's learn", string)` <br> 
> `fmt.Println("You have", remainingtickets, "of", tickets, "tickets")` <br> 
> `fmt.Printf("Ticket balance %v tickets of %v tickets", remainingtickets, tickets)`
> `fmt.Printf("Ticket balance: \n %T tickets of %T tickets", remainingtickets, tickets)`  // comments <br> 

Note :
 -  **%**type of variable 
 - **\n** for new line in string
_____
## 3.2 Initilize and Assign Values Later
> `var userName string` <br>
> `var pin int` <br>
> `var pin uint` <br> \\\ here uint protects from negetive i.e. it's an unsigned int <br> 
> `userName = "Bob"` <br>
> 
[datatype](https://www.geeksforgeeks.org/data-types-in-go/)

## 3.3 Accept values from user
> `fmt.Println("Enter Username & pin")` <br>
> `fmt.Scan(&userName)` <br>
> `fmt.Scan(&pin)` <br>

## 3.4 Input Vallidaton :  <br>
> `isValidUesr := len(userName) >= 4 && pin < 1000` <br>
> `isValidEmail := strings.Contains(email, "@")`  <br>
Note : IsValid saves a bool so we can use it in a if statement and handle these incorrect inputs
_____
# 3.4 Arthmatic Operations on variables
Same Datatypes
> `remainingtickets := totalTickets - usedTickets` <br>

Different datatypes:
> `totalCost := remainingtickets * cost // Error` <br> 
> `totalCost := float64(remainingtickets) * cost // No Error`

Note : only same datatype can be equated
_____
# 4 Array 
Predefied size and some default elements :
> `var array_letters = [50]string{"a", "b", "c"}` <br> 
> `fmt.Print("array_letters: ", array_letters)` <br> 

Predefined Empty array & Assigining values later and accessing them :
> `var bookings = [50]string{}` <br> 
> `bookings[0] = "Dear " + firstName` <br> 
> `bookings[1] = lastName` <br> 
> `fmt.Print(bookings[0])` <br>

Get the type of a item of an array:
> `fmt.Printf("Type : %T", bookings[0]) // type of item in array` <br> 
> `fmt.Printf("Type : %T \n", bookings) // type of the array` <br>
 
Get the length of :
> `fmt.Printf("length : %v", len(bookings[0])) // length of item in array` <br> 
> `fmt.Printf("length : %v \n ", len(bookings)) // length of array`
_____
# 5 Slice 
Abstraction of array but it doesnt need the size declared and its dynamic i.e. it has a size but it can be resized <br>
Initilizing a slice :
> `var array_alphanum = []string{}`<br> 

Append to slice : 
> `array_alphanum = append(array_alphanum, "a1")`<br> 

Print Slice & items :
> `print("\t", array_alphanum[0], "\n")`<br> 
> `print("\t", array_alphanum)`<br> 


_____
# 11 if & else
### (i) if :
Example 1 :
>`  if 20 > 18 {`<br>
>`    fmt.Println("20 is greater than 18")`<br>
?`  }`<br>
Example 2 
>`  x:= 20`<br>
>`  y:= 18`<br>
>`  if x > y {`<br>
>`    fmt.Println("x is greater than y")`<br>
>`  }`<br>

### (ii) if else :
> `  time := 20` <br>
> `  if (time < 18) {` <br>
> `    fmt.Println("Good day.")` <br>
> `  } else {` <br>
> `    fmt.Println("Good evening.")` <br>
> `  }` <br>

### (iii) else if :
>` time := 22`<br>
>` if time < 10 {`<br>
>`   fmt.Println("Good morning.")`<br>
>` } else if time < 20 {`<br>
>`   fmt.Println("Good day.")`<br>
>` } else {`<br>
>`   fmt.Println("Good evening.")`<br>
>` }`<br>

### (iv) nested if :
>` if num >= 10 {`<br>
>`   fmt.Println("Num is more than 10.")`<br>
>`   if num > 15 {`<br>
>`     fmt.Println("Num is also more than 15.")`<br>
>`    }`<br>
>` } else {`<br>
>`   fmt.Println("Num is less than 10.")`<br>
>` }`<br>

Note : Following are the Go operators :
- Comparison
  - Less than <
  - Less than or equal <=
  - Greater than >
  - Greater than or equal >=
  - Equal to ==
  - Not equal to !=
- Logical
  - Logical AND &&
  - Logical OR ||
  - Logical NOT !
- Arthmatic
  - Addition	x + y
  - Subtraction	x - y
  - Multiplication x * y
  - Division x / y
  - Modulus	Returnx % y
  - Increment	x++	
  -	Decrement x--
_____
# 12 Switch statement
> ` option := 1` <br>
> ` switch day {` <br>
> ` case 1:` <br>
> `   fmt.Println("Juice")` <br>
> ` case 2:` <br>
> `   fmt.Println("Chocolate")` <br>
> ` case 3,4: ` // Example for multi-case <br> 
> `   fmt.Println("Chips")` <br>
> ` case 5 :` <br>
> `   fmt.Println("Jelly")` <br>
> `default:` <br>
> `   fmt.Println("Selection was Invalid")` <br>
> ` }` <br>
` <br>

# 13 for loops
infinite loop : <br>
Explicite : 
> `for True {`  <br> 
     `logic` <br> 
     `}` <br>
  
Implicite : 
> `for {`  <br> 
     `logic` <br> 
     `}` <br>

Example 1 : Iterable and  index ignored using underscore 
>`for _, booking := range bookings {`  <br> 
>` 		var names = strings.Fields(booking)` <br> 
>` 		firstNames = append(firstNames, names[0])` <br> 
> `}` <br> 

Example 2 : set of items
> `for i:=0; i < 5; i++ {` <br> 
> `      fmt.Println(i)`<br> 
> `}`<br> 

Example 3 : 
> `fruits := [3]string{"apple", "orange", "banana"}`<br>
> `for idx, val := range fruits {`<br>
> `fmt.Printf("%v\t%v\n", idx, val)`<br>
> `}`<br>

Example 4 : With continue
> `func main() {` <br> 
> `  for i:=0; i < 5; i++ {` <br> 
> `    if i == 3 {` <br> 
> `      continue`//  skip 1 or more iterations in a loop  <br> 
> `    }` <br> 
> `   fmt.Println(i)` <br> 
> `  }` <br> 
> `}` <br> 
> 
Example 5 : With break
>`func main() {` <br>
>`  for i:=0; i < 5; i++ {` <br>
>`    if i == 3 {` <br>
>`      break` // to break/terminate the loop execution<br>
>`    }` <br>
>`   fmt.Println(i)` <br>
>`  }` <br>
>`}` <br>

Example 6  : Nested loop
> `func main() {` <br>
> `  adj := [2]string{"big", "tasty"}` <br>
> `  fruits := [3]string{"apple", "orange", "banana"}` <br>
> `  for i:=0; i < len(adj); i++ {` <br>
> `    for j:=0; j < len(fruits); j++ {` <br>
> `      fmt.Println(adj[i],fruits[j])` <br>
> `    }` <br>
> `  }` <br>
> `}` <br>


_____
# 14 Function
Create a Function : 
> `func myMessage() {` <br>
> `  fmt.Println("I just got executed!")` <br>
> `}` // this is the function<br> 
> `func main() {` <br>
> `  myMessage()` // This is the function <br>
> `}`  <br>

Parameters and Arguments : 
Example 1 : One paramater
> `func familyName(fname string) {` <br>
> `  fmt.Println("Hello", fname, "Refsnes")` <br>
> `}` <br>
> `func main() {` <br>
> `  familyName("Liam")` <br>
> `}` <br>

Example 2 : Multi paramater
> `func familyName(fname string, age int) {` <br>
> `  fmt.Println("Hello", age, "year old", fname, "Refsnes")` <br>
> `}` <br>
> `func main() {` <br>
> `  familyName("Liam", 3)` <br>
> `}` <br>

Return Values : 
Example 1 :   Return single value :
> `func myFunction(x int, y int) (result int) {` <br>
> `  result = x + y` <br>
> `  return` <br>
> `}` <br>

> `func main() {` <br>
> `fmt.Println(myFunction(1, 2))` <br>
> `total := myFunction(1, 2)` //Store the Return Value in a Variable <br>
> `fmt.Println(total)` <br>
> `}` <br>

Example 2 : Return multiple values :
> `func myFunction(x int, y string) (result int, txt1 string) {` <br>
> `  result = x + x` <br>
> `  txt1 = y + " World!"` <br>
> `  return` <br>
> `}` <br>

> `func main() {` <br>
> `  fmt.Println(myFunction(5, "Hello"))` <br>
> `  a, b := myFunction(5, "Hello")` // we store the two return values into two variables <br>
> `  fmt.Println(a, b)` <br>
> `}` <br>

Example 2 : Return multiple values :
> `func myFunction(x int, y string) (result int, txt1 string) {` <br>
> `  result = x + x` <br>
> `  txt1 = y + " World!"` <br>
> `  return` <br>
> `}` <br>

> `func main() {` <br>
> `  fmt.Println(myFunction(5, "Hello"))` <br>
> `  _, b := myFunction(5, "Hello")` // we store the two return values into two variables <br>
> `  fmt.Println(a, b)` <br>
> `}` <br>

Recursion Functions : 
> `func factorial_recursion(x float64) (y float64) {` <br>
> ` if x > 0 {` <br>
> `    y = x * factorial_recursion(x-1)` <br>
> ` } else {` <br>
> `    y = 1` <br>
> `}` <br>
> `  return` <br>
> `}` <br>

> `func main() {` <br>
> `  fmt.Println(factorial_recursion(4))` <br>
> `}` <br>


# 15 Global Variables or package level variables
> `package main` <br>
> `import (` <br>
> `	"fmt"` <br>
> `	"strings"` <br>
> `)` <br>
> // the following are package level variable <br>
> `const conferenceTickets int = 50` <br>
> `var remainingTickets uint = 50` <br>
> `var conferenceName = "Go Conference"` <br>
> `var bookings = []string{}` <br>
> // Main function
> `func main() {` <br>
> `main logic` <br>
> `}` <br>
_____
# 16 Multi package
- A Go application can have more than one file (See Example: helper.go)
  - They can be accessed by the main function just like a function call
  
- The go files can be stored in folders
  - If in folders we need to import it like any inbuilt package
  Example:
>`package main`<br>
>`import (` <br>
> `"booking-app/helper"` <br>
> `)`<br>

- Capatalizing
  - Function names - Exports the function to be accessable
  - Variable names - Will make it available globally
_____
# 16 Struct

Example 1 : 
> `type Person struct {` <br>
> `  name string` <br>
> `  age int` <br>
> `  job string` <br>
> `  salary int` <br>
> `}` <br><br>
> `func main() {`<br>
> `  var pers1 Person`<br>
> `  var pers2 Person`<br>
> `  // Pers1 specification`<br>
> `  pers1.name = "Hege"`<br>
> `  pers1.age = 45`<br>
> `  pers1.job = "Teacher"`<br>
> `  pers1.salary = 6000`<br> <br>
> `  // Pers2 specification`<br>
> `  pers2.name = "Cecilie"`<br>
> `  pers2.age = 24`<br>
> `  pers2.job = "Marketing"`<br>
> `  pers2.salary = 4500`<br> <br>
> `  // Access and print Pers1 info`<br>
> `  fmt.Println("Name: ", pers1.name)`<br>
> `  fmt.Println("Age: ", pers1.age)`<br>
> `  fmt.Println("Job: ", pers1.job)`<br>
> `  fmt.Println("Salary: ", pers1.salary)`<br> <br>
> `  // Access and print Pers2 info`<br>
> `  fmt.Println("Name: ", pers2.name)`<br>
> `  fmt.Println("Age: ", pers2.age)`<br>
> `  fmt.Println("Job: ", pers2.job)`<br>
> `  fmt.Println("Salary: ", pers2.salary)`<br>
> `}`<br>

Example 2 : 
> `package main` <br>
> `import ("fmt")` <br>

> `type Person struct {` <br>
> `  name string` <br>
> `  age int` <br>
> `  job string` <br>
> `  salary int` <br>
> `}` <br>

> `func main() {` <br>
> `  var pers1 Person` <br>
> `  var pers2 Person` <br>

> `  // Pers1 specification` <br>
> `  pers1.name = "Hege"` <br>
> `  pers1.age = 45` <br>
> `  pers1.job = "Teacher"` <br>
> `  pers1.salary = 6000` <br>

> `  // Pers2 specification` <br>
> `  pers2.name = "Cecilie"` <br>
> `  pers2.age = 24` <br>
> `  pers2.job = "Marketing"` <br>
> `  pers2.salary = 4500` <br>

> `  // Print Pers1 info by calling a function` <br>
> `  printPerson(pers1)` <br>
> `` <br>
> `  // Print Pers2 info by calling a function` <br>
> `  printPerson(pers2)` <br>
> `}` <br>

> `func printPerson(pers Person) {` <br>
> `  fmt.Println("Name: ", pers.name)` <br>
> `  fmt.Println("Age: ", pers.age)` <br>
> `  fmt.Println("Job: ", pers.job)` <br>
> `  fmt.Println("Salary: ", pers.salary)` <br>
> `}` <br>
_____
# 17 Maps
With values : 
> `var a = map[KeyType]ValueType{key1:value1, key2:value2,...}` <br>
> `b := map[KeyType]ValueType{key1:value1, key2:value2,...}` <br>

Empty map : 
>`var a map[KeyType]ValueType`<br>

Empty map with map function : 
> `var a = make(map[KeyType]ValueType)` <br>
> `b := make(map[KeyType]ValueType)` <br>

Accessing Map Elements :
> `value = map_name[key]`<br>

Updating and Adding Map Elements
> `map_name[key] = value`<br>

Deleting map item :
> `delete(map_name, key)`<br>

Check For Specific Elements in a Map :
val, ok :=map_name[key]

Iteration :

No order:
> `func main() {`<br>
> `  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}`<br>
> `  for k, v := range a {`<br>
> `    fmt.Printf("%v : %v, ", k, v)`<br>
> `  }`<br>
> `}`<br>
With order:
> `func main() {` <br>
> `  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}` <br>
> `` <br>
> `  var b = []string             // defining the order` <br>
> `  b = append(b, "one", "two", "three", "four")` <br>
> `` <br>
> `  for k, v := range a {        // loop with no order` <br>
> `    fmt.Printf("%v : %v, ", k, v)` <br>
> `  }` <br>
> `` <br>
> `  fmt.Println()` <br>

  for _, element := range b {  // loop with the defined order
    fmt.Printf("%v : %v, ", element, a[element])
  }
}

Note :
- Allowed key types = Booleans, Numbers, Strings, Arrays, Structs, Interfaces, Pointers , ~~Slices~~, ~~Maps~~ & ~~Functions~~)
- Allowed value types = any
- Maps are references to hash tables. - If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
 
_____
# 18 [Pointer](https://gobyexample.com/pointers)