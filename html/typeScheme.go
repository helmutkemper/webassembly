package html

type Scheme string

func (e Scheme) String() string {
	return string(e)
}

const (
	KSchemeBitcoin     = "bitcoin"
	KSchemeFtp         = "ftp"
	KSchemeFtps        = "ftps"
	KSchemeGeo         = "geo"
	KSchemeIm          = "im"
	KSchemeIrc         = "irc"
	KSchemeIrcs        = "ircs"
	KSchemeMagnet      = "magnet"
	KSchemeMailTo      = "mailto"
	KSchemeMatrix      = "matrix"
	KSchemeMms         = "mms"
	KSchemeNews        = "news"
	KSchemeNntp        = "nntp"
	KSchemeOpenPgp4Fpr = "openpgp4fpr"
	KSchemeSftp        = "sftp"
	KSchemeSip         = "sip"
	KSchemeSms         = "sms"
	KSchemeSmsTo       = "smsto"
	KSchemeSsh         = "ssh"
	KSchemeTel         = "tel"
	KSchemeUrn         = "urn"
	KSchemeWebCal      = "webcal"
	KSchemeWtai        = "wtai"
	KSchemeXmpp        = "xmpp"
)
