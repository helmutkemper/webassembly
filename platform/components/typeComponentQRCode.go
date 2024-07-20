package components

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/mathUtil"
	"github.com/helmutkemper/webassembly/qrcode"
	"image/color"
)

// __qrCodeOnInputEvent faz a captura de dados do event input
type __qrCodeOnInputEvent struct {
	Value string `wasmGet:"value"`
}

type QRCode struct {
	__value         string
	__change        *__qrCodeOnInputEvent
	__size          int
	__recoveryLevel qrcode.RecoveryLevel
	__color         color.Color
	__background    color.Color

	__canvasTag *html.TagCanvas `wasmPanel:"type:TagCanvas"`
}

func (e *QRCode) init() {
	if e.__value != "" {
		e.setValue(e.__value)
		e.__value = ""
	}
}

func (e *QRCode) setValue(value string) {
	e.__canvasTag.DrawQRCodeColor(e.__size, value, e.__recoveryLevel, e.__color, e.__background)
}

// SetValue Defines the content of the QR Code
//
//	Example formats:
//
//	  URL:
//	  Prefix: http:// or https://
//	  Example: https://www.example.com
//
//	  Phone number:
//	  Prefix: tel:
//	  Example: tel:+1234567890
//
//	  SMS:
//	  Prefix: sms:
//	  Example: sms:+1234567890?body=Hello
//
//	  Email:
//	  Prefix: mailto:
//	  Example: mailto:example@example.com
//
//	  Contact (vCard):
//	  Prefix: BEGIN:VCARD
//	  Example:
//	  `BEGIN:VCARD
//	  VERSION:3.0
//	  FN:John Doe
//	  TEL:+1234567890
//	  EMAIL:example@example.com
//	  END:VCARD`
//
//	  Event (iCalendar):
//	  Prefix: BEGIN:VEVENT
//	  Example:
//	  `BEGIN:VEVENT
//	  SUMMARY:Meeting
//	  DTSTART:20230701T120000Z
//	  DTEND:20230701T130000Z
//	  LOCATION:Conference Room
//	  END:VEVENT`
//
//	  Location (Geo URI):
//	  Prefix: geo:
//	  Example: geo:37.7749,-122.4194
//
//	  WiFi:
//	  Prefix: WIFI:
//	  Example: WIFI:T:WPA;S:NetworkName;P:Password;;
//
//	  Simple text:
//	  No prefix needed
//	  Example: Hello, world!
//
//	  Bitcoin:
//	  Prefix: bitcoin:
//	  Example: bitcoin:1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa
//
//	  PayPal:
//	  Prefix: paypal:
//	  Example: paypal:someone@example.com
//
//	  WhatsApp:
//	  Prefix: https://wa.me/
//	  Example: https://wa.me/1234567890
//
//	  MeCard (Contato):
//	  Prefix: MECARD:
//	  Example:
//	  `MECARD:N:John Doe;TEL:+1234567890;EMAIL:example@example.com;;`
func (e *QRCode) SetValue(value string) {
	if e.__canvasTag == nil {
		e.__value = value
		return
	}

	e.setValue(value)
}

func (e *QRCode) SetBackground(value any) (err error) {
	switch converted := value.(type) {
	case color.Color:
		e.__background = converted
	case string:
		e.__background, err = mathUtil.HexToColor(converted)
		if err != nil {
			e.__background = color.White
		}
	}
	return
}

func (e *QRCode) SetColor(value any) (err error) {
	switch converted := value.(type) {
	case color.Color:
		e.__color = converted
	case string:
		e.__color, err = mathUtil.HexToColor(converted)
		if err != nil {
			e.__color = color.Black
		}
	}
	return
}

func (e *QRCode) SetSize(value int) (err error) {
	e.__size = value
	return
}
