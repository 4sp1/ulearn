run "tostring_test" {
  variables {
    default = 5
  }

  assert {
    condition     = tostring(var.default) == "5"
    error_message = "itoa 5 should be \"5\""
  }

  assert {
    condition     = "5" != "${var.default}"
    error_message = "numbers in string templates are not strings"
  }

  assert {
    condition     = 5 == "${var.default}"
    error_message = "numbers in strings template are numbers"
  }

  assert {
    condition     = "instance-5" == "instance-${var.default}"
    error_message = "string has precedence on number with string template"
  }
}

run "automatic_conversion_test" {
  variables {
    true = true
  }

  assert {
    condition     = "this is ${var.true}" == "this is true"
    error_message = "boolean true should be automatically converted to \"true\""
  }

  assert {
    condition     = "${var.true}" == true
    error_message = "when there is only a string template with a bool it's not converted to string"
  }
}
