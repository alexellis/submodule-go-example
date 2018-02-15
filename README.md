# submodule-go-example

This example demonstrates GOPATH for functions.

## Sub-modules

The function exists within a `handler` folder in the root of the GOPATH. Within that folder is the function:

```
handler/function
```

Therefore any sub-modules within the function folder must take it as a prefix

I.e. a sub-module named `util` would be imported like this:

```
import handler/function/util
```

## Vendored dependencies

Vendored dependencies work exactly the same as in any other project.

Use a tool like `dep` to initialize and ensure dependencies.

