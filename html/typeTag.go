package html

type Tag string

func (e Tag) String() string {
	return string(e)
}

const (
	// KTagA
	//
	// English:
	//
	//  The Anchor element.
	//
	// The <a> HTML element (or anchor element), with its href attribute, creates a hyperlink to web
	// pages, files, email addresses, locations in the same page, or anything else a URL can address.
	//
	// Content within each <a> should indicate the link's destination. If the href attribute is present,
	// pressing the enter key while focused on the <a> element will activate it.
	//
	// Português
	//
	//  O elemento Âncora.
	//
	// O elemento HTML <a> (ou elemento âncora), com seu atributo href, cria um hiperlink para páginas
	// da web, arquivos, endereços de e-mail, locais na mesma página ou qualquer outra coisa que um URL
	// possa endereçar.
	//
	// O conteúdo de cada <a> deve indicar o destino do link. Se o atributo href estiver presente,
	// pressionar a tecla enter enquanto estiver focado no elemento <a> irá ativá-lo.
	KTagA Tag = "a"

	// KTagAbbr
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagAbbr Tag = "abbr"

	// KTagAddress
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagAddress Tag = "address"

	// KTagArea
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagArea Tag = "area"

	// KTagArticle
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagArticle Tag = "article"

	// KTagAside
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagAside Tag = "aside"

	// KTagAudio
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagAudio Tag = "audio"

	// KTagB
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagB Tag = "b"

	// KTagBase
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagBase Tag = "base"

	// KTagBdi
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagBdi Tag = "bdi"

	// KTagBdo
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagBdo Tag = "bdo"

	// KTagBlockquote
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagBlockquote Tag = "blockquote"

	// KTagBody
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagBody Tag = "body"

	// KTagBr
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagBr Tag = "br"

	// KTagButton
	//
	// English:
	//
	//  The <button> HTML element is an interactive element activated by a user with a mouse, keyboard, finger, voice command, or other assistive technology. Once activated, it then performs a programmable action, such as submitting a form or opening a dialog.
	//
	// By default, HTML buttons are presented in a style resembling the platform the user agent runs on, but you can change buttons' appearance with CSS.
	//
	// Português
	//
	//  O elemento HTML <button> é um elemento interativo ativado por um usuário com mouse, teclado,
	//  dedo, comando de voz ou outra tecnologia assistiva. Uma vez ativado, ele executa uma ação
	//  programável, como enviar um formulário ou abrir uma caixa de diálogo.
	//
	// Por padrão, os botões HTML são apresentados em um estilo semelhante à plataforma na qual o agente
	// do usuário é executado, mas você pode alterar a aparência dos botões com CSS.
	KTagButton Tag = "button"

	// KTagCanvas
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagCanvas Tag = "canvas"

	// KTagCaption
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagCaption Tag = "caption"

	// KTagCite
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagCite Tag = "cite"

	// KTagCode
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagCode Tag = "code"

	// KTagCol
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagCol Tag = "col"

	// KTagColgroup
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagColgroup Tag = "colgroup"

	// KTagData
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagData Tag = "data"

	// KTagDatalist
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagDatalist Tag = "datalist"

	// KTagDd
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagDd Tag = "dd"

	// KTagDel
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagDel Tag = "del"

	// KTagDetails
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagDetails Tag = "details"

	// KTagDfn
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagDfn Tag = "dfn"

	// KTagDialog
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagDialog Tag = "dialog"

	// KTagDiv
	//
	//
	// English:
	//
	//  The <div> tag defines a division or a section in an HTML document.
	//
	//   Note:
	//     * By default, browsers always place a line break before and after the <div> element;
	//     * The <div> tag is used as a container for HTML elements - which is then styled with CSS or
	//       manipulated with JavaScript;
	//     * The <div> tag is easily styled by using the class or id attribute;
	//     * Any sort of content can be put inside the <div> tag.
	//
	// Português:
	//
	//  A tag <div> define uma divisão ou uma seção em um documento HTML.
	//
	//   Nota:
	//     * Por padrão, os navegadores sempre colocam uma quebra de linha antes e depois do elemento
	//       <div>;
	//     * A tag <div> é usada como um contêiner para elementos HTML - que são estilizados com CSS ou
	//       manipulados com JavaScript
	//     * A tag <div> é facilmente estilizada usando o atributo class ou id;
	//     * Qualquer tipo de conteúdo pode ser colocado dentro da tag <div>.
	KTagDiv Tag = "div"

	// KTagDl
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagDl Tag = "dl"

	// KTagDt
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagDt Tag = "dt"

	// KTagEm
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagEm Tag = "em"

	// KTagEmbed
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagEmbed Tag = "embed"

	// KTagFieldset
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagFieldset Tag = "fieldset"

	// KTagFigcaption
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagFigcaption Tag = "figcaption"

	// KTagFigure
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagFigure Tag = "figure"

	// KTagFooter
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagFooter Tag = "footer"

	// KTagForm
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagForm Tag = "form"

	// KTagHead
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagHead Tag = "head"

	// KTagHeader
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagHeader Tag = "header"

	// KTagH1
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagH1 Tag = "h1"

	// KTagHr
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagHr Tag = "hr"

	// KTagHtml
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagHtml Tag = "html"

	// KTagI
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagI Tag = "i"

	// KTagIframe
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagIframe Tag = "iframe"

	// KTagImg
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagImg Tag = "img"

	// KTagInput
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagInput Tag = "input"

	// KTagIns
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagIns Tag = "ins"

	// KTagKbd
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagKbd Tag = "kbd"

	// KTagLabel
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagLabel Tag = "label"

	// KTagLegend
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagLegend Tag = "legend"

	// KTagLi
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagLi Tag = "li"

	// KTagLink
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagLink Tag = "link"

	// KTagMain
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagMain Tag = "main"

	// KTagMap
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagMap Tag = "map"

	// KTagMark
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagMark Tag = "mark"

	// KTagExperimental
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagExperimental Tag = "Experimental"

	// KTagMenu
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagMenu Tag = "menu"

	// KTagMeta
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagMeta Tag = "meta"

	// KTagMeter
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagMeter Tag = "meter"

	// KTagNav
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagNav Tag = "nav"

	// KTagNoscript
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagNoscript Tag = "noscript"

	// KTagObject
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagObject Tag = "object"

	// KTagOl
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagOl Tag = "ol"

	// KTagOptgroup
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagOptgroup Tag = "optgroup"

	// KTagOption
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagOption Tag = "option"

	// KTagOutput
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagOutput Tag = "output"

	// KTagP
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagP Tag = "p"

	// KTagParam
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagParam Tag = "param"

	// KTagPicture
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagPicture Tag = "picture"

	// KTagPortal
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagPortal Tag = "portal"

	// KTagPre
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagPre Tag = "pre"

	// KTagProgress
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagProgress Tag = "progress"

	// KTagQ
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagQ Tag = "q"

	// KTagRp
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagRp Tag = "rp"

	// KTagRt
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagRt Tag = "rt"

	// KTagRuby
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagRuby Tag = "ruby"

	// KTagS
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagS Tag = "s"

	// KTagSamp
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagSamp Tag = "samp"

	// KTagScript
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagScript Tag = "script"

	// KTagSection
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagSection Tag = "section"

	// KTagSelect
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagSelect Tag = "select"

	// KTagSlot
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagSlot Tag = "slot"

	// KTagSmall
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagSmall Tag = "small"

	// KTagSource
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagSource Tag = "source"

	// KTagSpan
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagSpan Tag = "span"

	// KTagStrong
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagStrong Tag = "strong"

	// KTagStyle
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagStyle Tag = "style"

	// KTagSub
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagSub Tag = "sub"

	// KTagSummary
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagSummary Tag = "summary"

	// KTagSup
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagSup Tag = "sup"

	// KTagTable
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTable Tag = "table"

	// KTagTbody
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTbody Tag = "tbody"

	// KTagTd
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTd Tag = "td"

	// KTagTemplate
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTemplate Tag = "template"

	// KTagTextarea
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTextarea Tag = "textarea"

	// KTagTfoot
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTfoot Tag = "tfoot"

	// KTagTh
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTh Tag = "th"

	// KTagThead
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagThead Tag = "thead"

	// KTagTime
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTime Tag = "time"

	// KTagTitle
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTitle Tag = "title"

	// KTagTr
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTr Tag = "tr"

	// KTagTrack
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagTrack Tag = "track"

	// KTagU
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagU Tag = "u"

	// KTagUl
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagUl Tag = "ul"

	// KTagVar
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagVar Tag = "var"

	// KTagVideo
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagVideo Tag = "video"

	// KTagWbr
	//
	// English:
	//
	//
	//
	// Português
	//
	//
	KTagWbr Tag = "wbr"
)
