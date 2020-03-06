package main

import "C"
import "github.com/weppos/dnscaa"

//export caa_dig
func caa_dig(hostname string) *C.char {
  records, err := dnscaa.Lookup(hostname)
  if err != nil {
		return C.CString("")
  } else {
	  out := records[0].Value
	  return(C.CString(out))
  }
}

func main() {}

