# Functions

## Type Conversion

[`main.tftest.hcl`](main.tftest.hcl) demonstrates a curiosity about string
templates one should not expect. It is an edge case and I can't see a situation
where this would happen. If your code is well covered by unit tests, you
shouldn't encounter this issue. Indeed string template `"${var.some_number}"`
will not be a `string` but a `number`. If you don't believe me run

```bash
terraform test
```

or 

```bash
tofu test
```

This is where explicitly using `tostring` can safeguard from akward results and
bugs that would be quite hard to inspect.

I recommend you read about [types](
https://opentofu.org/docs/language/expressions/types) and [types conversions](
https://opentofu.org/docs/language/expressions/types/#type-conversion).
