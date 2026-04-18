package charmapping

import models "github.com/SarthakSrivastava2/RomHacking/DecisionEngine/models"

func (pPokeVar *Pokemon) ValidateStatStage(statName models.StatName, delta models.StatBoostStage) bool {
	const minStage = -6.0
	const maxStage = 6.0

	var current models.StatBoostStage
	bs := pPokeVar.CurrStat

	switch statName {
	case models.StatNameHP:
		current = bs.HP.Boost
	case models.StatNameAttack:
		current = bs.Attack.Boost
	case models.StatNameDefense:
		current = bs.Defense.Boost
	case models.StatNameSpecialAttack:
		current = bs.SpecialAttack.Boost
	case models.StatNameSpecialDefense:
		current = bs.SpecialDefense.Boost
	case models.StatNameSpeed:
		current = bs.Speed.Boost
	default:
		return false
	}

	newStage := current + delta

	return newStage >= minStage && newStage <= maxStage
}

func (pPokeVar *Pokemon) IncrStatByStage(statName models.StatName, stage models.StatBoostStage) {
	if !pPokeVar.ValidateStatStage(statName, stage) {
		return
	}
	bs := pPokeVar.CurrStat
	switch statName {
	case models.StatNameHP:
		incrStat(&bs.HP.Val, stage)
		updStatBoostInData(&bs.HP.Boost, stage)
	case models.StatNameAttack:
		incrStat(&bs.Attack.Val, stage)
		updStatBoostInData(&bs.Attack.Boost, stage)
	case models.StatNameDefense:
		incrStat(&bs.Defense.Val, stage)
		updStatBoostInData(&bs.Defense.Boost, stage)
	case models.StatNameSpecialAttack:
		incrStat(&bs.SpecialAttack.Val, stage)
		updStatBoostInData(&bs.SpecialAttack.Boost, stage)
	case models.StatNameSpecialDefense:
		incrStat(&bs.SpecialDefense.Val, stage)
		updStatBoostInData(&bs.SpecialDefense.Boost, stage)
	case models.StatNameSpeed:
		incrStat(&bs.Speed.Val, stage)
		updStatBoostInData(&bs.Speed.Boost, stage)
	}
}

func incrStat(bsStatVal *models.StatValue, stage models.StatBoostStage) {
	var statVal = *bsStatVal
	switch stage {
	case models.StageMinusSix:
		statVal *= statVal / 4
	case models.StageMinusFive:
		statVal *= statVal * 2 / 7
	case models.StageMinusFour:
		statVal *= statVal / 3
	case models.StageMinusThree:
		statVal *= statVal * 2 / 5
	case models.StageMinusTwo:
		statVal *= statVal / 2
	case models.StageMinusOne:
		statVal = statVal * 3 / 4
	case models.StageZero:
		statVal *= statVal
	case models.StageOne:
		statVal *= statVal * 3 / 2
	case models.StageTwo:
		statVal *= statVal * 2
	case models.StageThree:
		statVal *= statVal * 5 / 2
	case models.StageFour:
		statVal *= statVal * 3
	case models.StageFive:
		statVal *= statVal * 7 / 2
	case models.StageSix:
		statVal *= statVal * 4
	}
	*bsStatVal = statVal
}

func updStatBoostInData(currStage *models.StatBoostStage, updStage models.StatBoostStage) {
	*currStage = updStage
}
