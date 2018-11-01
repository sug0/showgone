package showgone

import "io"

func (p *Pokemon) WriteTo(w io.Writer) (int64, error) {
    var wrote int64

    n, err := p.fieldSpecies(w)
    wrote += int64(n)
    if err != nil {
        return wrote, err
    }

    n, err = p.fieldAbility(w)
    wrote += int64(n)
    if err != nil {
        return wrote, err
    }

    n, err = p.fieldLevel(w)
    wrote += int64(n)
    if err != nil {
        return wrote, err
    }

    n, err = p.fieldShiny(w)
    wrote += int64(n)
    if err != nil {
        return wrote, err
    }

    n, err = p.fieldHappiness(w)
    wrote += int64(n)
    if err != nil {
        return wrote, err
    }

    n, err = p.fieldVs(w, 'I')
    wrote += int64(n)
    if err != nil {
        return wrote, err
    }

    n, err = p.fieldVs(w, 'E')
    wrote += int64(n)
    if err != nil {
        return wrote, err
    }

    n, err = p.fieldNature(w)
    wrote += int64(n)
    if err != nil {
        return wrote, err
    }

    n, err = p.fieldMoves(w)
    wrote += int64(n)
    if err != nil {
        return wrote, err
    }

    return wrote, nil
}
