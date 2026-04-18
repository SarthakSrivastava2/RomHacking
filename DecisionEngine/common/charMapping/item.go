package charmapping

import models "github.com/SarthakSrivastava2/RomHacking/DecisionEngine/models"

var InflictItemEffect map[models.HeldItem]func(*Pokemon, *Pokemon)
