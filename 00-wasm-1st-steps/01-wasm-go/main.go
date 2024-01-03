package main
// go
// add add (1) and hello (2) function

//export add
func add(x int, y int) int {
  return x + y
}

//export hello
func hello(name string) string {
  print(name)

  return "hello " + name
}


func main() {}
