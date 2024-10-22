package main

import "fmt"

func main() {
	// Use 'goto' to jump to the 'end' label.
	fmt.Println("Jumping to the end of the function!")
	goto end

	fmt.Println("This will be skipped")

end:
	fmt.Println("End of function reached")
}
