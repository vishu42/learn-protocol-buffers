package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vishu42/learn-protocol-buffers/tutorialpb"
	"go.starlark.net/lib/proto"
)

func main() {
	// phoneNumber := &tutorialpb.Person_PhoneNumber{
	// 	Number: "8395955922",
	// 	Type:   tutorialpb.Person_MOBILE,
	// }

	// vishalAddress := &tutorialpb.Person{
	// 	Name:   "vishal tewatia",
	// 	Id:     1234,
	// 	Phones: []*tutorialpb.Person_PhoneNumber{phoneNumber},
	// }

	// // Write the new address book back to disk.
	// out, err := proto.Marshal(vishalAddress)
	// if err != nil {
	// 	log.Fatalln("Failed to encode address book:", err)
	// }
	// if err := ioutil.WriteFile("book", out, 0o644); err != nil {
	// 	log.Fatalln("Failed to write address book:", err)
	// }

	// Read the existing address book.
	fname := "./book"
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &tutorialpb.Person{}
	prot, err := proto.Unmarshal(book.ProtoReflect().Descriptor(), in)
	if err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(prot)
}
