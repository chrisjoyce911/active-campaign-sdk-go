//go:build examples

package main

import "fmt"

// Example program demonstrating real client usage. This file is excluded from normal
// builds via the `examples` build tag. To build/run it, use `go run -tags=examples ./examples/contacts_real`.
func main() {
	fmt.Println("contacts_real example (build with -tags=examples)")
}
