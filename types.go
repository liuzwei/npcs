package npcs

type Power struct {
	Attack int
	Defense int
}

type Location struct {
	x float64
	y float64
	z float64
}

type NonPlayerCharacter struct {
	Name string
	Speed int
	HP int
	Power Power
	Loc Location
}
