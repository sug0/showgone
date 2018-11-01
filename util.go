package showgone

import (
    "io"
    "fmt"
    "reflect"
)

func convId(str string) reflect.Value {
    return reflect.ValueOf(String(str))
}

func convNum(num string) reflect.Value {
    var val uint8
    fmt.Sscan(num, &val)
    return reflect.ValueOf(val)
}

func convShiny(shiny string) reflect.Value {
    if shiny == "Yes" {
        return reflect.ValueOf(true)
    }
    return reflect.ValueOf(false)
}

func (p *Pokemon) fieldSpecies(w io.Writer) (int, error) {
    if p.Item != "" {
        if p.Gender != 0 {
            return fmt.Fprintf(w, "%s (%c) @ %s\n", p.Species, p.Gender, p.Item)
        }
        return fmt.Fprintf(w, "%s @ %s\n", p.Species, p.Item)
    }
    if p.Gender != 0 {
        return fmt.Fprintf(w, "%s (%c)\n", p.Species, p.Gender)
    }
    return fmt.Fprintf(w, "%s\n", p.Species)
}

func (p *Pokemon) fieldAbility(w io.Writer) (int, error) {
    if p.Ability != "" {
        return fmt.Fprintf(w, "Ability: %s\n", p.Ability)
    }
    return 0, nil
}

func (p *Pokemon) fieldLevel(w io.Writer) (int, error) {
    if p.Level != 100 {
        return fmt.Fprintf(w, "Level: %d\n", p.Level)
    }
    return 0, nil
}

func (p *Pokemon) fieldShiny(w io.Writer) (int, error) {
    if p.Shiny {
        return fmt.Fprintln(w, "Shiny: Yes")
    }
    return 0, nil
}

func (p *Pokemon) fieldHappiness(w io.Writer) (int, error) {
    if p.Happiness != 255 {
        return fmt.Fprintf(w, "Happiness: %d\n", p.Happiness)
    }
    return 0, nil
}

func (p *Pokemon) fieldVs(w io.Writer, v byte) (int, error) {
    var vs []uint8
    var def uint8
    var n, wrote int
    var err error
    var one bool

    if v == 'E' {
        vs = p.EVs[:]
        def = 0
    } else {
        vs = p.IVs[:]
        def = 31
    }

    for i := 0; i < 6; i++ {
        if vs[i] != def {
            if one {
                n, err = fmt.Fprintf(w, "/ %d %s ", vs[i], statStr[i])
            } else {
                n, err = fmt.Fprintf(w, "%cVs: %d %s ", v, vs[i], statStr[i])
                one = true
            }
            wrote += n

            if err != nil {
                return wrote, err
            }
        }
    }

    if one {
        n, err = w.Write([]byte("\n"))
        wrote += n

        if err != nil {
            return wrote, err
        }
    }
    return wrote, nil
}

func (p *Pokemon) fieldNature(w io.Writer) (int, error) {
    if p.Nature != 0 {
        return fmt.Fprintf(w, "%s Nature\n", p.Nature)
    }
    return 0, nil
}

func (p *Pokemon) fieldMoves(w io.Writer) (int, error) {
    var wrote int

    for i := 0; i < 4; i++ {
        if p.Moves[i] != "" {
            n, err := fmt.Fprintf(w, "- %s\n", p.Moves[i])
            wrote += n

            if err != nil {
                return wrote, err
            }
        }
    }

    return wrote, nil
}

func (n Nature) String() string {
    return natStr[n]
}

func (g Gender) String() string {
    switch g {
    default:
        return "Random"
    case 'M':
        return "Male"
    case 'F':
        return "Female"
    case 'G':
        return "Genderless"
    }
}

func (s String) String() string {
    if s == "" {
        return "None"
    }
    return string(s)
}
