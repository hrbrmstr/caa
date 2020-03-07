package main

/*
  #define USE_RINTERNALS
  #include <R.h>
  #include <Rinternals.h>

  SEXP MakeDF(int n, char** tag, char** val);

*/
import "C"
import "log"
import "io/ioutil"
import "unsafe"
import "github.com/weppos/dnscaa"

// import "fmt"

//export caa_dig
func caa_dig(hostname string) C.SEXP {

  records, err := dnscaa.Lookup(hostname)

  if ((err != nil) || (len(records) == 0)) {
		return(C.R_NilValue)
  } else {

    // go slices for our columns

    val := []string{}
    tag := []string{}

    for _, r := range records {
      val = append(val, r.Value)
      tag = append(tag, r.Tag)
    }

    // ugly setup & population code for turning those slices
    // into char** so we can pass them to the C code which
    // will make an R data frame

    // allocate the array of char* pointers
    val_arr := C.malloc(C.size_t(len(val)) * C.size_t(unsafe.Sizeof(uintptr(0))))
    val_go := (*[1<<30 - 1]*C.char)(val_arr) // https://stackoverflow.com/questions/48756732/what-does-1-30c-yourtype-do-exactly-in-cgo

    // populate the array
    for idx, elem := range val {
      val_go[idx] = C.CString(elem)
    }

    tag_arr := C.malloc(C.size_t(len(tag)) * C.size_t(unsafe.Sizeof(uintptr(0))))
    tag_go := (*[1<<30 - 1]*C.char)(tag_arr)

    for idx, elem := range tag {
      tag_go[idx] = C.CString(elem)
    }

    // how many things
    n := C.int(len(val)) // don't need to free https://blog.golang.org/c-go-cgo

    // clean up after calling the data frame maker

    defer C.free(unsafe.Pointer(val_arr))
    defer C.free(unsafe.Pointer(tag_arr))

    return(C.MakeDF(n, (**C.char)(tag_arr), (**C.char)(val_arr)));

  }

}

func init() {
  log.SetOutput(ioutil.Discard)
}

func main() {}

