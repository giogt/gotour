package interfaces

import (
	"fmt"
)

// Show returns a string representation for a type
type Show interface {
	Show() string
}

// The Vertex struct represents a 2D point with its coordinates
type Vertex struct {
	X, Y int
}

// Show implementation for Vertex
func (v *Vertex) Show() string {
	if v == nil {
		return "<nil>"
	}
	return fmt.Sprintf("[%v, %v]", v.X, v.Y)
}

/*
ShowToVertex provides access to the underlying concrete type Vertex from a Show interface

The type assertion happening in the function returns two values:
	- a pointer to the underlying value
	- a boolean value that reports whether the assertion succeeded or not

In case of success, it will return the tuple:
	- (a pointer to the underlying Vertex value, true)

In case of failure, it will return the tuple
	- (nil, true)
*/
func ShowToVertex(s Show) (*Vertex, bool) {
	v, ok := s.(*Vertex)
	return v, ok
}

/*
ShowToVertexUnsafe provides access to the underlying concrete type Vertex from a Show interface

The type assertion happening in the function returns the underlying Vertex
value, if the value is of Vertex type, otherwise it will panic.
*/
func ShowToVertexUnsafe(s Show) *Vertex {
	return s.(*Vertex)
}

// The Person struct represents a physical person with all its attributes
type Person struct {
	Name    string
	Surname string
	Age     uint
}

// Show implementation for Person
func (p *Person) Show() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("[name: %v, surname: %v, age: %v]", p.Name, p.Surname, p.Age)
}

/*
 * Describe any type of value.
 *
 * Empty interfaces (interface{}) may hold values of any type,
 * therefore they can be used by code that handles values of unknown type.
 */
func describe(i interface{}) string {
	var typeDescr string
	switch i.(type) {
	case uint:
		typeDescr = "unsigned integer"
	case Vertex:
		typeDescr = "2D vertex"
	case Person:
		typeDescr = "a person"
	default:
		typeDescr = "a nice type"
	}
	return fmt.Sprintf("(%v, %T) [%v]", i, i, typeDescr)
}
