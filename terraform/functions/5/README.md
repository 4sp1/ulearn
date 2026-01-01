# Functions

## Filesystem

You can find multiple tests in [`main.tftest.hcl`](main.tftest.hcl). Each test
is a `run` block followed by its name:
- `jackpot_test` reads [`gold.txt`](gold.txt) and
  [`gold_no_newline.txt`](gold_no_newline.txt) respectively containing the
  string `jackpot` and `jackpot\n`.  It also features the [`trimspace`](
  https://opentofu.org/docs/language/functions/trimspace/) function.
- `minerals_test` checks that the [`fileset`](
  https://opentofu.org/docs/language/functions/fileset/) defined in
  [`main.tf`](main.tf) only contains actual files.  It also feature
  [`trimsuffix`]( https://opentofu.org/docs/language/functions/trimsuffix/) to
  remove the `.txt` extension.
- `gem_test` checks that the [`fileset`](
  https://opentofu.org/docs/language/functions/fileset/) expanded from
  `local.gems` does not contains `emerald` as there is a voluntary typo in
  `local.gems`.  It also feature [`contains`](
  https://opentofu.org/docs/language/functions/contains/) function
- `confs_test` finally demonstrates the wildcard (`*`) pattern of [`fileset`](
  https://opentofu.org/docs/language/functions/fileset/) but you can find other
  supported patterns in the
  [docs](https://opentofu.org/docs/language/functions/fileset/)

In case you missed it
- the reason there is a `main.tf` file is to produce `output "jackpot"` (and
  `output "jackpot_line`) as the value of
  [`file`](https://opentofu.org/docs/language/functions/file/) is only
  available once state is computed. 
