package showgone

import "reflect"

const (
    HP = iota
    Atk
    Def
    SpA
    SpD
    Spe
)

const (
    Hardy Nature = iota
    Lonely
    Brave
    Adamant
    Naughty
    Bold
    Docile
    Relaxed
    Impish
    Lax
    Timid
    Hasty
    Serious
    Jolly
    Naive
    Modest
    Mild
    Quiet
    Bashful
    Rash
    Calm
    Gentle
    Sassy
    Careful
    Quirky
)

var statIdx = map[string]int{
    "HP": HP,
    "Atk": Atk,
    "Def": Def,
    "SpA": SpA,
    "SpD": SpD,
    "Spe": Spe,
}

var statStr = [...]string{
    "HP",
    "Atk",
    "Def",
    "SpA",
    "SpD",
    "Spe",
}

var convAttrib = map[string]func(string) reflect.Value{
    "Ability": convId,
    "Level": convNum,
    "Shiny": convShiny,
    "Happiness": convNum,
}

var natIdx = map[string]Nature{
    "Hardy": Hardy,
    "Lonely": Lonely,
    "Brave": Brave,
    "Adamant": Adamant,
    "Naughty": Naughty,
    "Bold": Bold,
    "Docile": Docile,
    "Relaxed": Relaxed,
    "Impish": Impish,
    "Lax": Lax,
    "Timid": Timid,
    "Hasty": Hasty,
    "Serious": Serious,
    "Jolly": Jolly,
    "Naive": Naive,
    "Modest": Modest,
    "Mild": Mild,
    "Quiet": Quiet,
    "Bashful": Bashful,
    "Rash": Rash,
    "Calm": Calm,
    "Gentle": Gentle,
    "Sassy": Sassy,
    "Careful": Careful,
    "Quirky": Quirky,
}

var natStr = [...]string{
    "Hardy",
    "Lonely",
    "Brave",
    "Adamant",
    "Naughty",
    "Bold",
    "Docile",
    "Relaxed",
    "Impish",
    "Lax",
    "Timid",
    "Hasty",
    "Serious",
    "Jolly",
    "Naive",
    "Modest",
    "Mild",
    "Quiet",
    "Bashful",
    "Rash",
    "Calm",
    "Gentle",
    "Sassy",
    "Careful",
    "Quirky",
}
