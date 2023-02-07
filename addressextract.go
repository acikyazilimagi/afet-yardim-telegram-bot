package main

import (
	"regexp"
)

type City string

const (
	ADANA      City = "Adana"
	ANTEP           = "Antep"
	BATMAN          = "Batman"
	BINGOL          = "Bingol"
	BITLIS          = "Bitlis"
	DIYARBAKIR      = "Diyarbakir"
	ELAZIG          = "Elazig"
	HAKKARI         = "Hakkari"
	HATAY           = "Hatay"
	KILIS           = "Kilis"
	MALATYA         = "Malatya"
	MARAS           = "Maras"
	MARDIN          = "Mardin"
	MUS             = "Mus"
	OSMANIYE        = "Osmaniye"
	SIIRT           = "Siirt"
	SIRNAK          = "Sirnak"
	VAN             = "Van"
	UNKNOWN         = "Bilinmiyor"
)

var (
	adanaPattern, _      = regexp.Compile("(?i)((Alada[gGğĞ])|(Ceyhan)|(Feke)|([ıIiİ]mamo[gGğĞ]lu)|(Karaisal[ıIiİ])|(Karata[sSşŞ])|(Kozan‎)|(Pozant[ıIiİ]‎)|(Saimbeyli)|(Tufanbeyli‎)|(Yumurtal[ıIiİ]k‎)|([cCçÇ]ukurova)|(Sar[ıIiİ][cCçÇ]am)|(Seyhan)|(Y[uUüÜ]re[gGğĞ]ir))")
	antepPattern, _      = regexp.Compile("(?i)((Araban)|([ıIiİ]slahiye)|(Karkam[ıIiİ][sSşŞ])|(Nizip)|(N[uüUÜ]rda[gGğĞ][ıIiİ])|(O[gGğĞ][uüUÜ]zeli)|(Yav[uüUÜ]zeli)|([sSşŞ]ahinbey)|([sSşŞ]ehitkamil))")
	batmanPattern, _     = regexp.Compile("(?i)((Be[sSşŞ]iri)|(Gerc[uUüÜ][sSşŞ])|(Hasankeyf)|(Kozluk)|(Sason))")
	bingolPattern, _     = regexp.Compile("(?i)((Adakl[ıIiİ])|(Gen[cCçÇ])|(Karl[ıIiİ]ova)|(Ki[gGğĞ][ıIiİ])|(Solhan)|(Yayladere)|(Yedisu))")
	bitlisPattern, _     = regexp.Compile("(?i)((Adilcevaz)|(Ahlat)|(G[uUüÜ]roymak)|(Hizan)|(Mutki)|(Tatvan))")
	diyarbakirPattern, _ = regexp.Compile("(?i)((Bismil)|([cCçÇ]ermik)|([cCçÇ][ıIiİ]nar)|([cCçÇ][uUüÜ]ng[uUüÜ][sSşŞ])|(Dicle)|(E[gGğĞ]il)|(Ergani)|(Hani)|(Hazro)|(Kocak[oOöÖ]y)|(Kulp)|(Lice)|(Silvan)|(Ba[gGğĞ]lar)|(Kayap[ıIiİ]nar)|(Sur)|(Yeni[sSşŞ]ehir))")
	elazigPattern, _     = regexp.Compile("(?i)((A[gGğĞ][ıIiİ]n)|(Alacakaya)|(Ar[ıIiİ]cak)|(Baskil)|(Karako[cCçÇ]an)|(Keban)|(Kovanc[ıIiİ]lar)|(Maden)|(Palu)|(Sivrice))")
	hakkariPattern, _    = regexp.Compile("(?i)(([cCçÇ]ukurca)|([sSşŞ]emdinli)|(Y[uUüÜ]ksekova)|(Derecik))")
	hatayPattern, _      = regexp.Compile("(?i)((Alt[ıIiİ]n[oOöÖ]z[uUüÜ])|(Arsuz)|(Belen‎)|(D[oOöÖ]rtyol‎)|(Erzin‎)|(Hassa‎)|([Iıİi]skenderun‎)|(K[ıIiİ]r[ıIiİ]khan)|(Kumlu)|(Payas)|(Reyhanl[ıIiİ]‎)|(Samanda[gGğĞ]‎)|(Yaylada[gGğĞ][ıIiİ]‎)|(Antakya‎)|(Defne))")
	kilisPattern, _      = regexp.Compile("(?i)((Elbeyli)|(Musabeyli)|(Polateli))")
	malatyaPattern, _    = regexp.Compile("(?i)((Ak[cCçÇ]ada[gGğĞ])|(Arapgir)|(Arguvan)|(Darende)|(Do[gGğĞ]an[sSşŞ]ehir)|(Do[gGğĞ]anyol)|(Hekimhan)|(Kale)|(Kuluncak)|(P[uUüÜ]t[uUüÜ]rge)|(Yaz[ıIiİ]han)|(Battalgazi)|(Ye[sSşŞ]ilyurt))")
	marasPattern, _      = regexp.Compile("(?i)((Af[sSşŞ]in)|(And[ıIiİ]r[ıIiİ]n)|([cCçÇ]a[gGğĞ]layancerit)|(Ekin[oOöÖ]z[uUüÜ])|(Elbistan)|(G[oOöÖ]ksun)|(Nurhak)|(Pazarc[ıIiİ]k)|(T[uUüÜ]rko[gGğĞ]lu)|(Dulkadiro[gGğĞ]lu)|(Oniki[sSşŞ]ubat))")
	mardinPattern, _     = regexp.Compile("(?i)((Artuklu)|(Darge[cCçÇ]it)|(Derik)|(K[ıIiİ]z[ıIiİ]ltepe)|(Maz[ıIiİ]da[gGğĞ][ıIiİ])|(Midyat)|(Nusaybin)|([oOöÖ]merli)|(Savur)|(Ye[sSşŞ]illi))")
	musPattern, _        = regexp.Compile("(?i)((Bulan[ıIiİ]k)|(Hask[oOöÖ]y)|(Korkut)|(Malazgirt)|(Varto))")
	osmaniyePattern, _   = regexp.Compile("(?i)((Bah[cCçÇ]e)|(D[uUüÜ]zi[cCçÇ]i)|(Hasanbeyli)|(Kadirli)|(Sumbas)|(Toprakkale))")
	siirtPattern, _      = regexp.Compile("(?i)((Baykan)|(Eruh)|(Kurtalan)|(Pervari)|([sSşŞ]irvan)|(Tillo))")
	sirnakPattern, _     = regexp.Compile("(?i)((Beyt[uUüÜ][sSşŞ][sSşŞ]ebap)|(Cizre)|(G[uUüÜ][cCçÇ]l[uUüÜ]konak)|([Iıİi]dil)|(Silopi)|(Uludere))")
	vanPattern, _        = regexp.Compile("(?i)((Bah[cCçÇ]esaray)|(Ba[sSşŞ]kale)|([cCçÇ]ald[ıIiİ]ran)|([cCçÇ]atak)|(Erci[sSşŞ])|(Geva[sSşŞ])|(G[uUüÜ]rp[ıIiİ]nar)|(Muradiye)|([oOöÖ]zalp)|(Saray)|(Edremit)|([Iıİi]pekyolu)|(Tu[sSşŞ]ba))")

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

	cityPattern, _ = regexp.Compile(`(?i)(?P<cities>(?P<Antep>gaz[ıiİI]antep)|(?P<Malatya>malatya)|(?P<Batman>batman)|(?P<Bingol>b[ıiIİ]ng[oöOÖ]l)|(?P<Elazig>elaz[Iİıi][gğ])|(?P<Kilis>kilis)|(?P<Diyarbakir>diyarbak[ıiIİ]r)|(?P<Mardin>mardin)|(?P<Siirt>siirt)|(?P<Sirnak>[SsŞş][ıiIİ]rnak)|(?P<Van>van)|(?P<Mus>mu[sşSŞ])|(?P<Bitlis>bitlis)|(?P<Hakkari>hakkari)|(?P<Adana>adana)|(?P<Osmaniye>osmaniye)|(?P<Hatay>hatay)|(?P<Maras>kahramanmara[sşSŞ])|(?P<Maras2>mara[SŞsş])|(?P<Antep2>antep))`)
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
	if city == UNKNOWN {
		return ""
	}

	regex := districtPatterns[city]
	return regex.FindString(s)
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
