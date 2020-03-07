#' Retrieve the CAA record values for a domain (if any)
#'
#' @param x domain name
#' @export
#' @return data frame (tibble)
#' @examples
#' # one record back
#' caa_dig("google.com")
#'
#' # multiple
#' caa::caa_dig("www.comodo.com")
#'
#' # none (lookup error)
#' caa::caa_dig("www.comodo.comm")
caa_dig <- function(x) {

  out <- .Call("R_caa_dig", x, PACKAGE = "caa")

  if (length(out) == 0) {

    data.frame(
      tag = character(0),
      value = character(0),
      stringsAsFactors = FALSE
    ) -> out

  }

  as_tibble(out)

}
