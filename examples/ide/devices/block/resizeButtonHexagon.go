package block

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/utilsDraw"
)

// ResizeButton
//
// English:
//
//	ResizeButton defines an interface for creating and manipulating a resizable polygon button.
//	It enables setting position, size, color, rotation, stroke, and provides access to the SVG element.
//
// Português:
//
//	ResizeButton define uma interface para criar e manipular um botão poligonal redimensionável.
//	Permite definir posição, tamanho, cor, rotação, borda e oferece acesso ao elemento SVG.
type ResizeButton interface {
	// SetCX
	//
	// English:
	//
	//  Sets the X coordinate (cx) of the hexagon's center.
	//  If the SVG element is already initialized, its X position is updated accordingly.
	//
	// Português:
	//
	//  Define a coordenada X (cx) do centro do hexágono.
	//  Se o elemento SVG já estiver inicializado, sua posição X é atualizada de forma correspondente.
	SetCX(cx rulesDensity.Density)

	// SetCY
	//
	// English:
	//
	//  Sets the Y coordinate (cy) of the hexagon's center.
	//  If the SVG element is already initialized, its Y position is updated accordingly.
	//
	// Português:
	//
	//  Define a coordenada Y (cy) do centro do hexágono.
	//  Se o elemento SVG já estiver inicializado, sua posição Y é atualizada de forma correspondente.
	SetCY(cy rulesDensity.Density)

	// SetSides
	//
	// English:
	//
	//  Sets the number of sides for the polygon.
	//
	// Português:
	//
	//  Define o número de lados do polígono.
	SetSides(sides int)

	// SetSize
	//
	// English:
	//
	//  Sets the radius (size from center to vertex).
	//
	// Português:
	//
	//  Define o raio (distância do centro ao vértice).
	SetSize(size rulesDensity.Density)

	// SetRotation
	//
	// English:
	//
	//  Sets the rotation angle in radians.
	//
	// Português:
	//
	//  Define o ângulo de rotação em radianos.
	SetRotation(rot float64)

	// SetFillColor
	//
	// English:
	//
	//  Sets the fill color for the polygon.
	//
	// Português:
	//
	//  Define a cor de preenchimento do polígono.
	SetFillColor(color any)

	// SetStrokeColor
	//
	// English:
	//
	//  Sets the stroke (border) color for the polygon.
	//
	// Português:
	//
	//  Define a cor da borda do polígono.
	SetStrokeColor(color any)

	// SetStrokeWidth
	//
	// English:
	//
	//  Sets the width of the stroke (border).
	//
	// Português:
	//
	//  Define a largura da borda do polígono.
	SetStrokeWidth(width rulesDensity.Density)

	// SetName
	//
	// English:
	//
	//  Defines the name used to identify the pressed button
	//
	// Português:
	//
	//  Define o nome usado para identificação do botão pressionado
	SetName(name string)

	// SetCursor
	//
	// English:
	//
	//  Sets the cursor style or behavior for the hexagon button. Updates the associated SVG element if initialized.
	//
	// Português:
	//
	//  Define o estilo ou comportamento do cursor para o botão hexagonal. Atualiza o elemento SVG associado se inicializado.
	SetCursor(cursor any)

	// SetVisible
	//
	// English:
	//
	//  Sets the visibility of the hexagon button. A value of true makes it visible, while false hides it.
	//
	// Português:
	//
	//  Define a visibilidade do botão hexagonal. Um valor verdadeiro o torna visível, enquanto falso o oculta.
	SetVisible(visible bool, father html.Compatible)

	// GetSvg
	//
	// English:
	//
	//  Returns the SVG container element representing the hexagon button.
	//
	// Português:
	//
	//  Retorna o elemento container SVG que representa o botão hexagonal.
	GetSvg() (svg *html.TagSvg)

	GetNew() (new *ResizeButtonHexagon)

	// GetSpace
	//
	// English:
	//
	//  Returns the space in pixels between the center of the resort image and the edge of the image
	//
	// Português:
	//
	//  Retorna o espaço em pixels entre o centro do da imagem resizer e a borda da imagem
	GetSpace() (space rulesDensity.Density)
}

// ResizeButtonHexagon
//
// English:
//
//	Represents a resizable hexagon-shaped SVG button with customizable sides,
//	size, rotation, fill color, stroke color, and stroke width.
//
// Português:
//
//	Representa um botão SVG redimensionável em forma de hexágono,
//	com lados, tamanho, rotação, cor de preenchimento, cor da borda e espessura da borda personalizáveis.
type ResizeButtonHexagon struct {
	svgAppended bool
	svg         *html.TagSvg     // SVG container element / Elemento container SVG
	button      *html.TagSvgPath // SVG path element / Elemento de caminho SVG

	colorFill   any                  // Fill color / Cor de preenchimento
	colorStroke any                  // Stroke color / Cor da borda
	StrokeWidth rulesDensity.Density // Stroke width / Espessura da borda

	name    string // Name used as an identifier
	cursor  any    // cursor represents the mouse pointer style or behavior for the element.
	visible bool   // visible indicates whether the hexagon button is visible or hidden.

	sides    int                  // Number of sides / Número de lados
	size     rulesDensity.Density // Radius (size from center to vertex) / Raio (distância do centro ao vértice)
	cx       rulesDensity.Density // Center X (not used here) / Centro X (não utilizado aqui)
	cy       rulesDensity.Density // Center Y (não usado aqui) / Centro Y (não utilizado aqui)
	space    rulesDensity.Density // Space between the center of the resize and the edge of the image
	rotation float64              // Rotation in radians / Rotação em radianos
}

func (e *ResizeButtonHexagon) GetNew() (newResize *ResizeButtonHexagon) {
	e.init()

	n := new(ResizeButtonHexagon)
	n.SetFillColor(e.colorFill)
	n.colorStroke = e.colorStroke
	n.StrokeWidth = e.StrokeWidth
	n.name = e.name
	n.cursor = e.cursor
	n.visible = e.visible
	n.sides = e.sides
	n.size = e.size
	n.SetCX(e.cx)
	n.SetCY(e.cy)
	n.rotation = e.rotation
	n.space = e.space
	n.init()
	return n
}

// SetSpace
//
// English:
//
//	Defines the space in pixels between the center of the resort image and the edge of the image
//
// Português:
//
//	Define o espaço em pixels entre o centro do da imagem resizer e a borda da imagem
func (e *ResizeButtonHexagon) SetSpace(space rulesDensity.Density) {
	e.space = space
}

// GetSpace
//
// English:
//
//	Returns the space in pixels between the center of the resort image and the edge of the image
//
// Português:
//
//	Retorna o espaço em pixels entre o centro do da imagem resizer e a borda da imagem
func (e *ResizeButtonHexagon) GetSpace() (space rulesDensity.Density) {
	return e.space
}

// SetName
//
// English:
//
//	Defines the name used to identify the pressed button
//
// Português:
//
//	Define o nome usado para identificação do botão pressionado
func (e *ResizeButtonHexagon) SetName(name string) {
	e.name = name

	if e.svg != nil {
		e.svg.DataKey("name", e.name)
	}
}

// SetCX
//
// English:
//
//	Sets the X coordinate (cx) of the hexagon's center.
//	If the SVG element is already initialized, its X position is updated accordingly.
//
// Português:
//
//	Define a coordenada X (cx) do centro do hexágono.
//	Se o elemento SVG já estiver inicializado, sua posição X é atualizada de forma correspondente.
func (e *ResizeButtonHexagon) SetCX(cx rulesDensity.Density) {
	e.cx = cx

	// Only update position if SVG is already initialized
	// Atualiza a posição apenas se o SVG já estiver inicializado
	if e.svg != nil {
		e.svg.Get().Call("removeAttribute", "transform")
		e.svg.X(cx.GetInt() - e.size.GetInt())
		e.svg.Transform(new(html.TransformFunctions).Rotate(e.rotation, e.cx.GetFloat(), e.cy.GetFloat()))
	}
}

// SetCY
//
// English:
//
//	Sets the Y coordinate (cy) of the hexagon's center.
//	If the SVG element is already initialized, its Y position is updated accordingly.
//
// Português:
//
//	Define a coordenada Y (cy) do centro do hexágono.
//	Se o elemento SVG já estiver inicializado, sua posição Y é atualizada de forma correspondente.
func (e *ResizeButtonHexagon) SetCY(cy rulesDensity.Density) {
	e.cy = cy

	// Only update position if SVG is already initialized
	// Atualiza a posição apenas se o SVG já estiver inicializado
	if e.svg != nil {
		e.svg.Get().Call("removeAttribute", "transform")
		e.svg.Y(cy.GetInt() - e.size.GetInt())
		e.svg.Transform(new(html.TransformFunctions).Rotate(e.rotation, e.cx.GetFloat(), e.cy.GetFloat()))
	}
}

// GetSvg
//
// English:
//
//	Returns the SVG container element representing the hexagon button.
//
// Português:
//
//	Retorna o elemento container SVG que representa o botão hexagonal.
func (e *ResizeButtonHexagon) GetSvg() (svg *html.TagSvg) {
	return e.svg
}

// SetSides
//
// English:
//
//	Sets the number of sides for the polygon.
//
// Português:
//
//	Define o número de lados do polígono.
func (e *ResizeButtonHexagon) SetSides(sides int) {
	e.sides = sides
}

// SetSize
//
// English:
//
//	Sets the radius (size from center to vertex).
//
// Português:
//
//	Define o raio (distância do centro ao vértice).
func (e *ResizeButtonHexagon) SetSize(size rulesDensity.Density) {
	e.size = size
}

// SetRotation
//
// English:
//
//	Sets the rotation angle in radians.
//
// Português:
//
//	Define o ângulo de rotação em radianos.
func (e *ResizeButtonHexagon) SetRotation(rot float64) {
	e.rotation = rot

	if e.svg != nil {
		e.svg.Transform(new(html.TransformFunctions).Rotate(e.rotation, e.cx.GetFloat(), e.cy.GetFloat()))
	}
}

// SetFillColor
//
// English:
//
//	Sets the fill color for the polygon.
//
// Português:
//
//	Define a cor de preenchimento do polígono.
func (e *ResizeButtonHexagon) SetFillColor(color any) {
	e.colorFill = color

	if e.svg != nil {
		e.button.Fill(e.colorFill)
	}
}

// SetStrokeColor
//
// English:
//
//	Sets the stroke (border) color for the polygon.
//
// Português:
//
//	Define a cor da borda do polígono.
func (e *ResizeButtonHexagon) SetStrokeColor(color any) {
	e.colorStroke = color

	if e.svg != nil {
		e.button.Stroke(e.colorStroke)
	}
}

// SetStrokeWidth
//
// English:
//
//	Sets the width of the stroke (border).
//
// Português:
//
//	Define a largura da borda do polígono.
func (e *ResizeButtonHexagon) SetStrokeWidth(width rulesDensity.Density) {
	e.StrokeWidth = width

	if e.svg != nil {
		e.button.StrokeWidth(e.StrokeWidth)
	}
}

// SetCursor
//
// English:
//
//	Sets the cursor style or behavior for the hexagon button. Updates the associated SVG element if initialized.
//
// Português:
//
//	Define o estilo ou comportamento do cursor para o botão hexagonal. Atualiza o elemento SVG associado se inicializado.
func (e *ResizeButtonHexagon) SetCursor(cursor any) {
	e.cursor = cursor

	if e.svg != nil {
		e.button.Cursor(e.cursor)
	}
}

// SetVisible
//
// English:
//
//	Sets the visibility of the hexagon button. A value of true makes it visible, while false hides it.
//
// Português:
//
//	Define a visibilidade do botão hexagonal. Um valor verdadeiro o torna visível, enquanto falso o oculta.
func (e *ResizeButtonHexagon) SetVisible(visible bool, father html.Compatible) {
	if visible {
		e.svgAppended = true
		father.Get().Call("appendChild", e.svg.Get())
	} else if e.svgAppended {
		e.svgAppended = false
		father.Get().Call("removeChild", e.svg.Get())
	}

	e.svg.AddStyleConditional(visible, "display", "block", "none")
}

// Init
//
// English:
//
//	Initializes the SVG hexagon button. If the number of sides is not set,
//	it defaults to 6 (hexagon). The polygon path is created and added to the SVG element.
//
// Português:
//
//	Inicializa o botão SVG em forma de hexágono. Se o número de lados não for definido,
//	o padrão é 6 (hexágono). O caminho do polígono é criado e adicionado ao elemento SVG.
func (e *ResizeButtonHexagon) init() {
	if e.sides == 0 {
		e.sides = 6
	}

	// Generate SVG path data for the polygon
	// Gera os dados de caminho SVG para o polígono
	path := utilsDraw.PolygonPath(e.sides, e.size, e.size, e.size, e.rotation)

	// Create SVG root element
	// Cria o elemento SVG raiz
	e.svg = factoryBrowser.NewTagSvg().
		Width(2*e.size.GetInt()).
		Height(2*e.size.GetInt()).
		X(e.cx.GetInt()-e.size.GetInt()).
		Y(e.cy.GetInt()).
		DataKey("name", e.name)

	// Create and configure the path element
	// Cria e configura o elemento de caminho SVG
	e.button = factoryBrowser.NewTagSvgPath().
		Fill(e.colorFill).
		Stroke(e.colorStroke).
		StrokeWidth(e.StrokeWidth.GetInt()).
		D(path)

	// Add the path to the SVG
	// Adiciona o caminho ao SVG
	e.svg.Append(e.button)
}
