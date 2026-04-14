package baseStat

type BaseStat struct {
	HP             Stat
	Attack         Stat
	Defense        Stat
	SpecialAttack  Stat
	SpecialDefense Stat
	Speed          Stat
}

type StatBoostStage float32

const (
	StageMinusSix   StatBoostStage = -6
	StageMinusFive  StatBoostStage = -5
	StageMinusFour  StatBoostStage = -4
	StageMinusThree StatBoostStage = -3
	StageMinusTwo   StatBoostStage = -2
	StageMinusOne   StatBoostStage = -1
	StageZero       StatBoostStage = 0
	StageOne        StatBoostStage = 1
	StageTwo        StatBoostStage = 2
	StageThree      StatBoostStage = 3
	StageFour       StatBoostStage = 4
	StageFive       StatBoostStage = 5
	StageSix        StatBoostStage = 6
)

type Stat struct {
	boost StatBoostStage
	Val   StatValue
}

type StatValue uint16

const (
	StatNameHP             StatName = "HP"
	StatNameAttack         StatName = "Attack"
	StatNameDefense        StatName = "Defense"
	StatNameSpecialAttack  StatName = "SpecialAttack"
	StatNameSpecialDefense StatName = "SpecialDefense"
	StatNameSpeed          StatName = "Speed"
)

type StatName string
