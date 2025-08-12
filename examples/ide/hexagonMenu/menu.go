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

package hexagonMenu

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/factoryFontFamily"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/examples/ide/manager"
	"github.com/helmutkemper/webassembly/examples/ide/rulesDensity"
	"github.com/helmutkemper/webassembly/examples/ide/rulesIcon"
	"github.com/helmutkemper/webassembly/examples/ide/rulesStage"
	"github.com/helmutkemper/webassembly/examples/ide/translate"
	"github.com/helmutkemper/webassembly/utilsDraw"
	"github.com/helmutkemper/webassembly/utilsText"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"image/color"
	"log"
	"syscall/js"
	"time"
)

const (
	//kFontAwesomeRepeat viewBox="0 0 512 512" source: https://fontawesome.com/icons/repeat?f=classic&s=solid
	kFontAwesomeRepeat = "M0 224c0 17.7 14.3 32 32 32s32-14.3 32-32c0-53 43-96 96-96l160 0 0 32c0 12.9 7.8 24.6 19.8 29.6s25.7 2.2 34.9-6.9l64-64c12.5-12.5 12.5-32.8 0-45.3l-64-64c-9.2-9.2-22.9-11.9-34.9-6.9S320 19.1 320 32l0 32L160 64C71.6 64 0 135.6 0 224zm512 64c0-17.7-14.3-32-32-32s-32 14.3-32 32c0 53-43 96-96 96l-160 0 0-32c0-12.9-7.8-24.6-19.8-29.6s-25.7-2.2-34.9 6.9l-64 64c-12.5 12.5-12.5 32.8 0 45.3l64 64c9.2 9.2 22.9 11.9 34.9 6.9s19.8-16.6 19.8-29.6l0-32 160 0c88.4 0 160-71.6 160-160z"

	//kFontAwesomeSquareRootVariable viewBox="0 0 576 512" source: https://fontawesome.com/icons/square-root-variable?f=classic&s=solid
	kFontAwesomeSquareRootVariable = "M282.6 78.1c8-27.3 33-46.1 61.4-46.1l200 0c17.7 0 32 14.3 32 32s-14.3 32-32 32L344 96 238.7 457c-3.6 12.3-14.1 21.2-26.8 22.8s-25.1-4.6-31.5-15.6L77.6 288 32 288c-17.7 0-32-14.3-32-32s14.3-32 32-32l45.6 0c22.8 0 43.8 12.1 55.3 31.8l65.2 111.8L282.6 78.1zM393.4 233.4c12.5-12.5 32.8-12.5 45.3 0L480 274.7l41.4-41.4c12.5-12.5 32.8-12.5 45.3 0s12.5 32.8 0 45.3L525.3 320l41.4 41.4c12.5 12.5 12.5 32.8 0 45.3s-32.8 12.5-45.3 0L480 365.3l-41.4 41.4c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3L434.7 320l-41.4-41.4c-12.5-12.5-12.5-32.8 0-45.3z"

	//kFontAwesomeScrewDriverWrench viewBox="0 0 512 512" source: https://fontawesome.com/icons/screwdriver-wrench?f=classic&s=solid
	kFontAwesomeScrewDriverWrench = "M78.6 5C69.1-2.4 55.6-1.5 47 7L7 47c-8.5 8.5-9.4 22-2.1 31.6l80 104c4.5 5.9 11.6 9.4 19 9.4l54.1 0 109 109c-14.7 29-10 65.4 14.3 89.6l112 112c12.5 12.5 32.8 12.5 45.3 0l64-64c12.5-12.5 12.5-32.8 0-45.3l-112-112c-24.2-24.2-60.6-29-89.6-14.3l-109-109 0-54.1c0-7.5-3.5-14.5-9.4-19L78.6 5zM19.9 396.1C7.2 408.8 0 426.1 0 444.1C0 481.6 30.4 512 67.9 512c18 0 35.3-7.2 48-19.9L233.7 374.3c-7.8-20.9-9-43.6-3.6-65.1l-61.7-61.7L19.9 396.1zM512 144c0-10.5-1.1-20.7-3.2-30.5c-2.4-11.2-16.1-14.1-24.2-6l-63.9 63.9c-3 3-7.1 4.7-11.3 4.7L352 176c-8.8 0-16-7.2-16-16l0-57.4c0-4.2 1.7-8.3 4.7-11.3l63.9-63.9c8.1-8.1 5.2-21.8-6-24.2C388.7 1.1 378.5 0 368 0C288.5 0 224 64.5 224 144l0 .8 85.3 85.3c36-9.1 75.8 .5 104 28.7L429 274.5c49-23 83-72.8 83-130.5zM56 432a24 24 0 1 1 48 0 24 24 0 1 1 -48 0z"

	//kFontAwesomeSliders viewBox="0 0 512 512" source: https://fontawesome.com/icons/sliders?f=classic&s=solid
	kFontAwesomeSliders = "M0 416c0 17.7 14.3 32 32 32l54.7 0c12.3 28.3 40.5 48 73.3 48s61-19.7 73.3-48L480 448c17.7 0 32-14.3 32-32s-14.3-32-32-32l-246.7 0c-12.3-28.3-40.5-48-73.3-48s-61 19.7-73.3 48L32 384c-17.7 0-32 14.3-32 32zm128 0a32 32 0 1 1 64 0 32 32 0 1 1 -64 0zM320 256a32 32 0 1 1 64 0 32 32 0 1 1 -64 0zm32-80c-32.8 0-61 19.7-73.3 48L32 224c-17.7 0-32 14.3-32 32s14.3 32 32 32l246.7 0c12.3 28.3 40.5 48 73.3 48s61-19.7 73.3-48l54.7 0c17.7 0 32-14.3 32-32s-14.3-32-32-32l-54.7 0c-12.3-28.3-40.5-48-73.3-48zM192 128a32 32 0 1 1 0-64 32 32 0 1 1 0 64zm73.3-64C253 35.7 224.8 16 192 16s-61 19.7-73.3 48L32 64C14.3 64 0 78.3 0 96s14.3 32 32 32l86.7 0c12.3 28.3 40.5 48 73.3 48s61-19.7 73.3-48L480 128c17.7 0 32-14.3 32-32s-14.3-32-32-32L265.3 64z"

	//kFontAwesomeWaveSquare viewBox="0 0 640 512" source: https://fontawesome.com/icons/wave-square?f=classic&s=solid
	kFontAwesomeWaveSquare = "M128 64c0-17.7 14.3-32 32-32l160 0c17.7 0 32 14.3 32 32l0 352 96 0 0-160c0-17.7 14.3-32 32-32l128 0c17.7 0 32 14.3 32 32s-14.3 32-32 32l-96 0 0 160c0 17.7-14.3 32-32 32l-160 0c-17.7 0-32-14.3-32-32l0-352-96 0 0 160c0 17.7-14.3 32-32 32L32 288c-17.7 0-32-14.3-32-32s14.3-32 32-32l96 0 0-160z"

	//kFontAwesomeHandHoldingDollar viewBox="0 0 576 512" source: https://fontawesome.com/icons/hand-holding-dollar?f=classic&s=solid
	kFontAwesomeHandHoldingDollar = "M312 24l0 10.5c6.4 1.2 12.6 2.7 18.2 4.2c12.8 3.4 20.4 16.6 17 29.4s-16.6 20.4-29.4 17c-10.9-2.9-21.1-4.9-30.2-5c-7.3-.1-14.7 1.7-19.4 4.4c-2.1 1.3-3.1 2.4-3.5 3c-.3 .5-.7 1.2-.7 2.8c0 .3 0 .5 0 .6c.2 .2 .9 1.2 3.3 2.6c5.8 3.5 14.4 6.2 27.4 10.1l.9 .3s0 0 0 0c11.1 3.3 25.9 7.8 37.9 15.3c13.7 8.6 26.1 22.9 26.4 44.9c.3 22.5-11.4 38.9-26.7 48.5c-6.7 4.1-13.9 7-21.3 8.8l0 10.6c0 13.3-10.7 24-24 24s-24-10.7-24-24l0-11.4c-9.5-2.3-18.2-5.3-25.6-7.8c-2.1-.7-4.1-1.4-6-2c-12.6-4.2-19.4-17.8-15.2-30.4s17.8-19.4 30.4-15.2c2.6 .9 5 1.7 7.3 2.5c13.6 4.6 23.4 7.9 33.9 8.3c8 .3 15.1-1.6 19.2-4.1c1.9-1.2 2.8-2.2 3.2-2.9c.4-.6 .9-1.8 .8-4.1l0-.2c0-1 0-2.1-4-4.6c-5.7-3.6-14.3-6.4-27.1-10.3l-1.9-.6c-10.8-3.2-25-7.5-36.4-14.4c-13.5-8.1-26.5-22-26.6-44.1c-.1-22.9 12.9-38.6 27.7-47.4c6.4-3.8 13.3-6.4 20.2-8.2L264 24c0-13.3 10.7-24 24-24s24 10.7 24 24zM568.2 336.3c13.1 17.8 9.3 42.8-8.5 55.9L433.1 485.5c-23.4 17.2-51.6 26.5-80.7 26.5L192 512 32 512c-17.7 0-32-14.3-32-32l0-64c0-17.7 14.3-32 32-32l36.8 0 44.9-36c22.7-18.2 50.9-28 80-28l78.3 0 16 0 64 0c17.7 0 32 14.3 32 32s-14.3 32-32 32l-64 0-16 0c-8.8 0-16 7.2-16 16s7.2 16 16 16l120.6 0 119.7-88.2c17.8-13.1 42.8-9.3 55.9 8.5zM193.6 384c0 0 0 0 0 0l-.9 0c.3 0 .6 0 .9 0z"

	//kFontAwesomeHandCommentsDollar viewBox="0 0 640 512" source: https://fontawesome.com/icons/comments-dollar?f=classic&s=solid
	kFontAwesomeHandCommentsDollar = "M416 176c0 97.2-93.1 176-208 176c-38.2 0-73.9-8.7-104.7-23.9c-7.5 4-16 7.9-25.2 11.4C59.8 346.4 37.8 352 16 352c-6.9 0-13.1-4.5-15.2-11.1s.2-13.8 5.8-17.9c0 0 0 0 0 0s0 0 0 0l.2-.2c.2-.2 .6-.4 1.1-.8c1-.8 2.5-2 4.3-3.7c3.6-3.3 8.5-8.1 13.3-14.3c5.5-7 10.7-15.4 14.2-24.7C14.7 250.3 0 214.6 0 176C0 78.8 93.1 0 208 0S416 78.8 416 176zM231.5 383C348.9 372.9 448 288.3 448 176c0-5.2-.2-10.4-.6-15.5C555.1 167.1 640 243.2 640 336c0 38.6-14.7 74.3-39.6 103.4c3.5 9.4 8.7 17.7 14.2 24.7c4.8 6.2 9.7 11 13.3 14.3c1.8 1.6 3.3 2.9 4.3 3.7c.5 .4 .9 .7 1.1 .8l.2 .2s0 0 0 0s0 0 0 0c5.6 4.1 7.9 11.3 5.8 17.9c-2.1 6.6-8.3 11.1-15.2 11.1c-21.8 0-43.8-5.6-62.1-12.5c-9.2-3.5-17.8-7.4-25.2-11.4C505.9 503.3 470.2 512 432 512c-95.6 0-176.2-54.6-200.5-129zM228 72c0-11-9-20-20-20s-20 9-20 20l0 14c-7.6 1.7-15.2 4.4-22.2 8.5c-13.9 8.3-25.9 22.8-25.8 43.9c.1 20.3 12 33.1 24.7 40.7c11 6.6 24.7 10.8 35.6 14l1.7 .5c12.6 3.8 21.8 6.8 28 10.7c5.1 3.2 5.8 5.4 5.9 8.2c.1 5-1.8 8-5.9 10.5c-5 3.1-12.9 5-21.4 4.7c-11.1-.4-21.5-3.9-35.1-8.5c-2.3-.8-4.7-1.6-7.2-2.4c-10.5-3.5-21.8 2.2-25.3 12.6s2.2 21.8 12.6 25.3c1.9 .6 4 1.3 6.1 2.1c0 0 0 0 0 0s0 0 0 0c8.3 2.9 17.9 6.2 28.2 8.4l0 14.6c0 11 9 20 20 20s20-9 20-20l0-13.8c8-1.7 16-4.5 23.2-9c14.3-8.9 25.1-24.1 24.8-45c-.3-20.3-11.7-33.4-24.6-41.6c-11.5-7.2-25.9-11.6-37.1-15l-.7-.2c-12.8-3.9-21.9-6.7-28.3-10.5c-5.2-3.1-5.3-4.9-5.3-6.7c0-3.7 1.4-6.5 6.2-9.3c5.4-3.2 13.6-5.1 21.5-5c9.6 .1 20.2 2.2 31.2 5.2c10.7 2.8 21.6-3.5 24.5-14.2s-3.5-21.6-14.2-24.5c-6.5-1.7-13.7-3.4-21.1-4.7L228 72z"

	//kFontAwesomeSackDollar viewBox="0 0 512 512" source: https://fontawesome.com/icons/sack-dollar?f=classic&s=solid
	kFontAwesomeSackDollar = "M320 96L192 96 144.6 24.9C137.5 14.2 145.1 0 157.9 0L354.1 0c12.8 0 20.4 14.2 13.3 24.9L320 96zM192 128l128 0c3.8 2.5 8.1 5.3 13 8.4C389.7 172.7 512 250.9 512 416c0 53-43 96-96 96L96 512c-53 0-96-43-96-96C0 250.9 122.3 172.7 179 136.4c0 0 0 0 0 0s0 0 0 0c4.8-3.1 9.2-5.9 13-8.4zm84 88c0-11-9-20-20-20s-20 9-20 20l0 14c-7.6 1.7-15.2 4.4-22.2 8.5c-13.9 8.3-25.9 22.8-25.8 43.9c.1 20.3 12 33.1 24.7 40.7c11 6.6 24.7 10.8 35.6 14l1.7 .5c12.6 3.8 21.8 6.8 28 10.7c5.1 3.2 5.8 5.4 5.9 8.2c.1 5-1.8 8-5.9 10.5c-5 3.1-12.9 5-21.4 4.7c-11.1-.4-21.5-3.9-35.1-8.5c-2.3-.8-4.7-1.6-7.2-2.4c-10.5-3.5-21.8 2.2-25.3 12.6s2.2 21.8 12.6 25.3c1.9 .6 4 1.3 6.1 2.1c0 0 0 0 0 0s0 0 0 0c8.3 2.9 17.9 6.2 28.2 8.4l0 14.6c0 11 9 20 20 20s20-9 20-20l0-13.8c8-1.7 16-4.5 23.2-9c14.3-8.9 25.1-24.1 24.8-45c-.3-20.3-11.7-33.4-24.6-41.6c-11.5-7.2-25.9-11.6-37.1-15c0 0 0 0 0 0l-.7-.2c-12.8-3.9-21.9-6.7-28.3-10.5c-5.2-3.1-5.3-4.9-5.3-6.7c0-3.7 1.4-6.5 6.2-9.3c5.4-3.2 13.6-5.1 21.5-5c9.6 .1 20.2 2.2 31.2 5.2c10.7 2.8 21.6-3.5 24.5-14.2s-3.5-21.6-14.2-24.5c-6.5-1.7-13.7-3.4-21.1-4.7l0-13.9z"

	//kFontAwesomeFloppyDisk viewBox="0 0 448 512" source: https://fontawesome.com/icons/floppy-disk?f=classic&s=solid FloppyDisk
	kFontAwesomeFloppyDisk = "M64 32C28.7 32 0 60.7 0 96L0 416c0 35.3 28.7 64 64 64l320 0c35.3 0 64-28.7 64-64l0-242.7c0-17-6.7-33.3-18.7-45.3L352 50.7C340 38.7 323.7 32 306.7 32L64 32zm0 96c0-17.7 14.3-32 32-32l192 0c17.7 0 32 14.3 32 32l0 64c0 17.7-14.3 32-32 32L96 224c-17.7 0-32-14.3-32-32l0-64zM224 288a64 64 0 1 1 0 128 64 64 0 1 1 0-128z"

	//kFontAwesomeDownload viewBox="0 0 512 512" source: https://fontawesome.com/icons/download?f=classic&s=solid
	kFontAwesomeDownload = "M288 32c0-17.7-14.3-32-32-32s-32 14.3-32 32l0 242.7-73.4-73.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l128 128c12.5 12.5 32.8 12.5 45.3 0l128-128c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L288 274.7 288 32zM64 352c-35.3 0-64 28.7-64 64l0 32c0 35.3 28.7 64 64 64l384 0c35.3 0 64-28.7 64-64l0-32c0-35.3-28.7-64-64-64l-101.5 0-45.3 45.3c-25 25-65.5 25-90.5 0L165.5 352 64 352zm368 56a24 24 0 1 1 0 48 24 24 0 1 1 0-48z"

	//kFontAwesomeShareNodes viewBox="0 0 448 512" source: https://fontawesome.com/icons/share-nodes?f=classic&s=solid
	kFontAwesomeShareNodes = "M352 224c53 0 96-43 96-96s-43-96-96-96s-96 43-96 96c0 4 .2 8 .7 11.9l-94.1 47C145.4 170.2 121.9 160 96 160c-53 0-96 43-96 96s43 96 96 96c25.9 0 49.4-10.2 66.6-26.9l94.1 47c-.5 3.9-.7 7.8-.7 11.9c0 53 43 96 96 96s96-43 96-96s-43-96-96-96c-25.9 0-49.4 10.2-66.6 26.9l-94.1-47c.5-3.9 .7-7.8 .7-11.9s-.2-8-.7-11.9l94.1-47C302.6 213.8 326.1 224 352 224z"

	//kFontAwesomeReTweet viewBox="0 0 576 512" source: https://fontawesome.com/icons/retweet?f=classic&s=solid
	kFontAwesomeReTweet = "M272 416c17.7 0 32-14.3 32-32s-14.3-32-32-32l-112 0c-17.7 0-32-14.3-32-32l0-128 32 0c12.9 0 24.6-7.8 29.6-19.8s2.2-25.7-6.9-34.9l-64-64c-12.5-12.5-32.8-12.5-45.3 0l-64 64c-9.2 9.2-11.9 22.9-6.9 34.9s16.6 19.8 29.6 19.8l32 0 0 128c0 53 43 96 96 96l112 0zM304 96c-17.7 0-32 14.3-32 32s14.3 32 32 32l112 0c17.7 0 32 14.3 32 32l0 128-32 0c-12.9 0-24.6 7.8-29.6 19.8s-2.2 25.7 6.9 34.9l64 64c12.5 12.5 32.8 12.5 45.3 0l64-64c9.2-9.2 11.9-22.9 6.9-34.9s-16.6-19.8-29.6-19.8l-32 0 0-128c0-53-43-96-96-96L304 96z"

	// viewBox="0 0 512 512" source: https://fontawesome.com/icons/server?f=classic&s=solid
	kFontAwesomeServer = "M64 32C28.7 32 0 60.7 0 96l0 64c0 35.3 28.7 64 64 64l384 0c35.3 0 64-28.7 64-64l0-64c0-35.3-28.7-64-64-64L64 32zm280 72a24 24 0 1 1 0 48 24 24 0 1 1 0-48zm48 24a24 24 0 1 1 48 0 24 24 0 1 1 -48 0zM64 288c-35.3 0-64 28.7-64 64l0 64c0 35.3 28.7 64 64 64l384 0c35.3 0 64-28.7 64-64l0-64c0-35.3-28.7-64-64-64L64 288zm280 72a24 24 0 1 1 0 48 24 24 0 1 1 0-48zm56 24a24 24 0 1 1 48 0 24 24 0 1 1 -48 0z"

	// viewBox="0 0 512 512" source: https://fontawesome.com/icons/upload?f=classic&s=solid
	kFontAwesomeUpload = "M288 109.3L288 352c0 17.7-14.3 32-32 32s-32-14.3-32-32l0-242.7-73.4 73.4c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3l128-128c12.5-12.5 32.8-12.5 45.3 0l128 128c12.5 12.5 12.5 32.8 0 45.3s-32.8 12.5-45.3 0L288 109.3zM64 352l128 0c0 35.3 28.7 64 64 64s64-28.7 64-64l128 0c35.3 0 64 28.7 64 64l0 32c0 35.3-28.7 64-64 64L64 512c-35.3 0-64-28.7-64-64l0-32c0-35.3 28.7-64 64-64zM432 456a24 24 0 1 0 0-48 24 24 0 1 0 0 48z"

	// viewBox="0 0 512 512" source: https://fontawesome.com/icons/rotate-left?f=classic&s=solid
	kFontAwesomeRotate = "M48.5 224L40 224c-13.3 0-24-10.7-24-24L16 72c0-9.7 5.8-18.5 14.8-22.2s19.3-1.7 26.2 5.2L98.6 96.6c87.6-86.5 228.7-86.2 315.8 1c87.5 87.5 87.5 229.3 0 316.8s-229.3 87.5-316.8 0c-12.5-12.5-12.5-32.8 0-45.3s32.8-12.5 45.3 0c62.5 62.5 163.8 62.5 226.3 0s62.5-163.8 0-226.3c-62.2-62.2-162.7-62.5-225.3-1L185 183c6.9 6.9 8.9 17.2 5.2 26.2s-12.5 14.8-22.2 14.8L48.5 224z"

	// viewBox="0 0 448 512" source: https://fontawesome.com/icons/bars?f=classic&s=solid
	kFontAwesomeBars = "M0 96C0 78.3 14.3 64 32 64l384 0c17.7 0 32 14.3 32 32s-14.3 32-32 32L32 128C14.3 128 0 113.7 0 96zM0 256c0-17.7 14.3-32 32-32l384 0c17.7 0 32 14.3 32 32s-14.3 32-32 32L32 288c-17.7 0-32-14.3-32-32zM448 416c0 17.7-14.3 32-32 32L32 448c-17.7 0-32-14.3-32-32s14.3-32 32-32l384 0c17.7 0 32 14.3 32 32z"

	// viewBox="0 0 512 512" source: https://fontawesome.com/icons/bug?f=classic&s=solid
	kFontAwesomeBug = "M256 0c53 0 96 43 96 96l0 3.6c0 15.7-12.7 28.4-28.4 28.4l-135.1 0c-15.7 0-28.4-12.7-28.4-28.4l0-3.6c0-53 43-96 96-96zM41.4 105.4c12.5-12.5 32.8-12.5 45.3 0l64 64c.7 .7 1.3 1.4 1.9 2.1c14.2-7.3 30.4-11.4 47.5-11.4l112 0c17.1 0 33.2 4.1 47.5 11.4c.6-.7 1.2-1.4 1.9-2.1l64-64c12.5-12.5 32.8-12.5 45.3 0s12.5 32.8 0 45.3l-64 64c-.7 .7-1.4 1.3-2.1 1.9c6.2 12 10.1 25.3 11.1 39.5l64.3 0c17.7 0 32 14.3 32 32s-14.3 32-32 32l-64 0c0 24.6-5.5 47.8-15.4 68.6c2.2 1.3 4.2 2.9 6 4.8l64 64c12.5 12.5 12.5 32.8 0 45.3s-32.8 12.5-45.3 0l-63.1-63.1c-24.5 21.8-55.8 36.2-90.3 39.6L272 240c0-8.8-7.2-16-16-16s-16 7.2-16 16l0 239.2c-34.5-3.4-65.8-17.8-90.3-39.6L86.6 502.6c-12.5 12.5-32.8 12.5-45.3 0s-12.5-32.8 0-45.3l64-64c1.9-1.9 3.9-3.4 6-4.8C101.5 367.8 96 344.6 96 320l-64 0c-17.7 0-32-14.3-32-32s14.3-32 32-32l64.3 0c1.1-14.1 5-27.5 11.1-39.5c-.7-.6-1.4-1.2-2.1-1.9l-64-64c-12.5-12.5-12.5-32.8 0-45.3z"

	// viewBox="0 0 640 640" source: https://fontawesome.com/icons/file-import?f=classic&s=solid
	kFontAwesomeFileImport = "M192 64C156.7 64 128 92.7 128 128L128 368L310.1 368L279.1 337C269.7 327.6 269.7 312.4 279.1 303.1C288.5 293.8 303.7 293.7 313 303.1L385 375.1C394.4 384.5 394.4 399.7 385 409L313 481C303.6 490.4 288.4 490.4 279.1 481C269.8 471.6 269.7 456.4 279.1 447.1L310.1 416.1L128 416.1L128 512.1C128 547.4 156.7 576.1 192 576.1L448 576.1C483.3 576.1 512 547.4 512 512.1L512 234.6C512 217.6 505.3 201.3 493.3 189.3L386.7 82.7C374.7 70.7 358.5 64 341.5 64L192 64zM453.5 240L360 240C346.7 240 336 229.3 336 216L336 122.5L453.5 240z"

	// https://fontawesome.com/icons/user-check?f=classic&s=solid
	// https://fontawesome.com/icons/user-minus?f=classic&s=solid
	// https://fontawesome.com/icons/user-plus?f=classic&s=solid
	// https://fontawesome.com/icons/user-slash?f=classic&s=solid
	// https://fontawesome.com/icons/user-doctor?f=classic&s=solid
	// https://fontawesome.com/icons/satellite-dish?f=classic&s=solid
	// https://fontawesome.com/icons/satellite?f=classic&s=solid
	// https://fontawesome.com/icons/shoe-prints?f=classic&s=solid

	//
	//
)

var Menu *makeIcon

func init() {
	Menu = new(makeIcon)
}

type makeIcon struct{}

type CategoryIcon struct {
	CategoryName string
	Menu         []MenuIcon
}
type MenuIcon struct {
	Col  int
	Row  int
	Name string
}

// -------------------------------------------------------------------------------------------------------
const doubleClickDelay = 500 * time.Millisecond // Tempo máximo entre dois cliques para duplo clique / Max time between two clicks for double click

// RegisterCanvasClickHandlers configura os eventos necessários no canvas para detectar click e double-click.
// RegisterCanvasClickHandlers sets up the necessary events on the canvas to detect click and double-click.
func RegisterCanvasClickHandlers(
	canvas js.Value,
	onClick func(x, y int),
	onDoubleClick func(x, y int),
) {

	tagCanvas := factoryBrowser.NewTagCanvas(800, 600).
		Import("canvas").
		SetWidth(800).
		SetHeight(900).
		SetZIndex(100)

	var pointerDown bool
	var clickTimer *time.Timer

	// evento de “pressionar” (mouse ou touch) inicia a flag pointerDown
	// “press” event (mouse or touch) sets pointerDown flag
	downCb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//args[0].Call("preventDefault")

		var font html.Font
		font.Family = factoryFontFamily.NewArial()
		font.Size = 17
		tagCanvas.ClearRect(0, 0, 100, 30).Font(font).FillText("down", 25, 25, -1)

		pointerDown = true
		return nil
	})

	// evento de “soltar” (mouse ou touch) trata como click se pointerDown == true
	// “release” event (mouse or touch) treated as click if pointerDown == true
	upCb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		//args[0].Call("preventDefault")

		var font html.Font
		font.Family = factoryFontFamily.NewArial()
		font.Size = 17
		tagCanvas.ClearRect(0, 0, 100, 30).Font(font).FillText("up", 25, 25, -1)

		if !pointerDown {
			return nil
		}
		pointerDown = false

		//ev := args[0]
		// calcula coordenadas relativas ao canvas
		// calculate coordinates relative to the canvas
		//rect := canvas.Call("getBoundingClientRect")
		//x := int(ev.Get("clientX").Float() - rect.Get("left").Float())
		//y := int(ev.Get("clientY").Float() - rect.Get("top").Float())
		x := 0
		y := 0

		if clickTimer == nil {
			// primeiro clique: aguarda tempo para ver se vem um segundo
			// first click: wait to see if a second click comes
			clickTimer = time.AfterFunc(doubleClickDelay, func() {
				onClick(x, y)
				clickTimer = nil
			})
		} else {
			// segundo clique dentro do intervalo: é um double click
			// second click within interval: it’s a double click
			clickTimer.Stop()
			clickTimer = nil
			onDoubleClick(x, y)
		}
		return nil
	})

	// associa listeners
	// attach listeners
	canvas.Call("addEventListener", "mousedown", downCb, map[string]any{"passive": true})
	canvas.Call("addEventListener", "mouseup", upCb, map[string]any{"passive": true})
	canvas.Call("addEventListener", "touchstart", downCb, map[string]any{"passive": true})
	canvas.Call("addEventListener", "touchend", upCb, map[string]any{"passive": true})
}

func (e makeIcon) Process(mainSvg *html.TagSvg) {
	e.register()

	size := rulesDensity.Density(50 + 6)
	hexMenu := new(rulesStage.Hexagon)
	hexMenu.Init(0, 0, size)

	canvas := factoryBrowser.NewTagCanvas(1200, 600).Import("canvas")
	//menuSvg := factoryBrowser.NewTagSvg().Width(rulesDensity.Density(600).GetInt()).Height(rulesDensity.Density(800).GetInt())
	//mainSvg.Append(menuSvg)

	menuOrder := []CategoryIcon{
		{
			CategoryName: "MainMenu",
			Menu: []MenuIcon{
				{
					Col:  1,
					Row:  1,
					Name: "SysFileLoad",
				},
				{
					Col:  1,
					Row:  3,
					Name: "SysMenu",
				},
				{
					Col:  1,
					Row:  5,
					Name: "SysBug",
				},
				{
					Col:  1,
					Row:  7,
					Name: "SysTools",
				},
				{
					Col:  1,
					Row:  9,
					Name: "SysBlank",
				},
			},
		},
		{
			CategoryName: "Menu",
			Menu: []MenuIcon{
				{
					Col:  2,
					Row:  2,
					Name: "SysMath",
				},
				{
					Col:  2,
					Row:  4,
					Name: "SysLoop",
				},
				{
					Col:  2,
					Row:  6,
					Name: "SysDonate",
				},
				{
					Col:  2,
					Row:  8,
					Name: "SysSave",
				},
				{
					Col:  2,
					Row:  10,
					Name: "SysUpload",
				},
			},
		},
		{
			CategoryName: "Loop",
			Menu: []MenuIcon{
				{
					Col:  2,
					Row:  12,
					Name: "Loop",
				},
				{
					Col:  2,
					Row:  14,
					Name: "SysGoBack",
				},
			},
		},
		{
			CategoryName: "Math",
			Menu: []MenuIcon{
				{
					Col:  5,
					Row:  7,
					Name: "SysGoBack",
				},
				{
					Col:  4,
					Row:  6,
					Name: "Add",
				},
				{
					Col:  4,
					Row:  8,
					Name: "Sub",
				},
				{
					Col:  6,
					Row:  6,
					Name: "Mul",
				},
				{
					Col:  6,
					Row:  8,
					Name: "Div",
				},
			},
		},
	}

	//"MainMenu": {
	//	"SysFileLoad": {1, 1},
	//	"SysMenu":     {1, 3},
	//	"SysBug":      {1, 5},
	//	"SysTools":    {1, 7},
	//	"SysBlank":    {5, 7},
	//},
	//"Menu": {
	//	"SysMath":   {1, 9},
	//	"SysLoop":   {4, 8},
	//	"SysDonate": {2, 2},
	//	"SysSave":   {2, 4},
	//	"SysUpload": {2, 6},
	//},
	//"Loop": {
	//	"SysLoop":   {2, 8},
	//	"Loop":      {3, 1},
	//	"SysGoBack": {3, 3},
	//},
	//"Math": {
	//	"SysMath":   {3, 5},
	//	"Add":       {3, 7},
	//	"Sub":       {3, 9},
	//	"Mul":       {4, 6},
	//	"Div":       {4, 2},
	//	"SysGoBack": {4, 4},
	//},
	//"Math": {
	//"SysMath": {5, 5},
	//"Add": {4, 6},
	//"Sub": {4, 8},
	//"Mul": {6, 6},
	//"Div": {6, 8},
	//"SysGoBack": {5, 9},
	//"SysBlank":  {5, 7},
	//"SysMenu": {5, 7},
	//"SysGoBack": {5, 7},
	//},
	//}

	// registra callbacks de click e double-click
	// register click and double-click callbacks
	//RegisterCanvasClickHandlers(canvas,
	//	func(x, y int) {
	//
	//		tagCanvas := factoryBrowser.NewTagCanvas(800, 600).
	//			Import("canvas").
	//			DisableSelection().
	//			SetWidth(800).
	//			SetHeight(600).
	//			SetZIndex(-100)
	//
	//		var font html.Font
	//		font.Family = factoryFontFamily.NewArial()
	//		font.Size = 17
	//		tagCanvas.ClearRect(0, 0, 100, 30).Font(font).FillText("click", 25, 25, -1)
	//	},
	//	func(x, y int) {
	//
	//		tagCanvas := factoryBrowser.NewTagCanvas(800, 600).
	//			Import("canvas").
	//			DisableSelection().
	//			SetWidth(800).
	//			SetHeight(600).
	//			SetZIndex(-100)
	//
	//		var font html.Font
	//		font.Family = factoryFontFamily.NewArial()
	//		font.Size = 17
	//		tagCanvas.ClearRect(0, 0, 100, 30).Font(font).FillText("dbl click", 25, 25, -1)
	//	},
	//)

	var test func()
	var status = 1

	go func() {
		icons := manager.Manager.GetMapIcons()

		for {
			//counter := 0.0
			canvas.ClearRect(0, 0, 800, 900)

			for _, categoryList := range menuOrder {
				category := categoryList.CategoryName
				for _, iconData := range categoryList.Menu {
					name := iconData.Name
					systemIcon, found := icons[category][name]

					if !found {
						systemIcon, found = icons["Main"][name]
					}

					if name == "Loop--" {
						//systemIcon.SetStatus(5)
						if test == nil {
							test = func() {
								time.Sleep(time.Millisecond * 1000 * 10)
								systemIcon.SetOpening(status)
								status *= -1
								defer test()
							}
							go test()
						}
					}

					if !found {
						continue
					}

					hexMenu.SetRowCol(iconData.Col, iconData.Row)
					x, y := hexMenu.GetCenter()
					canvas.DrawImage(systemIcon.GetIcon(), x.GetInt()+systemIcon.GetX(), y.GetInt()+systemIcon.GetY(), systemIcon.GetWidth(), systemIcon.GetHeight())
				}
			}
			time.Sleep(time.Microsecond)
		}
	}()

}

func (e makeIcon) getIcon(data rulesIcon.Data) (png js.Value) {

	data = rulesIcon.DataVerifySystemIcon(data)

	// icon body
	svgIcon := factoryBrowser.NewTagSvg().
		X(rulesIcon.Width.GetInt() / 2).
		Y(rulesIcon.Height.GetInt() / 2).
		Width(rulesIcon.Width.GetInt()).
		Height(rulesIcon.Height.GetInt())

	// hexagon maker
	hexPath := utilsDraw.PolygonPath(
		6,
		rulesIcon.Width/2,
		rulesIcon.Width/2,
		rulesIcon.Width/2,
		0,
	)

	// svg hexagon
	hexDraw := factoryBrowser.NewTagSvgPath().
		StrokeWidth(rulesIcon.BorderWidth.GetInt()).
		Stroke(data.ColorBorder).
		Fill(data.ColorBackground).
		D(hexPath)

	// icon svg font awesome, body
	icon := factoryBrowser.NewTagSvg().
		ViewBox(data.IconViewBox).
		X(data.IconX.GetInt()).
		Y(data.IconY.GetInt()).
		Width(data.IconWidth.GetInt()).
		Height(data.IconHeight.GetInt())

	// icon svg font awesome, path svg
	iconPath := factoryBrowser.NewTagSvgPath().
		Fill(data.ColorIcon).
		D(data.Path)
	icon.Append(iconPath)

	// calc width label
	widthLabel, _ := utilsText.GetTextSize(
		data.Label,
		rulesIcon.FontFamily,
		rulesIcon.FontWeight,
		rulesIcon.FontStyle,
		data.LabelFontSize.GetInt(),
	)

	// label, svg text
	label := factoryBrowser.NewTagSvgText().
		FontFamily(rulesIcon.FontFamily).
		FontWeight(rulesIcon.FontWeight).
		FontStyle(rulesIcon.FontStyle).
		FontSize(data.LabelFontSize.GetInt()).
		Text(data.Label).
		Fill(data.ColorLabel).
		X((rulesIcon.Width / 2).GetInt() - widthLabel/2).
		Y(data.LabelY.GetInt())
	svgIcon.Append(hexDraw, icon, label)

	w := rulesIcon.Width * rulesIcon.SizeRatio
	h := rulesIcon.Height * rulesIcon.SizeRatio
	return svgIcon.ToCanvas(
		html.CanvasData{
			Width:  w.GetInt(),
			Height: h.GetInt(),
		},
	)
}

func (e makeIcon) register() {
	manager.Manager.Register(e.getBlank())
	manager.Manager.Register(e.getBug())
	manager.Manager.Register(e.getMath())
	manager.Manager.Register(e.getLoop())
	manager.Manager.Register(e.getTools())
	manager.Manager.Register(e.getConfig())
	manager.Manager.Register(e.getGraph())
	manager.Manager.Register(e.getMenu())
	manager.Manager.Register(e.getDonate())
	manager.Manager.Register(e.getSave())
	manager.Manager.Register(e.getShare())
	manager.Manager.Register(e.getRetweet())
	manager.Manager.Register(e.getServer())
	manager.Manager.Register(e.getUpload())
	manager.Manager.Register(e.getGoBack())
	manager.Manager.Register(e.getFileImport())
}

func (e makeIcon) getBlank() (register *manager.RegisterIcon) {
	name := "SysBlank"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:          int(manager.KPipeLineNormal),
			Name:            name,
			Category:        category,
			ColorBackground: color.RGBA{R: 0xf8, G: 0xf8, B: 0xef, A: 0x8a},
			ColorBorder:     color.RGBA{R: 0x5F, G: 0x5F, B: 0x5F, A: 0x8a},
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:          int(manager.KPipeLineDisabled),
			Name:            name,
			Category:        category,
			ColorBackground: color.RGBA{R: 0xf8, G: 0xf8, B: 0xef, A: 0x8a},
			ColorBorder:     color.RGBA{R: 0x5F, G: 0x5F, B: 0x5F, A: 0x8a},
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:          int(manager.KPipeLineSelected),
			Name:            name,
			Category:        category,
			ColorBackground: color.RGBA{R: 0xf8, G: 0xf8, B: 0xef, A: 0x8a},
			ColorBorder:     color.RGBA{R: 0x5F, G: 0x5F, B: 0x5F, A: 0x8a},
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:          int(manager.KPipeLineAttention1),
			Name:            name,
			Category:        category,
			ColorBackground: color.RGBA{R: 0xf8, G: 0xf8, B: 0xef, A: 0x8a},
			ColorBorder:     color.RGBA{R: 0x5F, G: 0x5F, B: 0x5F, A: 0x8a},
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:          int(manager.KPipeLineAttention2),
			Name:            name,
			Category:        category,
			ColorBackground: color.RGBA{R: 0xf8, G: 0xf8, B: 0xef, A: 0x8a},
			//ColorBorder:     color.RGBA{R: 0x5F, G: 0x5F, B: 0x5F, A: 0x8a},
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getBug() (register *manager.RegisterIcon) {
	//var count int = 0
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID: "IconSystemBug",
			},
			//TemplateData: map[string]interface{}{
			//	"Name":  "Nick",
			//	"Count": count,
			//},
			//PluralCount: count,
		},
	)
	if err != nil {
		translated = "Bug"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysBug"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Label:    translated,
			Path:     kFontAwesomeBug,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Label:    translated,
			Path:     kFontAwesomeBug,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Label:    translated,
			Path:     kFontAwesomeBug,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Label:    translated,
			Path:     kFontAwesomeBug,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Label:    translated,
			Path:     kFontAwesomeBug,
			Name:     name,
			Category: category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getMath() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemMath",
				Other: "Math",
			},
		},
	)
	if err != nil {
		translated = "Math"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysMath"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineNormal),
			IconViewBox: []int{0, 0, 576, 512},
			Label:       translated,
			Path:        kFontAwesomeSquareRootVariable,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineDisabled),
			IconViewBox: []int{0, 0, 576, 512},
			Label:       translated,
			Path:        kFontAwesomeSquareRootVariable,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineSelected),
			IconViewBox: []int{0, 0, 576, 512},
			Label:       translated,
			Path:        kFontAwesomeSquareRootVariable,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineAttention1),
			IconViewBox: []int{0, 0, 576, 512},
			Label:       translated,
			Path:        kFontAwesomeSquareRootVariable,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineAttention2),
			IconViewBox: []int{0, 0, 576, 512},
			Label:       translated,
			Path:        kFontAwesomeSquareRootVariable,
			Name:        name,
			Category:    category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getLoop() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemLoop",
				Other: "Math",
			},
		},
	)
	if err != nil {
		translated = "Loop"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysLoop"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Label:    translated,
			Path:     kFontAwesomeRepeat,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Label:    translated,
			Path:     kFontAwesomeRepeat,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Label:    translated,
			Path:     kFontAwesomeRepeat,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Label:    translated,
			Path:     kFontAwesomeRepeat,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Label:    translated,
			Path:     kFontAwesomeRepeat,
			Name:     name,
			Category: category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getTools() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemTools",
				Other: "Tools",
			},
		},
	)
	if err != nil {
		translated = "Tools"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysTools"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Label:    translated,
			Path:     kFontAwesomeScrewDriverWrench,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Label:    translated,
			Path:     kFontAwesomeScrewDriverWrench,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Label:    translated,
			Path:     kFontAwesomeScrewDriverWrench,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Label:    translated,
			Path:     kFontAwesomeScrewDriverWrench,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Label:    translated,
			Path:     kFontAwesomeScrewDriverWrench,
			Name:     name,
			Category: category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getConfig() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemConfig",
				Other: "Config",
			},
		},
	)
	if err != nil {
		translated = "Config"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysConfig"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Label:    translated,
			Path:     kFontAwesomeSliders,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Label:    translated,
			Path:     kFontAwesomeSliders,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Label:    translated,
			Path:     kFontAwesomeSliders,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Label:    translated,
			Path:     kFontAwesomeSliders,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Label:    translated,
			Path:     kFontAwesomeSliders,
			Name:     name,
			Category: category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getGraph() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemGraph",
				Other: "Graph",
			},
		},
	)
	if err != nil {
		translated = "Graph"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysGraph"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineNormal),
			IconViewBox: []int{0, 0, 640, 512},
			Label:       translated,
			Path:        kFontAwesomeWaveSquare,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineDisabled),
			IconViewBox: []int{0, 0, 640, 512},
			Label:       translated,
			Path:        kFontAwesomeWaveSquare,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineSelected),
			IconViewBox: []int{0, 0, 640, 512},
			Label:       translated,
			Path:        kFontAwesomeWaveSquare,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineAttention1),
			IconViewBox: []int{0, 0, 640, 512},
			Label:       translated,
			Path:        kFontAwesomeWaveSquare,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineAttention2),
			IconViewBox: []int{0, 0, 640, 512},
			Label:       translated,
			Path:        kFontAwesomeWaveSquare,
			Name:        name,
			Category:    category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getMenu() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemMenu",
				Other: "Menu",
			},
		},
	)
	if err != nil {
		translated = "Menu"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysMenu"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineNormal),
			IconViewBox: []int{0, 0, 448, 512},
			Label:       translated,
			Path:        kFontAwesomeBars,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineDisabled),
			IconViewBox: []int{0, 0, 448, 512},
			Label:       translated,
			Path:        kFontAwesomeBars,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineSelected),
			IconViewBox: []int{0, 0, 448, 512},
			Label:       translated,
			Path:        kFontAwesomeBars,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineAttention1),
			IconViewBox: []int{0, 0, 448, 512},
			Label:       translated,
			Path:        kFontAwesomeBars,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineAttention2),
			IconViewBox: []int{0, 0, 448, 512},
			Label:       translated,
			Path:        kFontAwesomeBars,
			Name:        name,
			Category:    category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getDonate() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemDonate",
				Other: "Donate",
			},
		},
	)
	if err != nil {
		translated = "Donate"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysDonate"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Label:    translated,
			Path:     kFontAwesomeSackDollar,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Label:    translated,
			Path:     kFontAwesomeSackDollar,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Label:    translated,
			Path:     kFontAwesomeSackDollar,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Label:    translated,
			Path:     kFontAwesomeSackDollar,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Label:    translated,
			Path:     kFontAwesomeSackDollar,
			Name:     name,
			Category: category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getSave() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemSave",
				Other: "Save",
			},
		},
	)
	if err != nil {
		translated = "Save"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysSave"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Label:    translated,
			Path:     kFontAwesomeDownload,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Label:    translated,
			Path:     kFontAwesomeDownload,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Label:    translated,
			Path:     kFontAwesomeDownload,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Label:    translated,
			Path:     kFontAwesomeDownload,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Label:    translated,
			Path:     kFontAwesomeDownload,
			Name:     name,
			Category: category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getShare() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemShare",
				Other: "Share",
			},
		},
	)
	if err != nil {
		translated = "Share"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysShare"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Label:    translated,
			Path:     kFontAwesomeShareNodes,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Label:    translated,
			Path:     kFontAwesomeShareNodes,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Label:    translated,
			Path:     kFontAwesomeShareNodes,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Label:    translated,
			Path:     kFontAwesomeShareNodes,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Label:    translated,
			Path:     kFontAwesomeShareNodes,
			Name:     name,
			Category: category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getRetweet() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemRetweet",
				Other: "Retweet",
			},
		},
	)
	if err != nil {
		translated = "Retweet"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysRetweet"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineNormal),
			IconViewBox: []int{0, 0, 576, 512},
			Label:       translated,
			Path:        kFontAwesomeReTweet,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineDisabled),
			IconViewBox: []int{0, 0, 576, 512},
			Label:       translated,
			Path:        kFontAwesomeReTweet,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineSelected),
			IconViewBox: []int{0, 0, 576, 512},
			Label:       translated,
			Path:        kFontAwesomeReTweet,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineAttention1),
			IconViewBox: []int{0, 0, 576, 512},
			Label:       translated,
			Path:        kFontAwesomeReTweet,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineAttention2),
			IconViewBox: []int{0, 0, 576, 512},
			Label:       translated,
			Path:        kFontAwesomeReTweet,
			Name:        name,
			Category:    category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getServer() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemServer",
				Other: "Server",
			},
		},
	)
	if err != nil {
		translated = "Server"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysServer"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Label:    translated,
			Path:     kFontAwesomeServer,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Label:    translated,
			Path:     kFontAwesomeServer,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Label:    translated,
			Path:     kFontAwesomeServer,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Label:    translated,
			Path:     kFontAwesomeServer,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Label:    translated,
			Path:     kFontAwesomeServer,
			Name:     name,
			Category: category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getUpload() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemUpload",
				Other: "Upload",
			},
		},
	)
	if err != nil {
		translated = "Upload"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysUpload"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Label:    translated,
			Path:     kFontAwesomeUpload,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Label:    translated,
			Path:     kFontAwesomeUpload,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Label:    translated,
			Path:     kFontAwesomeUpload,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Label:    translated,
			Path:     kFontAwesomeUpload,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Label:    translated,
			Path:     kFontAwesomeUpload,
			Name:     name,
			Category: category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getGoBack() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemGoBack",
				Other: "Go back",
			},
		},
	)
	if err != nil {
		translated = "Back"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysGoBack"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineNormal),
			Label:    translated,
			Path:     kFontAwesomeRotate,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineDisabled),
			Label:    translated,
			Path:     kFontAwesomeRotate,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineSelected),
			Label:    translated,
			Path:     kFontAwesomeRotate,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention1),
			Label:    translated,
			Path:     kFontAwesomeRotate,
			Name:     name,
			Category: category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:   int(manager.KPipeLineAttention2),
			Label:    translated,
			Path:     kFontAwesomeRotate,
			Name:     name,
			Category: category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}

func (e makeIcon) getFileImport() (register *manager.RegisterIcon) {
	translated, err := translate.Localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "IconSystemLoad",
				Other: "Load",
			},
		},
	)
	if err != nil {
		translated = "Load"
		log.Printf("icon translation error: %v", err)
	}

	name := "SysFileLoad"
	category := "Main"
	iconPipeLine := make([]js.Value, 5)
	iconPipeLine[manager.KPipeLineNormal] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineNormal),
			IconViewBox: []int{0, 0, 640, 640},
			Label:       translated,
			Path:        kFontAwesomeFileImport,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineDisabled] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineDisabled),
			IconViewBox: []int{0, 0, 640, 640},
			Label:       translated,
			Path:        kFontAwesomeFileImport,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineSelected] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineSelected),
			IconViewBox: []int{0, 0, 640, 640},
			Label:       translated,
			Path:        kFontAwesomeFileImport,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention1] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineAttention1),
			IconViewBox: []int{0, 0, 640, 640},
			Label:       translated,
			Path:        kFontAwesomeFileImport,
			Name:        name,
			Category:    category,
		},
	)
	iconPipeLine[manager.KPipeLineAttention2] = e.getIcon(
		rulesIcon.Data{
			Status:      int(manager.KPipeLineAttention2),
			IconViewBox: []int{0, 0, 640, 640},
			Label:       translated,
			Path:        kFontAwesomeFileImport,
			Name:        name,
			Category:    category,
		},
	)

	register = new(manager.RegisterIcon)
	register.SetName(name)
	register.SetCategory(category)
	register.SetIcon(iconPipeLine)
	return register
}
