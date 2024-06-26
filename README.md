# &#129472; &#129472; &#129472; cheese-lang (Linguaggio Formaggio) &#129472; &#129472; &#129472; 

cheese-lang (or Linguaggio Formaggio in Italian)is a statically typed, interpreted language with a runtime
written in go.

## Language specification

### Types

there are only 5 types in cheese-lang, that are named after
famous Italian cheeses:

- `Parmesan` (equivalent to int)
- `Gorgonzola` (equivalent to float)
- `Mozzarella` (equivalent to string)
- `Milk` (equivalent to bool, can have ywo values: `fresh` and `spoiled` that are equivalent to `true` and `false`)
- `Ricotta` (equivalent to void)

### Variables
variables names must start with a letter and can contain letters, numbers and underscores.
variables names are also case sensitive, and are invalid if they are a keyword of the language.

the variables are declared with the following syntax:
```
<Type> <name> = <value>;
```

example:
```
Parmesan a = 5;
Gorgonzola b = 5.5;
Mozzarella c = "hello";
Milk d = fresh;
```

### Comments
 - `//` (single line comment)

### Control flow

 - `taste` (if)
    ```
    taste <condition>{
        <block>
    }
    ```
 - `recipe` (function)
    ```
    // declaration
    recipe <name>(<type param 1> <name param 1>, <type param 2> <name param 2>, ...) -> <return type 1>, <return type 2>, ... {
        <block>
        prepare <return value>;
    }
    
    // call
    <name>(<param 1>, <param 2>, ...);
    ``` 
 - `curdle` (loop)
    ```
    curdle {
        <block>
        taste <condition> {
            drain // break;
        }
    }
    ```
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
    - `!` (negation)
    - `&&` (and)
    - `||` (or)
    - `^` (exor)
### standard library

#### type conversion
 - `p_to_g(Parmesan p) -> Gorgonzola` (converts a `Parmesan` to a `Gorgonzola`)
 - `p_to_m(Parmesan p) -> Mozzarella` (converts a `Parmesan` to a `Mozzarella`)
 - `g_to_p(Gorgonzola g) -> Parmesan` (converts a `Gorgonzola` to a `Parmesan`)
 - `g_to_m(Gorgonzola g) -> Mozzarella` (converts a `Gorgonzola` to a `Mozzarella`)
 - `m_to_p(Mozzarella m) -> Parmesan, Milk` (converts a `Mozzarella` to a `Parmesan`, and return a milk to represent if the conversion was successful)
 - `m_to_g(Mozzarella m) -> Gorgonzola, Milk` (converts a `Mozzarella` to a `Gorgonzola`, and return a milk to represent if the conversion was successful)

#### input/output
 - `eat() -> Mozzarella` (reads a line from the standard input)
 - `serve(Mozzarella m)` (prints a string to the standard output)

#### Mozzarella manipulation
 - `weight(Mozzarella m) -> Parmesan` (returns the length of a string)
 - `slice(Mozzarella m, Parmesan start, Parmesan end) -> Mozzarella` (returns a substring of a string, from `start` to `end` (inclusive))
