// Package newmath 
package newmath

// Sqrt returns an approximation to the square root of x.
// here Sqrt is an exported method, S capital hence method is called exported and that makes id available outside the package
// n float64 : is the input paramter passed to the method. This is basically a 64 bit floating point number
// the next float64 is the return type of this method
func Sqrt(n float64) float64{
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - n) / (2 * z)
	}
	return z
}
