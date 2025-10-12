// src: https://go.dev/doc/tutorial/create-module
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
  // Return a greeting that embeds the name in a message
  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  // alternatively define message first with:
  // var message string
  // message := fmt.Sprintf("Hi, %v. Welcome!", name)
  return message
}
