package main

import (
  "fmt"
  "math"
  "math/rand"
  "time"
  "reflect"
)

type json_response_struct struct {
    Struct map[string]interface{}
    List []map[string]interface{}
  }

func testerFunc(ref *json_response_struct){
  fmt.Println("ref.List: ", (ref.List))
  fmt.Println("ref.Struct: ", (ref.Struct))
  fmt.Println("ref.List is null: ", (ref.List == nil))
  fmt.Println("ref.Struct is null: ", (ref.Struct == nil))
  ref.Struct = map[string]interface{}{"test":"blah2"}
}

func main() {
  // Basic operations
  fmt.Println("hello!")
  fmt.Println("testing " + "concatenation")
  fmt.Println("testing"," ","commas")
  fmt.Println("testing","no space commas") // fmt.Println automatically inserts space for commas
  fmt.Println("testing integer = ",5)
  fmt.Println("testing 5+37 = ",5+37)
  fmt.Println("testing 37/5 = ",37/5)
  fmt.Println("testing 37/5.0 = ",37/5.0)
  fmt.Println("testing 37.0/5 = ",37.0/5)
  fmt.Println("testing 37.0/5.0 = ",37.0/5.0)
  fmt.Println("testing 37 % 5 = ",37 % 5)
  fmt.Println("testing boolean = ",true)
  fmt.Println("testing true && false = ",true&&false)
  fmt.Println("testing true || false = ",true||false)
  // fmt.Println("testing 1 || 0 = ",1||0) // This is not allowed
  fmt.Println("testing !true = ",!true)
  // Random number generation
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  fmt.Println("random int in [0,100) = ",r1.Intn(100))
  fmt.Println("random Float64 = ",r1.Float64())

  // Variables
  var a = "string"
  b := "string"
  fmt.Println("a, b = ",a,b)
  fmt.Println("a = b ? ",a==b)
  var c int
  d := 0
  e := 0.0
  fmt.Println("c as uninitialized int = ",c)
  fmt.Println("c,d,e = ",c,d,e)
  fmt.Println("c = d ? ",c==d)
  // fmt.Println("c = e ? ",c==e) // This fails for mismatched types
  var f,g int = 5,6
  fmt.Println("f,g = ",f,g)
  var h = true
  fmt.Println("h = ",h)
  var i string
  fmt.Println("i as uninitialized string = ("+i+")")
  var j bool
  fmt.Println("j as uninitialized bool = ",j)

  // Constants
  const a2 = "string"
  fmt.Println("a2: ",a2)
  const b2 = 50000
  const c2 = 5.78
  const d2 = 3e20 / b2
  fmt.Println("b2 = ",b2)
  fmt.Println("c2 = ",c2)
  fmt.Println("d2 = ",d2)
  fmt.Println("int64(d2) = ",int64(d2))
  fmt.Println("sin(b2) = ",math.Sin(b2))


  // For loops
  i2 := 0
  for i2 >= 3 {
    fmt.Println(i2)
    i2 = i2 + 1
  }
  for j2 := 7; j2 <= 9; j2++ {
    fmt.Println(j2)
  }
  for {
    fmt.Println("loop!")
    break
  }
  for i2 = 0; i2 <= 6; i2++ {
    if ((i2 % 2) == 0){
      continue
    }
    fmt.Println(i2)
  }

  // If Else statements
  if 9 % 2 == 0 {
    fmt.Println("9 is even")
  } else {
    fmt.Println("9 is odd")
  }
  if 8 % 4 == 0 {
    fmt.Println("8 is divisible by 4")
  }
  if num := 9; num < 0 {
    fmt.Println(num,"is negative")
  } else if num < 10 {
    fmt.Println(num,"is less than 10")
  } else {
    fmt.Println(num,"has multiple digits")
  }

  // Switch statements
  var i3 int = 2
  fmt.Println("variable i3 is: ")
  switch i3 {
  case 1:
    fmt.Println("One")
  case 2:
    fmt.Println("Two")
  case 3:
    fmt.Println("Three")
  }
  
  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's a weekend")
  default:
    fmt.Println("It's a weekday")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }

  typeDeterminator := func(i interface{}) {
    switch i.(type) {
    case bool:
      fmt.Println("Input",i,"is a boolean")
    case int:
      fmt.Println("Input",i,"is a integer")
    case float64,float32:
      fmt.Println("Input",i,"is a float!")
    default:
      fmt.Println("Cannot tell what Input",i,"is!")
    }
  }
  typeDeterminator(true)
  typeDeterminator(10)
  typeDeterminator(5.7)
  typeDeterminator("blah!")
  fmt.Println(reflect.TypeOf(5.7))

  // Arrays
  var a3 [5] int
  fmt.Println("a3: ",a3)
  a3[4] = 100
  fmt.Println("a3: ",a3)
  fmt.Println("length of a3: ",len(a3))

  b3 := [5] int {10,15,20,4,7}
  fmt.Println("b3: ",b3)

  var c3 [2][3] int
  for i4:=0; i4<2; i4++ {
    for j4:=0; j4<3; j4++ {
      c3[i4][j4] = i4+j4
    }
  }
  fmt.Println("c3: ",c3)

  // Structs
  var jsonbody json_response_struct 
  jsonbody.Struct = map[string]interface{}{"test":"blah"}
  fmt.Println("jsonbody: ", jsonbody)
  fmt.Println("jsonbody.List is null: ", (jsonbody.List == nil))
  fmt.Println("jsonbody.Struct is null: ", (jsonbody.Struct == nil))

  testerFunc(&jsonbody)

  fmt.Println("jsonbody after testFunc call: ", jsonbody)

  // Maps
  test_input := map[string]interface{}{}
  test_input = map[string]interface{}{
    "test1": "value1",
    "test2": map[string]interface{}{
      "test2.1": "value2.1",
      "test2.2": map[string]interface{}{
        "test2.2.1": "value2.2.1",
      },
    },
  }
  fmt.Println("test_input: ",test_input)

  // Slices

  // Channels

  
}