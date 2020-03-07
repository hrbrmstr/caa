
res <- caa_dig("google.com")
expect_equal(res$value, "pki.goog")

expect_equal(nrow(caa::caa_dig("www.comodo.com")), 3)

expect_equal(nrow(caa::caa_dig("www.comodo.comm")), 0)
