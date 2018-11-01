package showgone

type (
    Nature int
    Gender byte
    String string
)

type (
    Move    = String
    Species = String
    Item    = String
    Ability = String
)

type Pokemon struct {
    Species   Species
    Gender    Gender
    Item      Item
    Ability   Ability
    Level     uint8
    Shiny     bool
    Happiness uint8
    EVs       [6]uint8
    IVs       [6]uint8
    Nature    Nature
    Moves     [4]Move
}
