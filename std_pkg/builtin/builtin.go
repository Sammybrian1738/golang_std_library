package builtin

import (
	"log"
)

func Example() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}   // slice is a dynamically sized flexible
	array1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // array has a fixed size

	/* The append built-in function appends elements to the end of a slice. */

	// 	slice = append(slice, elem1, elem2)
	// slice = append(slice, anotherSlice...)
	log.Print("\nTESTIND APPEND BUILTIN FUNCTION \n\n")
	slice1 = append(slice1, slice1...)
	log.Println("appended slice: ", slice1)

	/* The cap built-in function returns the capacity of v, according to its type: */

	log.Print("\n\n TESTING CAP BUILTIN FUNCTION \n\n")
	arrayCap := cap(array1) // the number of elements in v
	sliceCap := cap(slice1) // the maximum length the slice can reach when resliced
	log.Println("array capacity: ", arrayCap)
	log.Println("slice capacity: ", sliceCap)

	/*The close built-in function closes a channel, which must be either bidirectional or send-only.*/

	log.Print("\n\n TESTING CLOSE BUILTIN FUNCTION \n\n")
	chan1 := make(chan int)
	go func() {
		chan1 <- 6
	}()

	// close(chan1)

	value, ok := <-chan1 // if we close the channel, the value should be 0 and ok false, else value is 6 and ok is true
	log.Println("printing chan value after closing. Value: ", value, "OK: ", ok)

	/* The complex built-in function constructs a complex value from two floating-point values. The real(r) and imaginary(i) parts*/

	log.Print("\n\n TESTING COMPLEX BUILTIN FUNCTION \n\n")
	var complex1 complex64 = 2 + 4i
	var complex2 complex64 = complex(2, 4)

	log.Println("complex 1: ", complex1)
	log.Println("complex 2: ", complex2)

	/* The copy built-in function copies elements from a source slice into a destination slice. */
	// copy does not change the len of the slice.. only replaces corresponding elements as per the indices

	log.Print("\n\n TESTING COPY BUILTIN FUNCTION \n\n")
	slice2 := []int{11, 12, 13, 14, 15, 16}                                                                // len(slice2) = 6
	slice3 := []int{17, 18, 19, 20, 21, 22, 23, 24, 25}                                                    // len(slice3) = 9
	number_of_elements_copied := copy(slice2, slice3)                                                      // It returns the number of elements copied, which will be the minimum of len(dst) and len(src)
	log.Println("number of elements copied from slice3(src) to slice2(dst) : ", number_of_elements_copied) // prints 6
	log.Println("new contents of slice2 : ", slice2)                                                       // prints [17, 18, 19, 20, 21, 22]

	// Copy from a string to a byte slice (special case)
	var b = make([]byte, 7)
	number_of_elements_copied = copy(b, "Hello World")

	log.Println("number of bytes copied from string to byte array: ", number_of_elements_copied)
	log.Println("new contents of byte array: ", string(b))
}
