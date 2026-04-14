package typemapping

type PokemonType struct {
	primaryType   *Type
	secondaryType *Type
}

type Type string

const (
	Water    Type = "Water"
	Electric Type = "Electric"
	Fire     Type = "Fire"
	Grass    Type = "Grass"
	Rock     Type = "Rock"
	Bug      Type = "Bug"
	Ghost    Type = "Ghost"
	Psychic  Type = "Psychic"
	Fighting Type = "Fighting"
	Flying   Type = "Flying"
	Normal   Type = "Normal"
	Poison   Type = "Poison"
	Dark     Type = "Dark"
	Steel    Type = "Steel"
	Ice      Type = "Ice"
	Dragon   Type = "Dragon"
	Fairy    Type = "Fairy"
	Ground   Type = "Ground"
)

var typeArr []Type = []Type{
	Water,
	Electric,
	Fire,
	Grass,
	Rock,
	Bug,
	Ghost,
	Psychic,
	Fighting,
	Flying,
	Normal,
	Poison,
	Dark,
	Steel,
	Ice,
	Dragon,
	Fairy,
	Ground,
}
