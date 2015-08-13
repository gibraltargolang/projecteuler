// START OMIT
/***********************
 * Mission: Even Fibonacci numbers *
 ***********************
 *
 * Each new term in the Fibonacci sequence is generated by adding the previous two terms.
 * By starting with 1 and 2, the first 10 terms will be:
 *
 * 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
 *
 * By considering the terms in the Fibonacci sequence whose values do not exceed four million,
 * find the sum of the even-valued terms.
 *
 */
package main

import (
"fmt"
)

func main() {
  sum := 0
  // EDITABLE OMIT

  // Your code

  // UNEDITABLE OMIT
  fmt.Println(sum)
}

func fib(v1,v2 int) (int,int) {
  return v1+v2,v1
}
// END OMIT