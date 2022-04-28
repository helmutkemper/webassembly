package html

type Scheme string

func (e Scheme) String() string {
	return string(e)
}

const (
	KSchemeBitcoin     Scheme = "bitcoin"
	KSchemeFtp         Scheme = "ftp"
	KSchemeFtps        Scheme = "ftps"
	KSchemeGeo         Scheme = "geo"
	KSchemeIm          Scheme = "im"
	KSchemeIrc         Scheme = "irc"
	KSchemeIrcs        Scheme = "ircs"
	KSchemeMagnet      Scheme = "magnet"
	KSchemeMailTo      Scheme = "mailto"
	KSchemeMatrix      Scheme = "matrix"
	KSchemeMms         Scheme = "mms"
	KSchemeNews        Scheme = "news"
	KSchemeNntp        Scheme = "nntp"
	KSchemeOpenPgp4Fpr Scheme = "openpgp4fpr"
	KSchemeSftp        Scheme = "sftp"
	KSchemeSip         Scheme = "sip"
	KSchemeSms         Scheme = "sms"
	KSchemeSmsTo       Scheme = "smsto"
	KSchemeSsh         Scheme = "ssh"
	KSchemeTel         Scheme = "tel"
	KSchemeUrn         Scheme = "urn"
	KSchemeWebCal      Scheme = "webcal"
	KSchemeWtai        Scheme = "wtai"
	KSchemeXmpp        Scheme = "xmpp"
)
