package showgone

import (
    "fmt"
    "bufio"
    "strings"
    "unicode"
    "reflect"
)

func (p *Pokemon) FillDefaults() {
    p.Level = 100
    p.Happiness = 255

    for i := 0; i < 6; i++ {
        p.IVs[i] = 31
    }
}

func Parse(r *bufio.Reader) (*Pokemon, error) {
    var line string
    var err error
    var idx int

    var poke Pokemon
    poke.FillDefaults()

    pokeR := reflect.ValueOf(&poke).Elem()

    // skip spaces
    for {
        var rr rune
        rr, _, err = r.ReadRune()
        if err != nil {
            return nil, err
        }
        if !unicode.IsSpace(rr) {
            err = r.UnreadRune()
            if err != nil {
                return nil, err
            }
            break
        }
    }

    // get poke species
    line, err = r.ReadString('\n')
    if err != nil {
        return nil, err
    }

    idx = strings.Index(line, " @ ")
    if idx != -1 {
        poke.Species = String(line[:idx])
        poke.Item = String(strings.TrimRightFunc(line[idx+3:], unicode.IsSpace))
    } else {
        poke.Species = String(strings.TrimRightFunc(line, unicode.IsSpace))
    }

    // check if name includes sex
    sz := len(poke.Species)

    if sz > 3                 &&
    poke.Species[sz-4] == ' ' &&
    poke.Species[sz-3] == '(' &&
    poke.Species[sz-1] == ')' {
        poke.Gender = Gender(poke.Species[sz-2])
        poke.Species = poke.Species[:sz-4]
    }

    // rest of the fields...
    for move := 0;; {
        // get new line
        line, err = r.ReadString('\n')
        if err != nil || line[0] == '\r' || line[0] == '\n' {
            break
        }

        // check for moves
        if line[0] == '-' {
            if move > 3 {
                return nil, ErrInvalidFmt
            }
            poke.Moves[move] = String(strings.TrimRightFunc(line[2:], unicode.IsSpace))
            move++
            continue
        }

        // none apparently, determine what else
        // to do with the line...
        idx = strings.Index(line, ": ")
        if idx == -1 {
            // try to find Nature
            idx = strings.Index(line, " Nature")
            if idx == -1 {
                return nil, ErrInvalidFmt
            }

            // save found nature
            poke.Nature = natIdx[line[:idx]]
            continue
        }

        key := line[:idx]

        // either IVs or EVs
        if key[1] == 'V' {
            var vs []uint8

            if key[0] == 'E' {
                vs = poke.EVs[:]
            } else {
                vs = poke.IVs[:]
            }

            var val uint8
            var which string

            evtok := line[idx+2:]
            idx = -3

            for idx != -1 {
                evtok = evtok[idx+3:]
                fmt.Sscan(evtok, &val, &which)
                vs[statIdx[which]] = val
                idx = strings.Index(evtok, " / ")
            }
        } else {
            field := pokeR.FieldByName(key)
            value := strings.TrimRightFunc(line[idx+2:], unicode.IsSpace)
            conv := convAttrib[key]

            if conv == nil {
                return nil, ErrInvalidFmt
            }
            field.Set(conv(value))
        }
    }

    return &poke, nil
}
