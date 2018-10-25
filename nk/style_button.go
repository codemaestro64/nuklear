package nk

import "C"

type ButtonStyle struct {
	ctx   *Context
	style C.struct_nk_style_button
}

var (
	btnBackgroundNormal = C.struct_nk_style_item(NkStyleItemColor(NkRgba(0, 0, 0, 0)))
	btnBackgroundHover  = C.struct_nk_style_item(NkStyleItemColor(NkRgb(255, 165, 0)))
	btnBackgroundActive = C.struct_nk_style_item(NkStyleItemColor(NkRgb(220, 10, 0)))
	btnBorderColor      = C.struct_nk_color(NkRgb(255, 165, 0))
	btnTextBackground   = C.struct_nk_color(NkRgb(0, 0, 0))
	btnTextNormal       = C.struct_nk_color(NkRgb(255, 165, 0))
	btnTextHover        = C.struct_nk_color(NkRgb(28, 48, 62))
	btnTextActive       = C.struct_nk_color(NkRgb(28, 48, 62))
)

func NewButtonStyle(ctx *Context) *ButtonStyle {
	style := ctx.Style().button
	style.normal = btnBackgroundNormal
	style.hover = C.struct_nk_style_item(NkStyleItemColor(NkRgb(255, 255, 255)))
	style.active = C.struct_nk_style_item(NkStyleItemColor(NkRgb(255, 255, 255)))
	style.border_color = btnBorderColor
	style.text_background = btnTextBackground
	style.text_normal = btnTextNormal
	style.text_hover = btnTextHover
	style.text_active = btnTextActive

	return &ButtonStyle{
		style: style,
	}
}

func (b *ButtonStyle) SetBackgroundColor(color Color) {
	b.style.normal = C.struct_nk_style_item(NkStyleItemColor(color))
}

func (b *ButtonStyle) SetHoverBackgroundColor(color Color) {
	b.style.hover = C.struct_nk_style_item(NkStyleItemColor(color))
}

func (b *ButtonStyle) SetActiveBackgroundColor(color Color) {
	b.style.active = C.struct_nk_style_item(NkStyleItemColor(color))
}

func (b *ButtonStyle) SetBorderColor(color Color) {
	b.style.border_color = C.struct_nk_color(color)
}

func (b *ButtonStyle) SetTextBackground(color Color) {
	b.style.text_background = C.struct_nk_color(color)
}

func (b *ButtonStyle) SetTextNormal(color Color) {
	b.style.text_normal = C.struct_nk_color(color)
}

func (b *ButtonStyle) SetTextHover(color Color) {
	b.style.text_hover = C.struct_nk_color(color)
}

func (b *ButtonStyle) SetTextActive(color Color) {
	b.style.text_active = C.struct_nk_color(color)
}

func (b *ButtonStyle) GetStyle() *StyleButton {
	d := StyleButton(b.style)
	return &d
}
