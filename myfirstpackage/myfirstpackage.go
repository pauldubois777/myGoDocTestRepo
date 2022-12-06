package myfirstpackage

// myFirstType is the first type I created
type myFirstType struct {
	// fld1 is the first field in my first type
	fld1 string
	// fld2 is the second field in my first type
	fld2 string
}

// myFirstConst is my first constant
const myFirstConst string = "Hello"

// GetHello is my first function
//
// Returns a nice hello message to who is passed in
func GetHello(who string) string {
	return myFirstConst + " " + who
}
