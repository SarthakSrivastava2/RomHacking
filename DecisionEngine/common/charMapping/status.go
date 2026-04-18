package charmapping

import "github.com/SarthakSrivastava2/RomHacking/DecisionEngine/models"

var InflictStatusEffect map[models.PokeStatus]func(*Pokemon, *Pokemon)
