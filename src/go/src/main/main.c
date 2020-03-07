#include "_cgo_export.h"

#include <stdlib.h>
#include <string.h>

// define R bridge functions here

#define SHORT_VEC_LENGTH(x) (((VECSEXP) (x))->vecsxp.length)

SEXP MakeDF(int n, char** tag, char** val) {

  if (val != NULL) {

    SEXP value_vec = PROTECT(allocVector(STRSXP, n));
    for (int i=0; i<n; i++) {
      SET_STRING_ELT(value_vec, i, mkChar(val[i]));
      free(val[i]); // allocated on heap on the Go side
    }

    SEXP tag_vec = PROTECT(allocVector(STRSXP, n));
    for (int i=0; i<n; i++) {
      SET_STRING_ELT(tag_vec, i, mkChar(tag[i]));
      free(tag[i]); // allocated on heap on the Go side
    }

    // make column names
    SEXP colnames = PROTECT(allocVector(STRSXP, 2));
    SET_STRING_ELT(colnames, 0, mkChar("tag"));
    SET_STRING_ELT(colnames, 1, mkChar("value"));

    // make row names
    SEXP rownames = PROTECT(allocVector(INTSXP, 2));
    INTEGER(rownames)[0] = NA_INTEGER;
    INTEGER(rownames)[1] = -n;

    // the data frame (finally)
    SEXP out = PROTECT(allocVector(VECSXP, 2));
    setAttrib(out, R_ClassSymbol, ScalarString(mkChar("data.frame")));

    SET_VECTOR_ELT(out, 0, tag_vec);
    SET_VECTOR_ELT(out, 1, value_vec);

    setAttrib(out, R_RowNamesSymbol, rownames);
    setAttrib(out, R_NamesSymbol, colnames);

    UNPROTECT(5);

    return(out);

  } else {
    return(R_NilValue);
  }

}

SEXP R_caa_dig(SEXP x) {

  SEXP sx = STRING_ELT(x, 0);

  GoString h = { (char*) CHAR(sx), SHORT_VEC_LENGTH(sx) };

  SEXP out = caa_dig(h);

  return(out);

}
