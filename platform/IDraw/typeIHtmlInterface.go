package iotmaker_platform_IDraw

import "html"

type IHtml interface {
	NewImage(parent interface{}, propertiesList map[string]interface{}, waitLoad bool) html.Image
	Append(document, element interface{})
	Remove(document, element interface{})
	GetDocumentWidth(document interface{}) int
	GetDocumentHeight(document interface{}) int
}
