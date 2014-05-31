// Package newmath 
package newmath

// uses the lightweight test framework part of go language
import "testing"

// test the Sqrt function of newmath package
func TestSqrt(t *testing.T) {
	const input, output = 4, 2
	if x := Sqrt(input); x != output {
		t.Errorf("Sqrt(%v) = %v, and should be %v", input, x, output)
	}
}
