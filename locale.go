package lib

import (
	"regexp"		
)

type Langue struct {
	vieuxfrancais *regexp.Regexp
	vieuxanglais *regexp.Regexp
	anglais *regexp.Regexp
	francais *regexp.Regexp
	Arabic *regexp.Regexp
	Bulgarian *regexp.Regexp
	Cantonese *regexp.Regexp
	Catalan *regexp.Regexp
	Croatian *regexp.Regexp
	Czech *regexp.Regexp
	Danish *regexp.Regexp
	DutchBE *regexp.Regexp
	DutchNL *regexp.Regexp
	Farsi *regexp.Regexp
	Finnish *regexp.Regexp
	Flemmish *regexp.Regexp
	German *regexp.Regexp
	Greek *regexp.Regexp
	Hebrew *regexp.Regexp
	Hindi *regexp.Regexp
	Hungarian *regexp.Regexp
	Icelandic *regexp.Regexp
	Indonesian *regexp.Regexp
	Inuktitut *regexp.Regexp
	Italian *regexp.Regexp
	Japanese *regexp.Regexp
	Korean *regexp.Regexp
	Latvian *regexp.Regexp
	Lithuanian *regexp.Regexp
	Malay *regexp.Regexp
	Mandarin *regexp.Regexp
	Mandarintw *regexp.Regexp
	MandarinS *regexp.Regexp	
	Norwegian *regexp.Regexp
	Polish *regexp.Regexp
	BrazilianPortuguese *regexp.Regexp
	Portuguese *regexp.Regexp
	Romanian *regexp.Regexp
	Russian *regexp.Regexp
	Serbian *regexp.Regexp
	Slovak *regexp.Regexp
	Slovenian *regexp.Regexp
	LatinSpanish *regexp.Regexp
	SpanishMexican *regexp.Regexp
	CastilianSpanish *regexp.Regexp
	Swedish *regexp.Regexp
	Tamil *regexp.Regexp
	Telugu *regexp.Regexp
	Thai *regexp.Regexp
	Turkish *regexp.Regexp
	Ukrainian *regexp.Regexp
	Vietnamese *regexp.Regexp
}

func Locale(filename string) string {
	l := Langue{}	
	l.vieuxanglais = regexp.MustCompile("(?i)_en..-|_en-|_en_|_en.._")
	l.vieuxfrancais = regexp.MustCompile("(?i)_fr..-|_fr-|_fr_|_fr.._")
	l.anglais = regexp.MustCompile("(?i)[0-9]{1}._en..-|[0-9]{1}._en-|[0-9]{1}._en_|[0-9]{1}._en.._|ndf_en..")
	l.francais = regexp.MustCompile("(?i)[0-9]{1}._fr..-|[0-9]{1}._fr-|[0-9]{1}._fr_|[0-9]{1}._fr.._|ndf_fr..")
	l.Arabic = regexp.MustCompile("(?i)_ar-|_ar_")
	l.Bulgarian = regexp.MustCompile("(?i)_bg-|_bg_")
	l.Cantonese = regexp.MustCompile("(?i)_yue-|_yue_")
	l.Catalan = regexp.MustCompile("(?i)_ca-|_ca_")
	l.Croatian = regexp.MustCompile("(?i)_hr-|_hr_")
	l.Czech = regexp.MustCompile("(?i)_cs-|_cs_")
	l.Danish = regexp.MustCompile("(?i)_da-|_da_")
	l.DutchBE = regexp.MustCompile("(?i)_nlBE-|_nlBE_")
	l.DutchNL = regexp.MustCompile("(?i)_nlNL-|_nlNL_")
	l.Farsi = regexp.MustCompile("(?i)_fa-|_fa_")
	l.Finnish = regexp.MustCompile("(?i)_fi-|_fi_")
	l.Flemmish = regexp.MustCompile("(?i)_vl-|_vl_")
	l.German = regexp.MustCompile("(?i)_deDE-|_deDE_")
	l.Greek = regexp.MustCompile("(?i)_el-|_elCY-|_el_")
	l.Hebrew = regexp.MustCompile("(?i)_he-|_he_")
	l.Hindi = regexp.MustCompile("(?i)_hi-|_hi_")
	l.Hungarian = regexp.MustCompile("(?i)_hu-|_hu_")
	l.Icelandic = regexp.MustCompile("(?i)_is-|_is_")
	l.Indonesian = regexp.MustCompile("(?i)_id-|_id_")
	l.Inuktitut = regexp.MustCompile("(?i)_iu-|_iu_")
	l.Italian = regexp.MustCompile("(?i)_it-|_it_")
	l.Japanese = regexp.MustCompile("(?i)_ja-|_ja_")
	l.Korean = regexp.MustCompile("(?i)_ko-|_ko_")
	l.Latvian = regexp.MustCompile("(?i)_lv-|_lv_")
	l.Lithuanian = regexp.MustCompile("(?i)_lt-|_lt_")
	l.Malay = regexp.MustCompile("(?i)_ms-|_ms_")
	l.Mandarin = regexp.MustCompile("(?i)_cmn-|_cnm_")
	l.Mandarintw = regexp.MustCompile("(?i)_cmnTW-|_cnmtw_|_cmnHant_")
	l.MandarinS = regexp.MustCompile("(?i)_cmnHans_")
	l.Norwegian = regexp.MustCompile("(?i)_no-|_no_")
	l.Polish = regexp.MustCompile("(?i)_pl-|_pl_")
	l.BrazilianPortuguese = regexp.MustCompile("(?i)_ptBR-|_ptBR_")
	l.Portuguese = regexp.MustCompile("(?i)_ptPT-|_ptPT_")
	l.Romanian = regexp.MustCompile("(?i)_ro-|_ro_")
	l.Russian = regexp.MustCompile("(?i)_ru-|_ru_")
	l.Serbian = regexp.MustCompile("(?i)_sr-|_sr_")
	l.Slovak = regexp.MustCompile("(?i)_sk-|_sk_")
	l.Slovenian = regexp.MustCompile("(?i)_sl-|_sl_")
	l.LatinSpanish = regexp.MustCompile("(?i)_es419-|_es419_")
	l.SpanishMexican = regexp.MustCompile("(?i)_esMX-|_esmx_")
	l.CastilianSpanish = regexp.MustCompile("(?i)_esES-|_esES_")
	l.Swedish = regexp.MustCompile("(?i)_sv-|_sv_")
	l.Tamil = regexp.MustCompile("(?i)_ta-|_ta_")
	l.Telugu = regexp.MustCompile("(?i)_te-|_te_")
	l.Thai = regexp.MustCompile("(?i)_th-|_th_")
	l.Turkish = regexp.MustCompile("(?i)_tr-|_tr_")
	l.Ukrainian = regexp.MustCompile("(?i)_uk-|_uk_")
	l.Vietnamese = regexp.MustCompile("(?i)_vi-|_vi_")
	langue := ""

	anglais := l.anglais.MatchString(filename)
	if anglais {
		langue="English"
 	}	
	francais := l.francais.MatchString(filename)
 	if francais {
 		langue="French"
 	}
	
	if langue == "" {
		Arabic := l.Arabic.MatchString(filename)
		if Arabic  { 
			langue="Arabic"
		}		
		Bulgarian := l.Bulgarian.MatchString(filename)
		if Bulgarian   {
		
			langue="Bulgarian"
		}		
		Cantonese := l.Cantonese.MatchString(filename)
		if Cantonese   {
		
			langue="Cantonese"
		}		
		Catalan := l.Catalan.MatchString(filename)
		if Catalan   {
		
			langue="Catalan"
		}		
		Croatian := l.Croatian.MatchString(filename)
		if Croatian   {
		
			langue="Croatian"
		}		
		Czech := l.Czech.MatchString(filename)
		if Czech   {
		
			langue="Czech"
		}		
		Danish := l.Danish.MatchString(filename)
		if Danish   {
		
			langue="Danish"
		}		
		DutchBE := l.DutchBE.MatchString(filename)
		if DutchBE   {
		
			langue="Dutch"
		}		
		DutchNL := l.DutchNL.MatchString(filename)
		if DutchNL   {
		
			langue="Dutch"
		}		
		Farsi := l.Farsi.MatchString(filename)
		if Farsi   {
		
			langue="Farsi"
		}		
		Finnish := l.Finnish.MatchString(filename)
		if Finnish   {
		
			langue="Finnish"
		}		
		Flemmish := l.Flemmish.MatchString(filename)
		if Flemmish   {
		
			langue="Flemish"
		}		
		German := l.German.MatchString(filename)
		if German   {
		
			langue="German"
		}		
		Greek := l.Greek.MatchString(filename)
		if Greek   {
		
			langue="Greek"
		}		
		Hebrew := l.Hebrew.MatchString(filename)
		if Hebrew   {
		
			langue="Hebrew"
		}		
		Hindi := l.Hindi.MatchString(filename)
		if Hindi   {
		
			langue="Hindi"
		}		
		Hungarian := l.Hungarian.MatchString(filename)
		if Hungarian   {
		
			langue="Hungarian"
		}		
		Icelandic := l.Icelandic.MatchString(filename)
		if Icelandic   {
		
			langue="Icelandic"
		}		
		Indonesian := l.Indonesian.MatchString(filename)
		if Indonesian   {
		
			langue="Indonesian"
		}		
		Inuktitut := l.Inuktitut.MatchString(filename)
		if Inuktitut   {
		
			langue="Inuktitut"
		}		
		Italian := l.Italian.MatchString(filename)
		if Italian   {
		
			langue="Italian"
		}		
		Japanese := l.Japanese.MatchString(filename)
		if Japanese   {
		
			langue="Japanese"
		}		
		Korean := l.Korean.MatchString(filename)
		if Korean   {
		
			langue="Korean"
		}		
		Latvian := l.Latvian.MatchString(filename)
		if Latvian   {
		
			langue="Latvian"
		}		
		Lithuanian := l.Lithuanian.MatchString(filename)
		if Lithuanian   {
		
			langue="Lithuanian"
		}		
		Malay := l.Malay.MatchString(filename)
		if Malay   {
		
			langue="Malay"
		}		
		Mandarin := l.Mandarin.MatchString(filename)
		if Mandarin   {
		
			langue="Mandarin"
		}		
		Mandarintw := l.Mandarintw.MatchString(filename)
		if Mandarintw   {
			langue="Chinese-Mandarin-Traditional-Taiwan"
		}		
		MandarinS := l.MandarinS.MatchString(filename)
		if MandarinS   {
			langue="Chinese-Mandarin-Simplified-PRC"
		}	
		Norwegian := l.Norwegian.MatchString(filename)
		if Norwegian   {
		
			langue="Norwegian"
		}		
		Polish := l.Polish.MatchString(filename)
		if Polish   {
		
			langue="Polish"
		}		
		BrazilianPortuguese := l.BrazilianPortuguese.MatchString(filename)
		if BrazilianPortuguese   {
		
			langue="Brazilian Portuguese"
		}		
		Portuguese := l.Portuguese.MatchString(filename)
		if Portuguese   {
		
			langue="Portuguese"
		}		
		Romanian := l.Romanian.MatchString(filename)
		if Romanian   {
		
			langue="Romanian"
		}
		Russian := l.Russian.MatchString(filename)
		if Russian   {
		
			langue="Russian"
		}
		
		Serbian := l.Serbian.MatchString(filename)
		if Serbian   {
		
			langue="Serbian"
		}		
		Slovak := l.Slovak.MatchString(filename)
		if Slovak   {
		
			langue="Slovak"
		}
		
		Slovenian := l.Slovenian.MatchString(filename)
		if Slovenian   {
		
			langue="Slovenian"
		}
		
		LatinSpanish := l.LatinSpanish.MatchString(filename)
		if LatinSpanish   {
		
			langue="Latin Spanish"
		}
		
		SpanishMexican := l.SpanishMexican.MatchString(filename)
		if SpanishMexican   {
		
			langue="Spanish-Mexican"
		}
		
		CastilianSpanish := l.CastilianSpanish.MatchString(filename)
		if CastilianSpanish   {
		
			langue="Castilian Spanish"
		}		
		Swedish := l.Swedish.MatchString(filename)
		if Swedish   {
		
			langue="Swedish"
		}		
		Tamil := l.Tamil.MatchString(filename)
		if Tamil   {
		
			langue="Tamil"
		}		
		Telugu := l.Telugu.MatchString(filename)
		if Telugu   {
		
			langue="Telugu"
		}		
		Thai := l.Thai.MatchString(filename)
		if Thai   {
		
			langue="Thai"
		}		
		Turkish := l.Turkish.MatchString(filename)
		if Turkish   {
		
			langue="Turkish"
		}		
		Ukrainian := l.Ukrainian.MatchString(filename)
		if Ukrainian   {
		
			langue="Ukranian"
		}		
		Vietnamese := l.Vietnamese.MatchString(filename)
		if Vietnamese   {
		
			langue="Vietnamese"
		}	
	}
 	
 	return langue
}