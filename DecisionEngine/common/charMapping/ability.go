package charmapping

import models "github.com/SarthakSrivastava2/RomHacking/DecisionEngine/models"

var InflictAbilityEffect map[models.Ability]func(*Pokemon, *Pokemon)
