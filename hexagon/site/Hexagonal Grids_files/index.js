/*
 * From https://www.redblobgames.com/grids/hexagons/
 * Copyright 2018 Red Blob Games <redblobgames@gmail.com>
 * License: Apache v2.0 <http://www.apache.org/licenses/LICENSE-2.0.html>
 *
 * These are the diagrams on the index.html page. They use the code in
 * diagrams.js.
 */
'use strict';

/* global Hex, Point, OffsetCoord, DoubledCoord, makeRectangularShape,
 * makeHexagonalShape, makeDownTriangularShape, makeUpTriangularShape,
 * makeRhombusShape, hexLineFractional, hexRing, mix, bounds,
 * closedInterval, hex_equal, Diagram, layouts, orientations, mixins,
 * hexCenter, makeAnimatedLayout, AnimationMixin */

console.info("I'm happy to answer questions about hexagons or how this page works. For most of my pages, you can see the source code in the browser developer Sources/Debugger tab. The core hex algorithms used on this page are linked from the Implementation Guide. --redblobgames@gmail.com");
console.info("If you want to make your own interactive pages, I have a tutorial for that https://www.redblobgames.com/making-of/circle-drawing/");
console.info("Try scrolling down to a diagram and typing into the console:\n  orientations.global.layout = 'flat';");


Vue.component('hover-show', {
    props: ['show'],
    data: () => ({down: false}),
    // NOTE: I want the text to be selectable except when
    // interacting with it, so I set user-select:none only on
    // mousedown. I use @touchstart.prevent to prevent
    // scrolling on iOS, and to prevent the long-touch context
    // menu on Android. It does not prevent scrolling on
    // Android though. CSS touch-action:none does not help.
    template: `<dfn @touchstart.prevent=""
                    @pointerenter="set" @pointerleave="clear"
                    :style="{'user-select': down? 'none' : ''}">
                  <slot />
               </dfn>`,
    methods: {
        set(event) {
            this.down = true;
            this.$parent.show = this.show;
        },
        clear(event) {
            this.down = false;
            this.$parent.show = 'none';
        },
    },
});


function OffsetDiagram(id, layout, caption, converter) {
    Diagram(id, {
        template: "#template-diagram-grid-offset",
        mixins: [mixins.highlight],
        layout: layout,
        data: {
            caption,
            id: `diagram-${id}`,
            highlight: (() => { let hex = converter(new OffsetCoord(1, 1)); hex.col = 1; hex.row = 1; return hex; })(),
            grid: makeRectangularShape(0, layout == 'flat'? 7 : 6, 0, layout == 'flat'? 5 : 6, converter),
        },
        methods: {
            className(hex) {
                return {
                    highlight: hex_equal(hex, this.highlight),
                    ['q-axis-same']: hex.col === this.highlight.col,
                    ['r-axis-same']: hex.row === this.highlight.row,
                };
            },
        },
    });
}


Diagram('angles', {
    computed: {
        angleAdjust() {
            // The rotation allows a little bit of overshoot to make
            // the animation cooler but for the sample code we don't
            // want to allow the overshoot
            let angle = -Math.round(this.$root.rotation);
            if (angle > 30) angle = 30;
            if (angle < 0) angle = 0;
            return angle === 0 ? "" : ` - ${angle}°`;
        },
        cornerMarkers() {
            return [0, 60, 120, 180, 240, 300].map(angle => ({
                style: {transform: `rotate(${angle}deg) translate(100px,0px) rotate(${-angle}deg)`},
            }));
        },
        cornerLabels() {
            return [0, 30, 60, 90, 120, 150, 180, 210, 240, 270, 300, 330].map(angle => ({
                angle: Math.round(angle),
                style: {transform: `rotate(${angle}deg) translate(112px,0px) rotate(${-angle}deg)`},
            }));
        },
        dashedPath() {
            let corners = layouts.flat.polygonCorners(new Hex(0, 0, 0));
            let a = corners[5];
            let b = corners[0];
            return `M ${a.x.toFixed(0)} ${a.y.toFixed(0)} L 0 0 L ${b.x.toFixed(0)} ${b.y.toFixed(0)}`;
        },
        arcPath() {
            let corners = layouts.flat.polygonCorners(new Hex(0, 0, 0));
            let a = corners[1], b = corners[2], c = corners[3];
            let p = {x: mix(b.x, a.x, 0.3), y: mix(b.y, a.y, 0.3)};
            let q = {x: mix(b.x, c.x, 0.3), y: mix(b.y, c.y, 0.3)};
            return `M ${p.x.toFixed(0)} ${p.y.toFixed(0)} A 30 30 0 0 1 ${q.x.toFixed(0)} ${q.y.toFixed(0)}`;
        }
    },
});


Diagram('sizes', {
    data: {
        show: 'none',
        colors: {
            fade: "hsl(0, 0%, 85%)",
            gray: "hsl(0, 0%, 70%)",
            dark: "hsl(0, 0%, 60%)",
            red: "hsl(0, 40%, 60%)",
        },
        layouts: [
            {
                title: "Flat-top orientation", name: "flat", rotation: 0,
                height: Math.sqrt(3) * 50,
                width: 2 * 50,
            },
            {
                title: "Pointy-top orientation", name: "pointy", rotation: -30,
                height: 2 * 50,
                width: Math.sqrt(3) * 50,
            },
        ],
    },
    methods: {
        styleColors(layout) {
            const {colors, show} = this;
            // TODO: find a better way to express this logic
            return {
                '--heightColor': show === 'height' ? colors.red
                    : (layout.name === 'pointy' && show === 'circumradius')
                    || (layout.name === 'flat' && show === 'inradius')
                    ?  colors.dark : colors.gray,
                '--widthColor': show === 'width' ? colors.red
                    : (layout.name === 'flat' && show == 'circumradius')
                    || (layout.name === 'pointy' && show === 'inradius')
                    ? colors.dark : colors.gray,
                '--circumradiusColor': show === 'none' ? colors.red
                    : show === 'circumradius' ? colors.dark 
                    : (show === 'height' && layout.name === 'pointy')
                    || (show === 'width' && layout.name === 'flat')
                    ? colors.red : show === 'inradius'? colors.fade : colors.gray,
                '--inradiusColor': show === 'inradius' ? colors.dark : colors.fade,
            };
        },
        heightLabel(layout) {
            if (layout.name === 'pointy') {
                return this.show === 'circumradius' ? "height = 2 ✕ circumradius" : "height = 2 ✕ size";
            } else {
                return this.show === 'inradius' ? "height = 2 ✕ inradius" : "height = √3 ✕ size";
            }
        },
        widthLabel(layout) {
            if (layout.name === 'pointy') {
                return this.show === 'inradius' ? "width = 2 ✕ inradius" : "width = √3 ✕ size";
            } else {
                return this.show === 'circumradius' ? "width = 2 ✕ circumradius" : "width = 2 ✕ size";
            }
        },
    },
});


Diagram('spacing', {
    data: {
        hexes: [new Hex(-1, 0, 1), new Hex(0, -1, 1), new Hex(1, -1, 0), new Hex(0, 0, 0)],
        corners: [
            new Hex(-1/3, -1/3, 2/3), new Hex(1/3, -2/3, 1/3),
            new Hex(-5/3, 1/3, 4/3), new Hex(5/3, -4/3, -1/3),
            new Hex(-2/3, -2/3, 4/3), new Hex(2/3, -4/3, 2/3),
            new Hex(-4/3, -1/3, 5/3), new Hex(4/3, -5/3, 1/3),
            new Hex(-1/3, -4/3, 5/3), new Hex(1/3, -5/3, 4/3),
            new Hex(-1/3, 2/3, -1/3), new Hex(1/3, 1/3, -2/3),
            new Hex(-4/3, 2/3, 2/3), new Hex(4/3, -2/3, -2/3),
            new Hex(-2/3, 1/3, 1/3), new Hex(2/3, -1/3, -1/3),
            ],
    },
    components: {
        'a-measurement': {
            props: [
                'label',
                /* {a,b} - point a, point b
                   {w,h} - horizontal, vertical position
                   {p,f} - pointy, flat positions */
                'awp', 'ahp', 'bwp', 'bhp',
                'awf', 'ahf', 'bwf', 'bhf'
            ],
            template: `<g class="measurement">
               <a-arrow :w="3" :skip="0" :A="midpoint" :B="A" :withHead="true" />
               <a-arrow :w="3" :skip="0" :A="midpoint" :B="B" :withHead="true" />
               <text dy="0.5em" :transform="\`translate(\${midpoint.x.toFixed(0)},\${midpoint.y.toFixed(0)}) rotate(\${rotation}) translate(0,8)\`">{{label}}</text>
            </g>
            `,
            computed: {
                rotation() {
                    return 180 / Math.PI * Math.atan2(this.B.y - this.A.y, this.B.x - this.A.x);
                },
                A() {
                    const f = this.$root.rotation / -30;
                    const w = mix(50, Math.sqrt(3)/2*50, f);
                    const h = mix(Math.sqrt(3)/2*50, 50, f);
                    let aw = mix(this.awf, this.awp, f),
                        ah = mix(this.ahf, this.ahp, f);
                    return new Point(aw * w, ah * h);
                },
                B() {
                    const f = this.$root.rotation / -30;
                    const w = mix(50, Math.sqrt(3)/2*50, f);
                    const h = mix(Math.sqrt(3)/2*50, 50, f);
                    let bw = mix(this.bwf, this.bwp, f),
                        bh = mix(this.bhf, this.bhp, f);
                    return new Point(bw * w, bh * h);
                },
                midpoint() {
                    let {A, B} = this;
                    return new Point(mix(A.x, B.x, 0.5), mix(A.y, B.y, 0.5));
                },
            },
        },
    },
    computed: {
        layout() {
            return this.$root.orientation.layout;
        },
        lines() {
            function format_quarters(a) {
                // Format a/4 as a mixed numeral
                var suffix = ["", "¼", "½", "¾"][a % 4];
                var prefix = Math.floor(a/4);
                if (prefix == 0 && suffix != "") { prefix = ""; }
                return prefix + suffix;
            }
            
            const f = this.$root.rotation / -30;
            const w = mix(50, Math.sqrt(3)/2*50, f);
            const h = mix(Math.sqrt(3)/2*50, 50, f);
            
            let lines = [];

            // vertical
            for (let i = -5; i <= 5; i++) {
                lines.push({
                    x1: (i-f)*w, y1: (-6.5+f)*h,
                    x2: (i-f)*w, y2: 2*h,
                    label: format_quarters(i+5)+"w",
                    opacity: mix(1, i&1, f)
                });
            }
            // horizontal
            for (let i = -6; i <= 2; i++) {
                lines.push({
                    x1: (-5.5-f)*w, y1: (i+f)*h,
                    x2: (5-f)*w, y2: (i+f)*h,
                    label: format_quarters(i+6)+"h",
                    opacity: mix(1-(i&1), 1, f) * Math.min(1, 3-i-f)
                });
            }
            return lines;
        },
    },
    methods: {
        center: hexCenter,
    },
});


OffsetDiagram('grid-offset-odd-r', 'pointy', "“odd-r” horizontal layout<br/>shoves odd rows right",
              OffsetCoord.roffsetToCube.bind(null, OffsetCoord.ODD));
OffsetDiagram('grid-offset-even-r', 'pointy', "“even-r” horizontal layout<br/>shoves even rows right",
              OffsetCoord.roffsetToCube.bind(null, OffsetCoord.EVEN));
OffsetDiagram('grid-offset-odd-q', 'flat', "“odd-q” vertical layout<br/>shoves odd columns down",
              OffsetCoord.qoffsetToCube.bind(null, OffsetCoord.ODD));
OffsetDiagram('grid-offset-even-q', 'flat', "“even-q” vertical layout<br/>shoves even columns down",
              OffsetCoord.qoffsetToCube.bind(null, OffsetCoord.EVEN));


for (let {id, layout, converter, grid} of
     [{id: 'grid-doubled-r', layout: 'pointy', converter: 'rdoubledToCube', grid: makeRDoubledRectangularShape(0, 13, 0, 6)},
      {id: 'grid-doubled-q', layout: 'flat', converter: 'qdoubledToCube', grid: makeQDoubledRectangularShape(0, 7, 0, 11)}]) {
    Diagram(id, {
        mixins: [mixins.highlight],
        layout: layout,
        data: {
            highlight: (() => { let hex = new DoubledCoord(2, 2)[converter](); hex.col = 2; hex.row = 2; return hex; })(),
            grid: grid,
        },
        methods: {
            className(hex) {
                return {
                    highlight: hex_equal(hex, this.highlight),
                    ['q-axis-same']: hex.col === this.highlight.col,
                    ['r-axis-same']: hex.row === this.highlight.row,
                };
            },
        },
    });
}


for (let id of ['grid-cube', 'grid-axial']) {
    Diagram(id, {
        mixins: [mixins.highlight],
        data: {
            gridSize: 3,
        },
        computed: {
            layout() { return this.$root.orientation.layout; },
            grid()   { return makeHexagonalShape(this.gridSize); },
        },
        watch: {
            highlight(value) {
                let otherDiagram = id === 'grid-cube'? diagrams.grid_axial : diagrams.grid_cube;
                if (otherDiagram) otherDiagram.highlight = value;
            },
        },
        methods: {
            className(hex) {
                let result = {
                    origin: hex.len() === 0,
                    highlight: hex_equal(hex, this.highlight),
                };
                for (let axis of ['q', 'r', 's']) {
                    if (hex[axis] === this.highlight[axis]) {
                        result[`${axis}-axis-same`] = true;
                    }
                }
                return result;
            },
        },
    });
}


Diagram('conversions-axial', {
    mixins: [mixins.highlight],
    data: {
        grid: makeHexagonalShape(2),
    },
    methods: {
        className(hex) {
            return {
                origin: hex.len() === 0,
                highlight: hex_equal(hex, this.highlight),
            };
        },
    },
});


orientations.offsetConversions = makeAnimatedLayout();
Diagram('conversions-offset', {
    mixins: [mixins.highlight],
    layout: 'offsetConversions',
    data: {
        variant: 'odd_r',
        output: 'axial',
    },
    watch: {
        variant: {
            immediate: true,
            handler() {
                orientations.offsetConversions.layout = this.variant.slice(-1) === 'r'? 'pointy' : 'flat';
            },
        },
    },
    computed: {
        grid() {
            const layout = orientations.offsetConversions.layout;
            const oddeven = {o: OffsetCoord.ODD, e: OffsetCoord.EVEN}[this.variant[0]];
            const converter = OffsetCoord[this.variant.slice(-1)+'offsetToCube'].bind(null, oddeven);
            return makeRectangularShape(-2, 2, -2, 2, converter);
        },
        output_constructor() {
            return this.output === 'axial'? "Hex(q, r)" : "Cube(q, r, -q-r)";
        },
    },
    methods: {
        className(hex) {
            return {
                origin: hex.len() === 0,
                highlight: hex_equal(hex, this.highlight),
            };
        },
    },
});


['neighbors-cube', 'neighbors-axial'].forEach(id => Diagram(id, {
    mixins: [mixins.highlight],
    data: {
        grid: makeHexagonalShape(1),
        neighbors: Array.from(Array(6).keys()).map(Hex.direction),
    },
    methods: {
        className(hex) {
            return {
                origin: hex.len() === 0,
                highlight: hex.len() > 0 && hex_equal(hex, this.highlight),
            };
        },
    },
}));


Diagram('neighbors-diagonal', {
    mixins: [mixins.highlight],
    data: {
        neighbors: Array.from(Array(6).keys()).map(d => Hex.diagonals[d]),
    },
    computed: {
        grid() {
            return makeHexagonalShape(1).concat(this.neighbors);
        },
    },
    methods: {
        className(hex) {
            return {
                origin: hex.len() === 0,
                faded: hex.len() === 1,
                highlight: hex.len() > 1 && hex_equal(hex, this.highlight),
            };
        },
    },
});


orientations.offsetNeighbors = makeAnimatedLayout();
Diagram('neighbors-offset', {
    mixins: [mixins.highlight],
    layout: 'offsetNeighbors',
    data: {
        variant: 'odd_r',
        highlightParity: 1,
        grid: makeHexagonalShape(1),
        neighbors: Array.from(Array(6).keys()).map(Hex.direction),
    },
    watch: {
        variant: {
            immediate: true,
            handler() {
                orientations.offsetNeighbors.layout = this.variant.slice(-1) === 'r'? 'pointy' : 'flat';
            },
        },
    },
    computed: {
        parityField() {
            return this.variant.slice(-1) === 'r'? 'row' : 'col';
        },
    },
    methods: {
        xTranslate(parity) {
            return {
                o: [300, 900][parity],
                e: [900, 300][parity],
            }[this.variant[0]];
        },
        setHighlightWithParity(parity) {
            return (hex, fractionalHex) => {
                this.setHighlight(hex, fractionalHex);
                this.highlightParity = parity;
            };
        },
        neighbor(direction, parity) {
            function fmt(v) {
                if (v > 0) v = '+' + v;
                return ('   ' + v.toString()).slice(-2);
            }
            let rq = this.variant.slice(-1);
            let evenodd = OffsetCoord[this.variant.split('_')[0].toUpperCase()];
            let base = OffsetCoord[rq+'offsetToCube'](evenodd, new OffsetCoord(parity, parity));
            let neighbor = base.add(direction);
            let a = OffsetCoord[rq+'offsetFromCube'](evenodd, base);
            let b = OffsetCoord[rq+'offsetFromCube'](evenodd, neighbor);
            return `${fmt(b.col-a.col)}, ${fmt(b.row-a.row)}`;
        },
        className(hex, parity) {
            return {
                origin: hex.len() === 0,
                highlight:
                    this.highlightParity === parity
                    && hex.len() > 0
                    && hex_equal(hex, this.highlight),
            };
        },
    },
});
        

orientations.doubledNeighbors = makeAnimatedLayout();
Diagram('neighbors-doubled', {
    mixins: [mixins.highlight],
    layout: 'doubledNeighbors',
    data: {
        variant: 'width',
        grid: makeHexagonalShape(1),
        neighbors: Array.from(Array(6).keys()).map(Hex.direction),
    },
    watch: {
        variant: {
            immediate: true,
            handler() {
                orientations.doubledNeighbors.layout = this.variant === 'width'? 'pointy' : 'flat';
            },
        },
    },
    methods: {
        neighbor(direction) {
            function fmt(v) {
                if (v > 0) v = '+' + v;
                return ('   ' + v.toString()).slice(-2);
            }
            let hex = this.variant === 'width'
                ? DoubledCoord.rdoubledFromCube(direction)
                : DoubledCoord.qdoubledFromCube(direction);
            return `${fmt(hex.col)}, ${fmt(hex.row)}`;
        },
        className(hex) {
            return {
                origin: hex.len() === 0,
                highlight: hex.len() > 0
                    && hex_equal(hex, this.highlight),
            };
        },
    },
});



Diagram('distances', {
    mixins: [mixins.highlight],
    data: {
        highlight: new Hex(2, 0, -2),
    },
    computed: {
        grid() {
            return makeHexagonalShape(4);
        },
    },
    methods: {
        className(hex) {
            let result = {
                highlight: hex.len() === this.highlight.len(),
            };
            result['d'+hex.len()] = true;
            for (let axis of ['q', 'r', 's']) {
                if (Math.abs(hex[axis]) === hex.len()) {
                    result[`${axis}-axis-max`] = true;
                }
            }
            return result;
        },
    },
});


Diagram('line-drawing', {
    mixins: [mixins.highlight],
    data: {
        gridSize: 6,
        start: new Hex(-5, 0, 5),
        highlight: new Hex(4, 1, -5),
    },
    computed: {
        grid() {
            return makeHexagonalShape(this.gridSize);
        },
        N() {
            return this.start.distance(this.highlight);
        },
        interpolationPoints() {
            return hexLineFractional(this.start, this.highlight);
        },
        line() {
            let layout = layouts.flat;
            let p = layout.hexToPixel(this.start);
            let q = layout.hexToPixel(this.highlight);
            return `M ${p.x} ${p.y} L ${q.x} ${q.y}`;
        },
    },
    methods: {
        center: hexCenter,
        className(hex) {
            return {
                highlight: this.interpolationPoints.some(p => hex_equal(p.round(), hex)),
            };
        },
    },
});


Diagram('range-coordinate', {
    mixins: [mixins.highlight],
    data: {
        grid: makeHexagonalShape(5),
        highlight: new Hex(3, 0, -3),
        constraints: ["s ≥", "r ≤", "q ≥", "s ≤", "r ≥", "q ≤"],
    },
    computed: {
        distance() {
            return this.highlight.len();
        },
    },
    methods: {
        className(hex) {
            return {shadow: hex.len() > this.distance};
        },
        format(char, distance) {
            /* NOTE: this is a workaround, because my current
            xhtml+xslt+smartypants.pl+vue system doesn't let me put
            quotes into the templates that are defined in the xhtml */
            return char === '≥'? -distance : distance;
        },
    },
});


Diagram('range-intersection', {
    mixins: [mixins.highlight],
    data: {
        distance: 3,
        regionCenter: new Hex(-4, 4, 0),
        highlight: new Hex(1, 0, -1),
        grid: makeHexagonalShape(7),
    },
    methods: {
        className(hex) {
            return {
                highlight: hex_equal(this.highlight, hex),
                center: hex_equal(this.regionCenter, hex),
                regionA: this.highlight.subtract(hex).len() <= this.distance,
                regionB: this.regionCenter.subtract(hex).len() <= this.distance,
            };
        },
    },
});


['movement-range', 'pathfinding'].forEach(id => Diagram(id, {
    mixins: [mixins.highlight, mixins.pathfinding],
    data: {
        gridSize: 5,
        distanceLimit: id === 'pathfinding'? 1000 : 4,
        highlight: new Hex(4, 0, -4),
        walls: {
            [new Hex(2, -1, -1)]: true,
            [new Hex(2, 0, -2)]: true,
            [new Hex(0, 2, -2)]: true,
            [new Hex(-1, 2, -1)]: true,
            [new Hex(-1, 1, 0)]: true,
            [new Hex(1, -1, 0)]: true,
            [new Hex(1, 2, -3)]: true,
            [new Hex(1, -3, 2)]: true,
            [new Hex(0, -2, 2)]: true,
            [new Hex(-1, -1, 2)]: true,
            [new Hex(2, 1, -3)]: true,
            [new Hex(-2, 1, 1)]: true,
            [new Hex(-3, 2, 1)]: true,
            [new Hex(-4, 3, 1)]: true,
            [new Hex(-5, 4, 1)]: true,
        },
    },
    computed: {
        grid() {
            return makeHexagonalShape(this.gridSize);
        },
    },
    methods: {
        className(hex) {
            return {
                highlight: hex_equal(hex, this.highlight),
                path: this.reconstructedPath.find(h => hex_equal(h, hex)),
                origin: hex.len() === 0,
                wall: this.isWall(hex),
                shadow: this.isShadow(hex),
            };
        },
        isWall(hex) {
            return hex.len() > this.gridSize || !!this.walls[hex];
        },
        isShadow(hex) {
            return this.breadthFirstSearch.cost_so_far[hex] > this.distanceLimit;
        },
    },
}));


Diagram('rotation', {
    mixins: [mixins.highlight],
    data: {
        grid: makeHexagonalShape(5),
        highlight: new Hex(2, -3, 1),
    },
    computed: {
        origin() { return new Hex(0, 0, 0); },
        left1() { return this.highlight.rotateLeft(); },
        left2() { return this.left1.rotateLeft(); },
        right1() { return this.highlight.rotateRight(); },
        right2() { return this.right1.rotateRight(); },
        opposite() { return this.highlight.scale(-1); },
    },
    methods: {
        center: hexCenter,
        className(hex) {
            return {
                origin: hex.len() === 0,
                highlight: hex_equal(hex, this.highlight),
                left1: hex_equal(hex, this.left1),
                left2: hex_equal(hex, this.left2),
                right1: hex_equal(hex, this.right1),
                right2: hex_equal(hex, this.right2),
                opposite: hex_equal(hex, this.opposite),
            };
        },
    },
});


Diagram('reflection', {
    mixins: [mixins.highlight],
    data: {
        axis: 'r',
        grid: makeHexagonalShape(5),
        highlight: new Hex(-1, -3, 4),
    },
    computed: {
        primary() {
            switch (this.axis) {
            case 'q': return this.swappedCoordinate(this.highlight, 'r', 's');
            case 'r': return this.swappedCoordinate(this.highlight, 'q', 's');
            case 's': return this.swappedCoordinate(this.highlight, 'q', 'r');
            }
            return null;
        },
        secondary() { return this.primary.scale(-1); },
        axisRotation() { return {r: 30, s: 150, q: 270}[this.axis]; },
    },
    methods: {
        center: hexCenter,
        swappedCoordinate(hex, coord1, coord2) {
            let result = hex.scale(1);
            let a1 = hex[coord1],
                a2 = hex[coord2];
            result[coord1] = a2;
            result[coord2] = a1;
            return result;
        },
        className(hex) {
            const highlight = hex_equal(hex, this.highlight),
                  primary = hex_equal(hex, this.primary),
                  secondary = hex_equal(hex, this.secondary)
                  || hex_equal(hex, this.highlight.scale(-1));
            return {
                origin: hex.len() === 0,
                [this.axis + '-primary']: highlight || primary,
                [this.axis + '-secondary']: secondary,
                [this.axis + '-axis-same']:
                hex[this.axis] === 0 || highlight || primary || secondary,
            };
        },
    },
});


['rings', 'spiral'].forEach(id => Diagram(id, {
    mixins: [mixins.highlight],
    data: {
        grid: makeHexagonalShape(5),
        highlight: new Hex(3, 0, -3),
    },
    computed: {
        radius() {
            return this.highlight? this.highlight.len() : 0;
        },
        ring() {
            let results = [];
            for (let radius = (id === 'rings')? this.radius : 1; radius <= this.radius; radius++) {
                results = results.concat(hexRing(radius));
            }
            return results;
        },
        arrows() {
            return this.ring.map((hex, index) =>
                                 ({
                                     A: index === 0? new Hex(0, 0, 0) : this.ring[index-1],
                                     B: hex
                                 }));
        },
    },
    methods: {
        center: hexCenter,
        className(hex) {
            return {highlight:
                    (id === 'spiral' && hex.len() === 0)
                    || this.ring.some(h => hex_equal(hex, h))};
        },
    },
}));

Diagram('spiral-coordinates', {
    data: {
        radius: 5,
        show: 'none',
    },
    computed: {
        grid() {
            return makeHexagonalShape(5);
        },
        rings() {
            let results = [[new Hex(0, 0, 0)]];
            for (let radius = 1; radius <= this.radius; radius++) {
                results.push(hexRing(radius));
            }
            return results;
        },
        spiral() {
            return this.rings.flat(1);
        },
        ringStarts() {
            return this.rings.map((hexes) => hexes[0]);
        },
        ringSegmentStart() {
            let results = [];
            for (let segment = 0; segment < 6; segment++) {
                results.push([]);
                for (let radius = 1; radius <= this.radius; radius++) {
                    results[segment].push(hexRing(radius)[segment * radius]);
                }
            }
            return results;
        },
        ringSegmentFull() {
            let results = [];
            for (let radius = 1; radius <= this.radius; radius++) {
                let ring = hexRing(radius);
                for (let segment = 0; segment < 6; segment++) {
                    results.push(ring.slice(segment * radius, (segment + 1) * radius));
                }
            }
            return results;
        },
    },
    watch: {
        show(newValue, oldValue) {
            // Don't revert to the default setting; keep it at the new setting until explicitly changed
            if (newValue === 'none') {
                this.show = oldValue;
            }
        }
    },
    methods: {
        label(hex) {
            let spiral = this.spiralCoordinate(hex),
                radius = hex.len(),
                startIndex = this.radiusToStartIndex(radius),
                position = (spiral - startIndex) % radius;

            switch (this.show) {
                case 'radius': return radius;
                case 'position': return position || 0;
                default: return spiral;
            }
        },
        radiusToStartIndex(radius) {
            return radius === 0? 0 : 1 + 3 * radius * (radius - 1);
        },
        indexToRadius(index) {
            return Math.floor((Math.sqrt(12 * index - 3) + 3) / 6);
            // TODO: same as Math.round(Math.sqrt(12 * index - 3) / 6) ??
        },
        spiralCoordinate(hex) {
            let radius = hex.len();
            let ringHexes = hexRing(radius);
            let index = this.radiusToStartIndex(radius) + ringHexes.findIndex((h) => hex_equal(hex, h));
            return index;
        },
        path(...hexLists) {
            let d = [];
            for (let hexes of hexLists) {
                let command = 'M';
                for (let hex of hexes) {
                    let {x, y} = layouts.flat.hexToPixel(hex);
                    d.push(command, x, y);
                    command = 'L';
                }
            }
            return d.join(" ");
        },
    },
});

new Vue({
    el: "#calculation-spiral",
    data: { N: 4 },
    computed: {
        count() { return 1 + 3 * this.N * (this.N + 1); }
    },
});


Diagram('field-of-view', {
    mixins: [mixins.highlight],
    data: {
        start: new Hex(0, 0, 0),
        highlight: new Hex(3, -5, 2),
        grid: makeHexagonalShape(8),
        walls: {
            [new Hex(3, -3, 0)]: true,
            [new Hex(2, -3, 1)]: true,
            [new Hex(1, -3, 2)]: true,
            [new Hex(-2, 0, 2)]: true,
            [new Hex(-3, 0, 3)]: true,
            [new Hex(-3, 2, 1)]: true,
            [new Hex(-4, 2, 2)]: true,
            [new Hex(-3, 3, 0)]: true,
            [new Hex(-3, 4, -1)]: true,
            [new Hex(0, 2, -2)]: true,
            [new Hex(0, 3, -3)]: true,
            [new Hex(0, 4, -4)]: true,
            [new Hex(0, 5, -5)]: true,
            [new Hex(0, -3, 3)]: true,
            [new Hex(0, -4, 4)]: true,
            [new Hex(0, -5, 5)]: true,
            [new Hex(0, -6, 6)]: true,
            [new Hex(4, 2, -6)]: true,
            [new Hex(3, 3, -6)]: true,
            [new Hex(-5, 2, 3)]: true,
            [new Hex(-4, 0, 4)]: true,
            [new Hex(-5, 0, 5)]: true,
            [new Hex(-6, 0, 6)]: true,
            [new Hex(-7, 0, 7)]: true,
            [new Hex(-8, 0, 8)]: true,
        },
    },
    computed: {
        interpolationPoints() {
            let fracHexes = hexLineFractional(this.start, this.highlight);
            let points = fracHexes.map(frac => hexCenter(frac));
            let blocked = fracHexes.map(frac => this.blocked(frac));
            let firstBlocked = blocked.indexOf(true);
            return points.map((p, index) => ({
                point: p,
                blocked: blocked[index],
                visible: firstBlocked < 0 || index < firstBlocked
            }));
        },
    },
    methods: {
        center: hexCenter,
        className(hex) {
            return {
                origin: hex.len() === 0,
                highlight: hex_equal(hex, this.highlight),
                shadow: this.inShadow(hex),
                wall: this.walls[hex],
            };
        },
        inShadow(hex) {
            return hexLineFractional(this.start, hex)
                .map(frac => this.blocked(frac))
                .some(b => b);
        },
        blocked(frac) {
            /* I want to know if all perturbations of this fractional hex are blocked;
               in the case where the fractional hex just happened to land
               on an edge or corner, and ANY of the hexes are open, we'll say it's open */
            for (let direction = 0; direction < 6; direction++) {
                let rounded = frac.add(Hex.direction(direction).scale(1e-2)).round();
                if (!this.walls[rounded]) { return false; }
            }
            return true;
        },
    },
});
    

orientations.hexToPixel = makeAnimatedLayout();
Diagram('hex-to-pixel-axial', {
    layout: 'hexToPixel',
    data: {
        grid: [new Hex(0, 0, 0), new Hex(1, 0, -1), new Hex(0, 1, -1), new Hex(1, 1, -2)],
        layout: orientations.hexToPixel.layout,
    },
    watch: {
        layout() {
            orientations.hexToPixel.layout = this.layout;
        },
    },
    methods: {
        center: hexCenter,
        almost(hex) {
            let center = hexCenter(hex);
            return {x: center.x * 0.9, y: center.y * 0.9};
        },
    },
});

new Vue({
    el: '#code-hex-to-pixel-offset',
    template: "#template-code-hex-to-pixel-offset",
    data: { conversion: 'odd-r', },
});

new Vue({
    el: '#code-hex-to-pixel-mod-pixelsize',
    template: "#template-code-hex-to-pixel-mod-pixelsize",
    data: {
        desiredWidth: 17,
        desiredHeight: 24,
        layout: orientations.hexToPixel.layout,
    },
    watch: {
        layout() {
            orientations.hexToPixel.layout = this.layout;
        },
    },
});


orientations.pixelToHex = makeAnimatedLayout();
Diagram('pixel-to-hex', {
    mixins: [mixins.highlight],
    layout: 'pixelToHex',
    data: {
        grid: makeRectangularShape(
            -7, 7, -3, 3,
            OffsetCoord.qoffsetToCube.bind(null, OffsetCoord.ODD)
        ).concat(makeRectangularShape(
            0, 7, 4, 5,
            OffsetCoord.qoffsetToCube.bind(null, OffsetCoord.ODD)
        )).concat(makeRectangularShape(
            -7, 0, -5, -4,
            OffsetCoord.qoffsetToCube.bind(null, OffsetCoord.ODD)
        )),
        layout: orientations.pixelToHex.layout,
    },
    created() {
        this.setHighlight(new Hex(2, 1, -3), new Hex(2.2, 1.2, -3.4));
    },
    watch: {
        layout() {
            orientations.pixelToHex.layout = this.layout;
        },
    },
    computed: {
        axes() {
            const P = hexCenter;
            const {q, r} = this.highlightFractional;
            return [
                ['q', P(new Hex(-10, 0, 10)), P(new Hex(10, 0, -10))],
                ['r', P(new Hex(0, -10, 10)), P(new Hex(0, 10, -10))],
                ['q', P(new Hex(q, r, -q-r)), P(new Hex(0, r, -r))],
                ['r', P(new Hex(q, r, -q-r)), P(new Hex(q, 0, -q))],
            ];
        },
        transform() {
            let p = hexCenter(this.highlightFractional);
            return `translate(${p.x},${p.y})`;
        },
    },
    methods: {
        className(hex) {
            return {
                ['q-axis-same']: (hex.r === 0 || hex.r === this.highlight.r) && hex.q === this.highlight.q,
                ['r-axis-same']: (hex.q === 0 || hex.q === this.highlight.q) && hex.r === this.highlight.r,
                highlight: hex_equal(hex, this.highlight),
            };
        },
    },
});

new Vue({
    el: '#code-pixel-to-hex-mod-pixelsize',
    template: "#template-code-pixel-to-hex-mod-pixelsize",
    data: {
        desiredWidth: 17,
        desiredHeight: 24,
        layout: orientations.hexToPixel.layout,
    },
    watch: {
        layout() {
            orientations.hexToPixel.layout = this.layout;
        },
    },
});

Diagram('map-storage', {
    mixins: [mixins.highlight, AnimationMixin(500)],
    layout: 'pointy',
    data: {
        phase: 0, /* 1 means slide the rows */
        shape: 'hexagon',
        shapes: ['rectangle', 'hexagon', 'rhombus', 'down-triangle', 'up-triangle'],
        highlight: new Hex(3, 2, -5),
    },
    computed: {
        grid() {
            switch(this.shape) {
            case 'rectangle': return makeRectangularShape(
                0, 6, 0, 6,
                OffsetCoord.roffsetToCube.bind(null, OffsetCoord.ODD)
            );
            case 'hexagon': return makeHexagonalShape(3)
                    .map(hex => hex.add(new Hex(3, 3, -6)));
            case 'rhombus': return makeRhombusShape(7);
            case 'down-triangle': return makeDownTriangularShape(7);
            case 'up-triangle': return makeUpTriangularShape(7);
            }
            return [];
        },
        qRange() { return closedInterval(bounds(this.grid, 'q')); },
        rRange() { return closedInterval(bounds(this.grid, 'r')); },
        leftDiagramX() { return this.shape === 'down-triangle' ? 300 : this.shape === 'up-triangle' ? -250 : 0; },
        rightDiagramX() { return this.shape === 'rectangle'? 1350 : 1800; },
        styleHideWhenSliding() { return {'opacity': 1.0 - this.phaseInterpolation}; },
    },
    methods: {
        Hex: Hex,
        firstQOfRow(r) {
            return bounds(this.grid
                          .filter(h => h.r === r),
                          'q').min;
        },
        colX(q, r) {
            let leftQ = mix(this.qRange[0], this.firstQOfRow(r), this.phaseInterpolation);
            return 150 + 150 * (q - leftQ);
        },
        rowY(r) {
            let start = mix(150, 120, this.phaseInterpolation);
            let spacing = mix(150, 158, this.phaseInterpolation);
            return start + spacing * (r - this.rRange[0]);
        },
        wall(hex) {
            return !this.grid.some(h => hex_equal(h, hex));
        },
        styleRules(hex) {
            if (this.wall(hex)) {
                return this.styleHideWhenSliding;
            } else {
                return {};
            }
        },
        className(hex) {
            return {
                highlight: hex_equal(hex, this.highlight),
                firstcol: hex.q === this.firstQOfRow(hex.r),
                samerow: hex.r === this.highlight.r,
                wall: this.wall(hex),
            };
        },
    },
});


Diagram('wraparound', {
    mixins: [mixins.highlight],
    data: {
        highlight: (() => { let hex = new Hex(0, 0, 0); hex.base = hex; return hex; })(),
    },
    computed: {
        grid() {
            let base = makeHexagonalShape(2);
            base.forEach(hex => { hex.base = hex; hex.parity = 0; });

            let hexes = [].concat(base);
            let center = new Hex(5, -2, -3);
            for (let direction = 0; direction < 6; direction++, center = center.rotateRight()) {
                let mirror = makeHexagonalShape(2).map(hex => hex.add(center));
                mirror.forEach((hex, i) => { hex.base = base[i]; hex.parity = direction % 2; });
                hexes = hexes.concat(mirror);
            }

            return hexes;
        },
    },
    methods: {
        className(hex) {
            return {
                origin: hex.base.len() === 0,
                highlight: hex_equal(hex.base, this.highlight.base),
                wrapped: hex.len() > 2,
                parity: hex.parity,
            };
        },
    },
});


new Vue({
    el: "#diagram-cube-to-hex",
    template: "#template-diagram-cube-to-hex",
    mixins: [AnimationMixin(2500)],
    data: {phase: 0},
    computed: {
        cubes() {
            const scale = 100;
            const base_z = -3;
            const limit = 4;
            let cubes = [];
            for (let q = -2; q < 2; q++) {
                for (let r = -2; r < 2; r++) {
                    for (let s = -2; s < 2; s++) {
                        let z = q + r + s;
                        if (base_z <= z && z <= limit) {
                            let x = scale * (r-q) * Math.sqrt(3)/2,
                                y = scale * (0.5*(r+q) - s);
                            let startTime = 1 - (z-base_z) / (limit-base_z);
                            cubes.push({
                                q, r, s, x, y, z,
                                startTime,
                            });
                        }
                    }
                }
            }
            return cubes;
        },
        rhombusTop() { return this.rhombus(1); },
        rhombusLeft() { return this.rhombus(3); },
        rhombusRight() { return this.rhombus(5); },
        hexagonPoints() {
            return layouts.pointy
                .polygonCorners(new Hex(0,0,0))
                .map(p=>`${p.x.toFixed(0)},${p.y.toFixed(0)}`)
                .join(" ");
        },
    },
    methods: {
        localTime(cube) {
            return Math.max(0, 2 * this.phaseInterpolation - cube.startTime);
        },
        opacity(cube) {
            return Math.sqrt(Math.max(0, 1 - this.localTime(cube)));
        },
        position(cube) {
            let t = this.localTime(cube);
            if (cube.z === -3) { t = 0; }
            return `translate(${cube.x * (1+t)}, ${cube.y - 100 * t + 1000 * t * t})`;
        },
        /* Each cube has three visible faces, drawn as a rhombus */
        rhombus(direction) {
            let points = [0, 0];
            for (let i = 0; i < 3; i++) {
                let corner = layouts.pointy.hexCornerOffset((i + direction) % 6);
                points.push(corner.x.toFixed(0), corner.y.toFixed(0));
            }
            return points.join(' ');
        },
        /* Six axis lines */
        axis(direction /* 1 to 6, due to vue's counting */ ) {
            const r = 450,
                  a = (0.5+direction)/6 * 2*Math.PI,
                  a2 = (0.58+direction)/6 * 2*Math.PI;
            let x = r * Math.cos(a),
                y = r * Math.sin(a),
                x2 = (r-20) * Math.cos(a2),
                y2 = (r-20) * Math.sin(a2);
            return {
                x, y, x2, y2,
                axis: "rqsrqs"[direction-1],
                label:
                this.phase > 0.5
                    ? ['+r', '-q', '+s', '-r', '+q', '-s'][direction-1]
                    : ['+z', '-x', '+y', '-z', '+x', '-y'][direction-1]
            };
        },
    },
});
