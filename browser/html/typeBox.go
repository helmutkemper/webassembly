package html

type BoxModifier int

const ()

type Box struct {
	X          int
	Y          int
	Width      int
	Height     int
	HalfWidth  int
	HalfHeight int
}

func (e *Box) Update() {
	e.HalfWidth = e.X + e.Width/2
	e.HalfHeight = e.Y + e.Height/2
}

func (e *Box) Collision(element Box) (collision bool) {
	if e.X < element.X+element.Width &&
		e.X+e.Width > element.X &&
		e.Y < element.Y+element.Height &&
		e.Y+e.Height > element.Y {
		return true
	}

	return false
}

//
//    +-----+-----+
//    |     |     |
//    |  A  |  B  |
//    |     |     |
//    +-----+-----+
//    |     |     |
//    |  C  |  D  |
//    |     |     |
//    +-----+-----+
//

//        player
//    +------+------+
//    |      |      |
//    |      |      |
//    |      |      |
//    +------+------+
//    |      |      |
//    |      |   +--+----------+
//    |      |   |  |          |
//    +------+---+--+          |
//               |             |
//               |             |
//               |             |
//               |             |
//               |             |
//               +-------------+
//                  element

//  +------+------+       +------+------+
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  +------+------+       +------+------+
//  |      |      |       |      |      |
//  |      |   +--+---+---+--+   |      |
//  |      |   |  |   |   |  |   |      |
//  +------+---+--+   |   +--+---+------+
//             |      |      |
//             +------+------+
//             |      |      |
//  +------+---+--+   |   +--+---+------+
//  |      |   |  |   |   |  |   |      |
//  |      |   +--+---+---+--+   |      |
//  |      |      |       |      |      |
//  +------+------+       +------+------+
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  +------+------+       +------+------+
//

//  +------+------+       +------+------+
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  +------+------+       +------+------+
//  |      |      |       |      |      |
//  |      |   +--+---+---+--+   |      |
//  |      |   |  |   |   |  |   |      |
//  +------+---+--+   |   +--+---+------+
//             |      |      |
//             +------+------+
//             |      |      |
//  +------+---+--+   |   +--+---+------+
//  |      |   |  |   |   |  |   |      |
//  |      |   +--+---+---+--+   |      |
//  |      |      |       |      |      |
//  +------+------+       +------+------+
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  +------+------+       +------+------+
//

//         A                     B
//  +------+------+       +------+------+
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  +------+------+       +------+------+
//  |      |      |       |      |      |
//  |      |   +--+---+---+--+   |      |
//  |      |   |  |   |   |  |   |      |
//  +------+---+--+   |   +--+---+------+
//             |      |      |
//             +------+------+ Player
//             |      |      |
//  +------+---+--+   |   +--+---+------+
//  |      |   |  |   |   |  |   |      |
//  |      |   +--+---+---+--+   |      |
//  |      |      |       |      |      |
//  +------+------+       +------+------+
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  |      |      |       |      |      |
//  +------+------+       +------+------+
//         C                     D
//
//
//

//        player
//    +------+------+
//    |      |      |
//    |      |      |
//    |      |      |
//    +------+------+
//    |      |      |
//    |      |   +--+---+------+
//    |      |   |  |   |      |
//    +------+---+--+   |      |
//               |      |      |
//               +------+------+
//               |      |      |
//               |      |      |
//               |      |      |
//               +------+------+
//                   element

//
//    +-----+-----+
//    |     |     |
//    |  a  |  b  |
//    |     |     |
//    +-----+-----+
//    |     |     |
//    |  c  |  d  |
//    |     |     |
//    +-----+-----+
//

func (e *Box) CollisionA(element Box) (collision bool) {
	a := Box{}
	a.X = e.X
	a.Y = e.Y
	a.Width = e.Width / 2
	a.Height = e.Height / 2
	return a.Collision(element)
}

func (e *Box) CollisionB(element Box) (collision bool) {
	b := Box{}
	b.X = e.X + e.Width/2
	b.Y = e.Y
	b.Width = e.Width / 2
	b.Height = e.Height / 2
	return b.Collision(element)
}

func (e *Box) CollisionC(element Box) (collision bool) {
	c := Box{}
	c.X = e.X
	c.Y = e.Y + e.Height/2
	c.Width = e.Width / 2
	c.Height = e.Height / 2
	return c.Collision(element)
}

func (e *Box) CollisionD(element Box) (collision bool) {
	d := Box{}
	d.X = e.X + e.Width/2
	d.Y = e.Y + e.Height/2
	d.Width = e.Width / 2
	d.Height = e.Height / 2
	return d.Collision(element)
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
//
//
//
//
//
//
