package brainmapping

import (
	pokemonTypes "GoFolder/RomHacking/DecisionEngine/common/typeMapping"
)

type Pokemon struct {
	Type            pokemonTypes.PokemonType
	Exp             Experience
	Level           Level
	CurrHP          Stat
	BaseStat        *BaseStat
	CurrStat        *BaseStat
	Status          *PokeStatus
	Ability         *Ability
	IndiVal, EffVal *BaseStat
	HeldItem        *HeldItem
	Moveset         *Moveset
}
