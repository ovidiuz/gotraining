// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/oijJdRW3cD

// Declare an interface named speaker with a method named sayHello. Declare a struct
// named English that represents a person who speaks english and declare a struct named
// Chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the _address of_ a value of type English
// and call the method. Do it again for a value of type Chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the sayHello method on the interface value. Then create
// new values of each type and use the function.
package main

import "fmt"

// Add imports.

// Declare the speaker interface with a single method called sayHello.
type speaker interface {
	sayHello()
}
// Declare an empty struct type named english.

type english struct{}
// Declare a method named sayHello for the english type
// using a value receiver. "Hello World"

func (e english) sayHello() {
	fmt.Println("Hello World")
}

// Declare an empty struct type named chinese.
type chinese struct{}

// Declare a method named sayHello for the chinese type
// using a value receiver. "你好世界"
func (c chinese) sayHello() {
	fmt.Println("你好世界")
}
// sayHello accepts values of the speaker type.
func sayHello(s speaker) {
	// Call the sayHello() method from the speaker parameter.
	s.sayHello()
}

// main is the entry point for the application.
func main() {
	// Declare a variable of the speaker type set to its zero value.
	var sp speaker

	// Declare a variable of type english and assign it to
	// the speaker variable.
	var eng english
	sp = eng

	// Call the sayHello() method from the speaker variable.
	sp.sayHello()

	// Declare a variable of type chinese and assign it to
	// the speaker variable.
	// Call the sayHello() method from the speaker variable.
	var chi chinese
	sp = chi

	// Call the sayHello function passing each concrete type.
	sp.sayHello()
}
