package translate

import (
	"embed" // embedding dos arquivos de tradução
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n" // core do i18n
	"golang.org/x/text/language"            // definição de idiomas
	"io/fs"
	"log"
	"regexp"
	"syscall/js"
)

//go:embed locales/*.toml
var localeFS embed.FS
var localeFiles []string
var re *regexp.Regexp
var Localizer *i18n.Localizer

func init() {
	var err error
	if re, err = regexp.Compile("locales/([a-zA-Z0-9_\\-]+)\\.toml"); err != nil {
		log.Printf("translate: failed to compile regex: %s", err)
	}

	// ReadDir lista tudo dentro de "locales"
	entries, err := fs.ReadDir(localeFS, "locales")
	if err != nil {
		log.Printf("translate: failed to read dir: %s", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			localeFiles = append(localeFiles, "locales/"+entry.Name())
		}
	}
}

func Load() {
	var err error

	// todo: linguagem de escolha do usuário
	browserLang := js.Global().Get("navigator").Get("language").String()

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	langs := make([]string, 0)
	for _, file := range localeFiles {
		if _, err = bundle.LoadMessageFileFS(localeFS, file); err != nil {
			log.Printf("translate: failed to load %s: %s", file, err)
		}

		langs = append(langs, re.ReplaceAllString(file, "${1}"))
	}

	langs = append([]string{browserLang}, langs...)
	Localizer = i18n.NewLocalizer(bundle, langs...)
}

/*
goi18n extract \
-sourceLanguage en-US \
-bundleType toml \
-outdir ide/translate/locales \
.
*/
