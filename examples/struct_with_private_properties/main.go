package main

import (
	"fmt"
	"struct_with_private_properties/pets"
)

func main() {
	p := pets.Pet{}
	p.SetName("Dela")
	fmt.Println(p.Name())
}
