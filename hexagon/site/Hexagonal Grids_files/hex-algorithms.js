/*
 * From https://www.redblobgames.com/grids/hexagons/
 * Copyright 2018 Red Blob Games <redblobgames@gmail.com>
 * License: Apache v2.0 <http://www.apache.org/licenses/LICENSE-2.0.html>
 */
'use strict';

/* global Hex, OffsetCoord, DoubledCoord */


Hex.prototype.toString = function() {
    return `${this.q},${this.r},${this.s}`;
};


/* When t = 0, it's all a; when t = 1, it's all b */
function mix(a, b, t) {
    return a*(1-t) + b*t;
}


/* return min, max of one field of an array of objects */
function bounds(array, field) {
    let min = Infinity, max = -Infinity;
    for (let object of array) {
        let value = object[field];
        if (value < min) { min = value; }
        if (value > max) { max = value; }
    }
    return {min, max};
}


/* return integer sequence min <= x <= max (half-open)  */
function closedInterval({min, max}) {
    let results = [];
    for (let i = min; i <= max; i++) {
        results.push(i);
    }
    return results;
}


/* For whatever reason I don't have this in the generated code! */
function hex_equal(a, b) {
    return a.q == b.q && a.r == b.r && a.s == b.s;
}


/* Specifically for offset grid diagrams */
function makeRectangularShape(minCol, maxCol, minRow, maxRow, convert) {
    let results = [];
    for (let col = minCol; col <= maxCol; col++) {
        for (let row = minRow; row <= maxRow; row++) {
            let hex = convert(new OffsetCoord(col, row));
            hex.col = col;
            hex.row = row;
            results.push(hex);
        }
    }
    return results;
}


/* Specifically for doubled grid diagrams */
function makeRDoubledRectangularShape(minCol, maxCol, minRow, maxRow) {
    let results = [];
    for (let row = minRow; row <= maxRow; row++) {
        for (let col = minCol + (row & 1); col <= maxCol; col += 2) {
            let hex = new DoubledCoord(col, row).rdoubledToCube();
            hex.col = col;
            hex.row = row;
            results.push(hex);
        }
    }
    return results;
}


/* Specifically for doubled grid diagrams */
function makeQDoubledRectangularShape(minCol, maxCol, minRow, maxRow) {
    let results = [];
    for (let col = minCol; col <= maxCol; col++) {
        for (let row = minRow + (col & 1); row <= maxRow; row += 2) {
            let hex = new DoubledCoord(col, row).qdoubledToCube();
            hex.col = col;
            hex.row = row;
            results.push(hex);
        }
    }
    return results;
}


function makeHexagonalShape(N) {
    let results = [];
    for (let q = -N; q <= N; q++) {
        for (let r = -N; r <= N; r++) {
            let hex = new Hex(q, r, -q-r);
            if (hex.len() <= N) {
                results.push(hex);
            }
        }
    }
    return results;
}


function makeDownTriangularShape(N) {
    let results = [];
    for (let r = 0; r < N; r++) {
        for (let q = 0; q < N-r; q++) {
            results.push(new Hex(q, r, -q-r));
        }
    }
    return results;
}
    

function makeUpTriangularShape(N) {
    let results = [];
    for (let r = 0; r < N; r++) {
        for (let q = N-r-1; q < N; q++) {
            results.push(new Hex(q, r, -q-r));
        }
    }
    return results;
}
    

function makeRhombusShape(w, h) {
    if (!h) { h = w; }
    let results = [];
    for (let r = 0; r < h; r++) {
        for (let q = 0; q < w; q++) {
            results.push(new Hex(q, r, -q-r));
        }
    }
    return results;
}

    
/* Given a set of points, return the maximum extent
   {left, right, top, bottom} */
function pointSetBounds(points) {
    let left = Infinity, top = Infinity,
        right = -Infinity, bottom = -Infinity;
    for (let p of points) {
        if (p.x < left   ) { left = p.x;   }
        if (p.x > right  ) { right = p.x;  }
        if (p.y < top    ) { top = p.y;    }
        if (p.y > bottom ) { bottom = p.y; }
    }
    return {left, top, right, bottom};
}


/* Given a set of hexes, return the maximum extent
   {left, right, top, bottom} */
function hexSetBounds(layout, hexes) {
    let corners = [];
    for (let corner = 0; corner < 6; corner++) {
        corners.push(layout.hexCornerOffset(corner));
    }
    let cornerBounds = pointSetBounds(corners);
    
    let centerBounds = pointSetBounds(hexes.map(h => layout.hexToPixel(h)));

    return {
        left   : cornerBounds.left   + centerBounds.left,
        top    : cornerBounds.top    + centerBounds.top,
        right  : cornerBounds.right  + centerBounds.right,
        bottom : cornerBounds.bottom + centerBounds.bottom,
    };
}


function breadthFirstSearch(start, blocked) {
    /* see https://www.redblobgames.com/pathfinding/a-star/introduction.html */
    let cost_so_far = {}; cost_so_far[start] = 0;
    let came_from = {}; came_from[start] = null;
    let fringes = [[start]];
    for (let k = 0; fringes[k].length > 0; k++) {
        fringes[k+1] = [];
        for (let hex of fringes[k]) {
            for (let dir = 0; dir < 6; dir++) {
                let neighbor = hex.neighbor(dir);
                if (cost_so_far[neighbor] === undefined
                    && !blocked(neighbor)) {
                    cost_so_far[neighbor] = k+1;
                    came_from[neighbor] = hex;
                    fringes[k+1].push(neighbor);
                }
            }
        }
    }
    return {cost_so_far, came_from};
}


/* NOTE: this returns the *fractional* hexes between A and B; you need
   to call round() on them to get the hex tiles */
function hexLineFractional(A, B) {
    /* see https://www.redblobgames.com/grids/line-drawing.html */
    
    /* HACK: add a tiny offset to the start point to break ties,
     * because there are a lot of ties on a grid, and I want it to
     * always round the same direction for consistency. To demonstrate
     * the need for this hack, draw a line from Hex(-5, 0, +5) to
     * Hex(+5, -5, 0). Without the hack, there are points on the edge
     * that will sometimes be rounded one way and sometimes the other.
     * The hack will force them to be rounded consistently. */
    const offset = new Hex(1e-6, 2e-6, -3e-6);
    
    let N = A.subtract(B).len();
    let results = [];
    for (let i = 0; i <= N; i++) {
        results.push(A.lerp(B, i / Math.max(1, N)).add(offset));
    }
    return results;
}


function hexRing(radius) {
    if (radius === 0) { // special case because ring 0 has only one hex, not six
        return [new Hex(0, 0, 0)];
    }
    var results = [];
    var H = Hex.direction(4).scale(radius);
    for (var side = 0; side < 6; side++) {
        for (var step = 0; step < radius; step++) {
            results.push(H);
            H = H.neighbor(side);
        }
    }
    return results;
}
