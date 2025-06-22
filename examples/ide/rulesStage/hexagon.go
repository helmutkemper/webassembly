package rulesStage

import (
	"fmt"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/hexagon"
)

type GridAdjust interface {
	// AdjustCenter recalculates and returns the adjusted pixel coordinates (cx, cy) for the center of a hexagon based on inputs (x, y).
	AdjustCenter(x, y int) (cx, cy int)
}

// Hexagon represents a hexagonal grid cell in both pixel and coordinate space.
// It captures various properties such as position, layout, and cube coordinates.
type Hexagon struct {
	// Column and row in the doubled coordinate system
	col, row int

	// X and Y coordinates in 2D space
	cx, cy rulesDensity.Density

	// Layout defines how hexagons are positioned and transformed on the screen.
	//
	//   - Orientation specifies whether the hexes are pointy-topped or flat-topped,
	//     along with the transformation matrices for converting between hex and pixel space.
	//   - Size defines the radius (width and height) of a single hexagon in screen units (pixels).
	//   - Origin specifies the pixel coordinate where the hex grid origin (0,0,0) will be drawn.
	layout hexagon.Layout

	// Point represents a 2D position in pixel or screen space.
	//
	// It is typically used to store the result of hex-to-pixel conversion,
	// define layout sizes, or calculate hex corner positions for rendering.
	//
	//   - X: horizontal position.
	//   - Y: vertical position.
	point hexagon.Point

	// Hex represents a cube coordinate (q, r, s) used to model positions in a hexagonal grid.
	//
	// Cube coordinates satisfy the constraint q + r + s = 0, allowing for unambiguous positioning
	// and efficient computation of directions, distances, and neighbors.
	//
	//   - Q: corresponds to the "column" axis.
	//   - R: corresponds to the "row" axis.
	//   - S: the third axis, derived from Q and R (S = -Q - R).
	hex hexagon.Hex
}

// Init initializes the hexagon with its layout based on the specified origin (x, y) and size.
func (e *Hexagon) Init(x, y, size rulesDensity.Density) {
	e.layout = hexagon.Layout{
		Orientation: hexagon.LayoutFlat,
		Size:        hexagon.Point{X: float64(size), Y: float64(size)},
		Origin:      hexagon.Point{X: float64(x), Y: float64(y)},
	}
}

// GetColRow returns the column and row indices of the hexagon in the grid coordinate system.
func (e *Hexagon) GetColRow() (col, row int) {
	return e.col, e.row
}

// GetCenter returns the pixel coordinates (x, y) of the center of the hexagon.
func (e *Hexagon) GetCenter() (x, y rulesDensity.Density) {
	return e.cx, e.cy
}

// AdjustCenter recalculates and returns the adjusted pixel coordinates (cx, cy) for the center of a hexagon based on inputs (x, y).
func (e *Hexagon) AdjustCenter(x, y int) (cx, cy int) {
	hex := e.colHexToRow(hexagon.Point{X: float64(x), Y: float64(y)})
	point := hexagon.HexToPixel(e.layout, hex)
	return int(point.X), int(point.Y)
}

// SetPixelXY sets the hexagon's column and row based on the provided pixel coordinates (x, y).
func (e *Hexagon) SetPixelXY(x, y rulesDensity.Density) {
	hex := e.colHexToRow(hexagon.Point{X: float64(x), Y: float64(y)})
	cord := hexagon.QDoubledFromCube(hex)
	e.SetRowCol(cord.Col, cord.Row)
}

// SetRowCol updates the hexagon's column and row indices and triggers conversion of coordinates and layout adjustments.
func (e *Hexagon) SetRowCol(col, row int) {
	e.col = col
	e.row = row
	e.convertManager(e.col, e.row)
}

// colHexToRow converts a 2D pixel position (Point) to a hexagonal grid coordinate (Hex) based on the instance layout.
func (e *Hexagon) colHexToRow(point hexagon.Point) (hex hexagon.Hex) {
	return hexagon.PixelToHex(e.layout, point)
}

// colRowToHex converts a column and row from a doubled coordinate system to a cube coordinate (hexagon.Hex).
func (e *Hexagon) colRowToHex(col, row int) (hex hexagon.Hex) {
	return hexagon.QDoubledToCube(hexagon.DoubledCoordinate{Col: col, Row: row})
}

// convertManager recalculates and updates hexagon's attributes based on provided column and row indices in the grid system.
func (e *Hexagon) convertManager(col, row int) {
	e.col = col
	e.row = row
	e.hex = e.colRowToHex(col, row)
	e.point = hexagon.HexToPixel(e.layout, e.hex)
	e.cx = rulesDensity.Density(e.point.X)
	e.cy = rulesDensity.Density(e.point.Y)
}

// GetPath generates a path for the hexagon's outline as a series of SVG-compatible commands based on its corners.
func (e *Hexagon) GetPath() (path []string) {
	points := hexagon.PolygonCorners(e.layout, e.hex)
	for k, point := range points {
		if k == 0 {
			path = append(path, fmt.Sprintf("M %.2f,%.2f ", point.X, point.Y))
			continue
		}

		path = append(path, fmt.Sprintf("L %.2f,%.2f ", point.X, point.Y))
	}
	path = append(path, "z")
	return
}

// GetPoints returns the 2D integer coordinates of the hexagon's corners based on its layout and position.
func (e *Hexagon) GetPoints() (points [][2]rulesDensity.Density) {
	ps := hexagon.PolygonCorners(e.layout, e.hex)
	points = make([][2]rulesDensity.Density, len(ps))
	for k, point := range ps {
		points[k] = [2]rulesDensity.Density{rulesDensity.Density(point.X), rulesDensity.Density(point.Y)}
	}

	return
}
