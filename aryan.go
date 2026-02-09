package main
import (
"bufio"
"fmt"
"os"
"strings"
)
func main() {
// Output a prompt
fmt.Print("Enter your name: ")
// Use bufio.NewReader to read a full line of text, including spaces
reader := bufio.NewReader(os.Stdin)
name, _ := reader.ReadString('\n')
// Clean up the input string (remove the newline character)
name = strings.TrimSpace(name)
// Output the greeting
fmt.Println("Hello, " + name + "!")
// Example of reading a number with fmt.Scanf
var age int
fmt.Print("Enter your age: ")fmt.Scanf("%d", &age)
fmt.Println("You are", age, "years old.")
fmt.Println("Next year, you will be", age + 1)
}