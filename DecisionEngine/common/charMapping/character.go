package charmapping

import (
	pokemonTypes "github.com/SarthakSrivastava2/RomHacking/DecisionEngine/common/typeMapping"
	models "github.com/SarthakSrivastava2/RomHacking/DecisionEngine/models"
)

type Pokemon struct {
	Type            pokemonTypes.PokemonType
	Exp             models.Experience
	Level           models.Level
	CurrHP          models.Stat
	BaseStat        *models.BaseStat
	CurrStat        *models.BaseStat
	Status          *models.PokeStatus
	Ability         *models.Ability
	IndiVal, EffVal *models.BaseStat
	HeldItem        *models.HeldItem
	Moveset         *models.Moveset
	AbstractStats   *models.AbstractStats
}
