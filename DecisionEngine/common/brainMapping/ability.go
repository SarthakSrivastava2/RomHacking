package brainmapping

type Ability string

type AbilityEffect map[Ability]func(*Pokemon, *Pokemon)
