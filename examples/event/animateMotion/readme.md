# Dynamic event / Evento dinâmico

Este exemplo mostra como calcular a posição da imagem em uma animação.
![screen](./screen.png)

### O conceito de imagem

Normalmente, as imagens são definidas em dois grupos, as imagens formadas por píxeis e as imagens formadas por vetores.

No caso da imagem formada por píxeis, a imagem é formada por uma série de pontos, ou píxel, onde a soma de todos os 
pontos formam a imagem.

A principal característica dessa imagem, é o fato dela ter um tamanho fixo, dificultando o redimencionamento da mesma 
para um tamanho maior, onde a imagem sofreria perda de resolução.

No caso da imagem vetorial, ela é formada por vetores, informações matemáticas com as intruções de como desenhar a 
imagem.

Diferente das imagens formadas por píxeis, a imagem é formada matematicamente e ela não perde qualidade quando é 
redimencionada.

### O problema do svg

O problema está na forma como o navegador trata a imagem SVG, ele simplesmente estica a imagem ao maior tamanho 
possível para encaixar na tela do navegador, caso o tamanho do svg não tenha sido definido na criação da imagem. 

#### Veja dois exemplos abaixo

Nesse primeiro exemplo, foi definido o tamanho da **viewBox**, a caixa onde a imagem será contida, mas, como a imagem 
está sem tamanho definido, a **viewBox** apenas definirá a proporção da imagem em 2:1.

```go
  factoryBrowser.NewTagSvg().ViewBox([]float64{0,0,400,200}).Append( ... )
```

No segundo exemplo, o tamanho da imagem foi definido em 400:200 e isto forçará o navegador a respeitar o tamanho da 
imagem.

```go
  factoryBrowser.NewTagSvg().ViewBox([]float64{0,0,400,200}).Width(400).Height(200).Append( ... )
```

O problema dessas duas abordagens está na forma como o navegador retorna o comando **getBoundingClientRect()** usado de 
forma transparente por vários comandos, por exemplo, **GetRight()**.

Imagine um svg com um cículo dentro, onde foi definido um raio de 50 píxels, mas, o svg principal não tem tamanho 
definido, como no primeiro exemplo, ou seja, o navegador vai redimencionar o svg principal para encaixar na tela, mas, o 
comando **GetRight()**, relativo ao círculo, sempre retornará o tamanho definido para o círculo originalmente, 
independente de como ele foi redimencionada pelo navegador.

Por isto, o exemplo contém a seguinte construção abaixo:
```go
// Referencia a tag svg para uso posterior.
var container *html.TagSvg

// Fator de redimencionamento da imagem.
var factor = 1.0

// Largura do svg principal, usada nos cálculos.
var width = 400.0

// Forma de definir o tamanho do svg principal para poder se usar cálculos dinâmicos.
factoryBrowser.NewTagSvg().Reference(&container).ViewBox([]float64{0, 0, width, 200})...

// Adiciona uma função de cálculo de baixa prioridade, 10x por segundo.
stage.AddHighLatencyFunctions(func() {
	// captura o tamanho atual do svg principal e divide pelo tamanho previsto, criando o fator de tamanho da imagem
  factor = (container.GetRight() - container.GetX()) / width
})

// Adiciona uma função de cálculo para ser executada com alta prioridade, até 120x por seundo.
stage.AddDrawFunctions(func() {
	// calcula o ângulo entre o centro da seta e a bola vermelha da imagem,
	// tomando o cuidado de redimencionar x e y pelo fator de tamanho da imagem atual
	// nota: primeiro, multiplicação e divisão, depois soma e subtração.
  angle := math.Atan2(120-circle.GetY()/factor, 95-circle.GetX()/factor)
	
	// Go trabalha com ângulo em radianos e 'angle*180/math.Pi' transforma o ângulo em graus decimais
	// -90˚ corrige o local para onde a seta aponta. Dependendo da linguagem de programação, 0˚ pode ser o centro acima ou 
	// o centro a direita/esquerda da imagem.
	// nota: primeiro, multiplicação e divisão, depois soma e subtração.
  svgG.Transform(factoryBrowser.NewTransform().Rotate(angle*180/math.Pi-90, 25, 25))
  
	// Desenha uma linha entre o centro da bola vermelha e o centro da seta.
	// Caso a seta e a linha fiquem fora de alinhamento, o centro de um dos elementos foi definido erradamente.
	line.X1(100).Y1(125).X2(circle.GetX()/factor + 5).Y2(circle.GetY()/factor + 5)
})
```

