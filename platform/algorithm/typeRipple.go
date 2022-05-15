package algorithm

import (
	"math"
)

type ripple struct {
	Geometry
}

// GenerateRipple
//
// English:
//
//  Generates a sine ripple in the list of processed points, losing the original points
//
//   Input:
//     distance: maximum ripple distance, wavelength;
//     ripples: amounts of ripples;
//     processed: pointer of the processed points list.
//
// Português:
//
//  Gera uma ondulação senoidal na lista de pontos processados, perdendo os pontos originais
//
//   Entrada:
//     distance: distância máxima da ondulação, ou comprimento de onda;
//     ripples: quantidades de ondulações;
//     processed: ponteiro da lista de pontos processados.
func (e *ripple) generateRipple(distance float64, ripples int, processed *[]Point) {
	tmp := make([]Point, len(*processed))
	angle := 0.0

	l := len(*processed) - 1
	for i := 0; i != l; i += 1 {
		percent := float64(i) / float64(l)

		if i == l {
			angle = e.AngleBetweenTwoPoints((*processed)[0], (*processed)[l])
		} else {
			angle = e.AngleBetweenTwoPoints((*processed)[i], (*processed)[i+1])
		}

		angle += math.Pi / 2

		proportion := math.Pi * float64(ripples) * percent
		offset := math.Sin(proportion) * distance

		tmp[i] = e.NewPointByDistance((*processed)[i], offset, angle)
	}

	// esta linha carrega o último ponto, ou ele fica com (x,y) = (0,0)
	tmp[l] = e.NewPointByDistance((*processed)[l], 0, angle)
	copy(*processed, tmp)
}
