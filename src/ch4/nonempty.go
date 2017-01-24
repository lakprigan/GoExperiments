// Non Empty is an example of in-place slice algorithm

package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings
// the underlying array is modified during the call
func nonempty(strings []string) []string {
  i := 0
  for _,s := range strings {
    if s!= "" {
      strings[i] = s
      i++
    }
  }
  return strings [:i]
}
