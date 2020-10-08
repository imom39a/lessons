// Computes the greatest common divisor of p and q using Euclid's algorithm.
// If p > q, the gcd of p and q is the same as the gcd of q and p % q.

package main 

import "fmt"

func main(){

  fmt.Println("Euclid's algorithm for greatest common divisor")

  p := 120
  q := 36
  fmt.Println(gcd(p, q))
}

func gcd(p, q int) int {
  fmt.Println(p,q)
  if q == 0 {
    return p
  }
  d := p % q
return gcd(q, d)
}
