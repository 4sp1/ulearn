# Terraform

## Functions

- [Simple Sum Test](1/main.tftest.hcl) demonstrates the `sum` function.
- [Env multiplier](2/main.tftest.hcl) demonstrates `lookup` function and simple
  arithmetics.
- [Format date](3/main.tftest.hcl) features [`formatdate`](
  https://developer.hashicorp.com/terraform/language/functions/formatdate)
- [Indexing collections](4/main.tftest.hcl) featuring [`element`](
  https://opentofu.org/docs/v1.11/language/functions/element/) and [`lookup`](
  https://opentofu.org/docs/v1.11/language/functions/lookup/) functions.
- [Filesystem](5/README.md) explores [`file`](
  https://opentofu.org/docs/language/functions/fileset/) and [`fileset`](
  https://opentofu.org/docs/language/functions/file/).
- [Hash and Crypto](6/main.tftest.hcl) featuring
  [`sha256`](https://opentofu.org/docs/language/functions/sha256/) and
  [`filesha256`](https://opentofu.org/docs/language/functions/filesha256/).  We
  can notice that as well as in [Filesystem](5/README.md) example we need to
  proceed to plan or apply (the latter - apply - is not required) to evaluate
  [`filesha256`](https://opentofu.org/docs/language/functions/sha256/)
