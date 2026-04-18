package charmapping

import "github.com/SarthakSrivastava2/RomHacking/DecisionEngine/models"

func (pPokeVar *Pokemon) SetLevel(level models.Level) {
	(*pPokeVar).Level = level
}

func (pPokeVar *Pokemon) IncrLevel() {
	(*pPokeVar).Level += 1
}
