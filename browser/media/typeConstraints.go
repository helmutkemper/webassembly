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
	Media        Media
	Video        Video
	SharedScreen SharedScreen
}

func (e Constraints) Object() (object map[string]interface{}) {
	object = make(map[string]interface{})

	e.Video.mount(&object)

	return
}

// This is the list of ConstraintSets that the User Agent MUST attempt to satisfy, in order, skipping only those that cannot be satisfied. The order of these ConstraintSets is significant. In particular, when they are passed as an argument to applyConstraints, the User Agent MUST try to satisfy them in the order that is specified. Thus if advanced ConstraintSets C1 and C2 can be satisfied individually, but not together, then whichever of C1 and C2 is first in this list will be satisfied, and the other will not. The User Agent MUST attempt to satisfy all ConstraintSets in the list, even if some cannot be satisfied. Thus, in the preceding example, if constraint C3 is specified after C1 and C2, the User Agent will attempt to satisfy C3 even though C2 cannot be satisfied. Note that a given property name may occur only once in each ConstraintSet but may occur in more than one ConstraintSet.
// Esta é a lista de ConstraintSets que o User Agent DEVE tentar satisfazer, na ordem, pulando apenas aqueles que não podem ser satisfeitos. A ordem desses ConstraintSets é significativa. Em particular, quando eles são passados como um argumento para applyConstraints, o User Agent DEVE tentar satisfazê-los na ordem especificada. Assim, se ConstraintSets avançados C1 e C2 puderem ser satisfeitos individualmente, mas não juntos, então o que for C1 e C2 primeiro nesta lista será satisfeito, e o outro não. O User Agent DEVE tentar satisfazer todos os ConstraintSets na lista, mesmo que alguns não possam ser satisfeitos. Assim, no exemplo anterior, se a restrição C3 for especificada após C1 e C2, o User Agent tentará satisfazer C3 mesmo que C2 não possa ser satisfeito. Observe que um determinado nome de propriedade pode ocorrer apenas uma vez em cada ConstraintSet, mas pode ocorrer em mais de um ConstraintSet.

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
