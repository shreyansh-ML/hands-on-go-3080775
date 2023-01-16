// challenges/types-composite/begin/main.go
package main

import "fmt"

// define an author type with a name
//
type author struct {
	name string
}

// define a book type with a title and author
//
type book struct {
	//var title string
	//author
	books map[string]author
}

// define a library type to track a list of books
//
type library struct {
	b []book
}

// define addBook to add a book to the library
//
func (lib *library) addBook(bk book) {
	lib.b = append(lib.b, bk)
}

// define a lookupByAuthor function to find books by an author's name
//
func (lib *library) lookByAuthor(a author) (b []book) {
	var lst []book
	for _, bk := range lib.b {
		for _, v := range bk.books {
			if v == a {
				lst = append(lst, bk)
			}
		}
	}
	return lst
}
func main() {
	// create a new library
	//
	var lib library

	// add 2 books to the library by the same author
	//
	lib.addBook(book{map[string]author{"abc": {"rush"}}})
	lib.addBook(book{map[string]author{"def": {"rush"}}})
	// add 1 book to the library by a different author
	//
	lib.addBook(book{map[string]author{"abrakadabra": {"funny"}}})
	// dump the library with spew
	//

	// lookup books by known author in the library
	//
	temp := lib.lookByAuthor(author{"rush"})
	// print out the first book's title and its author's name
	//
	for _, j := range temp {
		for k, v := range j.books {
			fmt.Println(k, v.name)
		}

	}

}
