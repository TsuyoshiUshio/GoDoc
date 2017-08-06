package cmd

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employee", func() {
	Context("When I ask them their age", func() {
		It("returns their actual age + 1", func() {
			emp := Employee{
				"Tsuyoshi Ushio",
				46,
			}
			Expect(emp.GetAge()).To(Equal(47))
		})
	})
})

func ExampleEmployee() {
	emp := Employee{
		"Tsuyoshi Ushio",
		46,
	}
	fmt.Println(emp.Name)
	fmt.Println(emp.GetAge())
	// Output:
	// Tsuyoshi Ushio
	// 47
}

func ExampleEmployee_GetAge() {
	emp := Employee{
		"Tsuyoshi Ushio",
		46,
	}
	fmt.Println(emp.GetAge())
	// Output: 47
}
