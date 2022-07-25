package engine

type IEngine interface {

	// Init
	//
	// English:
	//
	// Start the engine
	//
	// Português:
	//
	// Inicializa à engine
	Init()

	// SetSleepFrame
	//
	// English:
	//
	// Sleep Frame can delay processing to prevent crashes.
	//
	// This functionality is experimental and comes from C++, where the loop can crash the system.
	//
	// Português:
	//
	// Sleep Frame pode dá um tempo no processamento para impedir travamentos.
	//
	// Esta funcionalidade é experimental e vem do C++, onde o laço pode travar o sistema.
	SetSleepFrame(value int)

	// GetSleepFrame
	//
	// English:
	//
	// Sleep Frame can delay processing to prevent crashes.
	//
	// This functionality is experimental and comes from C++, where the loop can crash the system.
	//
	// Português:
	//
	// Sleep Frame pode dá um tempo no processamento para impedir travamentos.
	//
	// Esta funcionalidade é experimental e vem do C++, onde o laço pode travar o sistema.
	GetSleepFrame() int

	// SetFpsMax
	//
	// English:
	//
	// Português:
	//
	// Define a quantidade máxima de FPS.
	//
	//   Notas:
	//     * A quantidade máxima de FPS pode fazer o navegador travar.
	SetFpsMax(value int)

	// SetFpsMin
	//
	// English:
	//
	// Sets the minimum amount of FPS.
	//
	//   Notes:
	//     * The FPS value automatically drops when the system is very busy.
	//
	// Português:
	//
	// Define a quantidade mínima de FPS.
	//
	//   Notas:
	//     * O valor do FPS cai de forma automática quando o sistema está muito ocupado.
	SetFpsMin(value int)

	// SetFPS
	//
	// English:
	//
	// Sets the maximum amount of FPS.
	//
	//   Notes:
	//     * The FPS value automatically goes up when the system is not overloaded.
	//
	// Português:
	//
	// Define a quantidade máxima de FPS.
	//
	//   Notas:
	//     * O valor do FPS sobe de forma automática quando o sistema não está sobrecarregado.
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

	// CursorAddDrawFunction
	//
	// English:
	//
	// Allows you to recreate the function that draws the cursor.
	//
	// Português:
	//
	// Permite recriar a função que desenha o cursor.
	CursorAddDrawFunction(runnerFunc func())

	// CursorRemoveDrawFunction
	//
	// English:
	//
	// Removes the role responsible for recreating the cursor.
	//
	// Português:
	//
	// Remove a função responssável por recria o cursor.
	CursorRemoveDrawFunction()

	// HighLatencyAddToFunctions
	//
	// English:
	//
	// Adds a high latency function, a low execution priority function.
	//
	//   Input:
	//     runnerFunc: function to be performed.
	//
	//   Output:
	//     UId: used to identify the function when removing.
	//     total: total number of functions running.
	//
	//   Notes:
	//     * High latency functions are secondary functions designed to run at a lower FPS rate.
	//
	// Português:
	//
	// Adiciona uma função de alta latencia, uma função de baixa prioridade de execussão.
	//
	//   Entrada:
	//     runnerFunc: função a ser executada.
	//
	//   Saída:
	//     UId da função, usado para identificar a função na hora de remover.
	//     total: quantidade total de funções em execução.
	//
	//   Notas:
	//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
	HighLatencyAddToFunctions(runnerFunc func()) (string, int)

	// HighLatencyDeleteFromFunctions
	//
	// English:
	//
	// Removes a high latency function added by the HighLatencyAddToFunctions() function.
	//
	//   Input:
	//     UId: ID returned by the HighLatencyAddToFunctions() function.
	//
	//   Notes:
	//     * High latency functions are secondary functions designed to run at a lower FPS rate.
	//
	// Português:
	//
	// Remove uma função de alta latencia adicionada pela função HighLatencyAddToFunctions().
	//
	//   Entrada:
	//     UId: ID retornado pela função HighLatencyAddToFunctions().
	//
	//   Notas:
	//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
	HighLatencyDeleteFromFunctions(UId string)

	// HighLatencySetZIndex
	//
	// English:
	//
	// Allows you to change the order of execution of the function, in the execution list.
	//
	//   Input:
	//     UId: ID returned by the HighLatencyAddToFunctions() function.
	//     index: 0 for the first function in the list
	//
	//   Notes:
	//     * High latency functions are secondary functions designed to run at a lower FPS rate.
	//
	// Português:
	//
	// Permite trocar a ordem de execução da função, na lista de execução.
	//
	//   Entrada:
	//     UId: ID retornado pela função HighLatencyAddToFunctions().
	//     index: 0 para a primeira função da lista
	//
	//   Notas:
	//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
	HighLatencySetZIndex(UId string, index int) int

	// HighLatencyGetZIndex
	//
	// English:
	//
	// Returns the function execution index in the list, where 0 is the first function to be executed.
	//
	//   Input:
	//     UId: ID returned by the HighLatencyAddToFunctions() function.
	//
	//   Output:
	//     index: Function execution order.
	//
	//   Notes:
	//     * High latency functions are secondary functions designed to run at a lower FPS rate.
	//
	// Português:
	//
	// Retorna o índice de execução da função na lista, onde 0 é a primera função a ser executada.
	//
	//   Entrada:
	//     UId: ID retornado pela função HighLatencyAddToFunctions().
	//
	//   Saída:
	//     index: Ordem de execução da função.
	//
	//   Notas:
	//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
	HighLatencyGetZIndex(UId string) int

	// HighLatencySetAsFistFunctionToRun
	//
	// English:
	//
	// Makes the function first on the execution list for high latency functions.
	//
	//   Input:
	//     UId: ID returned by the HighLatencyAddToFunctions() function.
	//
	//   Output:
	//     index: Function execution order.
	//
	//   Notes:
	//     * High latency functions are secondary functions designed to run at a lower FPS rate.
	//
	// Português:
	//
	// Faz a função ser a primeira da lista de execuções para funções de alta latencia.
	//
	//   Entrada:
	//     UId: ID retornado pela função HighLatencyAddToFunctions().
	//
	//   Saída:
	//     index: Ordem de execução da função.
	//
	//   Notas:
	//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
	HighLatencySetAsFistFunctionToRun(UId string) int

	// HighLatencySetAsLastFunctionToRun
	//
	// English:
	//
	// Makes the function last on the execution list for high latency functions.
	//
	//   Input:
	//     UId: ID returned by the HighLatencyAddToFunctions() function.
	//
	//   Entrada:
	//     index: Function execution order.
	//
	//   Notes:
	//     * High latency functions are secondary functions designed to run at a lower FPS rate.
	//
	// Português:
	//
	// Faz a função ser a útima da lista de execuções para funções de alta latencia.
	//
	//   Entrada:
	//     UId: ID retornado pela função HighLatencyAddToFunctions().
	//
	//   Saída:
	//     index: Ordem de execução da função.
	//
	//   Notas:
	//     * Funções de alta latência são funções secundárias feitas para rodarem em uma taxa de FPS menor.
	HighLatencySetAsLastFunctionToRun(UId string) int

	// SystemAddToFunctions
	//
	// English:
	//
	// Adds a function to the run list.
	//
	//   Input:
	//     runnerFunc: função a ser executada.
	//
	//   Output:
	//     UId: used to identify the function when removing.
	//     total: total number of functions running.
	//
	//   Notes:
	//     * System functions are the first functions in the list of executions and should be the system usage function.
	//
	// Português:
	//
	// Adiciona uma função a lista de execuções.
	//
	//   Entrada:
	//     runnerFunc: função a ser executada.
	//
	//   Saída:
	//     UId da função, usado para identificar a função na hora de remover.
	//     total: quantidade total de funções em execução.
	//
	//   Notas:
	//     * Funções de sistema são as primeiras funções da lista de execuções e devem ser as funções de uso do sistema.
	SystemAddToFunctions(runnerFunc func()) (string, int)

	// SystemDeleteFromFunctions
	//
	// English:
	//
	// Removes a function from the system functions list, added by the SystemAddToFunctions() function.
	//
	//   Input:
	//     UId: ID returned by the SystemAddToFunctions() function.
	//
	//   Notes:
	//     * System functions are the first functions in the list of executions and should be the system usage functions.
	//
	// Português:
	//
	// Remove uma função da lista de funções do sistema, adicionada pela função SystemAddToFunctions().
	//
	//   Entrada:
	//     UId: ID retornado pela função SystemAddToFunctions().
	//
	//   Notas:
	//     * Funções de sistema são as primeiras funções da lista de execuções e devem ser as funções de uso do sistema.
	SystemDeleteFromFunctions(UId string)

	// SystemSetZIndex
	//
	// English:
	//
	// Allows you to change the order of execution of the function, in the execution list.
	//
	//   Input:
	//     UId: ID returned by the SystemAddToFunctions() function.
	//     index: 0 for the first function in the list
	//
	//   Notes:
	//     * System functions are the first functions in the list of executions and should be the system usage functions.
	//
	// Português:
	//
	// Permite trocar a ordem de execução da função, na lista de execução.
	//
	//   Entrada:
	//     UId: ID retornado pela função SystemAddToFunctions().
	//     index: 0 para a primeira função da lista
	//
	//   Notas:
	//     * Funções de sistema são as primeiras funções da lista de execuções e devem ser as funções de uso do sistema.
	SystemSetZIndex(UId string, index int) int

	// SystemGetZIndex
	//
	// English:
	//
	// Returns the function execution index in the list, where 0 is the first function to be executed.
	//
	//   Input:
	//     UId: ID returned by the SystemAddToFunctions() function.
	//
	//   Entrada:
	//     index: Function execution order.
	//
	// Português:
	//
	// Retorna o índice de execução da função na lista, onde 0 é a primera função a ser executada.
	//
	//   Entrada:
	//     UId: ID retornado pela função SystemAddToFunctions().
	//
	//   Saída:
	//     index: Ordem de execução da função.
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

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
