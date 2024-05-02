package main;

import "fmt";
import "rsc.io/quote"; // Go packages are on https://pkg.go.dev/

func main() {
	fmt.Println("Hello world!"); // from import "fmt"
	fmt.Println(quote.Glass()); // from import "rsc.io/quote" => "quote"
}