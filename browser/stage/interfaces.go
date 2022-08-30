package stage

type Functions interface {
	AddHighLatencyFunctions(runnerFunc func()) (UId string, total int)
	DeleteHighLatencyFunctions(UId string)
	AddDrawFunctions(runnerFunc func()) (UId string, total int)
	DeleteDrawFunctions(UId string)
	AddMathFunctions(runnerFunc func()) (UId string, total int)
	DeleteMathFunctions(UId string)
}
