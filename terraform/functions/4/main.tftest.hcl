run "element_test" {
  variables {
    subnets = ["subnet-0", "subnet-1", "subnet-2", "subnet-3"]
    cidrs = {
      subnet_0 = "0.0.0.0/8"
      subnet_1 = "128.0.0.0/8"
    }
    index = 0
  }
  assert {
    condition     = "subnet-0" == element(var.subnets, var.index)
    error_message = "element[${var.index}] should be subnet-0"
  }
  assert {
    condition     = "subnet-3" == element(var.subnets, -1)
    error_message = "last element of subnets should be subnet-3"
  }
  assert {
    condition     = "0.0.0.0/8" == lookup(var.cidrs, "subnet_0", "")
    error_message = "cidrs[subnet_0] should be 0.0.0.0/8"
  }
  assert {
    condition     = "10.0.0.0/24" == lookup(var.cidrs, "subnet_local", "10.0.0.0/24")
    error_message = "default lookup value should be 10.0.0.0/24"
  }
}
