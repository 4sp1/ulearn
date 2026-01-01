run "sum_test" {
  variables {
    sum = [1, 2, 3]
  }
  assert {
    condition     = 6 == sum(var.sum)
    error_message = "sum should be 6"
  }
}
