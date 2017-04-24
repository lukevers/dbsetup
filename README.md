# DBSETUP

A CLI tool that reads an [HCL](https://github.com/hashicorp/hcl) configuration file and makes changes to databases.

See [example.hcl](example.hcl) for an example configuration file.

## Where Query

As of right now, `=` is the only type of where query supported.

## Values

### `NULL` Values

To use a `NULL` value in either a where or update map, use the string `"NULL"` and it will be used as `NULL`.
