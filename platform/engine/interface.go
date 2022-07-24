package engine

// fixme: projeto a parte?
type IEngine interface {
	Init()
	SetSleepFrame(value int)
	GetSleepFrame() int
	SetFPS(value int)

	// GetFPS
	//
	// English:
	//
	// Returns the amount of current FPS used in calculations.
	//
	// Português:
	//
	// Retorna a quantidade de FPS atual usado nos cálculos.
	GetFPS() int
	CursorAddDrawFunction(runnerFunc func()) string
	CursorRemoveDrawFunction(id string)
	HighLatencyAddToFunctions(runnerFunc func()) (string, int)
	HighLatencyDeleteFromFunctions(UId string)
	HighLatencySetZIndex(UId string, index int) int
	HighLatencyGetZIndex(UId string) int
	HighLatencySetAsFistFunctionToRun(UId string) int
	HighLatencySetAsLastFunctionToRun(UId string) int
	SystemAddToFunctions(runnerFunc func()) (string, int)
	SystemDeleteFromFunctions(UId string)
	SystemSetZIndex(UId string, index int) int
	SystemGetZIndex(UId string) int
	SystemSetAsFistFunctionToRun(UId string) int
	SystemSetAsLastFunctionToRun(UId string) int
	AfterSystemAddToFunctions(runnerFunc func()) (string, int)
	AfterSystemDeleteFromFunctions(UId string)
	AfterSystemSetZIndex(UId string, index int) int
	AfterSystemGetZIndex(UId string) int
	AfterSystemSetAsFistFunctionToRun(UId string) int
	AfterSystemSetAsLastFunctionToRun(UId string) int
	MathAddToFunctions(runnerFunc func()) (string, int)
	MathDeleteFromFunctions(UId string)
	MathSetZIndex(UId string, index int) int
	MathGetZIndex(UId string) int
	MathSetAsFistFunctionToRun(UId string) int
	MathSetAsLastFunctionToRun(UId string) int
	DrawAddToFunctions(runnerFunc func()) (string, int)
	DrawDeleteFromFunctions(UId string)
	DrawSetZIndex(UId string, index int) int
	DrawGetZIndex(UId string) int
	DrawSetAsFistFunctionToRun(UId string) int
	DrawSetAsLastFunctionToRun(UId string) int
}
