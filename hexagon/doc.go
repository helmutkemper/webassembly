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

//         ⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯
//      ⋰              ⋱
//    ⋰                  ⋱
//  ⋰                      ⋱
// ⋮                        ⋮
// ⋮                        ⋮
// ⋮                        ⋮
//  ⋱                      ⋰
//    ⋱                  ⋰
//      ⋱              ⋰
//         ⋯⋯⋯⋯⋯⋯⋯⋯⋯⋯

//       ⋯⋯⋯⋯⋯⋯⋯
//    ⋰           ⋱
//  ⋰               ⋱
// ⋮                 ⋮
// ⋮                 ⋮
//  ⋱               ⋰
//    ⋱           ⋰
//       ⋯⋯⋯⋯⋯⋯⋯

//     ⋯⋯⋯
//  ⋰       ⋱
// ⋮         ⋮
//  ⋱       ⋰
//     ⋯⋯⋯

//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱

//   ⋮ ⋯⋯⋯ ⋮   1,0   ⋮ ⋯⋯⋯ ⋮   3,0   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱
// ⋮   0,0   ⋮ ⋯⋯⋯ ⋮   2,0   ⋮ ⋯⋯⋯ ⋮   4,0   ⋮ ⋯⋯⋯ ⋮   6,0   ⋮
//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮   1,1   ⋮ ⋯⋯⋯ ⋮   3,1   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱
// ⋮   0,2   ⋮ ⋯⋯⋯ ⋮   2,2   ⋮ ⋯⋯⋯ ⋮   4,2   ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮   1,3   ⋮ ⋯⋯⋯ ⋮   3,3   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱
// ⋮   0,4   ⋮ ⋯⋯⋯ ⋮   2,4   ⋮ ⋯⋯⋯ ⋮   4,4   ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮   1,5   ⋮ ⋯⋯⋯ ⋮   3,5   ⋮ ⋯⋯⋯ ⋮   5,5   ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱
// ⋮   0,6   ⋮ ⋯⋯⋯ ⋮   2,6   ⋮ ⋯⋯⋯ ⋮   4,6   ⋮ ⋯⋯⋯ ⋮   6,6   ⋮
//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮   1,7   ⋮ ⋯⋯⋯ ⋮   3,7   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱
// ⋮   0,8   ⋮ ⋯⋯⋯ ⋮   2,8   ⋮ ⋯⋯⋯ ⋮   4,8   ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱       ⋰       ⋱       ⋰       ⋱       ⋰       ⋱       ⋰
//   ⋮ ⋯⋯⋯ ⋮   1,9   ⋮ ⋯⋯⋯ ⋮   3,9   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮

// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+4   -1⋰  -2   ⋱     -2⋰   0   ⋱     -3⋰  +2   ⋱+1   -4⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+3   -1⋰  -1   ⋱+2   -2⋰  +1   ⋱+1   -3⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+3    0⋰  -2   ⋱+2   -1⋰   0   ⋱+1   -2⋰  +2   ⋱0    -3⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+2    0⋰  -1   ⋱+1   -1⋰  +1   ⋱0    -2⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+2   +1⋰  -2   ⋱+1    0⋰   q   ⋱0    -1⋰  +2   ⋱-1   -2⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+1   +1⋰  -1   ⋱s     r⋰  +1   ⋱-1   -1⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+1   +2⋰  -2   ⋱0    +1⋰   0   ⋱-1    0⋰  +2   ⋱-2   -1⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱0    +2⋰  -1   ⋱-1   -1⋰  +1   ⋱-2    0⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱0    +3⋰  -2   ⋱-1   +2⋰   0   ⋱-2   +1⋰  +2   ⋱-3    0⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱-1   +3⋰  -1   ⋱-2   -2⋰  +1   ⋱-3   +1⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮

// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+4   -1⋰  -2   ⋱     -2⋰   0   ⋱     -3⋰  +2   ⋱+1   -4⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+3   -1⋰  -1   ⋱+2   -2⋰  +1   ⋱+1   -3⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+3    0⋰  -2   ⋱+2   -1⋰   0   ⋱+1   -2⋰  +2   ⋱0    -3⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+2    0⋰  -1   ⋱+1   -1⋰  +1   ⋱0    -2⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+2   +1⋰  -2   ⋱+1    0⋰   q   ⋱0    -1⋰  +2   ⋱-1   -2⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+1   +1⋰  -1   ⋱s     r⋰  +1   ⋱-1   -1⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+1   +2⋰  -2   ⋱0    +1⋰   0   ⋱-1    0⋰  +2   ⋱-2   -1⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱0    +2⋰  -1   ⋱-1   -1⋰  +1   ⋱-2    0⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱0    +3⋰  -2   ⋱-1   +2⋰   0   ⋱-2   +1⋰  +2   ⋱-3    0⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱-1   +3⋰  -1   ⋱-2   -2⋰  +1   ⋱-3   +1⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮

// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+4   -1⋰  -2   ⋱     -2⋰   0   ⋱     -3⋰  +2   ⋱+1   -4⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+3   -1⋰  -1   ⋱+2   -2⋰  +1   ⋱+1   -3⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+3    0⋰  -2   ⋱+2   -1⋰   0   ⋱+1   -2⋰  +2   ⋱0    -3⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+2    0⋰  -1   ⋱+1   -1⋰  +1   ⋱0    -2⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+2   +1⋰  -2   ⋱+1    0⋰   q   ⋱0    -1⋰  +2   ⋱-1   -2⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+1   +1⋰  -1   ⋱s     r⋰  +1   ⋱-1   -1⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+1   +2⋰  -2   ⋱0    +1⋰   0   ⋱-1    0⋰  +2   ⋱-2   -1⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱0    +2⋰  -1   ⋱-1   -1⋰  +1   ⋱-2    0⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱0    +3⋰  -2   ⋱-1   +2⋰   0   ⋱-2   +1⋰  +2   ⋱-3    0⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱-1   +3⋰  -1   ⋱-2   -2⋰  +1   ⋱-3   +1⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮

// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+4   -1⋰  -2   ⋱     -2⋰   0   ⋱     -3⋰  +2   ⋱+1   -4⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+3   -1⋰  -1   ⋱+2   -2⋰  +1   ⋱+1   -3⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+3    0⋰  -2   ⋱+2   -1⋰   0   ⋱+1   -2⋰  +2   ⋱0    -3⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+2    0⋰  -1   ⋱+1   -1⋰  +1   ⋱0    -2⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+2   +1⋰  -2   ⋱+1    0⋰   q   ⋱0    -1⋰  +2   ⋱-1   -2⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱+1   +1⋰  -1   ⋱s     r⋰  +1   ⋱-1   -1⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱+1   +2⋰  -2   ⋱0    +1⋰   0   ⋱-1    0⋰  +2   ⋱-2   -1⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱0    +2⋰  -1   ⋱-1   -1⋰  +1   ⋱-2    0⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
//  ⋱0    +3⋰  -2   ⋱-1   +2⋰   0   ⋱-2   +1⋰  +2   ⋱-3    0⋰
//   ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮
//  ⋰  -3   ⋱-1   +3⋰  -1   ⋱-2   -2⋰  +1   ⋱-3   +1⋰  +3   ⋱
// ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮ ⋯⋯⋯ ⋮         ⋮
