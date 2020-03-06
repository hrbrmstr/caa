package main

/*
  #define USE_RINTERNALS
  #include <R.h>
  #include <Rinternals.h>

*/
import "C"
//import "fmt"
import "unsafe"
import "github.com/weppos/dnscaa"

//export caa_dig
func caa_dig(hostname string, n *C.int) **C.char {

  records, err := dnscaa.Lookup(hostname)

  if err != nil {
		return nil
  } else {

    goResult := []string{}

    for _, r := range records {
      goResult = append(goResult, r.Value)
    }

    cArray := C.malloc(C.size_t(len(goResult)) * C.size_t(unsafe.Sizeof(uintptr(0))))
    a := (*[1<<30 - 1]*C.char)(cArray)

    for idx, substring := range goResult {
      a[idx] = C.CString(substring)
    }

    *n = C.int(len(goResult))

	  //out := records[0].Value
	  //return(C.CString(out))

	  return (**C.char)(cArray)

  }
}

func main() {}

