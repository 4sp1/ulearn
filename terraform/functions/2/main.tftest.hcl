run "env_multiplier_test" {
  variables {
    multiplier = {
      dev   = 1
      stage = 2
      prod  = 3
    }
    instance_count = 2
  }
  assert {
    condition     = 2 == var.instance_count * lookup(var.multiplier, "dev", 1) # defaults to 1
    error_message = "dev instance count should be 2"
  }
  assert {
    condition     = 4 == var.instance_count * lookup(var.multiplier, "stage", 1)
    error_message = "stage instance count should be  4"
  }
  assert {
    condition     = 6 == var.instance_count * lookup(var.multiplier, "prod", 1)
    error_message = "stage instance count should be  6"
  }
}
