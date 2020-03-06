#define USE_RINTERNALS
#include <R.h>
#include <Rinternals.h>
#include <stdlib.h>
#include <string.h>
#include "_cgo_export.h"

// define R bridge functions here

#define SHORT_VEC_LENGTH(x) (((VECSEXP) (x))->vecsxp.length)

SEXP R_caa_dig(SEXP x) {

  SEXP sx = STRING_ELT(x, 0);
  GoString h = { (char*) CHAR(sx), SHORT_VEC_LENGTH(sx) };

  char *res = caa_dig(h);

  if (res != NULL) {

    SEXP out = PROTECT(allocVector(STRSXP, 1));
    SET_STRING_ELT(out, 0, mkChar(res));
    UNPROTECT(1);

    free(res);

    return(out);

  } else {
    return(R_NilValue);
  }

}