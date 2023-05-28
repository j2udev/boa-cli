# Boa Generator

Boa generator is a simple tool for generating new CLI projects based on the
[Boa](https://github.com/j2udev/boa) pkg. It is loosely inspired by the
[Cobra-CLI](https://github.com/spf13/cobra-cli/tree/main) and heavily inspired
by [Bashly](https://github.com/DannyBen/bashly).

## Disclaimer

This project should be considered unstable until it is officially released. Use
at your own risk.

## Building from source (for now)

Until this project is officially released, just build from source:

```txt
go build -o /somewhere/on/your/path/boa
```

Then you can use it anywhere.

## Initialize

Much like the `bashly init` command, `boa init` is used to create a new boa
configuration file in the current directory. This config file, `boa.yml` is in
yaml by default, but can be changed to json or toml if desired. It has comments
that explain the structure of an example CLI. Modify it as you see fit.

## Generate

After defining your CLI in the `boa.yml`, just run `boa generate -i` and,
assuming your boa.yml is valid, you should have a working CLI! Just add your
business logic to the appropriate places under the `internal` package.

## "Day 2 Operations"

Unlike the bashly CLI, the boa generator is not meant to be continuously used to
generate overtop of your existing project (it _can_, but it doesn't come close
to extending the entirety of the boa package). The boa package should be easy
enough to continue building your project after the initial project generation.
The boa generator is simply meant to help with getting started and quickly
prototyping a new CLI project. It's expected and even encouraged to throw out
your `boa.yml/json/toml` once you've outgrown its capabilities.
