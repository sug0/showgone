# showgone

    Tapu Lele @ Fightinium Z
    Ability: Psychic Surge
    EVs: 252 SpA / 4 SpD / 252 Spe
    Timid Nature
    - Psychic
    - Moonblast
    - Taunt
    - Focus Blast

A parser for [Smogon](https://www.smogon.com/)'s pokémon format,
written in golang.

# Example

```go
package main

import (
    "os"
    "fmt"
    "bufio"

    "github.com/sugoiuguu/showgone"
)

func main() {
    r := bufio.NewReader(os.Stdin)
    for {
        poke, err := showgone.Parse(r)
        if err != nil {
            return
        }
        fmt.Println(poke.Species, "lvl", poke.Level, "@", poke.Item)
    }
}
```

More documentation available at [godoc](https://godoc.org/github.com/sugoiuguu/showgone).
