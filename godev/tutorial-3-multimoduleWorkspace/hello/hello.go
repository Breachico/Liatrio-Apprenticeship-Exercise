package main

import (
  "fmt"
  "golang.org/x/example/hello/reverse"
)

func main() {
  
  fmt.Println(reverse.String("Go Template."), reverse.Int(24601))
}

