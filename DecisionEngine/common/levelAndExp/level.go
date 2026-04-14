package levelandexp

type Level uint8

type Experience uint16

func (pokeLvl Level) SetLevel(level Level) {
	pokeLvl = level
}

func (pokeLvl Level) IncrLevel() {
	pokeLvl += 1
}
