import "fmt"

func main() {

	var input int

	fmt.Println("This program converts all decimal numbers to binary")
	fmt.Println("Kindly enter the number you want to convert: ")
	fmt.Scan(&input)
	fmt.Printf("The binary equivalent of the number %d is %b", input, input)
}
