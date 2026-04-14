package typemapping

import "fmt"

var typeEffectivenessMap map[Type]map[Type]float32

func InitTypeMapping() {

	typeEffectivenessMap = map[Type]map[Type]float32{
		Water: {
			Water:  0.5,
			Grass:  0.5,
			Dragon: 0.5,

			Rock:   2.0,
			Fire:   2.0,
			Ground: 2.0,
		},
		Electric: {
			Grass:    0.5,
			Electric: 0.5,
			Dragon:   0.5,

			Ground: 0.0,

			Flying: 2.0,
			Water:  2.0,
		},
		Fire: {
			Rock:   0.5,
			Fire:   0.5,
			Water:  0.5,
			Dragon: 0.5,

			Steel: 2.0,
			Grass: 2.0,
			Bug:   2.0,
			Ice:   2.0,
		},
		Grass: {
			Flying: 0.5,
			Poison: 0.5,
			Grass:  0.5,
			Dragon: 0.5,
			Steel:  0.5,
			Fire:   0.5,
			Bug:    0.5,

			Water:  2.0,
			Rock:   2.0,
			Ground: 2.0,
		},
		Rock: {
			Fighting: 0.5,
			Ground:   0.5,
			Steel:    0.5,

			Flying: 2.0,
			Bug:    2.0,
			Fire:   2.0,
			Ice:    2.0,
		},
		Bug: {
			Flying:   0.5,
			Fighting: 0.5,
			Poison:   0.5,
			Ghost:    0.5,
			Steel:    0.5,
			Fire:     0.5,
			Fairy:    0.5,

			Grass:   2.0,
			Psychic: 2.0,
			Dark:    2.0,
		},
		Ghost: {
			Dark: 0.5,

			Normal: 0.0,

			Ghost:   2.0,
			Psychic: 2.0,
		},
		Psychic: {
			Steel:   0.5,
			Psychic: 0.5,

			Dark:     0.0,
			Fighting: 2.0,
			Poison:   2.0,
		},
		Fighting: {
			Flying:  0.5,
			Poison:  0.5,
			Bug:     0.5,
			Fairy:   0.5,
			Psychic: 0.5,

			Ghost: 0.0,

			Rock:   2.0,
			Normal: 2.0,
			Dark:   2.0,
			Steel:  2.0,
			Ice:    2.0,
		},
		Flying: {
			Rock:     0.5,
			Steel:    0.5,
			Electric: 0.5,

			Fighting: 2.0,
			Bug:      2.0,
			Grass:    2.0,
		},
		Normal: {
			Rock:  0.5,
			Steel: 0.5,

			Ghost: 0.0,
		},
		Poison: {
			Poison: 0.5,
			Rock:   0.5,
			Ground: 0.5,
			Ghost:  0.5,

			Steel: 0.0,

			Grass: 2.0,
			Fairy: 2.0,
		},
		Dark: {
			Fighting: 0.5,
			Dark:     0.5,
			Fairy:    0.5,

			Ghost:   2.0,
			Psychic: 2.0,
		},
		Steel: {
			Steel:    0.5,
			Fire:     0.5,
			Water:    0.5,
			Electric: 0.5,

			Rock:  2.0,
			Ice:   2.0,
			Fairy: 2.0,
		},
		Ice: {
			Steel: 0.5,
			Fire:  0.5,
			Water: 0.5,
			Ice:   0.5,

			Flying: 2.0,
			Ground: 2.0,
			Grass:  2.0,
			Dragon: 2.0,
		},
		Dragon: {
			Steel: 0.5,

			Fairy: 0.0,

			Dragon: 2.0,
		},
		Fairy: {
			Poison: 0.5,
			Steel:  0.5,
			Fire:   0.5,

			Fighting: 2.0,
			Dragon:   2.0,
			Dark:     2.0,
		},
		Ground: {
			Bug:   0.5,
			Grass: 0.5,

			Flying: 0.0,

			Poison:   2.0,
			Ground:   2.0,
			Steel:    2.0,
			Fire:     2.0,
			Electric: 2.0,
		},
	}
	setUnaffectedMoveInMap()
}

func setUnaffectedMoveInMap() {
	for offenseTypeName, typeSubMap := range typeEffectivenessMap {
		for _, defenseTypeName := range typeArr {
			if effectiveness, ok := typeSubMap[defenseTypeName]; !ok {
				effectiveness = 1.0
				typeEffectivenessMap[offenseTypeName][defenseTypeName] = effectiveness
			}
		}
	}
}

func printMapStr() {
	fmt.Printf("Type chart\n")

	for offenseTypeName, typeSubMap := range typeEffectivenessMap {
		for defenseTypeName, effectiveness := range typeSubMap {
			fmt.Printf("%s does %vx damage to %s", offenseTypeName, effectiveness, defenseTypeName)
		}
	}
}
