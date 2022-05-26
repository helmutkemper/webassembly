package html

type MeetOrSliceReference string

func (e MeetOrSliceReference) String() string {
	return string(e)
}

const (
	// KMeetOrSliceReferenceMeet
	//
	// English:
	//
	//  (default) Scale the graphic such that:
	//
	// Aspect ratio is preserved
	//
	// The entire viewBox is visible within the viewport
	//
	// The viewBox is scaled up as much as possible, while still meeting the other criteria
	//
	// In this case, if the aspect ratio of the graphic does not match the viewport, some of the viewport will extend
	// beyond the bounds of the viewBox (i.e., the area into which the viewBox will draw will be smaller than the
	// viewport).
	//
	// Português:
	//
	//  (default) Dimensione o gráfico de tal forma que:
	//
	// Proporção é preservada
	//
	// Toda a viewBox é visível dentro da viewport
	//
	// O viewBox é ampliado o máximo possível, enquanto ainda atende aos outros critérios
	//
	// Nesse caso, se a proporção do gráfico não corresponder à janela de visualização, parte da janela de visualização se
	// estenderá além dos limites da caixa de visualização (ou seja, a área na qual a caixa de visualização será desenhada
	// será menor que a porta de visualização).
	KMeetOrSliceReferenceMeet MeetOrSliceReference = "meet"

	// KMeetOrSliceReferenceSlice
	//
	// English:
	//
	//  Scale the graphic such that:
	//
	// Aspect ratio is preserved
	//
	// The entire viewport is covered by the viewBox
	//
	// The viewBox is scaled down as much as possible, while still meeting the other criteria
	//
	// In this case, if the aspect ratio of the viewBox does not match the viewport, some of the viewBox will extend
	// beyond the bounds of the viewport (i.e., the area into which the viewBox will draw is larger than the viewport).
	//
	// Português:
	//
	//  Dimensione o gráfico de tal forma que:
	//
	// Proporção é preservada
	//
	// A viewport inteira é coberta pela viewBox
	//
	// A viewBox é reduzida o máximo possível, enquanto ainda atende aos outros critérios
	//
	// Nesse caso, se a proporção da viewBox não corresponder à viewport, parte da viewBox se estenderá além dos limites
	// da viewport (ou seja, a área na qual a viewBox desenhará é maior que a viewport).
	KMeetOrSliceReferenceSlice MeetOrSliceReference = "slice"
)
