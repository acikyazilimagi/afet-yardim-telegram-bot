package main

import (
	"regexp"
)

type City string

const (
	ADANA      City = "Adana"
	ANTEP      City = "Antep"
	BATMAN     City = "Batman"
	BINGOL     City = "Bingol"
	BITLIS     City = "Bitlis"
	DIYARBAKIR City = "Diyarbakir"
	ELAZIG     City = "Elazig"
	HAKKARI    City = "Hakkari"
	HATAY      City = "Hatay"
	KILIS      City = "Kilis"
	MALATYA    City = "Malatya"
	MARAS      City = "Maras"
	MARDIN     City = "Mardin"
	MUS        City = "Mus"
	OSMANIYE   City = "Osmaniye"
	SIIRT      City = "Siirt"
	SIRNAK     City = "Sirnak"
	VAN        City = "Van"
	UNKNOWN    City = "Bilinmiyor"
)

var (
	adanaPattern      = regexp.MustCompile("(?i)((Alada[gGğĞ])|(Ceyhan)|(Feke)|([ıIiİ]mamo[gGğĞ]lu)|(Karaisal[ıIiİ])|(Karata[sSşŞ])|(Kozan‎)|(Pozant[ıIiİ]‎)|(Saimbeyli)|(Tufanbeyli‎)|(Yumurtal[ıIiİ]k‎)|([cCçÇ]ukurova)|(Sar[ıIiİ][cCçÇ]am)|(Seyhan)|(Y[uUüÜ]re[gGğĞ]ir))")
	antepPattern      = regexp.MustCompile("(?i)((Araban)|([ıIiİ]slahiye)|(Karkam[ıIiİ][sSşŞ])|(Nizip)|(N[uüUÜ]rda[gGğĞ][ıIiİ])|(O[gGğĞ][uüUÜ]zeli)|(Yav[uüUÜ]zeli)|([sSşŞ]ahinbey)|([sSşŞ]ehitkamil))")
	batmanPattern     = regexp.MustCompile("(?i)((Be[sSşŞ]iri)|(Gerc[uUüÜ][sSşŞ])|(Hasankeyf)|(Kozluk)|(Sason))")
	bingolPattern     = regexp.MustCompile("(?i)((Adakl[ıIiİ])|(Gen[cCçÇ])|(Karl[ıIiİ]ova)|(Ki[gGğĞ][ıIiİ])|(Solhan)|(Yayladere)|(Yedisu))")
	bitlisPattern     = regexp.MustCompile("(?i)((Adilcevaz)|(Ahlat)|(G[uUüÜ]roymak)|(Hizan)|(Mutki)|(Tatvan))")
	diyarbakirPattern = regexp.MustCompile("(?i)((Bismil)|([cCçÇ]ermik)|([cCçÇ][ıIiİ]nar)|([cCçÇ][uUüÜ]ng[uUüÜ][sSşŞ])|(Dicle)|(E[gGğĞ]il)|(Ergani)|(Hani)|(Hazro)|(Kocak[oOöÖ]y)|(Kulp)|(Lice)|(Silvan)|(Ba[gGğĞ]lar)|(Kayap[ıIiİ]nar)|(Sur)|(Yeni[sSşŞ]ehir))")
	elazigPattern     = regexp.MustCompile("(?i)((A[gGğĞ][ıIiİ]n)|(Alacakaya)|(Ar[ıIiİ]cak)|(Baskil)|(Karako[cCçÇ]an)|(Keban)|(Kovanc[ıIiİ]lar)|(Maden)|(Palu)|(Sivrice))")
	hakkariPattern    = regexp.MustCompile("(?i)(([cCçÇ]ukurca)|([sSşŞ]emdinli)|(Y[uUüÜ]ksekova)|(Derecik))")
	hatayPattern      = regexp.MustCompile("(?i)((Alt[ıIiİ]n[oOöÖ]z[uUüÜ])|(Arsuz)|(Belen‎)|(D[oOöÖ]rtyol‎)|(Erzin‎)|(Hassa‎)|([Iıİi]skenderun‎)|(K[ıIiİ]r[ıIiİ]khan)|(Kumlu)|(Payas)|(Reyhanl[ıIiİ]‎)|(Samanda[gGğĞ]‎)|(Yaylada[gGğĞ][ıIiİ]‎)|(Antakya‎)|(Defne))")
	kilisPattern      = regexp.MustCompile("(?i)((Elbeyli)|(Musabeyli)|(Polateli))")
	malatyaPattern    = regexp.MustCompile("(?i)((Ak[cCçÇ]ada[gGğĞ])|(Arapgir)|(Arguvan)|(Darende)|(Do[gGğĞ]an[sSşŞ]ehir)|(Do[gGğĞ]anyol)|(Hekimhan)|(Kale)|(Kuluncak)|(P[uUüÜ]t[uUüÜ]rge)|(Yaz[ıIiİ]han)|(Battalgazi)|(Ye[sSşŞ]ilyurt))")
	marasPattern      = regexp.MustCompile("(?i)((Af[sSşŞ]in)|(And[ıIiİ]r[ıIiİ]n)|([cCçÇ]a[gGğĞ]layancerit)|(Ekin[oOöÖ]z[uUüÜ])|(Elbistan)|(G[oOöÖ]ksun)|(Nurhak)|(Pazarc[ıIiİ]k)|(T[uUüÜ]rko[gGğĞ]lu)|(Dulkadiro[gGğĞ]lu)|(Oniki[sSşŞ]ubat))")
	mardinPattern     = regexp.MustCompile("(?i)((Artuklu)|(Darge[cCçÇ]it)|(Derik)|(K[ıIiİ]z[ıIiİ]ltepe)|(Maz[ıIiİ]da[gGğĞ][ıIiİ])|(Midyat)|(Nusaybin)|([oOöÖ]merli)|(Savur)|(Ye[sSşŞ]illi))")
	musPattern        = regexp.MustCompile("(?i)((Bulan[ıIiİ]k)|(Hask[oOöÖ]y)|(Korkut)|(Malazgirt)|(Varto))")
	osmaniyePattern   = regexp.MustCompile("(?i)((Bah[cCçÇ]e)|(D[uUüÜ]zi[cCçÇ]i)|(Hasanbeyli)|(Kadirli)|(Sumbas)|(Toprakkale))")
	siirtPattern      = regexp.MustCompile("(?i)((Baykan)|(Eruh)|(Kurtalan)|(Pervari)|([sSşŞ]irvan)|(Tillo))")
	sirnakPattern     = regexp.MustCompile("(?i)((Beyt[uUüÜ][sSşŞ][sSşŞ]ebap)|(Cizre)|(G[uUüÜ][cCçÇ]l[uUüÜ]konak)|([Iıİi]dil)|(Silopi)|(Uludere))")
	vanPattern        = regexp.MustCompile("(?i)((Bah[cCçÇ]esaray)|(Ba[sSşŞ]kale)|([cCçÇ]ald[ıIiİ]ran)|([cCçÇ]atak)|(Erci[sSşŞ])|(Geva[sSşŞ])|(G[uUüÜ]rp[ıIiİ]nar)|(Muradiye)|([oOöÖ]zalp)|(Saray)|(Edremit)|([Iıİi]pekyolu)|(Tu[sSşŞ]ba))")
)

var (
	districtPatterns = map[City]*regexp.Regexp{
		ADANA:      adanaPattern,
		ANTEP:      antepPattern,
		BATMAN:     batmanPattern,
		BINGOL:     bingolPattern,
		BITLIS:     bitlisPattern,
		DIYARBAKIR: diyarbakirPattern,
		ELAZIG:     elazigPattern,
		HAKKARI:    hakkariPattern,
		HATAY:      hatayPattern,
		KILIS:      kilisPattern,
		MALATYA:    malatyaPattern,
		MARAS:      marasPattern,
		MARDIN:     mardinPattern,
		MUS:        musPattern,
		OSMANIYE:   osmaniyePattern,
		SIIRT:      siirtPattern,
		SIRNAK:     sirnakPattern,
		VAN:        vanPattern,
	}
)

var (
	cityPattern = regexp.MustCompile(`(?i)(?P<cities>(?P<Antep>gaz[ıiİI]antep)|(?P<Malatya>malatya)|(?P<Batman>batman)|(?P<Bingol>b[ıiIİ]ng[oöOÖ]l)|(?P<Elazig>elaz[Iİıi][gğ])|(?P<Kilis>kilis)|(?P<Diyarbakir>diyarbak[ıiIİ]r)|(?P<Mardin>mardin)|(?P<Siirt>siirt)|(?P<Sirnak>[SsŞş][ıiIİ]rnak)|(?P<Van>van)|(?P<Mus>mu[sşSŞ])|(?P<Bitlis>bitlis)|(?P<Hakkari>hakkari)|(?P<Adana>adana)|(?P<Osmaniye>osmaniye)|(?P<Hatay>hatay)|(?P<Maras>kahramanmara[sşSŞ])|(?P<Maras2>mara[SŞsş])|(?P<Antep2>antep))`)
)

func ExtractCity(s string) City {
	cityName := getCityName(s)
	if cityName == "Maras2" {
		return MARAS
	}

	if cityName == "Antep2" {
		return ANTEP
	}

	if cityName == "" {
		return UNKNOWN
	}

	return City(cityName)
}

func ExtractDistrict(city City, s string) string {
	if match, ok := districtPatterns[city]; ok {
		return match.FindString(s)
	}

	return ""
}

func getCityName(s string) string {
	match := cityPattern.FindStringSubmatch(s)[2:]

	cityNames := cityPattern.SubexpNames()[2:]
	for i, name := range cityNames {
		if i <= len(match) {
			if len(match[i]) > 0 && len(name) > 0 {
				return name
			}
		}
	}
	return ""
}
