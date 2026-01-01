locals {
  path = "${path.module}/../5/ruby.txt"
}

output "filesha256" {
  value = filesha256(local.path)
}
