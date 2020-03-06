#' Retrieve the CAA record values for a domain (if any)
#'
#' @param x domain name
#' @export
#' @examples
#' caa_dig("google.com")
caa_dig <- function(x) {
  .Call("R_caa_dig", x, PACKAGE = "caa")
}
