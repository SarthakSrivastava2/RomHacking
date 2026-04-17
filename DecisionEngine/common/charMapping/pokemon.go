package brainmapping

import (
	pokemonAbility "GoFolder/RomHacking/DecisionEngine/common/ability"
	pokemonStats "GoFolder/RomHacking/DecisionEngine/common/baseStat"
	pokemonItems "GoFolder/RomHacking/DecisionEngine/common/item"
	pokemonLvlAndXp "GoFolder/RomHacking/DecisionEngine/common/levelAndExp"
	pokemonMoveset "GoFolder/RomHacking/DecisionEngine/common/moveset"
	pokemonStatus "GoFolder/RomHacking/DecisionEngine/common/status"
	pokemonTypes "GoFolder/RomHacking/DecisionEngine/common/typeMapping"
)

type Pokemon struct {
	Type            pokemonTypes.PokemonType
	Exp             pokemonLvlAndXp.Experience
	Level           pokemonLvlAndXp.Level
	CurrHP          pokemonStats.Stat
	BaseStat        *pokemonStats.BaseStat
	CurrStat        *pokemonStats.BaseStat
	Status          *pokemonStatus.PokeStatus
	Ability         *pokemonAbility.Ability
	IndiVal, EffVal *pokemonStats.BaseStat
	HeldItem        *pokemonItems.HeldItem
	Moveset         *pokemonMoveset.Moveset
}
