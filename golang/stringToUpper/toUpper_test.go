package stringToUpper

import "testing"
import "fmt"
import "os"

func TestWrite(t *testing.T) {
	fmt.Fprintln(&UpperWrite{os.Stdout}, "hello, world")
}