package media

// Constraints
//
// English:
//
// Specify one or more exact values from which one must be the parameter's value, or a set of ideal values which should
// be used if possible. You can also specify a single value (or an array of values) which the user agent will do its
// best to match once all more stringent constraints have been applied.
//
// Português:
//
// Especifique um ou mais valores exatos dos quais um deve ser o valor do parâmetro, ou um conjunto de valores ideais
// que devem ser usados se possível. Você também pode especificar um valor único (ou uma matriz de valores) que o agente
// do usuário fará o possível para corresponder assim que todas as restrições mais rigorosas forem aplicadas.
type Constraints struct {
	Media        Media        `js:"media"`
	Audio        Audio        `js:"audio"`
	Image        Image        `js:"image"`
	Video        Video        `js:"video"`
	SharedScreen SharedScreen `js:""`
}
