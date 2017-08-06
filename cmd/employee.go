// Package cmd will command and control you.
package cmd

// Employee contains Name and Age with several functions.
// You can use it if you like
type Employee struct {
	Name string
	Age  int
}

// GetAge enable you to get fake age. This function returns an age which is actual age plus one.
func (c *Employee) GetAge() int {
	return c.Age + 1
}

// GetRealAge get the actual age
//
// Deprecated: use GetAge instead
func (c *Employee) GetRealAge() int {
	return c.Age
}
