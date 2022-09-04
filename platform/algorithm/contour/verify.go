package contour

func (e *Contour) Verify() {
	for {
		pass := false

		for y := 0; y != len(*e.matrix)-1; y += 1 {
			for x := 0; x != len((*e.matrix)[y])-1; x += 1 {
				if e.verified[y][x] == false && e.verifyFunction(e.matrix, x, y) == true {
					pass = true
					e.verifyPoint(x, y)
					return
				}
			}
		}

		if pass == false {
			break
		}
	}
}
