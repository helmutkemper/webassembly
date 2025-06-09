package block

// max Returns the maximum value
func (e *Block) max(a, b int) (max int) {
	if a > b {
		return a
	}
	return b
}
