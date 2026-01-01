package main
import "fmt"

const (
    Low = iota    // 0
    Medium        // 1
    High          // 2
)


func main() {
	fmt.Println(Low, Medium, High)
}