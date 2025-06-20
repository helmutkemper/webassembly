// Package hexagon provides a complete implementation of hexagonal grid math,
// including cube, offset, and doubled coordinate systems, distance and direction
// calculations, layout transformations, and support for rendering via pixel-space
// conversion.
//
// This package is based on the Red Blob Games guide to hex grids:
// https://www.redblobgames.com/grids/hexagons/
//
// It supports both pointy-topped and flat-topped hex layouts and is suitable
// for games, simulations, map editors, and visualizations.
package hexagon

import (
	"errors"
	"math"
)

// Point represents a 2D position in pixel or screen space.
//
// It is typically used to store the result of hex-to-pixel conversion,
// define layout sizes, or calculate hex corner positions for rendering.
//
//   - X: horizontal position.
//   - Y: vertical position.
type Point struct {
	X float64 // X coordinate in 2D space
	Y float64 // Y coordinate in 2D space
}

// Hex represents a cube coordinate (q, r, s) used to model positions in a hexagonal grid.
//
// Cube coordinates satisfy the constraint q + r + s = 0, allowing for unambiguous positioning
// and efficient computation of directions, distances, and neighbors.
//
//   - Q: corresponds to the "column" axis.
//   - R: corresponds to the "row" axis.
//   - S: the third axis, derived from Q and R (S = -Q - R).
type Hex struct {
	Q int // Column coordinate
	R int // Row coordinate
	S int // Derived third coordinate (enforced by q + r + s = 0)
}

// NewHex creates a new Hex using the given integer cube coordinates (q, r, s).
//
// It validates that the cube coordinate constraint q + r + s == 0 is satisfied,
// which is a requirement for valid hexagonal grid positions.
//
// If the constraint is violated, the function returns an error.
func NewHex(q, r, s int) (hex Hex, err error) {
	if q+r+s != 0 {
		return Hex{}, errors.New("q + r + s must be 0")
	}
	return Hex{Q: q, R: r, S: s}, nil
}

// FractionalHex represents a cube coordinate using floating-point values.
//
// This type is used for intermediate or in-between positions on the hex grid,
// such as during interpolation (e.g., when drawing a line between two hexes).
//
// It maintains the same constraint as integer cube coordinates: q + r + s ≈ 0,
// but allows for non-integer values, enabling smooth transitions and animations.
type FractionalHex struct {
	Q, R, S float64 // Cube coordinates (non-integer)
}

// NewFractionalHex creates a new FractionalHex with the given floating-point cube coordinates (q, r, s).
//
// It validates that the coordinates approximately satisfy the constraint q + r + s ≈ 0,
// which is required for all valid hex coordinates in cube space.
//
// This function is useful for intermediate values (e.g., during interpolation),
// which are later rounded to the nearest integer hex using HexRound().
func NewFractionalHex(q, r, s float64) (FractionalHex, error) {
	if math.Round(q+r+s) != 0 {
		return FractionalHex{}, errors.New("q + r + s must be 0")
	}
	return FractionalHex{Q: q, R: r, S: s}, nil
}

// OffsetCoordinate represents a hex coordinate in the offset coordinate system.
//
// This system uses a standard (row, column) grid with staggered rows or columns,
// depending on whether the offset is "even" or "odd" (defined by constants EVEN/ODD).
//
// Common offset layouts include:
//   - even-q / odd-q: columns are aligned, and rows are offset.
//   - even-r / odd-r: rows are aligned, and columns are offset.
//
// This format is useful for integrating hex grids with typical 2D array data structures.
type OffsetCoordinate struct {
	Col, Row int // Column and row in the offset coordinate system
}

// DoubledCoordinate represents a hex coordinate using the doubled coordinate system.
//
// This system separates adjacent rows or columns by two units instead of one,
// allowing for simple indexing and storage of hex grids in a 2D array.
//
// There are two types of doubled coordinates:
//   - q-doubled: even columns (Col) spaced by 2, rows (Row) unchanged.
//   - r-doubled: even rows (Row) spaced by 2, columns (Col) unchanged.
type DoubledCoordinate struct {
	Col, Row int // Column and row in the doubled coordinate system
}

// AxialCoordinate represents axial coordinates (q, r)
type AxialCoordinate struct {
	Q, R int
}

// Orientation defines the transformation matrices and angle offset
// used to convert between hex coordinates (cube) and 2D pixel space.
//
//   - F0–F3 define the forward transformation matrix (hex → pixel).
//   - B0–B3 define the inverse transformation matrix (pixel → hex).
//   - StartAngle defines the angle (in 1/6ths of a full circle) to rotate the hex,
//     used when computing corner positions for drawing (e.g., 0.5 = 30 degrees).
type Orientation struct {
	F0, F1, F2, F3 float64 // Forward transform coefficients (hex → pixel)
	B0, B1, B2, B3 float64 // Inverse transform coefficients (pixel → hex)
	StartAngle     float64 // Starting angle in 1/6 turns (used for corner calculation)
}

// Layout defines how hexagons are positioned and transformed on the screen.
//
//   - Orientation specifies whether the hexes are pointy-topped or flat-topped,
//     along with the transformation matrices for converting between hex and pixel space.
//   - Size defines the radius (width and height) of a single hexagon in screen units (pixels).
//   - Origin specifies the pixel coordinate where the hex grid origin (0,0,0) will be drawn.
type Layout struct {
	Orientation  Orientation // Determines the orientation and transformation matrices
	Size, Origin Point       // Size of each hexagon and position of the grid origin
}

const (
	// EVEN (1): Used for even-q or even-r offset coordinates.
	//
	// EVEN and ODD are constants for offset grid types.
	//
	// Constants representing the offset layout modes for hex grids.
	//
	// These are used in functions that convert between offset and cube coordinates.
	EVEN = 1

	// ODD  (-1): Used for odd-q or odd-r offset coordinates.
	//
	// EVEN and ODD are constants for offset grid types.
	//
	// Constants representing the offset layout modes for hex grids.
	//
	// These are used in functions that convert between offset and cube coordinates.
	ODD = -1
)

// LayoutPointy defines the orientation matrix for pointy-topped hexagons.
//
// The F* values form the forward transformation matrix (hex → pixel).
//
// The B* values form the inverse transformation matrix (pixel → hex).
//
// StartAngle defines the initial rotation offset (in hex corner generation).
//
// This orientation is used when hexagons are arranged with a vertex (point) facing up.
var LayoutPointy = Orientation{
	F0: math.Sqrt(3),     // horizontal scale factor for q
	F1: math.Sqrt(3) / 2, // horizontal scale factor for r
	F3: 3.0 / 2.0,        // vertical scale factor for r

	B0: math.Sqrt(3) / 3, // inverse horizontal factor for q
	B1: -1.0 / 3.0,       // inverse horizontal factor for r
	B3: 2.0 / 3.0,        // inverse vertical factor for r

	StartAngle: 0.5, // offset used when computing hex corner positions (in radians)
}

// LayoutFlat defines the orientation matrix for flat-topped hexagons.
//
// The F* values form the forward transformation matrix (hex → pixel).
//
// The B* values form the inverse transformation matrix (pixel → hex).
//
// This orientation is used when hexagons are arranged with their flat sides on the top and bottom.
var LayoutFlat = Orientation{
	F0: 3.0 / 2.0, // horizontal scale factor for q
	F1: 0.0,
	F2: math.Sqrt(3) / 2.0, // vertical scale factor for q
	F3: math.Sqrt(3),       // vertical scale factor for r

	B0:         2.0 / 3.0, // inverse horizontal factor for q
	B1:         0.0,
	B2:         -1.0 / 3.0,         // inverse horizontal factor for r
	B3:         math.Sqrt(3) / 3.0, // inverse vertical factor for r
	StartAngle: 0.0,
}

// HexAdd returns the result of adding two hex coordinates `a` and `b`.
//
// This is a vector addition in cube coordinates and is used for operations
// like moving to neighbors or composing movement vectors.
func HexAdd(a, b Hex) Hex {
	return Hex{Q: a.Q + b.Q, R: a.R + b.R, S: a.S + b.S}
}

// HexSubtract returns the result of subtracting hex `b` from hex `a`.
//
// This produces a vector (in cube coordinates) that represents the offset from `b` to `a`.
//
// Commonly used to compute distance or direction between two hexes.
func HexSubtract(a, b Hex) Hex {
	return Hex{Q: a.Q - b.Q, R: a.R - b.R, S: a.S - b.S}
}

// HexScale multiplies each component of the hex coordinate `a` by a scalar factor `k`.
//
// This scales the hex outward from the origin, effectively moving it `k` steps in its direction.
//
// Useful for extending a direction vector or generating hex rings.
func HexScale(a Hex, k int) Hex {
	return Hex{Q: a.Q * k, R: a.R * k, S: a.S * k}
}

// HexRotateLeft rotates a hex coordinate 60 degrees counter-clockwise (to the left) around the origin.
//
// This is done by cyclically permuting and negating the cube coordinates (q, r, s),
// while maintaining the invariant q + r + s = 0.
func HexRotateLeft(a Hex) Hex {
	return Hex{Q: -a.S, R: -a.Q, S: -a.R}
}

// HexRotateRight rotates a hex coordinate 60 degrees clockwise around the origin.
//
// This operation permutes and negates the cube coordinates to achieve a rotation,
// preserving the constraint q + r + s = 0.
func HexRotateRight(a Hex) Hex {
	return Hex{Q: -a.R, R: -a.S, S: -a.Q}
}

// HexDirections defines the six primary movement directions in a hex grid using cube coordinates.
//
// Each entry represents a unit vector pointing to one of the six neighboring hexes,
// ordered typically in a clockwise or counter-clockwise fashion starting from the right.
//
// These directions are used for navigation, pathfinding, and neighbor lookup.
var HexDirections = []Hex{
	{Q: 1, S: -1}, // Direction 0: east
	{Q: 1, R: -1}, // Direction 1: northeast
	{R: -1, S: 1}, // Direction 2: northwest
	{Q: -1, S: 1}, // Direction 3: west
	{Q: -1, R: 1}, // Direction 4: southwest
	{R: 1, S: -1}, // Direction 5: southeast
}

// HexDirection returns the unit vector (as a Hex) representing the direction specified by the index (0–5).
//
// Each direction corresponds to one of the six adjacent directions in a hex grid using cube coordinates.
//
// The returned Hex can be added to another hex to move in that direction.
func HexDirection(direction int) Hex {
	return HexDirections[direction]
}

// HexNeighbor returns the neighboring hex of a given hex in one of the six adjacent directions.
//
// The `direction` parameter should be an integer from 0 to 5, corresponding to one of the six cardinal directions.
//
// It works by adding the appropriate directional vector (from HexDirection) to the current hex.
func HexNeighbor(hex Hex, direction int) Hex {
	return HexAdd(hex, HexDirection(direction))
}

// HexDiagonals defines the six possible diagonal directions in a hex grid using cube coordinates.
//
// Unlike direct neighbors (which are adjacent), diagonal directions jump over one hex.
//
// These vectors can be added to a hex to find its diagonal neighbors.
var HexDiagonals = []Hex{
	{Q: 2, R: -1, S: -1}, {Q: 1, R: -2, S: 1}, {Q: -1, R: -1, S: 2},
	{Q: -2, R: 1, S: 1}, {Q: -1, R: 2, S: -1}, {Q: 1, R: 1, S: -2},
}

// HexDiagonalNeighbor returns the diagonal neighbor of a hex in a given direction.
//
// Direction is an index (0–5) into the predefined HexDiagonals array,
// which contains the six possible diagonal offset vectors.
func HexDiagonalNeighbor(hex Hex, direction int) Hex {
	return HexAdd(hex, HexDiagonals[direction])
}

// HexLength computes the distance from the origin (0,0,0) to the given hex in cube coordinates.
//
// The formula sums the absolute values of q, r, and s, and divides by 2.
//
// This works because in a hex grid, the sum of the three coordinates is always zero (q + r + s = 0),
// so the length of a hex vector is half the total displacement.
func HexLength(hex Hex) int {
	return (AbsInt(hex.Q) + AbsInt(hex.R) + AbsInt(hex.S)) / 2
}

// HexDistance calculates the distance between two hexes `a` and `b`.
//
// It does this by subtracting `b` from `a` to get the offset vector,
// then computing the length of that vector in cube space.
//
// The result is the minimum number of steps between the two hexes.
func HexDistance(a, b Hex) int {
	return HexLength(HexSubtract(a, b))
}

// HexLerp performs linear interpolation between two fractional hex coordinates (a and b).
//
// The parameter `t` should be between 0.0 and 1.0, where 0.0 returns `a`, 1.0 returns `b`,
// and values in between return a point along the line connecting them.
//
// This is commonly used for animations or drawing paths in a hex grid.
func HexLerp(a, b FractionalHex, t float64) FractionalHex {
	return FractionalHex{
		Q: a.Q*(1.0-t) + b.Q*t,
		R: a.R*(1.0-t) + b.R*t,
		S: a.S*(1.0-t) + b.S*t,
	}
}

// HexRound converts a fractional hex coordinate (FractionalHex) to the nearest integer hex (Hex).
//
// It first rounds each component (q, r, s) independently to the nearest integer.
//
// Then, it adjusts the component with the largest rounding difference to ensure q + r + s = 0,
// preserving the invariant required by cube coordinates.
func HexRound(h FractionalHex) Hex {
	qi := int(math.Round(h.Q))
	ri := int(math.Round(h.R))
	si := int(math.Round(h.S))

	qDiff := math.Abs(float64(qi) - h.Q)
	rDiff := math.Abs(float64(ri) - h.R)
	sDiff := math.Abs(float64(si) - h.S)

	if qDiff > rDiff && qDiff > sDiff {
		qi = -ri - si
	} else if rDiff > sDiff {
		ri = -qi - si
	} else {
		si = -qi - ri
	}

	return Hex{qi, ri, si}
}

// HexLineDraw returns a list of hexes that form a straight line from hex `a` to hex `b`.
//
// It works by interpolating between `a` and `b` in fractional hex space, then rounding each step to the nearest hex.
//
// Small nudges (1e-6) are added to both endpoints to avoid rounding errors that can occur when points lie exactly on
// hex boundaries.
func HexLineDraw(a, b Hex) []Hex {
	N := HexDistance(a, b)
	aNudge := FractionalHex{float64(a.Q) + 1e-6, float64(a.R) + 1e-6, float64(a.S) - 2e-6}
	bNudge := FractionalHex{float64(b.Q) + 1e-6, float64(b.R) + 1e-6, float64(b.S) - 2e-6}
	results := []Hex{}
	step := 1.0 / math.Max(float64(N), 1)
	for i := 0; i <= N; i++ {
		results = append(results, HexRound(HexLerp(aNudge, bNudge, step*float64(i))))
	}
	return results
}

// QOffsetFromCube converts a cube coordinate (Hex) to a column-based offset coordinate (OffsetCoordinate).
//
// The `offset` parameter must be either EVEN (+1) or ODD (-1), indicating the layout's staggered column mode.
//
// This function is used for "even-q" or "odd-q" layouts where columns are aligned vertically,
// and rows are staggered up or down based on the column parity.
func QOffsetFromCube(offset int, h Hex) OffsetCoordinate {
	if offset != EVEN && offset != ODD {
		panic("offset must be EVEN (+1) or ODD (-1)")
	}
	col := h.Q
	row := h.R + (h.Q+offset*(h.Q&1))/2
	return OffsetCoordinate{col, row}
}

// QOffsetToCube converts an offset coordinate (column-based) to a cube coordinate (Hex).
//
// The `offset` parameter must be either EVEN (+1) or ODD (-1), defining the offset type.
//
// This conversion is commonly used for "even-q" or "odd-q" layouts where columns are offset vertically.
func QOffsetToCube(offset int, h OffsetCoordinate) Hex {
	if offset != EVEN && offset != ODD {
		panic("offset must be EVEN (+1) or ODD (-1)")
	}
	q := h.Col
	r := h.Row - (h.Col+offset*(h.Col&1))/2
	s := -q - r
	return Hex{q, r, s}
}

// ROffsetFromCube converts a cube coordinate (Hex) to a row-based offset coordinate (OffsetCoordinate).
//
// The `offset` parameter must be either EVEN (+1) or ODD (-1), specifying the type of offset used.
//
// This conversion is typically used for "even-r" or "odd-r" layouts where rows are offset horizontally.
func ROffsetFromCube(offset int, h Hex) OffsetCoordinate {
	if offset != EVEN && offset != ODD {
		panic("offset must be EVEN (+1) or ODD (-1)")
	}
	col := h.Q + (h.R+offset*(h.R&1))/2
	row := h.R
	return OffsetCoordinate{col, row}
}

// ROffsetToCube converts a row-based offset coordinate (OffsetCoordinate) to a cube coordinate (Hex).
//
// The `offset` parameter must be either EVEN (+1) or ODD (-1), indicating the row offset mode.
//
// This is used in "even-r" or "odd-r" layouts where horizontal rows are staggered.
func ROffsetToCube(offset int, h OffsetCoordinate) Hex {
	if offset != EVEN && offset != ODD {
		panic("offset must be EVEN (+1) or ODD (-1)")
	}
	q := h.Col - (h.Row+offset*(h.Row&1))/2
	r := h.Row
	s := -q - r
	return Hex{q, r, s}
}

// QDoubledFromCube converts a cube coordinate (Hex) to a column-based doubled coordinate (DoubledCoordinate).
//
// This is used in "doubled-q" layouts where columns are doubled and rows are continuous.
//
// The formula shifts rows based on the column (q) to flatten the staggered structure.
func QDoubledFromCube(h Hex) DoubledCoordinate {
	return DoubledCoordinate{h.Q, 2*h.R + h.Q}
}

// QDoubledToCube converts a column-based doubled coordinate (DoubledCoordinate) back to a cube coordinate (Hex).
//
// This function reverses the transformation applied by QDoubledFromCube.
// It is used in "doubled-q" layouts where columns are doubled and rows are linear.
func QDoubledToCube(h DoubledCoordinate) Hex {
	q := h.Col
	r := (h.Row - h.Col) / 2
	s := -q - r
	return Hex{Q: q, R: r, S: s}
}

// RDoubledFromCube converts a cube coordinate (Hex) to a row-based doubled coordinate (DoubledCoordinate).
//
// This is used in "doubled-r" layouts where rows are doubled and columns are continuous.
//
// The formula shifts columns based on the row (r) to linearize the staggered layout.
func RDoubledFromCube(h Hex) DoubledCoordinate {
	return DoubledCoordinate{2*h.Q + h.R, h.R}
}

// RDoubledToCube converts a row-based doubled coordinate (DoubledCoordinate) back to a cube coordinate (Hex).
//
// This function reverses the transformation performed by RDoubledFromCube.
//
// It is used in "doubled-r" layouts where rows are doubled and columns are linear.
func RDoubledToCube(h DoubledCoordinate) Hex {
	q := (h.Col - h.Row) / 2
	r := h.Row
	s := -q - r
	return Hex{q, r, s}
}

// HexToPixel converts a hex coordinate (Hex) to a 2D pixel position (Point).
//
// It uses the layout's orientation, size, and origin to compute the pixel location.
//
// This is commonly used to render hexagons on screen by mapping grid positions to pixel space (center).
func HexToPixel(layout Layout, h Hex) Point {
	M := layout.Orientation
	size := layout.Size
	origin := layout.Origin
	x := (M.F0*float64(h.Q) + M.F1*float64(h.R)) * size.X
	y := (M.F2*float64(h.Q) + M.F3*float64(h.R)) * size.Y
	return Point{x + origin.X, y + origin.Y}
}

// PixelToHex converts a 2D pixel position (Point) to the nearest hex coordinate (Hex).
//
// It delegates to PixelToHexFractional and then rounds the result to the nearest integer hex.
//
// This is useful for detecting which hex was clicked or hovered in a graphical interface.
func PixelToHex(layout Layout, p Point) Hex {
	return HexRound(PixelToHexFractional(layout, p))
}

// PixelToHexFractional converts a 2D pixel position (Point) to a fractional hex coordinate (FractionalHex).
//
// It applies the inverse transformation defined by the layout's orientation, size, and origin.
//
// This is useful when precise hex positions are needed (e.g., interpolation or smooth cursor tracking).
func PixelToHexFractional(layout Layout, p Point) FractionalHex {
	M := layout.Orientation
	size := layout.Size
	origin := layout.Origin
	pt := Point{(p.X - origin.X) / size.X, (p.Y - origin.Y) / size.Y}
	q := M.B0*pt.X + M.B1*pt.Y
	r := M.B2*pt.X + M.B3*pt.Y
	return FractionalHex{q, r, -q - r}
}

// HexCornerOffset calculates the 2D offset from the center of a hex to one of its corners.
//
// The corner index should be in the range [0, 5].
//
// This function uses the layout's orientation and size to compute the angle and length of the corner vector.
func HexCornerOffset(layout Layout, corner int) Point {
	M := layout.Orientation
	size := layout.Size
	angle := 2.0 * math.Pi * (M.StartAngle - float64(corner)) / 6.0
	return Point{size.X * math.Cos(angle), size.Y * math.Sin(angle)}
}

// PolygonCorners returns the 2D pixel coordinates of the six corners of a given hex.
//
// It first calculates the center of the hex using HexToPixel, then adds the offset for each corner.
//
// This is typically used to draw the hexagon shape in a 2D rendering context (e.g., SVG or canvas).
func PolygonCorners(layout Layout, h Hex) []Point {
	corners := []Point{}
	center := HexToPixel(layout, h)
	for i := 0; i < 6; i++ {
		offset := HexCornerOffset(layout, i)
		corners = append(corners, Point{center.X + offset.X, center.Y + offset.Y})
	}
	return corners
}

// AbsInt returns the absolute value of an integer.
//
// If x is negative, it returns -x; otherwise, it returns x as-is.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// DoubleWidthDistance calculates the distance between two hexes
//
// in a grid with double-width horizontal layout (even-q / odd-q)
func DoubleWidthDistance(a, b OffsetCoordinate) int {
	dcol := abs(a.Col - b.Col) // difference in columns
	drow := abs(a.Row - b.Row) // difference in rows
	return drow + max(0, (dcol-drow)/2)
}

// DoubleHeightDistance calculates the distance between two hexes
//
// in a grid with double-height vertical layout (even-r / odd-r)
func DoubleHeightDistance(a, b OffsetCoordinate) int {
	dcol := abs(a.Col - b.Col) // difference in columns
	drow := abs(a.Row - b.Row) // difference in rows
	return dcol + max(0, (drow-dcol)/2)
}

// OffsetDistance converts two offset hex coordinates to axial,
// and returns the distance between them using axial distance formula.
func OffsetDistance(a, b OffsetCoordinate) int {
	ac := OffsetToAxial(a)
	bc := OffsetToAxial(b)
	return AxialDistance(ac, bc)
}

// OffsetToAxial converts offset coordinates to axial coordinates.
// This example assumes "odd-r" layout (rows are offset).
func OffsetToAxial(o OffsetCoordinate) AxialCoordinate {
	col := o.Col
	row := o.Row

	q := col - (row-(row&1))/2
	r := row

	return AxialCoordinate{Q: q, R: r}
}

// Predefined axial direction vectors in clockwise order (starting from east)
var axialDirectionVectors = []Hex{
	{Q: +1, R: 0}, {Q: +1, R: -1}, {Q: 0, R: -1},
	{Q: -1, R: 0}, {Q: -1, R: +1}, {Q: 0, R: +1},
}

// AxialDirection returns the unit vector for a given direction (0–5)
func AxialDirection(direction int) Hex {
	return axialDirectionVectors[direction]
}

// AxialAdd adds two axial coordinates (hex + vec)
func AxialAdd(a, b Hex) Hex {
	return Hex{
		Q: a.Q + b.Q,
		R: a.R + b.R,
	}
}

// AxialNeighbor returns the neighbor of a hex in a given axial direction (0–5)
func AxialNeighbor(hex Hex, direction int) Hex {
	return AxialAdd(hex, AxialDirection(direction))
}

// Predefined cube direction vectors in clockwise order (starting from +q)
var cubeDirectionVectors = []Cube{
	{Q: +1, R: 0, S: -1},
	{Q: +1, R: -1, S: 0},
	{Q: 0, R: -1, S: +1},
	{Q: -1, R: 0, S: +1},
	{Q: -1, R: +1, S: 0},
	{Q: 0, R: +1, S: -1},
}

// CubeDirection returns the unit vector for the given direction (0–5)
func CubeDirection(direction int) Cube {
	return cubeDirectionVectors[direction]
}

// CubeNeighbor returns the neighboring cube in the specified direction (0–5)
func CubeNeighbor(cube Cube, direction int) Cube {
	return CubeAdd(cube, CubeDirection(direction))
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// AxialSubtract returns the vector difference between two axial coordinates
func AxialSubtract(a, b AxialCoordinate) AxialCoordinate {
	return AxialCoordinate{
		Q: a.Q - b.Q,
		R: a.R - b.R,
	}
}

// AxialDistance returns the distance between two axial coordinates
// using the standard axial distance formula
func AxialDistance(a, b AxialCoordinate) int {
	vec := AxialSubtract(a, b)

	dq := abs(vec.Q)
	dr := abs(vec.R)
	ds := abs(vec.Q + vec.R) // since s = -q - r, so ds = abs(-dq - dr) = abs(q + r)

	return (dq + dr + ds) / 2
}

// Cube represents cube coordinates for a hex (q, r, s) where q + r + s == 0
type Cube Hex

// CubeSubtract returns the vector difference between two cube coordinates
func CubeSubtract(a, b Cube) Cube {
	return Cube{
		Q: a.Q - b.Q,
		R: a.R - b.R,
		S: a.S - b.S,
	}
}

// CubeDistance returns the distance between two hexes in cube coordinates.
// The distance is the maximum of the absolute differences along each axis.
func CubeDistance(a, b Cube) int {
	vec := CubeSubtract(a, b)

	return maxCube(abs(vec.Q), abs(vec.R), abs(vec.S))

	// Alternatively:
	// return max(
	//     abs(a.Q - b.Q),
	//     abs(a.R - b.R),
	//     abs(a.S - b.S),
	// )
}

// Predefined cube diagonal directions (there are 6 diagonal directions)
var cubeDiagonalVectors = []Cube{
	{+2, -1, -1},
	{+1, -2, +1},
	{-1, -1, +2},
	{-2, +1, +1},
	{-1, +2, -1},
	{+1, +1, -2},
}

// CubeAdd returns the result of adding two cube coordinates
func CubeAdd(a, b Cube) Cube {
	return Cube{
		Q: a.Q + b.Q,
		R: a.R + b.R,
		S: a.S + b.S,
	}
}

// CubeDiagonalNeighbor returns the cube coordinate in a diagonal direction
//
// direction must be in the range [0, 5]
func CubeDiagonalNeighbor(cube Cube, direction int) (cume Cube, err error) {
	if direction < 0 || direction >= len(cubeDiagonalVectors) {
		return cube, errors.New("invalid cube direction")
	}

	return CubeAdd(cube, cubeDiagonalVectors[direction]), nil
}

// Predefined direction vectors for double-width layout
var doubleWidthDirectionVectors = []DoubledCoordinate{
	{+2, 0}, {+1, -1}, {-1, -1},
	{-2, 0}, {-1, +1}, {+1, +1},
}

// Predefined direction vectors for double-height layout
var doubleHeightDirectionVectors = []DoubledCoordinate{
	{+1, +1}, {+1, -1}, {0, -2},
	{-1, -1}, {-1, +1}, {0, +2},
}

// DoubleWidthAdd returns the result of adding two double-width hex coordinates
func DoubleWidthAdd(a, b DoubledCoordinate) DoubledCoordinate {
	return DoubledCoordinate{
		Col: a.Col + b.Col,
		Row: a.Row + b.Row,
	}
}

// DoubleWidthNeighbor returns the neighbor of a double-width hex in a given direction
// direction must be in [0..5]
func DoubleWidthNeighbor(hex DoubledCoordinate, direction int) DoubledCoordinate {
	vec := doubleWidthDirectionVectors[direction]
	return DoubleWidthAdd(hex, vec)
}

// DoubleHeightAdd returns the result of adding two double-height hex coordinates
func DoubleHeightAdd(a, b DoubledCoordinate) DoubledCoordinate {
	return DoubledCoordinate{
		Col: a.Col + b.Col,
		Row: a.Row + b.Row,
	}
}

// DoubleHeightNeighbor returns the neighbor of a double-height hex in a given direction
// direction must be in [0..5]
func DoubleHeightNeighbor(hex DoubledCoordinate, direction int) DoubledCoordinate {
	vec := doubleHeightDirectionVectors[direction]
	return DoubleHeightAdd(hex, vec)
}

// OddROffsetNeighbor OddR offset neighbor (odd-r layout: row parity affects direction)
func OddROffsetNeighbor(hex OffsetCoordinate, direction int) OffsetCoordinate {
	parity := hex.Row & 1
	diff := oddRDirectionDifferences[parity][direction]
	return OffsetCoordinate{
		Col: hex.Col + diff[0],
		Row: hex.Row + diff[1],
	}
}

// EvenROffsetNeighbor EvenR offset neighbor (even-r layout: row parity affects direction)
func EvenROffsetNeighbor(hex OffsetCoordinate, direction int) OffsetCoordinate {
	parity := hex.Row & 1
	diff := evenRDirectionDifferences[parity][direction]
	return OffsetCoordinate{
		Col: hex.Col + diff[0],
		Row: hex.Row + diff[1],
	}
}

// OddQOffsetNeighbor OddQ offset neighbor (odd-q layout: column parity affects direction)
func OddQOffsetNeighbor(hex OffsetCoordinate, direction int) OffsetCoordinate {
	parity := hex.Col & 1
	diff := oddQDirectionDifferences[parity][direction]
	return OffsetCoordinate{
		Col: hex.Col + diff[0],
		Row: hex.Row + diff[1],
	}
}

// EvenQOffsetNeighbor EvenQ offset neighbor (even-q layout: column parity affects direction)
func EvenQOffsetNeighbor(hex OffsetCoordinate, direction int) OffsetCoordinate {
	parity := hex.Col & 1
	diff := evenQDirectionDifferences[parity][direction]
	return OffsetCoordinate{
		Col: hex.Col + diff[0],
		Row: hex.Row + diff[1],
	}
}

// DoubleHeightToAxial Converts a double-height hex coordinate to an axial hex coordinate
func DoubleHeightToAxial(hex DoubledCoordinate) Hex {
	q := hex.Col
	r := (hex.Row - hex.Col) / 2
	return Hex{Q: q, R: r}
}

// AxialToDoubleHeight Converts an axial hex coordinate to a double-height hex coordinate
func AxialToDoubleHeight(hex Hex) DoubledCoordinate {
	col := hex.Q
	row := 2*hex.R + hex.Q
	return DoubledCoordinate{Col: col, Row: row}
}

// DoubleWidthToAxial Converts a double-width hex coordinate to an axial hex coordinate
func DoubleWidthToAxial(hex DoubledCoordinate) Hex {
	q := (hex.Col - hex.Row) / 2
	r := hex.Row
	return Hex{Q: q, R: r}
}

// AxialToDoubleWidth Converts an axial hex coordinate to a double-width hex coordinate
func AxialToDoubleWidth(hex Hex) DoubledCoordinate {
	col := 2*hex.Q + hex.R
	row := hex.R
	return DoubledCoordinate{Col: col, Row: row}
}

var oddRDirectionDifferences = [2][6][2]int{
	// even rows
	{
		{+1, 0}, {0, -1}, {-1, -1},
		{-1, 0}, {-1, +1}, {0, +1},
	},
	// odd rows
	{
		{+1, 0}, {+1, -1}, {0, -1},
		{-1, 0}, {0, +1}, {+1, +1},
	},
}

var evenRDirectionDifferences = [2][6][2]int{
	// even rows
	{
		{+1, 0}, {+1, -1}, {0, -1},
		{-1, 0}, {0, +1}, {+1, +1},
	},
	// odd rows
	{
		{+1, 0}, {0, -1}, {-1, -1},
		{-1, 0}, {-1, +1}, {0, +1},
	},
}

var oddQDirectionDifferences = [2][6][2]int{
	// even cols
	{
		{+1, 0}, {+1, -1}, {0, -1},
		{-1, -1}, {-1, 0}, {0, +1},
	},
	// odd cols
	{
		{+1, +1}, {+1, 0}, {0, -1},
		{-1, 0}, {-1, +1}, {0, +1},
	},
}

var evenQDirectionDifferences = [2][6][2]int{
	// even cols
	{
		{+1, +1}, {+1, 0}, {0, -1},
		{-1, 0}, {-1, +1}, {0, +1},
	},
	// odd cols
	{
		{+1, 0}, {+1, -1}, {0, -1},
		{-1, -1}, {-1, 0}, {0, +1},
	},
}

// maxCube returns the largest of three integers
func maxCube(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

// max returns the larger of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AxialToOddR(hex Hex) OffsetCoordinate {
	col := hex.Q + (hex.R-(hex.R&1))/2
	row := hex.R
	return OffsetCoordinate{Col: col, Row: row}
}

func OddRToAxial(offset OffsetCoordinate) Hex {
	q := offset.Col - (offset.Row-(offset.Row&1))/2
	r := offset.Row
	return Hex{Q: q, R: r}
}

func AxialToEvenR(hex Hex) OffsetCoordinate {
	col := hex.Q + (hex.R+(hex.R&1))/2
	row := hex.R
	return OffsetCoordinate{Col: col, Row: row}
}

func EvenRToAxial(offset OffsetCoordinate) Hex {
	q := offset.Col - (offset.Row+(offset.Row&1))/2
	r := offset.Row
	return Hex{Q: q, R: r}
}

func AxialToOddQ(hex Hex) OffsetCoordinate {
	col := hex.Q
	row := hex.R + (hex.Q-(hex.Q&1))/2
	return OffsetCoordinate{Col: col, Row: row}
}

func OddQToAxial(offset OffsetCoordinate) Hex {
	q := offset.Col
	r := offset.Row - (offset.Col-(offset.Col&1))/2
	return Hex{Q: q, R: r}
}

func AxialToEvenQ(hex Hex) OffsetCoordinate {
	col := hex.Q
	row := hex.R + (hex.Q+(hex.Q&1))/2
	return OffsetCoordinate{Col: col, Row: row}
}

func EvenQToAxial(offset OffsetCoordinate) Hex {
	q := offset.Col
	r := offset.Row - (offset.Col+(offset.Col&1))/2
	return Hex{Q: q, R: r}
}

func CubeToOddR(cube Cube) OffsetCoordinate {
	col := cube.Q + (cube.R-(cube.R&1))/2
	row := cube.R
	return OffsetCoordinate{Col: col, Row: row}
}

func OddRToCube(offset OffsetCoordinate) Cube {
	q := offset.Col - (offset.Row-(offset.Row&1))/2
	r := offset.Row
	s := -q - r
	return Cube{Q: q, R: r, S: s}
}

func CubeToEvenR(cube Cube) OffsetCoordinate {
	col := cube.Q + (cube.R+(cube.R&1))/2
	row := cube.R
	return OffsetCoordinate{Col: col, Row: row}
}

func EvenRToCube(offset OffsetCoordinate) Cube {
	q := offset.Col - (offset.Row+(offset.Row&1))/2
	r := offset.Row
	s := -q - r
	return Cube{Q: q, R: r, S: s}
}

func CubeToOddQ(cube Cube) OffsetCoordinate {
	col := cube.Q
	row := cube.R + (cube.Q-(cube.Q&1))/2
	return OffsetCoordinate{Col: col, Row: row}
}

func OddQToCube(offset OffsetCoordinate) Cube {
	q := offset.Col
	r := offset.Row - (offset.Col-(offset.Col&1))/2
	s := -q - r
	return Cube{Q: q, R: r, S: s}
}

func CubeToEvenQ(cube Cube) OffsetCoordinate {
	col := cube.Q
	row := cube.R + (cube.Q+(cube.Q&1))/2
	return OffsetCoordinate{Col: col, Row: row}
}

func EvenQToCube(offset OffsetCoordinate) Cube {
	q := offset.Col
	r := offset.Row - (offset.Col+(offset.Col&1))/2
	s := -q - r
	return Cube{Q: q, R: r, S: s}
}

// CubeToAxial converts a Cube coordinate to an Axial (Hex) coordinate
func CubeToAxial(cube Cube) Hex {
	return Hex{
		Q: cube.Q,
		R: cube.R,
	}
}

// AxialToCube converts an Axial (Hex) coordinate to a Cube coordinate
func AxialToCube(hex Hex) Cube {
	q := hex.Q
	r := hex.R
	s := -q - r
	return Cube{
		Q: q,
		R: r,
		S: s,
	}
}

//// Teste simples
//func main() {
//	a, _ := NewHex(3, -7, 4)
//	b, _ := NewHex(0, 0, 0)
//	fmt.Println("Hex Distance:", HexDistance(a, b))
//
//	line := HexLineDraw(b, a)
//	for i, h := range line {
//		fmt.Printf("Step %d: %+v\n", i, h)
//	}
//}
