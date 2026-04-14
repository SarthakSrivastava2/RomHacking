package typemapping

func TypeEffectiveChecker(offenseTypeName, defenseTypeName Type) float32 {
	return typeEffectivenessMap[offenseTypeName][defenseTypeName]
}

func FullTypeEffectiveChecker(offenseMoveType Type, defender PokemonType) float32 {
	var result float32 = 1.0
	if defender.primaryType != nil {
		result *= TypeEffectiveChecker(offenseMoveType, *defender.primaryType)
	}
	if defender.secondaryType != nil {
		result *= TypeEffectiveChecker(offenseMoveType, *defender.secondaryType)
	}
	return result
}
