package brainmapping

type VolatileStatus string
type PermStatus string
type PokeStatus struct {
	VolatileStatus *[]VolatileStatus
	PermStatus     *PermStatus
}
