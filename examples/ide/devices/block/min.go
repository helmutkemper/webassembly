package block

// min Returns the minimum value
func (e *Block) min(a, b int) (min int) {
	if a < b {
		return a
	}
	return b
}
