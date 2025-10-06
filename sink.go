// Following this Tutorial: https://www.geeksforgeeks.org/go-language/go-programming-language-introduction/
// This program is to illustrate "keywords" in go
// They seem like they are more or less variables in others, but capture some other stuff. Like "import as" in Python.

// AH scratch that. Keywords are the generic terms used to define or call other things.
// The correct case for what I was talking about before are "identifiers"

// Also I have changed my mind. This is the Go "stuf" sandbox file

// package is a keyword
package main

// import is a keyword 
import {
  "fmt"
  "math" // ERROR Missing Import Path? // Come back here on return.
}

// fmt is a (sub?)"package" in the go standard library package.
// It's more or less a simpler version of C's io functions printf and scanf.
// ref: https://pkg.go.dev/fmt

func main() {

  /*
   * keywords()
   *
   */

  var hexen = 0xFee
    //fmt.printf(0xFeeL, "\n")
  fmt.Println(hexen , "\n")

  var sciBrain = 13e7
  fmt.Println(sciBrain , "\n")

 // var galaxyBrain = 123.456e30, 0755775)
 // fmt.Println(galaxyBrain, "\n");
  var bigIf = true
  if (sciBrain > 0) { // Only bools
    fmt.Println(bigIf , "\n")
  }

}

func keywords() {
  //var Aname("Go GoLang Go!")
  // Go doesn't like the above assignment syntax
  var Aname = "Go GoLang Go!"
  var Bname = "Go Desiderius Erasmus Roterodamus Go!"
  // I would asume the "fmt.FUNCTIONi()" is akin to the cpp std:: --This makes me wonder if there is something
  // like 'using namespace' in Golang...
  fmt.Printf("We are using this space to: %s\n", Aname) // NO! NO SEMICOLONS!!
  fmt.Printf("Now on a bicycle bulit for two: %1s and %2s!\n", Aname, Bname)
}

