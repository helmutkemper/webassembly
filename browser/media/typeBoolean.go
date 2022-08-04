package media

// BOOLEAN
//
// English:
//
// BOOLEAN helps distinguish between true (1), not set (0) and false (-1).
//
// Português:
//
// BOOLEAN ajuda a distinguir entre true (1), not set (0) e false (-1).
type BOOLEAN int

func (e BOOLEAN) IsSet() (set bool) {
	return e != 0
}

func (e BOOLEAN) Bool() (value bool) {
	return e == 1
}
