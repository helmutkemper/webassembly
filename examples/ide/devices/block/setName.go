package block

// SetName Defines a unique name for the device [compulsory]
func (e *Block) SetName(name string) (err error) {
	e.name, err = e.SequentialId.GetId(name)
	return
}
