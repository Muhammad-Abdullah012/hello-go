package main;

import "fmt";
import "rsc.io/quote"; // Go packages are on https://pkg.go.dev/
import "greetings"; // 

func main() {
	fmt.Println("Hello world!"); // from import "fmt"
	fmt.Println(quote.Glass()); // from import "rsc.io/quote" => "quote"
	fmt.Println(greetings.Hello("Abdullah")); // from import "greetings" => "greetings"
}