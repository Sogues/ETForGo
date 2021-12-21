package battle

type (
	ChaProperty struct {
		hp int

		attack int

		moveSpeed int

		actionSpeed int

		ammo int

		bodyRadius float64

		hitRadius float64

		moveType MoveType
	}

	SChaResource struct {
		hp int

		ammo int

		stamina int
	}
	ChaResource *SChaResource
)