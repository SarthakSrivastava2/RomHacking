package models

import (
	pokemonTypes "github.com/SarthakSrivastava2/RomHacking/DecisionEngine/common/typeMapping"
)

type MoveName string
type MoveAttribute string

const (
	MoveAttributePhysical MoveAttribute = "PHYSICAL"
	MoveAttributeSpecial  MoveAttribute = "SPECIAL"
	MoveAttributeStatus   MoveAttribute = "STATUS"
)

type Move struct {
	Name      MoveName
	Attribute MoveAttribute
	Accuracy  StatValue
	Damage    StatValue
	Type      pokemonTypes.Type
}
type Moveset [4]Move
