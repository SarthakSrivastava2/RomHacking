package baseStat

func (bs *BaseStat) ValidateStatStage(statName StatName, delta StatBoostStage) bool {
	const minStage = -6.0
	const maxStage = 6.0

	var current StatBoostStage

	switch statName {
	case StatNameHP:
		current = bs.HP.boost
	case StatNameAttack:
		current = bs.Attack.boost
	case StatNameDefense:
		current = bs.Defense.boost
	case StatNameSpecialAttack:
		current = bs.SpecialAttack.boost
	case StatNameSpecialDefense:
		current = bs.SpecialDefense.boost
	case StatNameSpeed:
		current = bs.Speed.boost
	default:
		return false
	}

	newStage := current + delta

	return newStage >= minStage && newStage <= maxStage
}

func (bs *BaseStat) IncrStatByStage(statName StatName, stage StatBoostStage) {
	if !bs.ValidateStatStage(statName, stage) {
		return
	}
	switch statName {
	case StatNameHP:
		incrStat(&bs.HP.Val, stage)
		updStatBoostInData(&bs.HP.boost, stage)
	case StatNameAttack:
		incrStat(&bs.Attack.Val, stage)
		updStatBoostInData(&bs.Attack.boost, stage)
	case StatNameDefense:
		incrStat(&bs.Defense.Val, stage)
		updStatBoostInData(&bs.Defense.boost, stage)
	case StatNameSpecialAttack:
		incrStat(&bs.SpecialAttack.Val, stage)
		updStatBoostInData(&bs.SpecialAttack.boost, stage)
	case StatNameSpecialDefense:
		incrStat(&bs.SpecialDefense.Val, stage)
		updStatBoostInData(&bs.SpecialDefense.boost, stage)
	case StatNameSpeed:
		incrStat(&bs.Speed.Val, stage)
		updStatBoostInData(&bs.Speed.boost, stage)
	}
}

func incrStat(bsStatVal *StatValue, stage StatBoostStage) {
	var statVal = *bsStatVal
	switch stage {
	case StageMinusSix:
		statVal *= statVal / 4
	case StageMinusFive:
		statVal *= statVal * 2 / 7
	case StageMinusFour:
		statVal *= statVal / 3
	case StageMinusThree:
		statVal *= statVal * 2 / 5
	case StageMinusTwo:
		statVal *= statVal / 2
	case StageMinusOne:
		statVal = statVal * 3 / 4
	case StageZero:
		statVal *= statVal
	case StageOne:
		statVal *= statVal * 3 / 2
	case StageTwo:
		statVal *= statVal * 2
	case StageThree:
		statVal *= statVal * 5 / 2
	case StageFour:
		statVal *= statVal * 3
	case StageFive:
		statVal *= statVal * 7 / 2
	case StageSix:
		statVal *= statVal * 4
	}
	*bsStatVal = statVal
}

func updStatBoostInData(currStage *StatBoostStage, updStage StatBoostStage) {
	*currStage = updStage
}
