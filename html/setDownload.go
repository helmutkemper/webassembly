package html

import (
	"log"
)

// SetDownload
//
// English:
//
//  Causes the browser to treat the linked URL as a download. Can be used with or without a value
//
//   Note:
//     * Without a value, the browser will suggest a filename/extension, generated from various
//       sources:
//         The Content-Disposition HTTP header;
//         The final segment in the URL path;
//         The media type (from the Content-Type header, the start of a data: URL, or Blob.type for a
//         blob: URL).
//     * Defining a value suggests it as the filename. / and \ characters are converted to
//       underscores (_). Filesystems may forbid other characters in filenames, so browsers will
//       adjust the suggested name if necessary;
//     * Download only works for same-origin URLs, or the blob: and data: schemes;
//     * How browsers treat downloads varies by browser, user settings, and other factors. The user
//       may be prompted before a download starts, or the file may be saved automatically, or it may
//       open automatically, either in an external application or in the browser itself;
//     * If the Content-Disposition header has different information from the download attribute,
//       resulting behavior may differ:
//         * If the header specifies a filename, it takes priority over a filename specified in the
//           download attribute;
//         * If the header specifies a disposition of inline, Chrome and Firefox prioritize the
//           attribute and treat it as a download. Old Firefox versions (before 82) prioritize the
//           header and will display the content inline.
//
// Português:
//
//  Faz com que o navegador trate a URL vinculada como um download. Pode ser usado com ou sem valor
//
//   Nota:
//     * Sem um valor, o navegador sugerirá uma extensão de nome de arquivo, gerada a partir de várias
//       fontes:
//         O cabeçalho HTTP Content-Disposition;
//         O segmento final no caminho do URL;
//         O tipo de mídia (do cabeçalho Content-Type, o início de um data: URL ou Blob.type para um
//         blob: URL).
//     * Definir um valor sugere-o como o nome do arquivo. / e \ caracteres são convertidos em
//       sublinhados (_). Os sistemas de arquivos podem proibir outros caracteres em nomes de
//       arquivos, portanto, os navegadores ajustarão o nome sugerido, se necessário;
//     * O download funciona apenas para URLs de mesma origem, ou os esquemas blob: e data: schemes;
//     * A forma como os navegadores tratam os downloads varia de acordo com o navegador, as
//       configurações do usuário e outros fatores. O usuário pode ser avisado antes do início de um
//       download, ou o arquivo pode ser salvo automaticamente, ou pode ser aberto automaticamente,
//       seja em um aplicativo externo ou no próprio navegador;
//     * Se o cabeçalho Content-Disposition tiver informações diferentes do atributo download, o
//       comportamento resultante pode ser diferente:
//         * Se o cabeçalho especificar um nome de arquivo, ele terá prioridade sobre um nome de
//           arquivo especificado no atributo download;
//         * Se o cabeçalho especificar uma disposição de inline, o Chrome e o Firefox priorizarão o
//           atributo e o tratarão como um download. Versões antigas do Firefox (antes de 82)
//           priorizam o cabeçalho e exibirão o conteúdo inline.
func (e *GlobalAttributes) SetDownload(download string) (ref *GlobalAttributes) {
	switch e.tag {
	case KTagA:
	default:
		log.Printf("tag " + e.tag.String() + " does not support download property")
	}

	e.selfElement.Set("download", download)
	return e
}
