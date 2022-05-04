package iotmaker_platform_IDraw

import "github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"

type Primitive interface {
  GetPlatform() IDraw
  GetScratchPad() IDraw
  GetId() string
  GetDimensions() genericTypes.Dimensions
  GetInk() genericTypes.Ink
}

