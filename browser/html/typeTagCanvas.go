package html

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/css"
	"github.com/helmutkemper/iotmaker.webassembly/browser/event/generic"
	"image/color"
	"log"
	"strconv"
	"syscall/js"
)

// todo: transformar em exemplo
// https://github.com/mdn/content/blob/main/files/en-us/web/api/canvas_api/tutorial/compositing/example/index.md?plain=1

// TagCanvas
//
// English:
//
// Use the HTML <canvas> element with either the canvas scripting API or the WebGL API to draw graphics and animations.
//
// Português:
//
// Use o elemento HTML <canvas> com a API de script de tela ou a API WebGL para desenhar gráficos e animações.
type TagCanvas struct {
	commonEvents commonEvents

	tag         Tag
	id          string
	selfElement js.Value
	cssClass    *css.Class

	// stage
	//
	// English:
	//
	//  Browser main document reference captured at startup.
	//
	// Português:
	//
	//  Referencia do documento principal do navegador capturado na inicialização.
	stage js.Value

	context js.Value
	width   int
	height  int

	gradient  js.Value
	pattern   js.Value
	transform js.Value

	collisionDataFunction func(red, green, blue, alpha uint8) bool
}

// Reference
//
// English:
//
// Pass the object reference to an external variable.
//
// Português:
//
// Passa a referencia do objeto para uma variável externa.
//
//	Example: / Exemplo:
//	  var circle *html.TagSvgCircle
//	  factoryBrowser.NewTagSvgCircle().Reference(&circle).R(5).Fill(factoryColor.NewRed())
//	  log.Printf("x: %v, y: %v", circle.GetX(), circle.GetY())
func (el *TagCanvas) Reference(reference **TagCanvas) (ref *TagCanvas) {
	*reference = el
	return el
}

// Init
//
// English:
//
//	Initializes the object correctly.
//
// Português:
//
//	Inicializa o objeto corretamente.
func (el *TagCanvas) Init(width, height int) (ref *TagCanvas) {
	//e.listener = new(sync.Map)
	el.CreateElement("canvas", width, height)

	el.collisionDataFunction = func(r, g, b, a uint8) bool {
		if a != 0 {
			return true
		}

		return false
	}

	return el
}

// Id
//
// English:
//
//	Specifies a unique id for an element
//
// The id attribute specifies a unique id for an HTML element (the value must be unique within the
// HTML document).
//
// The id attribute is most used to point to a style in a style sheet, and by JavaScript (via the HTML
// DOM) to manipulate the element with the specific id.
//
// Português:
//
//	Especifica um ID exclusivo para um elemento
//
// O atributo id especifica um id exclusivo para um elemento HTML (o valor deve ser exclusivo no
// documento HTML).
//
// O atributo id é mais usado para apontar para um estilo em uma folha de estilo, e por JavaScript
// (através do HTML DOM) para manipular o elemento com o id específico.
func (el *TagCanvas) Id(id string) (ref *TagCanvas) {
	el.id = id
	el.selfElement.Set("id", id)
	return el
}

// CreateElement
//
// English:
//
//	In an HTML document, the Document.createElement() method creates the specified HTML element or an
//	HTMLUnknownElement if the given element name is not known.
//
// Português:
//
//	Em um documento HTML, o método Document.createElement() cria o elemento HTML especificado ou um
//	HTMLUnknownElement se o nome do elemento dado não for conhecido.
func (el *TagCanvas) CreateElement(tag Tag, width, height int) (ref *TagCanvas) {
	el.selfElement = js.Global().Get("document").Call("createElement", tag.String())
	if el.selfElement.IsUndefined() == true || el.selfElement.IsNull() == true {
		log.Print(KNewElementIsUndefined)
		return
	}
	el.tag = tag

	el.context = el.selfElement.Call("getContext", "2d")
	el.selfElement.Set("width", width)
	el.selfElement.Set("height", height)

	el.width = width
	el.height = height

	return el
}

func (el *TagCanvas) Get() (element js.Value) {
	return el.selfElement
}

// DrawImage
//
// English:
//
//		Draws a preloaded image on the canvas element.
//
//		 Input:
//		   image: js.Value object with a preloaded image.
//
//	 Notes:
//	   * Make sure the image is preloaded before calling this function.
//
// Português:
//
//		Desenha uma imagem pre carregada no elemento canvas.
//
//		 Entrada:
//		   image: objeto js.Value com uma imagem pré carregada.
//
//	 Notas:
//	   * Garanta o pré carregamento da imagem antes de chamar esta função.
func (el *TagCanvas) DrawImage(image interface{}) (ref *TagCanvas) {
	switch converted := image.(type) {
	case Compatible:
		el.context.Call("drawImage", converted.Get(), 0, 0, el.width, el.height)
	case js.Value:
		el.context.Call("drawImage", converted, 0, 0, el.width, el.height)
	default:
		log.Printf("canvas.DrawImage(image).err: image must be a js.Value image or TagImg object.")
	}

	return el
}

// DrawImageResize
//
// English:
//
//		Draws a preloaded image on the canvas element.
//
//		 Input:
//		   image: js.Value object with a preloaded image.
//	    x: Point x of the image to be copied;
//	    y: Point y of the image to be copied;
//	    width: Horizontal size of the image to be copied;
//	    height: Vertical size of the image to be copied;
//
//	 Notes:
//	   * Make sure the image is preloaded before calling this function.
//
// Português:
//
//		Desenha uma imagem pre carregada no elemento canvas.
//
//		 Entrada:
//		   image: objeto js.Value com uma imagem pré carregada;
//	    x: Ponto x da imagem a ser copiada;
//	    y: Ponto y da imagem a ser copiada;
//	    width: Tamanho horizontal da imagem a ser copiada;
//	    height: Tamanho vertical da imagem a ser copiada;
//
//	 Notas:
//	   * Garanta o pré carregamento da imagem antes de chamar esta função.
func (el *TagCanvas) DrawImageResize(image interface{}, x, y, width, height int) (ref *TagCanvas) {
	if converted, ok := image.(Compatible); ok {
		el.context.Call("drawImage", converted.Get(), x, y, width, height)
		return el
	}

	el.context.Call("drawImage", image, x, y, width, height)
	return el
}

// DrawImageTile
//
// English:
//
//		Draws a preloaded image on the canvas element.
//
//		 Input:
//		   image: js.Value object with a preloaded image.
//	    x: Point x of the image to be copied;
//	    y: Point y of the image to be copied;
//	    width: Horizontal size of the image to be copied;
//	    height: Vertical size of the image to be copied;
//
//	 Notes:
//	   * Make sure the image is preloaded before calling this function.
//
// Português:
//
//		Desenha uma imagem pre carregada no elemento canvas.
//
//		 Entrada:
//		   image: objeto js.Value com uma imagem pré carregada;
//	    x: Ponto x da imagem a ser copiada;
//	    y: Ponto y da imagem a ser copiada;
//	    width: Tamanho horizontal da imagem a ser copiada;
//	    height: Tamanho vertical da imagem a ser copiada;
//
//	 Notas:
//	   * Garanta o pré carregamento da imagem antes de chamar esta função.
func (el *TagCanvas) DrawImageTile(image interface{}, dx, dy, x, y, width, height int) (ref *TagCanvas) {
	if converted, ok := image.(Compatible); ok {
		el.context.Call("drawImage", converted.Get(), dx, dy, width, height, x, y, width, height)
		return el
	}

	el.context.Call("drawImage", image, dx, dy, width, height, x, y, width, height)
	return el
}

// DrawImageComplete
//
// English:
//
// Draws a preloaded image on the canvas element.
//
//	Input:
//	  sx: The x-axis coordinate of the top left corner of the sub-rectangle of the source image to draw into the
//	    destination context.
//	  sy: The y-axis coordinate of the top left corner of the sub-rectangle of the source image to draw into the
//	    destination context.
//	  sWidth: The width of the sub-rectangle of the source image to draw into the destination context.
//	  sHeight: The height of the sub-rectangle of the source image to draw into the destination context.
//	    A negative value will flip the image.
//	  dx: The x-axis coordinate in the destination canvas at which to place the top-left corner of the source image.
//	  dy: The y-axis coordinate in the destination canvas at which to place the top-left corner of the source image.
//	  dWidth: The width to draw the image in the destination canvas. This allows scaling of the drawn image.
//	  dHeight: The height to draw the image in the destination canvas. This allows scaling of the drawn image.
//
// Português:
//
// Desenha uma imagem pré-carregada no elemento canvas.
//
//	Entrada:
//	  sx: A coordenada do eixo x do canto superior esquerdo do sub-retângulo da imagem de origem para desenhar no
//	    contexto de destino.
//	  sy: A coordenada do eixo y do canto superior esquerdo do sub-retângulo da imagem de origem para desenhar no
//	    contexto de destino.
//	  sWidth: A largura do sub-retângulo da imagem de origem a ser desenhada no contexto de destino. Um valor negativo
//	    inverterá à imagem.
//	  sHeight: A altura do sub-retângulo da imagem de origem a ser desenhada no contexto de destino. Um valor negativo
//	    inverterá à imagem.
//	  dx: A coordenada do eixo x na tela de destino na qual colocar o canto superior esquerdo da imagem de origem.
//	  dy: A coordenada do eixo y na tela de destino na qual colocar o canto superior esquerdo da imagem de origem.
//	  dWidth: A largura para desenhar à imagem na tela de destino. Isso permite dimensionar à imagem desenhada.
//	  dHeight: A altura para desenhar à imagem na tela de destino. Isso permite dimensionar à imagem desenhada.
func (el *TagCanvas) DrawImageComplete(image interface{}, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight int) (ref *TagCanvas) {
	if converted, ok := image.(Compatible); ok {
		el.context.Call("drawImage", converted.Get(), sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight)
		return el
	}

	el.context.Call("drawImage", image, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight)
	return el
}

// AppendById
//
// English:
//
//	Adds a node to the end of the list of children of a specified parent node. If the node already
//	exists in the document, it is removed from its current parent node before being added to the
//	new parent.
//
//	 Input:
//	   appendId: id of parent element.
//
//	 Note:
//	   * The equivalent of:
//	       var p = document.createElement("p");
//	       document.body.appendChild(p);
//
// Português:
//
//	Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no
//	documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//	 Entrada:
//	   appendId: id do elemento pai.
//
//	 Nota:
//	   * Equivale a:
//	       var p = document.createElement("p");
//	       document.body.appendChild(p);
func (el *TagCanvas) AppendById(appendId string) (ref *TagCanvas) {

	toAppend := js.Global().Get("document").Call("getElementById", appendId)
	if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
		log.Print(KIdToAppendNotFound, appendId)
		return el
	}

	toAppend.Call("appendChild", el.selfElement)
	return el
}

// Append
//
// English:
//
//	Adds a node to the end of the list of children of a specified parent node. If the node already
//	exists in the document, it is removed from its current parent node before being added to the new
//	parent.
//
//	 Input:
//	   append: element in js.Value format.
//
//	 Note:
//	   * The equivalent of:
//	       var p = document.createElement("p");
//	       document.body.appendChild(p);
//
// Português:
//
//	Adiciona um nó ao final da lista de filhos de um nó pai especificado. Se o nó já existir no
//	documento, ele é removido de seu nó pai atual antes de ser adicionado ao novo pai.
//
//	 Entrada:
//	   appendId: elemento no formato js.Value.
//
//	 Nota:
//	   * Equivale a:
//	       var p = document.createElement("p");
//	       document.body.appendChild(p);
func (el *TagCanvas) Append(append interface{}) (ref *TagCanvas) {
	switch converted := append.(type) {
	case *TagCanvas:
		el.selfElement.Call("appendChild", converted.selfElement)
	case js.Value:
		el.selfElement.Call("appendChild", converted)
	case string:
		toAppend := js.Global().Get("document").Call("getElementById", converted)
		if toAppend.IsUndefined() == true || toAppend.IsNull() == true {
			log.Print(KIdToAppendNotFound, converted)
			return el
		}

		toAppend.Call("appendChild", el.selfElement)
	}

	return el
}

// CollisionDataFunction
//
// English:
//
// Determines a new function for calculating the default collision matrix, return ALPHA != 0
//
// Português:
//
// Determina uma nova função para o cálculo da matriz de colisão padrão, return ALPHA != 0
func (el *TagCanvas) CollisionDataFunction(f func(red, green, blue, alpha uint8) bool) (ref *TagCanvas) {
	el.collisionDataFunction = f
	return ref
}

// GetCollisionData
//
// English:
//
//	Returns an array (x,y) with a boolean indicating transparency.
//
// The collision map is a quick way to preload data about the coordinates of where there are parts of
// the image.
//
//	Output:
//	  data: [x][y]transparent
//
// The collision matrix is a Boolean map of (x,y) coordinates based on the image box, where (0,0) represents the (0,0)
// coordinate of the image box, like this:
//
//	, , , , , , , , , , ,
//	, , , , ,1,1, , , , ,
//	, , ,1,1,1,1,1,1, , ,
//	, , ,1,1,1,1,1,1, , ,
//	,1,1,1,1,1,1,1,1,1,1,
//	,1,1,1,1,1,1,1,1,1,1,
//	, , , , , , , , , , ,
//
// Português:
//
//	Retorna uma array (x,y) com um booleano indicando transparência.
//
// O mapa de colisão é uma forma rápida de deixar um dado pre carregado sobre as coordenadas de onde
// há partes da imagem.
//
//	Saída:
//	  data: [x][y]transparente
//
// A matrix de colisão é um mapa booleano de coordenadas (x,y) baseado na caixa da imagem, onde (0,0) representa a
// coordenada (0,0) da caixa da imagem, assim:
//
//	, , , , , , , , , , ,
//	, , , , ,1,1, , , , ,
//	, , ,1,1,1,1,1,1, , ,
//	, , ,1,1,1,1,1,1, , ,
//	,1,1,1,1,1,1,1,1,1,1,
//	,1,1,1,1,1,1,1,1,1,1,
//	, , , , , , , , , , ,
//
// todo: fazer exemplo
func (el *TagCanvas) GetCollisionData() (data [][]bool) {

	dataInterface := el.context.Call("getImageData", 0, 0, el.width, el.height)
	dataJs := dataInterface.Get("data")

	var rgbaLength = 4

	var i = 0
	var x int
	var y int

	var red uint8
	var green uint8
	var blue uint8
	var alpha uint8

	// [x][y]bool
	data = make([][]bool, el.height)
	for y = 0; y != el.height; y += 1 {
		data[y] = make([]bool, el.width)
		for x = 0; x != el.width; x += 1 {

			red = uint8(dataJs.Index(i + 0).Int())
			green = uint8(dataJs.Index(i + 1).Int())
			blue = uint8(dataJs.Index(i + 2).Int())
			alpha = uint8(dataJs.Index(i + 3).Int())

			data[y][x] = el.collisionDataFunction(red, green, blue, alpha)

			i += rgbaLength
		}
	}

	return
}

func (el *TagCanvas) GetCollisionDataFromImageData(dataJs js.Value, width, height int) (data [][]bool) {
	dataJs = dataJs.Get("data")

	var rgbaLength = 4

	var i = 0
	var x int
	var y int

	var red uint8
	var green uint8
	var blue uint8
	var alpha uint8

	data = make([][]bool, height)
	for y = 0; y != height; y += 1 {
		data[y] = make([]bool, width)
		for x = 0; x != width; x += 1 {

			red = uint8(dataJs.Index(i + 0).Int())
			green = uint8(dataJs.Index(i + 1).Int())
			blue = uint8(dataJs.Index(i + 2).Int())
			alpha = uint8(dataJs.Index(i + 3).Int())

			data[y][x] = el.collisionDataFunction(red, green, blue, alpha)

			i += rgbaLength
		}
	}

	return
}

func (el *TagCanvas) AdjustBox(dx, dy int) {
	//todo: implementar
}

//func (el *TagCanvas) GetCollisionBox() (box Box) {
//
//	var xMin = math.MaxInt
//	var yMin = math.MaxInt
//	var xMax = math.MinInt
//	var yMax = math.MinInt
//
//	collisionData := el.GetCollisionData()
//	for y, dataY := range collisionData {
//		for x, value := range dataY {
//			if value == true {
//				xMin = el.min(xMin, x+1)
//				yMin = el.min(yMin, y+1)
//				xMax = el.max(xMax, x)
//				yMax = el.max(yMax, y)
//			}
//		}
//	}
//
//	return Box{
//		X:         xMin,
//		Y:         yMin,
//		Width:     xMax - xMin,
//		Height:    yMax - yMin,
//		WidthImg:  el.width,
//		HeightImg: el.height,
//	}
//}

//func (el *TagCanvas) GetCollisionBoxFromImageData(dataJs js.Value, width, height int) (box Box) {
//
//	var xMin = math.MaxInt
//	var yMin = math.MaxInt
//	var xMax = math.MinInt
//	var yMax = math.MinInt
//
//	collisionData := el.GetCollisionDataFromImageData(dataJs, width, height)
//	for y, dataY := range collisionData {
//		for x, value := range dataY {
//			if value == true {
//				xMin = el.min(xMin, x+1)
//				yMin = el.min(yMin, y+1)
//				xMax = el.max(xMax, x)
//				yMax = el.max(yMax, y)
//			}
//		}
//	}
//
//	return Box{
//		X:         xMin,
//		Y:         yMin,
//		Width:     xMax - xMin,
//		Height:    yMax - yMin,
//		WidthImg:  width,
//		HeightImg: height,
//	}
//}

// flipData
//
// English:
//
// Reverses the direction of the image, both vertically and horizontally.
//
// Português:
//
// Reverte o sentido da imagem, tanto na vertical quanto na horizontal.
func (el *TagCanvas) flipData(dataJs, flipped js.Value, width, height int) {
	var x, y int

	var rgbaLength = 4

	var i = 0
	x = 0
	y = 0

	m := height * width * 4
	for y = 0; y != height; y += 1 {
		for x = 0; x != width; x += 1 {

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			flipped.SetIndex(m-i+0, uint8(dataJs.Index(i+0).Int()))
			flipped.SetIndex(m-i+1, uint8(dataJs.Index(i+1).Int()))
			flipped.SetIndex(m-i+2, uint8(dataJs.Index(i+2).Int()))
			flipped.SetIndex(m-i+3, uint8(dataJs.Index(i+3).Int()))

			i += rgbaLength
		}
	}
}

// flipDataHorizontal
//
// English:
//
// Reverses the direction of the image horizontally.
//
// Português:
//
// Inverte o sentido da imagem na horizontal.
func (el *TagCanvas) flipDataHorizontal(dataJs, flipped js.Value, width, height int) {
	var x, y, m, i int

	var rgbaLength = 4

	for y = 0; y != height; y += 1 {
		for x = 0; x != width; x += 1 {
			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			flipped.SetIndex(m+0+(width-x-1)*4, uint8(dataJs.Index(i+0).Int()))
			flipped.SetIndex(m+1+(width-x-1)*4, uint8(dataJs.Index(i+1).Int()))
			flipped.SetIndex(m+2+(width-x-1)*4, uint8(dataJs.Index(i+2).Int()))
			flipped.SetIndex(m+3+(width-x-1)*4, uint8(dataJs.Index(i+3).Int()))

			i += rgbaLength
		}
		m = m + width*4
	}
}

// flipDataVertical
//
// English:
//
// Reverses the direction of the image vertically.
//
// Português:
//
// Inverte o sentido da imagem na vertical.
func (el *TagCanvas) flipDataVertical(dataJs, flipped js.Value, width, height int) {
	var x, y, m, i int

	var rgbaLength = 4

	for y = 0; y != height; y += 1 {

		// English: Shifts the data position in the matrix to an ordinary two-dimensional matrix.
		// Português: Desloca a posição do dado na matriz para uma matriz bidimencional comum.
		m = (height*el.width - el.width*(y+1)) + 1

		for x = 0; x != width; x += 1 {
			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			// English:   +x shifts the data in the common matrix
			//            *4 shifts the data in the js image array
			// Português: +x desloca o dado na matriz comum
			//            *4 desloca o dado na matriz de imagem js
			flipped.SetIndex((m+x)*4+0, uint8(dataJs.Index(i+0).Int()))
			flipped.SetIndex((m+x)*4+1, uint8(dataJs.Index(i+1).Int()))
			flipped.SetIndex((m+x)*4+2, uint8(dataJs.Index(i+2).Int()))
			flipped.SetIndex((m+x)*4+3, uint8(dataJs.Index(i+3).Int()))

			i += rgbaLength
		}

	}
}

// GetImageData
//
// English:
//
// Returns the image data in raw js format.
//
// Português:
//
// Retorna os dados da imagem no formato js bruto.
func (el *TagCanvas) GetImageData(x, y, width, height int, flipVertical, flipHorizontal bool) (data js.Value) {
	imageData := el.context.Call("getImageData", x, y, width, height)

	if flipVertical == false && flipHorizontal == false {
		return imageData
	}

	data1 := imageData.Get("data")

	var imageDataFlipped = el.context.Call("getImageData", x, y, width, height)
	data2 := imageDataFlipped.Get("data")

	if flipVertical == false && flipHorizontal == true {
		el.flipDataHorizontal(data1, data2, width, height)
		return imageDataFlipped

	} else if flipVertical == true && flipHorizontal == false {
		el.flipDataVertical(data1, data2, width, height)
		return imageDataFlipped

	} else { //flipVertical == true && flipHorizontal == true
		el.flipData(data1, data2, width, height)
		return imageDataFlipped
	}
}

// GetImageDataMatrix
//
// English:
//
//	Returns an array copy of the image.
//
//	 Input:
//	   x: x position of the image;
//	   y: y position of the image;
//	   width: image width;
//	   height: image height.
//
//	 Output:
//	   data: image in matrix format.
//	     [y][x][0]: red color value between 0 and 255
//	     [y][x][1]: green color value between 0 and 255
//	     [y][x][2]: blue color value between 0 and 255
//	     [y][x][3]: alpha color value between 0 and 255
//
// Português:
//
//	Retorna uma cópia matricial da imagem.
//
//	 Entrada:
//	   x: Posição x da imagem;
//	   y: Posição y da imagem;
//	   width: comprimento da imagem;
//	   height: altura da imagem.
//
//	 Saída:
//	   data: imagem em formato matricial.
//	     [y][x][0]: valor da cor vermelha entre 0 e 255
//	     [y][x][1]: valor da cor verde entre 0 e 255
//	     [y][x][2]: valor da cor azul entre 0 e 255
//	     [y][x][3]: valor da cor alpha entre 0 e 255
//
// todo: fazer exemplo
func (el *TagCanvas) GetImageDataMatrix(x, y, width, height int) (data [][][]uint8) {

	dataInterface := el.context.Call("getImageData", x, y, width, height)
	dataJs := dataInterface.Get("data")

	var rgbaLength = 4

	var i = 0
	x = 0
	y = 0

	// [x][y][4-channel]
	data = make([][][]uint8, height)
	for y = 0; y != height; y += 1 {
		data[y] = make([][]uint8, width)
		for x = 0; x != width; x += 1 {
			data[y][x] = make([]uint8, 4)

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			data[y][x][0] = uint8(dataJs.Index(i + 0).Int())
			data[y][x][1] = uint8(dataJs.Index(i + 1).Int())
			data[y][x][2] = uint8(dataJs.Index(i + 2).Int())
			data[y][x][3] = uint8(dataJs.Index(i + 3).Int())

			i += rgbaLength
		}
	}

	return
}

func (el *TagCanvas) GetAverageGrayScale(x, y, width, height int) (gray [][][]uint8) {

	dataInterface := el.context.Call("getImageData", x, y, width, height)
	dataJs := dataInterface.Get("data")

	var rgbaLength = 4
	var data [][][]uint8
	var i = 0
	x = 0
	y = 0

	// [x][y][4-channel]
	data = make([][][]uint8, height)
	for y = 0; y != height; y += 1 {
		data[y] = make([][]uint8, width)
		for x = 0; x != width; x += 1 {
			data[y][x] = make([]uint8, 4)

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			data[y][x][0] = uint8(dataJs.Index(i + 0).Int())
			data[y][x][1] = uint8(dataJs.Index(i + 1).Int())
			data[y][x][2] = uint8(dataJs.Index(i + 2).Int())
			data[y][x][3] = uint8(dataJs.Index(i + 3).Int())

			i += rgbaLength
		}
	}

	gray = make([][][]uint8, height)
	for y = 0; y != height; y += 1 {
		gray[y] = make([][]uint8, width)
		for x = 0; x != width; x += 1 {
			gray[y][x] = make([]uint8, 4)

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			gray[y][x][0] = (data[y][x][0] + data[y][x][1] + data[y][x][2]) / 3
			gray[y][x][1] = (data[y][x][0] + data[y][x][1] + data[y][x][2]) / 3
			gray[y][x][2] = (data[y][x][0] + data[y][x][1] + data[y][x][2]) / 3
			gray[y][x][3] = data[y][x][3]
		}
	}

	return
}

// ImageData
//
// English:
//
//	Transform an array of data into an image.
//
//	 Input:
//	   imgData: data array with the new image;
//	     [y][x][0]: red color value between 0 and 255;
//	     [y][x][1]: green color value between 0 and 255;
//	     [y][x][2]: blue color value between 0 and 255;
//	     [y][x][3]: alpha color value between 0 and 255.
//	   width: image width;
//	   height: image height.
//
// Português:
//
//	Transforma uma matrix de dados em uma imagem.
//
//	 Entrada:
//	   imgData: array de dados com o a nova imagem;
//	     [y][x][0]: valor da cor vermelha entre 0 e 255;
//	     [y][x][1]: valor da cor verde entre 0 e 255;
//	     [y][x][2]: valor da cor azul entre 0 e 255;
//	     [y][x][3]: valor da cor alpha entre 0 e 255.
//	   width: comprimento da imagem;
//	   height: altura da imagem.
//
// todo: fazer exemplo
func (el *TagCanvas) ImageData(imgData [][][]uint8, width, height int) (ref *TagCanvas) {

	dataJs := el.context.Call("createImageData", width, height)

	var rgbaLength = 4

	var i = 0
	var x int
	var y int
	for y = 0; y != height; y += 1 {
		for x = 0; x != width; x += 1 {

			//Red:   uint8(dataJs.Index(i + 0).Int()),
			//Green: uint8(dataJs.Index(i + 1).Int()),
			//Blue:  uint8(dataJs.Index(i + 2).Int()),
			//Alpha: uint8(dataJs.Index(i + 3).Int()),

			dataJs.Get("data").SetIndex(i+0, imgData[y][x][0])
			dataJs.Get("data").SetIndex(i+1, imgData[y][x][1])
			dataJs.Get("data").SetIndex(i+2, imgData[y][x][2])
			dataJs.Get("data").SetIndex(i+3, imgData[y][x][3])

			i += rgbaLength
		}
	}

	el.context.Call("putImageData", dataJs, 0, 0, 0, 0, width, height)
	return el
}

func (el *TagCanvas) PutImageDataComplete(imageData js.Value, dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight int) (ref *TagCanvas) {
	el.context.Call("putImageData", imageData, dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
	return el
}

func (el *TagCanvas) PutImageData(imageData js.Value, dx, dy int) (ref *TagCanvas) {
	el.context.Call("putImageData", imageData, dx, dy)
	return el
}

// AddColorStopPosition
//
// English:
//
//	Specifies the colors and stop positions in a gradient object
//
//	   Input:
//	     stopPosition: A value between 0.0 and 1.0 that represents the position between start (0%) and
//	       end (100%) in a gradient;
//	     color: A color RGBA value to display at the stop position. You can use
//	       factoryColor.NewColorName() to make it easier;
//
//	 Note:
//	   * Before using this function, you need to generate a gradient with the CreateLinearGradient()
//	     or CreateRadialGradient() functions;
//	   * You can call the AddColorStopPosition() method multiple times to change a gradient. If you
//	     omit this method for gradient objects, the gradient will not be visible. You need to create
//	     at least one color stop to have a visible gradient.
//
// Português:
//
//	Especifica a cor e a posição final para a cor dentro do gradiente
//	 Entrada:
//	   stopPosition: Um valor entre 0.0 e 1.0 que representa a posição entre o início (0%) e o fim
//	     (100%) dentro do gradiente;
//	   color: Uma cor no formato RGBA para ser mostrada na posição determinada. Você pode usar
//	     factoryColor.NewColorName() para facilitar.
//
//	 Nota:
//	   * Antes de usar esta função, você necessita gerar um gradiente com as funções
//	     CreateLinearGradient() ou CreateRadialGradient();
//	   * Você pode chamar o método AddColorStopPosition() várias vezes para adicionar várias cores ao
//	     gradiente, porém, se você omitir o método, o gradiente não será visível. Você tem à obrigação
//	     de chamar o método pelo menos uma vez com uma cor para que o gradiente seja visível.
func (el *TagCanvas) AddColorStopPosition(stopPosition float64, color color.RGBA) (ref *TagCanvas) {
	el.gradient.Call("addColorStop", stopPosition, RGBAToJs(color))
	return el
}

// Arc
//
// English:
//
//	Creates an arc/curve (used to create circles, or parts of circles).
//
//	 Input:
//	   x: The horizontal coordinate of the arc's center;
//	   y: The vertical coordinate of the arc's center;
//	   radius: The arc's radius. Must be positive;
//	   startAngle: The angle at which the arc starts in radians, measured from the positive x-axis;
//	   endAngle: The angle at which the arc ends in radians, measured from the positive x-axis;
//	   anticlockwise: An optional Boolean. If true, draws the arc counter-clockwise between the start
//	     and end angles. The default is false (clockwise).
//
// Português:
//
//	Creates an arc/curve (used to create circles, or parts of circles).
//
//	 Input:
//	   x: A coordenada horizontal do centro do arco;
//	   y: A coordenada vertical do centro do arco;
//	   radius: O raio do arco. Deve ser positivo;
//	   startAngle: O ângulo no qual o arco começa em radianos, medido a partir do eixo x positivo;
//	   endAngle: O ângulo no qual o arco termina em radianos, medido a partir do eixo x positivo;
//	   anticlockwise: Um booleano opcional. Se true, desenha o arco no sentido anti-horário entre os
//	     ângulos inicial e final. O padrão é false (sentido horário).
func (el *TagCanvas) Arc(x, y int, radius, startAngle, endAngle float64, anticlockwise bool) (ref *TagCanvas) {
	el.context.Call("arc", x, y, radius, startAngle, endAngle, anticlockwise)
	return el
}

// ArcTo
//
// English:
//
//	Creates an arc/curve between two tangents.
//
//	 Input:
//	   x1: The x-axis coordinate of the first control point.
//	   y1: The y-axis coordinate of the first control point.
//	   x2: The x-axis coordinate of the second control point.
//	   y2: The y-axis coordinate of the second control point.
//	   radius: The arc's radius. Must be non-negative.
//
// Português:
//
//	Cria um arco / curva entre duas tangentes.
//
//	 Input:
//	   x1: A coordenada do eixo x do primeiro ponto de controle;
//	   y1: A coordenada do eixo y do primeiro ponto de controle;
//	   x2: A coordenada do eixo x do segundo ponto de controle;
//	   y2: A coordenada do eixo y do segundo ponto de controle;
//	   radius: O raio do arco. Deve ser não negativo.
func (el *TagCanvas) ArcTo(x1, y1, x2, y2 int, radius int) (ref *TagCanvas) {
	el.context.Call("arcTo", x1, y1, x2, y2, radius)
	return el
}

// BeginPath
//
// English:
//
//		Begins a path, or resets the current path
//
//	  Note:
//	    * Use MoveTo(), LineTo(), QuadricCurveTo(), BezierCurveTo(), ArcTo(), and Arc(), to
//	       create paths;
//	    * Use the Stroke() method to actually draw the path on the canvas.
//
// Português:
//
//	Inicia ou reinicializa uma nova rota no desenho
//
//	 Nota:
//	   * Dica: Use MoveTo(), LineTo(), QuadricCurveTo(), BezierCurveTo(), ArcTo(), e Arc(), para
//	     criar uma nova rota no desenho;
//	   * Use o método Stroke() para desenhar a rota no elemento canvas.
func (el *TagCanvas) BeginPath() (ref *TagCanvas) {
	el.context.Call("beginPath")
	return el
}

// BezierCurveTo
//
// English:
//
//	Creates a cubic Bézier curve.
//
//	 Input:
//
//	   cp1x: The x-axis coordinate of the first control point;
//	   cp1y: The y-axis coordinate of the first control point;
//	   cp2x: The x-axis coordinate of the second control point;
//	   cp2y: The y-axis coordinate of the second control point;
//	   x: The x-axis coordinate of the end point;
//	   y: The y-axis coordinate of the end point.
//
// Português:
//
//	Cria uma curva de Bézier cúbica.
//
//	 Entrada:
//
//	   cp1x: A coordenada do eixo x do primeiro ponto de controle;
//	   cp1y: A coordenada do eixo y do primeiro ponto de controle;
//	   cp2x: A coordenada do eixo x do segundo ponto de controle;
//	   cp2y: A coordenada do eixo y do segundo ponto de controle;
//	   x: A coordenada do eixo x do ponto final;
//	   y: A coordenada do eixo y do ponto final.
func (el *TagCanvas) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y int) (ref *TagCanvas) {
	el.context.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
	return el
}

// ClearRect
//
// English:
//
//	Clears the specified pixels within a given rectangle.
//	   x: The x-coordinate of the upper-left corner of the rectangle to clear
//	   y: The y-coordinate of the upper-left corner of the rectangle to clear
//	   width: The width of the rectangle to clear, in pixels
//	   height: The height of the rectangle to clear, in pixels
//
//	   The ClearRect() method clears the specified pixels within a given rectangle.
//
// Português:
//
//	Limpa os pixels especificados em um determinado retângulo.
//	   x: A coordenada x do canto superior esquerdo do retângulo para limpar;
//	   y: A coordenada y do canto superior esquerdo do retângulo para limpar;
//	   width: A largura do retângulo a ser limpo, em pixels;
//	   height: A altura do retângulo a ser limpo, em pixels.
//
//	   O método ClearRect() limpa os pixels especificados em um determinado retângulo.
func (el *TagCanvas) ClearRect(x, y, width, height int) (ref *TagCanvas) {
	el.context.Call("clearRect", x, y, width, height)
	return el
}

// Clip
//
// English:
//
//	Clips a region of any shape and size from the original canvas.
//
// The Clip() method clips a region of any shape and size from the original canvas.
//
//	Note:
//	  * Once a region is clipped, all future drawing will be limited to the clipped region (no
//	    access to other regions on the canvas). You can however save the current canvas region using
//	    the Save() method before using the Clip() method, and restore it (with the Restore() method)
//	    any time in the future.
//
// Português:
//
//	Recorta uma região de qualquer forma e tamanho do canvas original.
//
// O método Clip() recorta uma região de qualquer forma e tamanho do canvas original.
//
//	Nota:
//	  * Uma vez recortada uma região, todos os desenhos futuros serão limitados à região recortada
//	    (sem acesso a outras regiões do canvas). No entanto, você pode salvar a região do canvas
//	    atual usando o método Save() antes de usar o método Clip() e restaurá-la, com o método
//	    Restore(), a qualquer momento no futuro.
//
// todo: https://developer.mozilla.org/pt-BR/docs/Web/API/CanvasRenderingContext2D/clip
// todo: https://developer.mozilla.org/pt-BR/docs/Web/API/Canvas_API/Tutorial/Compositing
func (el *TagCanvas) Clip(x, y int) (ref *TagCanvas) {
	el.context.Call("clip", x, y)
	return el
}

// ClosePath
//
// English:
//
//	Creates a path from the current point back to the starting point.
//
// The ClosePath() method creates a path from the current point back to the starting point.
//
//	Note:
//	  * Use the Stroke() method to actually draw the path on the canvas;
//	  * Use the Fill() method to fill the drawing, black is default. Use the FillStyle() function to
//	    fill with another color/gradient.
//
// Português:
//
//	Cria um caminho do ponto atual de volta ao ponto inicial.
//
// O método ClosePath() cria um caminho do ponto atual de volta ao ponto inicial.
//
//	Nota:
//	  * Use o método Stroke() para realmente desenhar o caminho no canvas;
//	  * Use o método Fill() para preencher o desenho, preto é o padrão. Use a função
//	    FillStyle() para preencher com outra cor / gradiente.
//
// todo: fazer exemplo
func (el *TagCanvas) ClosePath(x, y int) (ref *TagCanvas) {
	el.context.Call("closePath", x, y)
	return el
}

// CreateLinearGradient
//
// English:
//
//	Creates a linear gradient.
//
//	 Input:
//	   x0: The x-coordinate of the start point of the gradient;
//	   y0: The y-coordinate of the start point of the gradient;
//	   x1: The x-coordinate of the end point of the gradient;
//	   y1: The y-coordinate of the end point of the gradient.
//
// The CreateLinearGradient() method creates a linear gradient object.
//
// The gradient can be used to fill rectangles, circles, lines, text, etc.
//
//	Note:
//	  * Use this object as the value to the strokeStyle or fillStyle properties; //todo: rever documentação
//	  * Use the AddColorStopPosition() method to specify different colors, and where to position the
//	    colors in the gradient object.
//
// Português:
//
//	Cria um gradiente linear.
//
//	 Entrada:
//	   x0: A coordenada x do ponto inicial do gradiente;
//	   y0: A coordenada y do ponto inicial do gradiente;
//	   x1: A coordenada x do ponto final do gradiente;
//	   y1: A coordenada y do ponto final do gradiente.
//
// O método CreateLinearGradient() cria um objeto gradiente linear.
//
// O gradiente pode ser usado para preencher retângulos, círculos, linhas, texto, etc.
//
//	Nota:
//	  * Use esta objeto como o valor para as propriedades StrokeStyle() ou FillStyle(); //todo: rever documentação
//	  * Use o método AddColorStopPosition() para especificar cores diferentes e onde posicionar as
//	    cores no objeto gradiente.
func (el *TagCanvas) CreateLinearGradient(x0, y0, x1, y1 int) (ref *TagCanvas) {
	el.gradient = el.context.Call("createLinearGradient", x0, y0, x1, y1)
	return el
}

// CreatePattern
//
// English:
//
//	Repeats a specified element in the specified direction.
//
//	 Input:
//	   image: Specifies the image, canvas, or video element of the pattern to use
//	   repeatedElement
//	   repeatRule: Image repetition rule.
//	     KRepeatRuleRepeat: (Default) The pattern repeats both horizontally and vertically
//	     KRepeatRuleRepeatX: The pattern repeats only horizontally
//	     KRepeatRuleRepeatY: The pattern repeats only vertically
//	     KRepeatRuleNoRepeat: The pattern will be displayed only once (no repeat)
//
// The CreatePattern() method repeats the specified element in the specified direction.
// The element can be an image, video, or another <canvas> element.
// The repeated element can be used to draw/fill rectangles, circles, lines etc.
//
// Português:
//
//	Repete um elemento especificado na direção especificada.
//
//	 Entrada:
//	   image: Especifica a imagem, tela ou elemento de vídeo do padrão para usar repeatElement;
//	   repeatRule: Regra de repetição da imagem.
//	     KRepeatRuleRepeat: (Padrão) O padrão se repete horizontal e verticalmente;
//	     KRepeatRuleRepeatX: O padrão se repete apenas horizontalmente;
//	     KRepeatRuleRepeatY: O padrão se repete apenas verticalmente;
//	     KRepeatRuleNoRepeat: O padrão será exibido apenas uma vez (sem repetição).
//
// O método CreatePattern() repete o elemento especificado na direção especificada.
//
// O elemento pode ser uma imagem, vídeo ou outro elemento <canvas>.
//
// O elemento repetido pode ser usado para desenhar retângulos, círculos, linhas etc.
func (el *TagCanvas) CreatePattern(image interface{}, repeatRule CanvasRepeatRule) (ref *TagCanvas) {
	if converted, ok := image.(*TagImg); ok {
		el.pattern = el.context.Call("createPattern", converted.Get(), repeatRule.String())
		return el
	}

	el.pattern = el.context.Call("createPattern", image, repeatRule.String())
	return el
}

// CreateRadialGradient
//
// English:
//
//	Creates a radial/circular gradient (to use on canvas content)
//
//	 Input:
//	   x0: The x-coordinate of the starting circle of the gradient;
//	   y0: The y-coordinate of the starting circle of the gradient;
//	   r0: The radius of the starting circle;
//	   x1: The x-coordinate of the ending circle of the gradient;
//	   y1: The y-coordinate of the ending circle of the gradient;
//	   r1: The radius of the ending circle.
//
// Português:
//
//	Cria um gradiente radial/circular (para usar no conteúdo do canvas)
//
//	 Entrada:
//	   x0: A coordenada x do círculo inicial do gradiente;
//	   y0: A coordenada y do círculo inicial do gradiente;
//	   r0: O raio do círculo inicial;
//	   x1: A coordenada x do círculo final do gradiente;
//	   y1: A coordenada y do círculo final do gradiente;
//	   r1: O raio do círculo final.
func (el *TagCanvas) CreateRadialGradient(x0, y0, r0, x1, y1 int, r1 float64) (ref *TagCanvas) {
	el.gradient = el.context.Call("createRadialGradient", x0, y0, r0, x1, y1, r1)
	return el
}

// FillRect
//
// English:
//
//	Draws a "filled" rectangle
//
//	 Input:
//	   x: The x-coordinate of the upper-left corner of the rectangle;
//	   y: The y-coordinate of the upper-left corner of the rectangle;
//	   width: The width of the rectangle, in pixels;
//	   height: The height of the rectangle, in pixels.
//
// The FillRect() method draws a "filled" rectangle. The default color of the fill is black.
//
//	Note:
//	  * Use the FillStyle() function to set a color, FillStyleGradient() to set a gradient or
//	    FillStylePattern() to set a pattern used to fill the drawing.
//
//	Example:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    FillRect(20, 20, 150, 100).
//	    AppendToStage()
//
// Português:
//
//	Desenha um retângulo "preenchido"
//
//	 Entrada:
//	   x: A coordenada x do canto superior esquerdo do retângulo;
//	   y: A coordenada y do canto superior esquerdo do retângulo;
//	   width: A largura do retângulo, em pixels;
//	   height: A altura do retângulo, em pixels.
//
// O método FillRect() desenha um retângulo "preenchido". A cor padrão do preenchimento é preto.
//
//	Nota:
//	  * Use a função FillStyle() para definir uma cor, FillStyleGradient() para definir um gradiente
//	    ou FillStylePattern() para definir um padrão usado para preencher o desenho.
//
//	Example:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    FillRect(20, 20, 150, 100).
//	    AppendToStage()
func (el *TagCanvas) FillRect(x, y, width, height int) (ref *TagCanvas) {
	el.context.Call("fillRect", x, y, width, height)
	return el
}

// FillText
//
// English:
//
//	Draws "filled" text on the canvas
//
//	 Input:
//	   text: Specifies the text that will be written on the canvas;
//	   x: The x coordinate where to start painting the text (relative to the canvas);
//	   y: The y coordinate where to start painting the text (relative to the canvas);
//	   maxWidth: Optional. The maximum allowed width of the text, in pixels.
//
// The FillText() method draws filled text on the canvas.
//
//	Note:
//	  * Use the Font() function to specify font and font size, and use the FillStyle() function to
//	    render the text in another color/gradient;
//	  * The default color of the text is black.
//
// Português:
//
//	Desenha o texto "preenchido" no canvas.
//
//	 Entrada:
//	   text: Especifica o texto que será escrito no canvas;
//	   x: A coordenada x onde começar a pintar o texto (em relação ao canvas)
//	   y: A coordenada y onde começar a pintar o texto (em relação ao canvas)
//	   maxWidth: A largura máxima permitida do texto, em pixels.
//
// O método FillText() desenha texto preenchido no canvas.
//
//	Nota:
//	  * Use a função Font() para especificar a fonte e o tamanho da fonte e use a função FillStyle()
//	    para renderizar o texto em outra cor/gradiente;
//	  * A cor padrão do texto é preto.
func (el *TagCanvas) FillText(text string, x, y, maxWidth int) (ref *TagCanvas) {
	el.context.Call("fillText", text, x, y, maxWidth)
	return el
}

// Font
//
// English:
//
//	Sets the current font properties for text content.
//
//	  Input:
//	    font-style: Specifies the font style. Use the factory factoryFontStyle.
//	      E.g.: factoryFontStyle.NewItalic()
//	    font-variant: Specifies the font variant. Use the factory factoryFontVariant.
//	      E.g.: factoryFontVariant.NewSmallCaps()
//	    font-weight: Specifies the font weight. Use the factory factoryFontWeight.
//	      E.g.: factoryFontWeight.NewBold()
//	    font-size/line-height: Specifies the font size and the line-height in pixels
//	    font-family: Specifies the font family. Use the factory factoryFontFamily.
//	      E.g.: factoryFontFamily.NewArial()
//	    caption: Use the font captioned controls (like buttons, drop-downs, etc.)
//	    icon: Use the font used to label icons
//	    menu: Use the font used in menus (drop-down menus and menu lists)
//	    message-box: Use the font used in dialog boxes
//	    small-caption: Use the font used for labeling small controls
//	    status-bar: Use the fonts used in window status bar
//
// The Font() function sets the current font properties for text content on the canvas.
//
// Default value: 10px sans-serif
//
// Português:
//
//	Define as propriedades de fonte atuais para conteúdo de texto.
//
//	  Entrada:
//	    font-style: Especifica o estilo da fonte. Usar a fábrica factoryFontStyle.
//	      Ex.: factoryFontStyle.NewItalic()
//	    font-variant: Especifica a variante da fonte. Usar a fábrica factoryFontVariant.
//	      Ex.: factoryFontVariant.NewSmallCaps()
//	    font-weight: Especifica o peso da fonte. Usar a fábrica factoryFontWeight.
//	      Ex.: factoryFontWeight.NewBold()
//	    font-size/line-height: Especifica o tamanho da fonte e a altura da linha em pixels
//	    font-family: Especifica a família de fontes. Usar a fábrica factoryFontFamily.
//	      Ex.: factoryFontFamily.NewArial()
//	    caption: Use os controles legendados de fonte (como botões, menus suspensos etc.)
//	    icon: Use a fonte usada para rotular os ícones
//	    menu: Use a fonte usada nos menus (menus suspensos e listas de menus)
//	    message-box: Use a fonte usada nas caixas de diálogo
//	    small-caption: Use a fonte usada para rotular pequenos controles
//	    status-bar: Use as fontes usadas na barra de status da janela
//
// # A função Font() define ou retorna as propriedades de fonte atuais para conteúdo de texto no canvas
//
// A propriedade font usa a mesma sintaxe que a propriedade font CSS.
//
// Valor padrão: 10px sem serifa
func (el *TagCanvas) Font(font Font) (ref *TagCanvas) {
	el.context.Set("font", font.String())
	return el
}

// GlobalAlpha
//
// English:
//
//	Sets the current alpha or transparency value of the drawing
//	   value: The transparency value. Must be a number between 0.0 (fully transparent) and 1.0
//	     (no transparency)
//
//	   Default value: 1.0
//
// Português:
//
//	Define o valor alfa ou transparência atual do desenho
//	   value: O valor da transparência. Deve ser um número entre 0.0 (totalmente transparente) e
//	     1.0 (sem transparência)
//
//	   Valor padrão: 1.0
func (el *TagCanvas) GlobalAlpha(value float64) (ref *TagCanvas) {
	el.context.Set("globalAlpha", value)
	return el
}

// GlobalCompositeOperation
//
// English:
//
//	Sets how a new image are drawn onto an existing image
//
//	 Input:
//	   value: how a source (new) image are drawn onto a destination image.
//
// The GlobalCompositeOperation() function sets how a source (new) image are drawn onto a destination
// (existing) image.
//
//	Note:
//	  * source image = drawings you are about to place onto the canvas;
//	  * destination image = drawings that are already placed onto the canvas;
//	  * Default value: source-over.
//
// Português:
//
//	Define como uma nova imagem é desenhada em uma imagem existente.
//
//	 Entrada:
//	   value: como uma imagem de origem (nova) é desenhada em uma imagem de destino.
//
// A função GlobalCompositeOperation() define como uma imagem de origem (nova) é desenhada em uma
// imagem de destino (existente).
//
//	Nota:
//	  * imagem de origem = desenhos que você está prestes a colocar no canvas;
//	  * imagem de destino = desenhos que já estão colocados no canvas;
//	  * Valor padrão: source-over.
func (el *TagCanvas) GlobalCompositeOperation(value CompositeOperationsRule) (ref *TagCanvas) {
	el.context.Set("globalCompositeOperation", value.String())
	return el
}

// Height
//
// English:
//
//	Returns the height of an ImageData object.
//
//	 Output:
//	   height: returns the height of an ImageData object, in pixels.
//
// Português:
//
//	Retorna a altura de um objeto ImageData.
//
//	 Saída:
//	   height: retorna à altura de um objeto ImageData, em pixels.
func (el *TagCanvas) Height() (height int) {
	return el.context.Get("height").Int()
}

// IsPointInPath
//
// English:
//
//	Returns true if the specified point is in the current path, otherwise false
//
//	 Input:
//	   x: The x-axis coordinate of the point to check.
//	   y: The y-axis coordinate of the point to check.
//	   fillRule: The algorithm by which to determine if a point is inside or outside the path.
//	        "nonzero": The non-zero winding rule. Default rule.
//	        "evenodd": The even-odd winding rule.
//
//	 Example:
//
//	   canvas := factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Rect(20, 20, 150, 100).
//	     AppendToStage()
//	     if canvas.IsPointInPath(20, 51, html.KFillRuleNonZero) {
//	       canvas.Stroke()
//	     }
//
// Português:
//
//	Retorna true se o ponto especificado estiver no caminho atual, caso contrário, false
//
//	 Entrada:
//	   x: A coordenada do eixo x do ponto a ser verificado.
//	   y: A coordenada do eixo y do ponto a ser verificado.
//	   fillRule: O algoritmo pelo qual determinar se um ponto está dentro ou fora do caminho.
//	        "nonzero": A regra de enrolamento diferente de zero. Regra padrão.
//	        "evenodd": A regra do enrolamento par-ímpar.
//
//	 Exemplo:
//
//	   canvas := factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Rect(20, 20, 150, 100).
//	     AppendToStage()
//	     if canvas.IsPointInPath(20, 51, html.KFillRuleNonZero) {
//	       canvas.Stroke()
//	     }
func (el *TagCanvas) IsPointInPath(x, y int, fillRule FillRule) (isPointInPath bool) {
	return el.context.Call("isPointInPath", x, y, fillRule.String()).Bool()
}

// LineCap
//
// English::
//
//	Sets the style of the end caps for a line
//
//	 Input:
//	   PlatformBasicType: style of the end caps for a line
//
//	 Note:
//	   * The value "round" and "square" make the lines slightly longer.
//	   * Default value: butt
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     BeginPath().
//	     LineCap(html.KCapRuleRound).
//	     LineWidth(5).
//	     MoveTo(20, 20).
//	     LineTo(20, 200).
//	     Stroke().
//	     AppendToStage()
//
// Português::
//
//	Define o estilo das terminações de uma linha
//
//	 Entrada:
//	   PlatformBasicType: estilo das tampas de extremidade para uma linha
//
//	 Nota:
//	   * O valor "redondo" e "quadrado" tornam as linhas um pouco mais longas.
//	   * Valor padrão: butt
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     BeginPath().
//	     LineCap(html.KCapRuleRound).
//	     LineWidth(5).
//	     MoveTo(20, 20).
//	     LineTo(20, 200).
//	     Stroke().
//	     AppendToStage()
func (el *TagCanvas) LineCap(value CapRule) (ref *TagCanvas) {
	el.context.Set("lineCap", value.String())
	return el
}

// LineJoin
//
// English:
//
//	Sets the type of corner created, when two lines meet
//
//	 Input:
//	   PlatformBasicType: type of corner created
//
// The LineJoin() function sets the type of corner created, when two lines meet.
//
//	Note:
//	  * The KJoinRuleMiter value is affected by the MiterLimit() function.
//	  * Default value: miter
//
//	Example:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    BeginPath().
//	    LineWidth(5).
//	    LineJoin(html.KJoinRuleRound).
//	    MoveTo(20, 20).
//	    LineTo(100, 50).
//	    LineTo(20, 100).
//	    Stroke().
//	    AppendToStage()
//
// Português:
//
//	Define o tipo de canto criado, quando duas linhas se encontram
//
//	 Entrada:
//	   PlatformBasicType: tipo de canto criado
//
// A função LineJoin() define o tipo de canto criado, quando duas linhas se encontram.
//
//	Nota:
//	  * O valor KJoinRuleMiter é afetado pela função MiterLimit().
//	  * Valor padrão: miter
//
//	Exemplo:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    BeginPath().
//	    LineWidth(5).
//	    LineJoin(html.KJoinRuleRound).
//	    MoveTo(20, 20).
//	    LineTo(100, 50).
//	    LineTo(20, 100).
//	    Stroke().
//	    AppendToStage()
func (el *TagCanvas) LineJoin(value JoinRule) (ref *TagCanvas) {
	el.context.Set("lineJoin", value.String())
	return el
}

// Stroke
//
// English:
//
//	The Stroke() method actually draws the path you have defined with all those MoveTo() and
//	LineTo() methods.
//
// The default color is black.
//
//	Note:
//	  * Use the StrokeStyle() function to draw with another color/gradient.
//
//	Example:
//
//	factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	  BeginPath().
//	  Arc(100, 75, 50, 0, 2 * math.Pi, false).
//	  Stroke().
//	  AppendToStage()
//
// Português:
//
//	A função Stroke() na verdade desenha o caminho que você definiu com todos os métodos MoveTo()
//	e LineTo().
//
// A cor padrão é preto.
//
//	Nota:
//	  * Use a Função StrokeStyle() para desenhar com outra cor/gradiente
//
//	Exemplo:
//
//	factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	  BeginPath().
//	  Arc(100, 75, 50, 0, 2 * math.Pi, false).
//	  Stroke().
//	  AppendToStage()
func (el *TagCanvas) Stroke() (ref *TagCanvas) {
	el.context.Call("stroke")
	return el
}

// MoveTo
//
// English:
//
//	Moves the path to the specified point in the canvas, without creating a line
//
//	 Input:
//	   x: The x-coordinate of where to move the path to
//	   y: The y-coordinate of where to move the path to
//
// The MoveTo() function moves the path to the specified point in the canvas, without creating a line.
//
//	Note:
//	  * Use the stroke() method to actually draw the path on the canvas.
//
//	Example:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    BeginPath().
//	    LineWidth(5).
//	    LineJoin(html.KJoinRuleRound).
//	    MoveTo(20, 20).
//	    LineTo(100, 50).
//	    LineTo(20, 100).
//	    Stroke().
//	    AppendToStage()
//
// Português:
//
//	Move o caminho para o ponto especificado no canvas, sem criar uma linha
//
//	 Entrada:
//	   x: A coordenada x de onde mover o caminho para
//	   y: A coordenada y de onde mover o caminho para
//
// A função MoveTo() move o caminho para o ponto especificado no canvas, sem criar uma linha.
//
//	Nota:
//	  * Use o método Stroke() para realmente desenhar o caminho no canvas.
//
//	Exemplo:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    BeginPath().
//	    LineWidth(5).
//	    LineJoin(html.KJoinRuleRound).
//	    MoveTo(20, 20).
//	    LineTo(100, 50).
//	    LineTo(20, 100).
//	    Stroke().
//	    AppendToStage()
func (el *TagCanvas) MoveTo(x, y int) (ref *TagCanvas) {
	el.context.Call("moveTo", x, y)
	return el
}

// LineTo
//
// English:
//
//	Adds a new point and creates a line from that point to the last specified point in the canvas
//
//	 Input:
//	   x: The x-coordinate of where to create the line to
//	   y: The y-coordinate of where to create the line to
//
//	 Note:
//	   * This method does not draw the line.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     BeginPath().
//	     LineWidth(5).
//	     LineJoin(html.KJoinRuleRound).
//	     MoveTo(20, 20).
//	     LineTo(100, 50).
//	     LineTo(20, 100).
//	     Stroke().
//	     AppendToStage()
//
// Português:
//
//	Adiciona um novo ponto e cria uma linha desse ponto até o último ponto especificado no canvas
//
//	 Entrada:
//	   x: A coordenada x de onde criar a linha para
//	   y: A coordenada y de onde criar a linha para
//
//	 Nota:
//	   * Este método não desenha a linha.
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     BeginPath().
//	     LineWidth(5).
//	     LineJoin(html.KJoinRuleRound).
//	     MoveTo(20, 20).
//	     LineTo(100, 50).
//	     LineTo(20, 100).
//	     Stroke().
//	     AppendToStage()
func (el *TagCanvas) LineTo(x, y int) (ref *TagCanvas) {
	el.context.Call("lineTo", x, y)
	return el
}

// LineWidth
//
// English:
//
//	Sets the current line width
//	   value: The current line width, in pixels
//
//	 Note:
//	   * Default value: 1
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     BeginPath().
//	     LineWidth(5).
//	     LineJoin(html.KJoinRuleRound).
//	     MoveTo(20, 20).
//	     LineTo(100, 50).
//	     LineTo(20, 100).
//	     Stroke().
//	     AppendToStage()
//
// Português:
//
//	Define a largura da linha atual
//	   value: A largura da linha atual, em pixels
//
//	 Nota:
//	   * Valor padrão: 1
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     BeginPath().
//	     LineWidth(5).
//	     LineJoin(html.KJoinRuleRound).
//	     MoveTo(20, 20).
//	     LineTo(100, 50).
//	     LineTo(20, 100).
//	     Stroke().
//	     AppendToStage()
func (el *TagCanvas) LineWidth(value int) (ref *TagCanvas) {
	el.context.Set("lineWidth", value)
	return el
}

// MeasureText
//
// English:
//
//	Returns an object that contains the width of the specified text
//
//	 Input:
//	   text: The text to be measured
//
//	 Example:
//
//	   var font html.Font
//	     font.Family = factoryFontFamily.NewArial()
//	     font.Size = 20
//
//	   canvas := factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(font)
//	     w := canvas.MeasureText("Hello Word!")
//	     wText := strconv.FormatInt(int64(w), 10)
//	     canvas.FillText("width:"+wText, 10, 50, 300).
//	     AppendToStage()
//
// Português:
//
//	Retorna um objeto que contém a largura do texto especificado
//
//	 Entrada:
//	   text: O texto a ser medido
//
//	 Exemplo:
//
//	   var font html.Font
//	     font.Family = factoryFontFamily.NewArial()
//	     font.Size = 20
//
//	   canvas := factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(font)
//	     w := canvas.MeasureText("Hello Word!")
//	     wText := strconv.FormatInt(int64(w), 10)
//	     canvas.FillText("width:"+wText, 10, 50, 300).
//	     AppendToStage()
func (el *TagCanvas) MeasureText(text string) (width int) {
	return el.context.Call("measureText", text).Get("width").Int()
}

// MiterLimit
//
// English:
//
//	Sets the maximum miter length
//
//	 Input:
//	   value: A positive number that specifies the maximum miter length.
//
//	 Note:
//	   * If the current miter length exceeds the MiterLimit(), the corner will display as
//	     LineJoin(KJoinRuleBevel);
//	   * The miter length is the distance between the inner corner and the outer corner where two
//	     lines meet;
//	   * Default value: 10
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     LineWidth(10).
//	     LineJoin(html.KJoinRuleMiter).
//	     MiterLimit(5).
//	     MoveTo(20, 20).
//	     LineTo(50, 27).
//	     LineTo(20, 34).
//	     Stroke().
//	     AppendToStage()
//
// Português:
//
//	Define ou retorna o comprimento máximo da mitra
//
//	 Entrada:
//	   value: Um número positivo que especifica o comprimento máximo da mitra.
//
//	 Nota:
//	   * Se o comprimento da mitra atual exceder o MiterLimit(), o canto será exibido como
//	     LineJoin(KJoinRuleBevel);
//	   * O comprimento da mitra é a distância entre o canto interno e o canto externo onde duas
//	     linhas se encontram;
//	   * Valor padrão: 10.
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     LineWidth(10).
//	     LineJoin(html.KJoinRuleMiter).
//	     MiterLimit(5).
//	     MoveTo(20, 20).
//	     LineTo(50, 27).
//	     LineTo(20, 34).
//	     Stroke().
//	     AppendToStage()
func (el *TagCanvas) MiterLimit(value int) (ref *TagCanvas) {
	el.context.Set("miterLimit", value)
	return el
}

// QuadraticCurveTo
//
// English:
//
//	Creates a quadratic Bézier curve.
//
//	 Input:
//	   cpx: The x-axis coordinate of the control point;
//	   cpy: The y-axis coordinate of the control point;
//	   x: The x-axis coordinate of the end point;
//	   y: The y-axis coordinate of the end point.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     BeginPath().
//	     MoveTo(20, 20).
//	     QuadraticCurveTo(20, 100, 200, 20).
//	     Stroke().
//	     AppendToStage()
//
// Português:
//
//	Cria uma curva Bézier quadrática.
//
//	 Entrada:
//	   cpx: A coordenada do eixo x do ponto de controle;
//	   cpy: A coordenada do eixo y do ponto de controle;
//	   x: A coordenada do eixo x do ponto final;
//	   y: A coordenada do eixo y do ponto final.
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     BeginPath().
//	     MoveTo(20, 20).
//	     QuadraticCurveTo(20, 100, 200, 20).
//	     Stroke().
//	     AppendToStage()
func (el *TagCanvas) QuadraticCurveTo(cpx, cpy, x, y int) (ref *TagCanvas) {
	el.context.Call("quadraticCurveTo", cpx, cpy, x, y)
	return el
}

// Rect
//
// English:
//
//	Creates a rectangle.
//
//	 Input:
//	   x: The x-coordinate of the upper-left corner of the rectangle;
//	   y: The y-coordinate of the upper-left corner of the rectangle;
//	   width: The width of the rectangle, in pixels;
//	   height: The height of the rectangle, in pixels.
//
//	 Note:
//	   * Use the Stroke() or Fill() functions to actually draw the rectangle on the canvas.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Rect(20, 20, 150, 100).
//	     Stroke().
//	     AppendToStage()
//
// Português:
//
//	Cria um retângulo.
//
//	 Entrada:
//	   x: A coordenada x do canto superior esquerdo do retângulo;
//	   y: A coordenada y do canto superior esquerdo do retângulo;
//	   width: A largura do retângulo, em pixels;
//	   height: A altura do retângulo, em pixels.
//
//	 Nota:
//	   * Use as funções Stroke() ou Fill() para realmente desenhar o retângulo no canvas.
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Rect(20, 20, 150, 100).
//	     Stroke().
//	     AppendToStage()
func (el *TagCanvas) Rect(x, y, width, height int) (ref *TagCanvas) {
	el.context.Call("rect", x, y, width, height)
	return el
}

// Restore
//
// English:
//
//	Returns previously saved path state and attributes.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(fontA).
//	     FillText("Hello World!", 10, 50, 300).
//	     Save().
//	     CreateLinearGradient(0, 0, 160, 0).
//	     AddColorStopPosition(0.0, factoryColor.NewMagenta()).
//	     AddColorStopPosition(0.5, factoryColor.NewBlue()).
//	     AddColorStopPosition(1.0, factoryColor.NewRed()).
//	     FillStyleGradient().
//	     Font(fontB).
//	     FillText("Big smile!", 10, 90, 300).
//	     Restore().
//	     FillText("Same font used before save", 10, 120, 300).
//	     AppendToStage()
//
// Português:
//
//	Retorna o estado e os atributos do caminho salvos anteriormente.
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(fontA).
//	     FillText("Hello World!", 10, 50, 300).
//	     Save().
//	     CreateLinearGradient(0, 0, 160, 0).
//	     AddColorStopPosition(0.0, factoryColor.NewMagenta()).
//	     AddColorStopPosition(0.5, factoryColor.NewBlue()).
//	     AddColorStopPosition(1.0, factoryColor.NewRed()).
//	     FillStyleGradient().
//	     Font(fontB).
//	     FillText("Big smile!", 10, 90, 300).
//	     Restore().
//	     FillText("Same font used before save", 10, 120, 300).
//	     AppendToStage()
func (el *TagCanvas) Restore() (ref *TagCanvas) {
	el.context.Call("restore")
	return el
}

// Rotate
//
// English:
//
//	Rotates the current drawing
//
//	 Input:
//	   angle: The rotation angle, in radians.
//
//	 Note:
//	   * To calculate from degrees to radians: degrees*math.Pi/180.
//	     Example: to rotate 5 degrees, specify the following: 5*math.Pi/180.
//	   * The rotation will only affect drawings made AFTER the rotation is done.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Rect(50, 20, 150, 100).
//	     Stroke().
//	     Rotate(20*math.Pi/180).
//	     Rect(50, 20, 150, 100).
//	     Stroke().
//	     AppendToStage()
//
// Português:
//
//	Gira o desenho atual.
//
//	 Entrada:
//	   angle: O ângulo de rotação, em radianos.
//
//	 Nota:
//	   * Para calcular de graus para radianos: graus*math.Pi/180.
//	     Exemplo: para girar 5 graus, especifique o seguinte: 5*math.Pi/180.
//	   * A rotação só afetará os desenhos feitos APÓS a rotação.
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Rect(50, 20, 150, 100).
//	     Stroke().
//	     Rotate(20*math.Pi/180).
//	     Rect(50, 20, 150, 100).
//	     Stroke().
//	     AppendToStage()
func (el *TagCanvas) Rotate(angle float64) (ref *TagCanvas) {
	el.context.Call("rotate", angle)
	return el
}

// Save
//
// English:
//
//	Saves the state of the current context.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(fontA).
//	     FillText("Hello World!", 10, 50, 300).
//	     Save().
//	     CreateLinearGradient(0, 0, 160, 0).
//	     AddColorStopPosition(0.0, factoryColor.NewMagenta()).
//	     AddColorStopPosition(0.5, factoryColor.NewBlue()).
//	     AddColorStopPosition(1.0, factoryColor.NewRed()).
//	     FillStyleGradient().
//	     Font(fontB).
//	     FillText("Big smile!", 10, 90, 300).
//	     Restore().
//	     FillText("Same font used before save", 10, 120, 300).
//	     AppendToStage()
//
// Português:
//
//	Salva o estado do contexto atual.
//
//	Exemplo:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    Font(fontA).
//	    FillText("Hello World!", 10, 50, 300).
//	    Save().
//	    CreateLinearGradient(0, 0, 160, 0).
//	    AddColorStopPosition(0.0, factoryColor.NewMagenta()).
//	    AddColorStopPosition(0.5, factoryColor.NewBlue()).
//	    AddColorStopPosition(1.0, factoryColor.NewRed()).
//	    FillStyleGradient().
//	    Font(fontB).
//	    FillText("Big smile!", 10, 90, 300).
//	    Restore().
//	    FillText("Same font used before save", 10, 120, 300).
//	    AppendToStage()
func (el *TagCanvas) Save() (ref *TagCanvas) {
	el.context.Call("save")
	return el
}

// Scale
//
// English:
//
//	Scales the current drawing bigger or smaller.
//
//	 Input:
//	   scaleWidth: Scales the width of the current drawing (1.0=100%, 0.5=50%, 2.0=200%, etc.)
//	   scaleHeight: Scales the height of the current drawing (1.0=100%, 0.5=50%, 2.0=200%, etc.)
//
//	 Note:
//	   * If you scale a drawing, all future drawings will also be scaled;
//	   * The positioning will also be scaled. If you scale(2.0,2.0); drawings will be positioned
//	     twice as far from the left and top of the canvas as you specify.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     StrokeRect(5, 5, 25, 15).
//	     Scale(2.0, 6.0).
//	     StrokeRect(5, 5, 25, 15).
//	     AppendToStage()
//
// Português:
//
//	Dimensiona o desenho atual para maior ou menor.
//
//	 Entrada:
//	   scaleWidth: Dimensiona a largura do desenho atual (1.0=100%, 0.5=50%, 2.0=200%, etc.)
//	   scaleHeight: Dimensiona a altura do desenho atual (1.0=100%, 0.5=50%, 2.0=200%, etc.)
//
//	 Nota:
//	   * Se você dimensionar um desenho, todos os desenhos futuros também serão dimensionados;
//	   * O posicionamento também será dimensionado. Se você dimensionar (2.0, 2.0); os desenhos serão
//	     posicionados duas vezes mais distantes da esquerda e do topo do canvas conforme
//	     você especificar.
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     StrokeRect(5, 5, 25, 15).
//	     Scale(2.0, 6.0).
//	     StrokeRect(5, 5, 25, 15).
//	     AppendToStage()
func (el *TagCanvas) Scale(scaleWidth, scaleHeight float64) (ref *TagCanvas) {
	el.context.Call("scale", scaleWidth, scaleHeight)
	return el
}

// SetTransform
//
// English:
//
//	Resets the current transform to the identity matrix.
//
//	 Input:
//	   a: Scales the drawings horizontally;
//	   b: Skews the drawings horizontally;
//	   c: Skews the drawings vertically;
//	   d: Scales the drawings vertically;
//	   e: Moves the the drawings horizontally;
//	   f: Moves the the drawings vertically.
//
//	 Note:
//	   * Each object on the canvas has a current transformation matrix.
//	     The SetTransform() function resets the current transform to the identity matrix, and then
//	     put transform data on GetLastTransform() function.
//	     In other words, the SetTransform() function lets you scale, rotate, move, and skew the
//	     current context.
//	   * The transformation will only affect drawings made after the SetTransform() function is
//	     called.
//	   * You can use the Save() and Restore() functions to archive the original transform parameters.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     FillStyle(factoryColor.NewYellow()).
//	     FillRect(50, 50, 250, 100).
//	     SetTransform(1.0, 0.5, -0.5, 1.0, 30.0, 10.0).
//	     FillStyle(factoryColor.NewRed()).
//	     FillRect(50, 50, 250, 100).
//	     SetTransform(1.0, 0.5, -0.5, 1.0, 30.0, 10.0).
//	     FillStyle(factoryColor.NewBlue()).
//	     FillRect(50, 50, 230, 70).
//	     AppendToStage()
//
// Português:
//
//	Redefine a transformação atual para a matriz de identidade.
//
//	 Entrada:
//	   a: Dimensiona os desenhos horizontalmente;
//	   b: Inclina os desenhos horizontalmente;
//	   c: Inclina os desenhos verticalmente;
//	   d: Dimensiona os desenhos verticalmente;
//	   e: Move os desenhos horizontalmente;
//	   f: Move os desenhos verticalmente.
//
//	 Nota:
//	   * Cada objeto no canvas tem uma matriz de transformação atual.
//	     A função SetTransform() redefine a transformação atual para a matriz de identidade e, em
//	     seguida, coloca os dados de transformação na função GetLastTransform().
//	     Em outras palavras, o método SetTransform() permite dimensionar, girar, mover e inclinar o
//	     contexto atual.
//	   * A transformação só afetará os desenhos feitos após a chamada da função SetTransform().
//	   * Você pode usar as funções Save() e Restore() para arquivar os parâmetros de transform
//	     originais.
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     FillStyle(factoryColor.NewYellow()).
//	     FillRect(50, 50, 250, 100).
//	     SetTransform(1.0, 0.5, -0.5, 1.0, 30.0, 10.0).
//	     FillStyle(factoryColor.NewRed()).
//	     FillRect(50, 50, 250, 100).
//	     SetTransform(1.0, 0.5, -0.5, 1.0, 30.0, 10.0).
//	     FillStyle(factoryColor.NewBlue()).
//	     FillRect(50, 50, 230, 70).
//	     AppendToStage()
func (el *TagCanvas) SetTransform(a, b, c, d, e, f float64) (ref *TagCanvas) {
	el.transform = el.context.Call("setTransform", a, b, c, d, e, f)
	return el
}

// StrokeRect
//
// English:
//
//	Draws a rectangle (no fill)
//
//	 Input:
//	   x: The x-coordinate of the upper-left corner of the rectangle;
//	   y: The y-coordinate of the upper-left corner of the rectangle;
//	   width: The width of the rectangle, in pixels;
//	   height: The height of the rectangle, in pixels;
//
//	 Note:
//	   * The default color of the stroke is black;
//	   * Use the StrokeStyle() function to set a color, CreateRadialGradient(), CreateLinearGradient()
//	     or CreatePattern() to style the stroke.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     StrokeRect(5, 5, 25, 15).
//	     AppendToStage()
//
// Português:
//
//	Desenha um retângulo (sem preencher)
//
//	 entrada:
//	   x: A coordenada x do canto superior esquerdo do retângulo;
//	   y: A coordenada y do canto superior esquerdo do retângulo;
//	   width: A largura do retângulo, em pixels;
//	   height: A altura do retângulo, em pixels;
//
//	 Nota:
//	   * A cor padrão do traço é preto;
//	   * Use a função StrokeStyle() para definir uma cor, CreateRadialGradient(),
//	     CreateLinearGradient() ou CreatePattern() para estilizar o traço.
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     StrokeRect(5, 5, 25, 15).
//	     AppendToStage()
func (el *TagCanvas) StrokeRect(x, y, width, height int) (ref *TagCanvas) {
	el.context.Call("strokeRect", x, y, width, height)
	return el
}

// StrokeStyle
//
// English:
//
//	Sets the color, gradient, or pattern used for strokes.
//
//	 Input:
//	   value: The style must be the color name in textual form, such as "red" or "green", or a color.RGBA value.
//
//	 Note:
//	   * The default color is black.
//
//	 Example:
//
//	   var colorArc color.RGBA
//	   canvas := factoryBrowser.NewTagCanvas("canvas_0", 800, 600)
//	   for i := 0.0; i != 6.0; i += 1.0 {
//	     for j := 0.0; j != 6.0; j += 1.0 {
//	       colorArc.R = 0
//	       colorArc.G = uint8(255.0 - 42.5*i)
//	       colorArc.B = uint8(255.0 - 42.5*j)
//	       colorArc.A = 255
//	       canvas.StrokeStyle(colorArc).
//	         BeginPath().
//	         Arc(int(12.5+j*25.0), int(12.5+i*25.0), 10.0, 0.0, math.Pi*2.0, true).
//	         Stroke()
//	     }
//	   }
//	   canvas.AppendToStage()
//
// Português:
//
//	Define a cor, gradiente ou padrão usado para traçados.
//
//	 Entrada:
//	   value: O estilo deve ser o nome da cor na forma textual, como "red" ou "green", ou um valor color.RGBA.
//
//	 Note:
//	   * A cor padrão é preta.
//
//	 Exemplo:
//
//	   var colorArc color.RGBA
//	   canvas := factoryBrowser.NewTagCanvas("canvas_0", 800, 600)
//	   for i := 0.0; i != 6.0; i += 1.0 {
//	     for j := 0.0; j != 6.0; j += 1.0 {
//	       colorArc.R = 0
//	       colorArc.G = uint8(255.0 - 42.5*i)
//	       colorArc.B = uint8(255.0 - 42.5*j)
//	       colorArc.A = 255
//	       canvas.StrokeStyle(colorArc).
//	         BeginPath().
//	         Arc(int(12.5+j*25.0), int(12.5+i*25.0), 10.0, 0.0, math.Pi*2.0, true).
//	         Stroke()
//	     }
//	   }
//	   canvas.AppendToStage()
func (el *TagCanvas) StrokeStyle(value interface{}) (ref *TagCanvas) {
	if converted, ok := value.(color.RGBA); ok {
		el.context.Set("strokeStyle", RGBAToJs(converted))
		return el
	}

	el.context.Set("strokeStyle", value)
	return el
}

// StrokeStyleGradient
//
//	Sets javascript's strokeStyle property after using CreateLinearGradient() or
//	CreateRadialGradient() functions.
//
//	 Example:
//
//	   var fontA html.Font
//	   fontA.Family = factoryFontFamily.NewArial()
//	   fontA.Style = factoryFontStyle.NewItalic()
//	   fontA.Size = 20
//
//	   var fontB html.Font
//	   fontB.Family = factoryFontFamily.NewVerdana()
//	   fontB.Size = 35
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(fontA).
//	     StrokeText("Hello World!", 10, 50, 300).
//	     CreateLinearGradient(0, 0, 160, 0).
//	     AddColorStopPosition(0.0, factoryColor.NewMagenta()).
//	     AddColorStopPosition(0.5, factoryColor.NewBlue()).
//	     AddColorStopPosition(1.0, factoryColor.NewRed()).
//	     StrokeStyleGradient().
//	     Font(fontB).
//	     StrokeText("Big smile!", 10, 90, 300).
//	     AppendToStage()
//
// Português:
//
//	Define a propriedade strokeStyle do javascript depois de usar as funções CreateLinearGradient()
//	ou CreateRadialGradient().
//
//	 Exemplo:
//
//	   var fontA html.Font
//	   fontA.Family = factoryFontFamily.NewArial()
//	   fontA.Style = factoryFontStyle.NewItalic()
//	   fontA.Size = 20
//
//	   var fontB html.Font
//	   fontB.Family = factoryFontFamily.NewVerdana()
//	   fontB.Size = 35
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(fontA).
//	     StrokeText("Hello World!", 10, 50, 300).
//	     CreateLinearGradient(0, 0, 160, 0).
//	     AddColorStopPosition(0.0, factoryColor.NewMagenta()).
//	     AddColorStopPosition(0.5, factoryColor.NewBlue()).
//	     AddColorStopPosition(1.0, factoryColor.NewRed()).
//	     StrokeStyleGradient().
//	     Font(fontB).
//	     StrokeText("Big smile!", 10, 90, 300).
//	     AppendToStage()
func (el *TagCanvas) StrokeStyleGradient() (ref *TagCanvas) {
	el.context.Set("strokeStyle", el.gradient)
	return el
}

// StrokeText
//
// English:
//
//	Draws text on the canvas (no fill)
//
//	 Input:
//	   text: Specifies the text that will be written on the canvas
//	   x: The x coordinate where to start painting the text (relative to the canvas)
//	   y: The y coordinate where to start painting the text (relative to the canvas)
//	   maxWidth: The maximum allowed width of the text, in pixels
//
//	 Note:
//	   * The default color of the text is black.
//	   * Use the Font() function to specify font and font size, and use the StrokeStyle() function to
//	     render the text in another color/gradient.
//
//	 Example:
//
//	   var fontA html.Font
//	   fontA.Family = factoryFontFamily.NewArial()
//	   fontA.Style = factoryFontStyle.NewItalic()
//	   fontA.Size = 20
//
//	   var fontB html.Font
//	   fontB.Family = factoryFontFamily.NewVerdana()
//	   fontB.Size = 35
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(fontA).
//	     StrokeText("Hello World!", 10, 50, 300).
//	     CreateLinearGradient(0, 0, 160, 0).
//	     AddColorStopPosition(0.0, factoryColor.NewMagenta()).
//	     AddColorStopPosition(0.5, factoryColor.NewBlue()).
//	     AddColorStopPosition(1.0, factoryColor.NewRed()).
//	     StrokeStyleGradient().
//	     Font(fontB).
//	     StrokeText("Big smile!", 10, 90, 300).
//	     AppendToStage()
//
// Português:
//
//	Desenha o texto no canvas (sem preenchimento)
//
//	 Entrada:
//	   text: Especifica o texto a ser desenhado no canvas;
//	   x: A coordenada X de onde iniciar a pintura do texto (relativa ao canvas)
//	   y: A coordenada Y de onde iniciar a pintura do texto (relativa ao canvas)
//	   maxWidth: A largura máxima permitida do texto, em pixels.
//
//	 Nota:
//	   * A cor padrão é preto.
//	   * Use a função Font() para especificar a fonte e o tamanho do texto, e use a função
//	     StrokeStyle() para renderizar o texto em outra cor/gradiente.
//
//	 Exemplo:
//
//	   var fontA html.Font
//	   fontA.Family = factoryFontFamily.NewArial()
//	   fontA.Style = factoryFontStyle.NewItalic()
//	   fontA.Size = 20
//
//	   var fontB html.Font
//	   fontB.Family = factoryFontFamily.NewVerdana()
//	   fontB.Size = 35
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(fontA).
//	     StrokeText("Hello World!", 10, 50, 300).
//	     CreateLinearGradient(0, 0, 160, 0).
//	     AddColorStopPosition(0.0, factoryColor.NewMagenta()).
//	     AddColorStopPosition(0.5, factoryColor.NewBlue()).
//	     AddColorStopPosition(1.0, factoryColor.NewRed()).
//	     StrokeStyleGradient().
//	     Font(fontB).
//	     StrokeText("Big smile!", 10, 90, 300).
//	     AppendToStage()
func (el *TagCanvas) StrokeText(text string, x, y, maxWidth int) (ref *TagCanvas) {
	el.context.Call("strokeText", text, x, y, maxWidth)
	return el
}

// TextAlign
//
// English:
//
//	Sets the current alignment for text content
//
//	 Input:
//	   value: the anchor point.
//
// Normally, the text will START in the position specified, however, if you set TextAlign(html.KFontAlignRuleRight) and place the text in position 150, it means that the text should END in position 150.
//
//	Note:
//	  * Use the FillText() or the StrokeText() function to actually draw and position the text on the canvas.
//	  * Default value: html.KFontAlignRuleStart
//
//	Example:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    StrokeStyle(factoryColor.NewRed()).
//	    MoveTo(150, 20).
//	    LineTo(150, 170).
//	    Stroke().
//	    Font(font).
//	    TextAlign(html.KFontAlignRuleStart).
//	    FillText("textAlign = start", 150, 60, 400).
//	    TextAlign(html.KFontAlignRuleEnd).
//	    FillText("textAlign = end", 150, 80, 400).
//	    TextAlign(html.KFontAlignRuleEnd).
//	    FillText("textAlign = end", 150, 80, 400).
//	    TextAlign(html.KFontAlignRuleLeft).
//	    FillText("textAlign = left", 150, 100, 400).
//	    TextAlign(html.KFontAlignRuleCenter).
//	    FillText("textAlign = center", 150, 120, 400).
//	    TextAlign(html.KFontAlignRuleRight).
//	    FillText("textAlign = right", 150, 140, 400).
//	    AppendToStage()
//
// Português:
//
//	Define o alinhamento atual do texto.
//
//	 Entrada:
//	   value: o ponto da âncora.
//
// Normalmente, o texto COMEÇARÁ na posição especificada, no entanto, se você definir
// TextAlign(html.KFontAlignRuleRight) e colocar o texto na posição 150, significa que o texto deve
// TERMINAR na posição 150.
//
//	NotA:
//	  * Use a função FillText() ou StrokeText() para realmente desenhar e posicionar o texto no
//	    canvas;
//	  * Valor padrão: html.KFontAlignRuleStart
//
//	Exemplo:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    StrokeStyle(factoryColor.NewRed()).
//	    MoveTo(150, 20).
//	    LineTo(150, 170).
//	    Stroke().
//	    Font(font).
//	    TextAlign(html.KFontAlignRuleStart).
//	    FillText("textAlign = start", 150, 60, 400).
//	    TextAlign(html.KFontAlignRuleEnd).
//	    FillText("textAlign = end", 150, 80, 400).
//	    TextAlign(html.KFontAlignRuleEnd).
//	    FillText("textAlign = end", 150, 80, 400).
//	    TextAlign(html.KFontAlignRuleLeft).
//	    FillText("textAlign = left", 150, 100, 400).
//	    TextAlign(html.KFontAlignRuleCenter).
//	    FillText("textAlign = center", 150, 120, 400).
//	    TextAlign(html.KFontAlignRuleRight).
//	    FillText("textAlign = right", 150, 140, 400).
//	    AppendToStage()
func (el *TagCanvas) TextAlign(value FontAlignRule) (ref *TagCanvas) {
	el.context.Set("textAlign", value.String())
	return el
}

// TextBaseline
//
// English:
//
//	Sets the current text baseline used when drawing text.
//
//	 Input:
//	   PlatformBasicType: text baseline used when drawing text.
//
//	 Note:
//	   * The FillText() and StrokeText() functions will use the specified TextBaseline() value when positioning the text on the canvas.
//	   Default value: alphabetic
//
//	 Example:
//
//	   var font html.Font
//	   font.Family = factoryFontFamily.NewArial()
//	   font.Size = 20
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     StrokeStyle(factoryColor.NewRed()).
//	     MoveTo(5, 100).
//	     LineTo(395, 100).
//	     Stroke().
//	     Font(font).
//	     TextBaseline(html.KTextBaseLineRuleTop).
//	     FillText("Top", 5, 100, 300).
//	     TextBaseline(html.KTextBaseLineRuleBottom).
//	     FillText("Bottom", 50, 100, 300).
//	     TextBaseline(html.KTextBaseLineRuleMiddle).
//	     FillText("Middle", 120, 100, 300).
//	     TextBaseline(html.KTextBaseLineRuleAlphabetic).
//	     FillText("Alphabetic", 190, 100, 300).
//	     TextBaseline(html.KTextBaseLineRuleHanging).
//	     FillText("Hanging", 290, 100, 300).
//	     AppendToStage()
//
// Português:
//
//	Define a linha de base usada para desenhar o texto.
//
//	 Entrada:
//	   PlatformBasicType: linha de base usada para desenhar o texto.
//
//	 Nota:
//	   * As funções FillText() e StrokeText() vão usar a linha de base especificada pela função TextBaseline() antes de posicionar o texto no canvas.
//	   * Valor padrão: html.KTextBaseLineRuleAlphabetic
//
//	 Exemplo:
//
//	   var font html.Font
//	   font.Family = factoryFontFamily.NewArial()
//	   font.Size = 20
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     StrokeStyle(factoryColor.NewRed()).
//	     MoveTo(5, 100).
//	     LineTo(395, 100).
//	     Stroke().
//	     Font(font).
//	     TextBaseline(html.KTextBaseLineRuleTop).
//	     FillText("Top", 5, 100, 300).
//	     TextBaseline(html.KTextBaseLineRuleBottom).
//	     FillText("Bottom", 50, 100, 300).
//	     TextBaseline(html.KTextBaseLineRuleMiddle).
//	     FillText("Middle", 120, 100, 300).
//	     TextBaseline(html.KTextBaseLineRuleAlphabetic).
//	     FillText("Alphabetic", 190, 100, 300).
//	     TextBaseline(html.KTextBaseLineRuleHanging).
//	     FillText("Hanging", 290, 100, 300).
//	     AppendToStage()
func (el *TagCanvas) TextBaseline(value TextBaseLineRule) (ref *TagCanvas) {
	el.context.Set("textBaseline", value.String())
	return el
}

// Transform
//
// English:
//
//	Replaces the current transformation matrix for the drawing
//
//	 Input:
//	   a: Scales the drawing horizontally;
//	   b: Skew the the drawing horizontally;
//	   c: Skew the the drawing vertically;
//	   d: Scales the drawing vertically;
//	   e: Moves the the drawing horizontally;
//	   f: Moves the the drawing vertically.
//
// Each object on the canvas has a current transformation matrix.
// The Transform() method replaces the current transformation matrix. It multiplies the current
// transformation matrix with the matrix described by:
//
//	 a | c | e
//	---+---+---
//	 b | d | f
//	---+---+---
//	 0 | 0 | 1
//
// In other words, the Transform() method lets you scale, rotate, move, and skew the current context.
//
//	Note:
//	  * The transformation will only affect drawings made after the Transform() method is called;
//	  * The Transform() function behaves relatively to other transformations made by Rotate(),
//	    Scale(), Translate(), or Transform().
//	      Example: If you already have set your drawing to scale by two, and the Transform() method
//	      scales your drawings by two, your drawings will now scale by four.
//	  * Check out the SetTransform() method, which does not behave relatively to other
//	    transformations.
//	  * You can use the Save() and Restore() functions to archive the original transform parameters.
//
//	Example:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    FillStyle(factoryColor.NewYellow()).
//	    FillRect(50, 50, 250, 100).
//	    Transform(1.0, 0.5, -0.5, 1.0, 30.0, 10.0).
//	    FillStyle(factoryColor.NewRed()).
//	    FillRect(50, 50, 250, 100).
//	    Transform(1.0, 0.5, -0.5, 1.0, 30.0, 10.0).
//	    FillStyle(factoryColor.NewBlue()).
//	    FillRect(50, 50, 230, 70).
//	    AppendToStage()
//
// Português:
//
//	Substitui a matriz de transformação atual para o desenho
//
//	 Entrada:
//	   a: Dimensiona o desenho horizontalmente;
//	   b: Inclinar o desenho horizontalmente;
//	   c: Inclinar o desenho verticalmente;
//	   d: Dimensiona o desenho verticalmente;
//	   e: Move o desenho horizontalmente;
//	   f: Move o desenho verticalmente.
//
// Cada objeto no canvas tem uma matriz de transformação atual.
// O método Transform() substitui a matriz de transformação atual. Ele multiplica a matriz de
// transformação atual com a matriz descrita por:
//
//	 a | c | e
//	---+---+---
//	 b | d | f
//	---+---+---
//	 0 | 0 | 1
//
// Em outras palavras, o método Transform() permite dimensionar, girar, mover e inclinar o contexto
// atual.
//
//	Nota:
//	  * A transformação só afetará os desenhos feitos depois que o método Transform() for chamado;
//	  * A função Transform() se comporta relativamente a outras transformações feitas por Rotate(),
//	    Scale(), Translate() ou Transform().
//	    Exemplo: Se você já configurou seu desenho para dimensionar em dois, e o método Transform()
//	      dimensiona seus desenhos em dois, seus desenhos agora serão dimensionados em quatro.
//	  * Confira o método SetTransform(), que não se comporta em relação a outras transformações.
//	  * Você pode usar as funções Save() e Restore() para arquivar os parâmetros de transform
//	    originais.
//
//	Exemplo:
//
//	  factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	    FillStyle(factoryColor.NewYellow()).
//	    FillRect(50, 50, 250, 100).
//	    Transform(1.0, 0.5, -0.5, 1.0, 30.0, 10.0).
//	    FillStyle(factoryColor.NewRed()).
//	    FillRect(50, 50, 250, 100).
//	    Transform(1.0, 0.5, -0.5, 1.0, 30.0, 10.0).
//	    FillStyle(factoryColor.NewBlue()).
//	    FillRect(50, 50, 230, 70).
//	    AppendToStage()
func (el *TagCanvas) Transform(a, b, c, d, e, f float64) (ref *TagCanvas) {
	el.context.Call("transform", a, b, c, d, e, f)
	return el
}

// Translate
//
// English:
//
//	Remaps the (0,0) position on the canvas
//
//	 Input:
//	   x: The value to add to horizontal (x) coordinates;
//	   y: The value to add to vertical (y) coordinates.
//
//	 Note:
//	   * When you call a method like FillRect() after Translate(), the value is added to the x and y
//	     coordinate values for all new canvas elements.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     FillRect(10, 10, 100, 50).
//	     Translate(70, 70).
//	     FillRect(10, 10, 100, 50).
//	     AppendToStage()
//
// Português:
//
//	Remapeia a posição (0,0) no canvas
//
//	 Entrada:
//	   x: O valor a ser adicionado às coordenadas horizontais (x);
//	   y: O valor a ser adicionado às coordenadas verticais (y).
//
//	 Nota:
//	   * Quando você chama um método como FillRect() após Translate(), o valor é adicionado aos
//	     valores das coordenadas x e y para todos os novos elementos do canvas.
//
//	 Exemplo:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     FillRect(10, 10, 100, 50).
//	     Translate(70, 70).
//	     FillRect(10, 10, 100, 50).
//	     AppendToStage()
func (el *TagCanvas) Translate(x, y int) (ref *TagCanvas) {
	el.context.Call("translate", x, y)
	return el
}

// FillStyle
//
// English:
//
//	Sets the color, gradient, or pattern used to fill the drawing.
//
//	 Input:
//	   value: the color color.RGBA. You can use factoryColor.NewColorName() to make it easier;
//
//	 Note:
//	   * The default color is black.
//
// Português:
//
//	Define a cor, gradiente ou padrão usado para preencher o desenho.
//
//	 Entrada:
//	   value: a cor color.RGBA. Você pode usar factoryColor.NewColorName() para facilitar;
//
//	 Nota:
//	   * A cor padrão é preto.
func (el *TagCanvas) FillStyle(value interface{}) (ref *TagCanvas) {
	if converted, ok := value.(color.RGBA); ok {
		el.context.Set("fillStyle", RGBAToJs(converted))
		return el
	}

	el.context.Set("fillStyle", value)
	return el
}

// FillStyleGradient
//
// English:
//
//	Sets javascript's fillStyle property after using CreateLinearGradient() or CreateRadialGradient()
//	functions.
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(fontA).
//	     FillText("Hello World!", 10, 50, 300).
//	     Save().
//	     CreateLinearGradient(0, 0, 160, 0).
//	     AddColorStopPosition(0.0, factoryColor.NewMagenta()).
//	     AddColorStopPosition(0.5, factoryColor.NewBlue()).
//	     AddColorStopPosition(1.0, factoryColor.NewRed()).
//	     FillStyleGradient().
//	     Font(fontB).
//	     FillText("Big smile!", 10, 90, 300).
//	     Restore().
//	     FillText("Same font used before save", 10, 120, 300).
//	     AppendToStage()
//
// Português:
//
//	Define a propriedade fillStyle do javascript depois de usar as funções CreateLinearGradient() ou
//	CreateRadialGradient().
//
//	 Example:
//
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     Font(fontA).
//	     FillText("Hello World!", 10, 50, 300).
//	     Save().
//	     CreateLinearGradient(0, 0, 160, 0).
//	     AddColorStopPosition(0.0, factoryColor.NewMagenta()).
//	     AddColorStopPosition(0.5, factoryColor.NewBlue()).
//	     AddColorStopPosition(1.0, factoryColor.NewRed()).
//	     FillStyleGradient().
//	     Font(fontB).
//	     FillText("Big smile!", 10, 90, 300).
//	     Restore().
//	     FillText("Same font used before save", 10, 120, 300).
//	     AppendToStage()
func (el *TagCanvas) FillStyleGradient() (ref *TagCanvas) {
	el.context.Set("fillStyle", el.gradient)
	return el
}

// FillStylePattern
//
// English:
//
//	Sets the javascript's fillStyle property after using the CreatePattern() function.
//
//	 Example
//
//	   var img = factoryBrowser.NewTagImage(
//	     "spacecraft",
//	     "./small.png",
//	     29,
//	     50,
//	     true,
//	   )
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     CreatePattern(img, html.KRepeatRuleRepeat).
//	     Rect(0, 0, 300, 300).
//	     FillStylePattern().
//	     Fill().
//	     AppendToStage()
//
// Português:
//
//	Define a propriedade fillStyle do javascript após o uso da função CreatePattern().
//
//	 Exemplo:
//
//	   var img = factoryBrowser.NewTagImage(
//	     "spacecraft",
//	     "./small.png",
//	     29,
//	     50,
//	     true,
//	   )
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     CreatePattern(img, html.KRepeatRuleRepeat).
//	     Rect(0, 0, 300, 300).
//	     FillStylePattern().
//	     Fill().
//	     AppendToStage()
func (el *TagCanvas) FillStylePattern() (ref *TagCanvas) {
	el.context.Set("fillStyle", el.pattern)
	return el
}

// Fill
//
// English:
//
//	Fills the current drawing (path)
//
//	 Note:
//	   * The default color is black.
//	   * Use the FillStyle() function to fill with another color/gradient.
//	   * If the path is not closed, the Fill() method will add a line from the last point to the
//	     startpoint of the path to close the path (like ClosePath()), and then fill the path.
//
//	 Example:
//
//	   var img = factoryBrowser.NewTagImage(
//	     "spacecraft",
//	     "./small.png",
//	     29,
//	     50,
//	     true,
//	   )
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     CreatePattern(img, html.KRepeatRuleRepeat).
//	     Rect(0, 0, 300, 300).
//	     FillStylePattern().
//	     Fill().
//	     AppendToStage()
//
// Português:
//
//	Preenche o desenho atual (caminho)
//
//	 Nota:
//	   * A cor padrão é preto.
//	   * Use a função FillStyle() para preencher com outro gradiente de cor.
//	   * Se o caminho não estiver fechado, o método Fill() adicionará uma linha do último ponto ao
//	     ponto inicial do caminho para fechar o caminho (como ClosePath()) e, em seguida, preencherá
//	     o caminho.
//
//	 Exemplo:
//
//	   var img = factoryBrowser.NewTagImage(
//	     "spacecraft",
//	     "./small.png",
//	     29,
//	     50,
//	     true,
//	   )
//	   factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
//	     CreatePattern(img, html.KRepeatRuleRepeat).
//	     Rect(0, 0, 300, 300).
//	     FillStylePattern().
//	     Fill().
//	     AppendToStage()
func (el *TagCanvas) Fill() (ref *TagCanvas) {
	el.context.Call("fill")
	return el
}

// SetXY #replicar
//
// English:
//
//	Sets the X and Y axes in pixels.
//
// Português:
//
//	Define os eixos X e Y em pixels.
func (el *TagCanvas) SetXY(x, y int) (ref *TagCanvas) {
	// fixme: copiado parcialmente da tag div
	px := strconv.FormatInt(int64(x), 10) + "px"
	py := strconv.FormatInt(int64(y), 10) + "px"

	el.selfElement.Get("style").Set("left", px)
	el.selfElement.Get("style").Set("top", py)
	el.selfElement.Get("style").Set("position", "absolute")

	return el
}

func (el *TagCanvas) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (el *TagCanvas) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (e *TagCanvas) AddListenerAbort(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAbort(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerAbort() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerAbort()
	return e
}

func (e *TagCanvas) AddListenerAuxclick(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAuxclick(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerAuxclick() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerAuxclick()
	return e
}

func (e *TagCanvas) AddListenerBeforeinput(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeinput(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerBeforeinput() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerBeforeinput()
	return e
}

func (e *TagCanvas) AddListenerBeforematch(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforematch(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerBeforematch() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerBeforematch()
	return e
}

func (e *TagCanvas) AddListenerBeforetoggle(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforetoggle(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerBeforetoggle() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerBeforetoggle()
	return e
}

func (e *TagCanvas) AddListenerCancel(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCancel(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerCancel() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerCancel()
	return e
}

func (e *TagCanvas) AddListenerCanplay(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplay(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerCanplay() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerCanplay()
	return e
}

func (e *TagCanvas) AddListenerCanplaythrough(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCanplaythrough(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerCanplaythrough() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerCanplaythrough()
	return e
}

func (e *TagCanvas) AddListenerChange(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerChange(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerChange() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerChange()
	return e
}

func (e *TagCanvas) AddListenerClick(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClick(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerClick() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerClick()
	return e
}

func (e *TagCanvas) AddListenerClose(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerClose(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerClose() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerClose()
	return e
}

func (e *TagCanvas) AddListenerContextlost(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextlost(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerContextlost() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerContextlost()
	return e
}

func (e *TagCanvas) AddListenerContextmenu(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextmenu(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerContextmenu() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerContextmenu()
	return e
}

func (e *TagCanvas) AddListenerContextrestored(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerContextrestored(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerContextrestored() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerContextrestored()
	return e
}

func (e *TagCanvas) AddListenerCopy(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCopy(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerCopy() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerCopy()
	return e
}

func (e *TagCanvas) AddListenerCuechange(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCuechange(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerCuechange() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerCuechange()
	return e
}

func (e *TagCanvas) AddListenerCut(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerCut(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerCut() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerCut()
	return e
}

func (e *TagCanvas) AddListenerDblclick(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDblclick(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerDblclick() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerDblclick()
	return e
}

func (e *TagCanvas) AddListenerDrag(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrag(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerDrag() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerDrag()
	return e
}

func (e *TagCanvas) AddListenerDragend(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragend(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerDragend() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerDragend()
	return e
}

func (e *TagCanvas) AddListenerDragenter(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragenter(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerDragenter() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerDragenter()
	return e
}

func (e *TagCanvas) AddListenerDragleave(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragleave(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerDragleave() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerDragleave()
	return e
}

func (e *TagCanvas) AddListenerDragover(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragover(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerDragover() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerDragover()
	return e
}

func (e *TagCanvas) AddListenerDragstart(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDragstart(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerDragstart() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerDragstart()
	return e
}

func (e *TagCanvas) AddListenerDrop(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDrop(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerDrop() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerDrop()
	return e
}

func (e *TagCanvas) AddListenerDurationchange(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerDurationchange(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerDurationchange() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerDurationchange()
	return e
}

func (e *TagCanvas) AddListenerEmptied(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEmptied(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerEmptied() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerEmptied()
	return e
}

func (e *TagCanvas) AddListenerEnded(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerEnded(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerEnded() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerEnded()
	return e
}

func (e *TagCanvas) AddListenerFormdata(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFormdata(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerFormdata() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerFormdata()
	return e
}

func (e *TagCanvas) AddListenerInput(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInput(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerInput() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerInput()
	return e
}

func (e *TagCanvas) AddListenerInvalid(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerInvalid(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerInvalid() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerInvalid()
	return e
}

func (e *TagCanvas) AddListenerKeydown(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeydown(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerKeydown() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerKeydown()
	return e
}

func (e *TagCanvas) AddListenerKeypress(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeypress(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerKeypress() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerKeypress()
	return e
}

func (e *TagCanvas) AddListenerKeyup(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerKeyup(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerKeyup() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerKeyup()
	return e
}

func (e *TagCanvas) AddListenerLoadeddata(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadeddata(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerLoadeddata() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerLoadeddata()
	return e
}

func (e *TagCanvas) AddListenerLoadedmetadata(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadedmetadata(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerLoadedmetadata() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerLoadedmetadata()
	return e
}

func (e *TagCanvas) AddListenerLoadstart(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoadstart(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerLoadstart() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerLoadstart()
	return e
}

func (e *TagCanvas) AddListenerMousedown(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousedown(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerMousedown() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerMousedown()
	return e
}

func (e *TagCanvas) AddListenerMouseenter(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseenter(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerMouseenter() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerMouseenter()
	return e
}

func (e *TagCanvas) AddListenerMouseleave(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseleave(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerMouseleave() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerMouseleave()
	return e
}

func (e *TagCanvas) AddListenerMousemove(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMousemove(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerMousemove() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerMousemove()
	return e
}

func (e *TagCanvas) AddListenerMouseout(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseout(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerMouseout() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerMouseout()
	return e
}

func (e *TagCanvas) AddListenerMouseover(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseover(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerMouseover() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerMouseover()
	return e
}

func (e *TagCanvas) AddListenerMouseup(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMouseup(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerMouseup() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerMouseup()
	return e
}

func (e *TagCanvas) AddListenerPaste(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPaste(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerPaste() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerPaste()
	return e
}

func (e *TagCanvas) AddListenerPause(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPause(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerPause() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerPause()
	return e
}

func (e *TagCanvas) AddListenerPlay(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlay(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerPlay() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerPlay()
	return e
}

func (e *TagCanvas) AddListenerPlaying(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPlaying(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerPlaying() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerPlaying()
	return e
}

func (e *TagCanvas) AddListenerProgress(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerProgress(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerProgress() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerProgress()
	return e
}

func (e *TagCanvas) AddListenerRatechange(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRatechange(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerRatechange() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerRatechange()
	return e
}

func (e *TagCanvas) AddListenerReset(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReset(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerReset() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerReset()
	return e
}

func (e *TagCanvas) AddListenerScrollend(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScrollend(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerScrollend() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerScrollend()
	return e
}

func (e *TagCanvas) AddListenerSecuritypolicyviolation(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSecuritypolicyviolation(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerSecuritypolicyviolation() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerSecuritypolicyviolation()
	return e
}

func (e *TagCanvas) AddListenerSeeked(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeked(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerSeeked() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerSeeked()
	return e
}

func (e *TagCanvas) AddListenerSeeking(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSeeking(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerSeeking() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerSeeking()
	return e
}

func (e *TagCanvas) AddListenerSelect(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSelect(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerSelect() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerSelect()
	return e
}

func (e *TagCanvas) AddListenerSlotchange(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSlotchange(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerSlotchange() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerSlotchange()
	return e
}

func (e *TagCanvas) AddListenerStalled(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStalled(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerStalled() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerStalled()
	return e
}

func (e *TagCanvas) AddListenerSubmit(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSubmit(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerSubmit() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerSubmit()
	return e
}

func (e *TagCanvas) AddListenerSuspend(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerSuspend(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerSuspend() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerSuspend()
	return e
}

func (e *TagCanvas) AddListenerTimeupdate(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerTimeupdate(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerTimeupdate() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerTimeupdate()
	return e
}

func (e *TagCanvas) AddListenerToggle(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerToggle(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerToggle() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerToggle()
	return e
}

func (e *TagCanvas) AddListenerVolumechange(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVolumechange(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerVolumechange() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerVolumechange()
	return e
}

func (e *TagCanvas) AddListenerWaiting(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWaiting(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerWaiting() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerWaiting()
	return e
}

func (e *TagCanvas) AddListenerWebkitanimationend(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationend(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerWebkitanimationend() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerWebkitanimationend()
	return e
}

func (e *TagCanvas) AddListenerWebkitanimationiteration(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationiteration(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerWebkitanimationiteration() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerWebkitanimationiteration()
	return e
}

func (e *TagCanvas) AddListenerWebkitanimationstart(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkitanimationstart(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerWebkitanimationstart() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerWebkitanimationstart()
	return e
}

func (e *TagCanvas) AddListenerWebkittransitionend(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWebkittransitionend(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerWebkittransitionend() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerWebkittransitionend()
	return e
}

func (e *TagCanvas) AddListenerWheel(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerWheel(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerWheel() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerWheel()
	return e
}

func (e *TagCanvas) AddListenerBlur(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBlur(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerBlur() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerBlur()
	return e
}

func (e *TagCanvas) AddListenerError(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerError(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerError() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerError()
	return e
}

func (e *TagCanvas) AddListenerFocus(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerFocus(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerFocus() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerFocus()
	return e
}

func (e *TagCanvas) AddListenerLoad(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLoad(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerLoad() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerLoad()
	return e
}

func (e *TagCanvas) AddListenerResize(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerResize(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerResize() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerResize()
	return e
}

func (e *TagCanvas) AddListenerScroll(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerScroll(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerScroll() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerScroll()
	return e
}

func (e *TagCanvas) AddListenerAfterprint(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerAfterprint(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerAfterprint() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerAfterprint()
	return e
}

func (e *TagCanvas) AddListenerBeforeprint(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeprint(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerBeforeprint() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerBeforeprint()
	return e
}

func (e *TagCanvas) AddListenerBeforeunload(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerBeforeunload(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerBeforeunload() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerBeforeunload()
	return e
}

func (e *TagCanvas) AddListenerHashchange(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerHashchange(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerHashchange() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerHashchange()
	return e
}

func (e *TagCanvas) AddListenerLanguagechange(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerLanguagechange(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerLanguagechange() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerLanguagechange()
	return e
}

func (e *TagCanvas) AddListenerMessage(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessage(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerMessage() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerMessage()
	return e
}

func (e *TagCanvas) AddListenerMessageerror(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerMessageerror(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerMessageerror() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerMessageerror()
	return e
}

func (e *TagCanvas) AddListenerOffline(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOffline(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerOffline() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerOffline()
	return e
}

func (e *TagCanvas) AddListenerOnline(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerOnline(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerOnline() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerOnline()
	return e
}

func (e *TagCanvas) AddListenerPageswap(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageswap(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerPageswap() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerPageswap()
	return e
}

func (e *TagCanvas) AddListenerPagehide(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagehide(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerPagehide() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerPagehide()
	return e
}

func (e *TagCanvas) AddListenerPagereveal(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPagereveal(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerPagereveal() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerPagereveal()
	return e
}

func (e *TagCanvas) AddListenerPageshow(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPageshow(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerPageshow() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerPageshow()
	return e
}

func (e *TagCanvas) AddListenerPopstate(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerPopstate(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerPopstate() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerPopstate()
	return e
}

func (e *TagCanvas) AddListenerRejectionhandled(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerRejectionhandled(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerRejectionhandled() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerRejectionhandled()
	return e
}

func (e *TagCanvas) AddListenerStorage(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerStorage(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerStorage() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerStorage()
	return e
}

func (e *TagCanvas) AddListenerUnhandledrejection(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnhandledrejection(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerUnhandledrejection() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerUnhandledrejection()
	return e
}

func (e *TagCanvas) AddListenerUnload(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerUnload(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerUnload() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerUnload()
	return e
}

func (e *TagCanvas) AddListenerReadystatechange(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerReadystatechange(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerReadystatechange() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerReadystatechange()
	return e
}

func (e *TagCanvas) AddListenerVisibilitychange(genericEvent chan generic.Data) (ref *TagCanvas) {
	e.commonEvents.selfElement = &e.selfElement
	e.commonEvents.AddListenerVisibilitychange(genericEvent)
	return e
}

func (e *TagCanvas) RemoveListenerVisibilitychange() (ref *TagCanvas) {
	e.commonEvents.RemoveListenerVisibilitychange()
	return e
}
