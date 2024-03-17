# cheese-lang-&#129472;

cheese-lang is a statically typed, interpreted language with a runtime
written in go.

## Language specification

### Types

there are only 4 types in cheese-lang, that are named after
famous Italian cheeses:

- `Parmesan` (equivalent to int)
- `Gorgonzola` (equivalent to float)
- `Mozzarella` (equivalent to string)
- `Milk` (equivalent to bool, can have ywo values: `fresh` and `spoiled` that are equivalent to `true` and `false`)

### Type operations

- `Parmesan`:
  - `+` (addition)
  - `-` (subtraction)
  - `*` (multiplication)
  - `/` (division)
  - `%` (modulo)
  - `==` (equality)
  - `!=` (inequality)
  - `>` (greater than)
  - `<` (less than)
  - `>=` (greater than or equal to)
  - `<=` (less than or equal to)

- `Gorgonzola`:
    - `+` (addition)
    - `-` (subtraction)
    - `*` (multiplication)
    - `/` (division)
    - `==` (equality)
    - `!=` (inequality)
    - `>` (greater than)
    - `<` (less than)
    - `>=` (greater than or equal to)
    - `<=` (less than or equal to)

- `Mozzarella`:
    - `+` (concatenation)
    - `==` (equality)
    - `!=` (inequality)

- `Ricotta`:
    - `==` (equality)
    - `!=` (inequality)

### standard library

