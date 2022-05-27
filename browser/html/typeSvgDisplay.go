package html

type SvgDisplay string

func (e SvgDisplay) String() string {
	return string(e)
}

const (
	KSvgDisplayBlock             SvgDisplay = "block"
	KSvgDisplayInline            SvgDisplay = "inline"
	KSvgDisplayRunIn             SvgDisplay = "run-in"
	KSvgDisplayFlow              SvgDisplay = "flow"
	KSvgDisplayFlowRoot          SvgDisplay = "flow-root"
	KSvgDisplayTable             SvgDisplay = "table"
	KSvgDisplayFlex              SvgDisplay = "flex"
	KSvgDisplayGrid              SvgDisplay = "grid"
	KSvgDisplayRuby              SvgDisplay = "ruby"
	KSvgDisplayListItem          SvgDisplay = "list-item"
	KSvgDisplayTableRowGroup     SvgDisplay = "table-row-group"
	KSvgDisplayTableHeaderGroup  SvgDisplay = "table-header-group"
	KSvgDisplayTableFooterGroup  SvgDisplay = "table-footer-group"
	KSvgDisplayTableRow          SvgDisplay = "table-row"
	KSvgDisplayTableCell         SvgDisplay = "table-cell"
	KSvgDisplayTableColumnGroup  SvgDisplay = "table-column-group"
	KSvgDisplayTableColumn       SvgDisplay = "table-column"
	KSvgDisplayTableCaption      SvgDisplay = "table-caption"
	KSvgDisplayRubyBase          SvgDisplay = "ruby-base"
	KSvgDisplayRubyText          SvgDisplay = "ruby-text"
	KSvgDisplayRubyBaseContainer SvgDisplay = "ruby-base-container"
	KSvgDisplayRubyTextContainer SvgDisplay = "ruby-text-container"
	KSvgDisplayContents          SvgDisplay = "contents"
	KSvgDisplayNone              SvgDisplay = "none"
	KSvgDisplayInlineBlock       SvgDisplay = "inline-block"
	KSvgDisplayInlineListItem    SvgDisplay = "inline-list-item"
	KSvgDisplayInlineTable       SvgDisplay = "inline-table"
	KSvgDisplayInlineFlex        SvgDisplay = "inline-flex"
	KSvgDisplayInlineGrid        SvgDisplay = "inline-grid"
)
