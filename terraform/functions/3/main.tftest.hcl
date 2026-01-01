run "formatdate_test" {
  variables {
    date   = "2026-01-01T00:00:00Z" # RFC339
    format = "YYYYMMDD"
  }
  assert {
    condition     = "20260101" == formatdate(var.format, var.date)
    error_message = "formatdate($format, $date) should be 20260101"
  }
}
