package contour

func (e *Contour) pAdjust(p int) int {
	p = p % (pMax + 1)
	if p < 0 {
		p *= -1
		p -= 1
		p = pMax - p
	}
	return p
}
