package main

import (
	"github.com/resty"	
	"log"
	"encoding/xml"	
	"html/template"
	"defaults"
	"bytes"
	"encoding/json"
	"go-wkhtmltopdf"
	"regexp"		
	"strings"
	"os"
	"strconv"
	"path/filepath"
	"rapport-qc/lib"
	"io/ioutil"	

	//"os/exec"
	//"net/http"
	// "reflect"
	//"rapport-qc/lib"
	// "encoding/xml"
	// "time"	
	// "fmt"
)

var (    
	MetaVers	string = "0.3"
	Build   	string = ""
	tmpl		string = "/qc_report.tmpl"
	VideoScan	string = "p"
)

type Tpl_Version struct {
	XMLName 	xml.Name	`xml:"html"`
	Meta		Meta		`xml:"meta"`
}

type Meta struct {
	Version		string	 	`xml:"version,attr"`
}

type Resolution struct {
	Nom			string
	Court		string
}

type CatRest struct {
	Client		*resty.Client
	ClipIDs		[]string
	Endpoint	string
}

type Fields struct {
	Titre							string
	Distributeur					string
	Version							string
	Saison							string
	Episode							string
	Qc_type							string		`json:"qc.type"`
	Provider						string
	Validation_type					string		`json:"validation.type"`
	Date							string
	Production_year					string		`json:"production.year"`
	Qc_operator						string		`json:"qc.operator"`
	Facility						string		
	Source_filename					string		`json:"source.filename"`
	Coordinator						string		
	Asset_type						string		`json:"asset.type"`
//	Creation_de_rapports			[]string	`json:"création.de.rapports"`
	Baton_taskid					string		`json:"baton.taskid"`
	Narrative_text_presence_1		string		`json:"narrative.text.presence.1"`
	Burnedin_dialogs_presence		string		`json:"burned-in.dialogs.presence"`
	Text_language					string		`json:"text.language"`
	Nom_original					string		`json:"nom.original"`
	Commentaires_generaux			string		`json:"commentaires.généraux"`
	Commentaires_packaging			string		`json:"commentaires.packaging"`
	Aspect_ratio					string		`json:"aspect.ratio"`
	Source_audio_ch_1				string		`json:"source.audio.ch.1"`
	Source_audio_ch_2				string		`json:"source.audio.ch.2"`
	Source_audio_ch_3				string		`json:"source.audio.ch.3"`
	Source_audio_ch_4				string		`json:"source.audio.ch.4"`
	Source_audio_ch_5				string		`json:"source.audio.ch.5"`
	Source_audio_ch_6				string		`json:"source.audio.ch.6"`
	Source_audio_ch_7				string		`json:"source.audio.ch.7"`
	Source_audio_ch_8				string		`json:"source.audio.ch.8"`
}	

type Data struct  {
	Name		string
	Fields		Fields
}

type CatDV struct  {
	Data		Data
}

type StartTimeCodeSMPTE struct {
	Value			string				`xml:"value,attr"`
}

type DurationSMPTE struct {
	Value			string				`xml:"value,attr"`
}

type DropFrameFlag struct {
	Value			bool				`xml:"value,attr"`
}

type ChromaBitDepth struct {
	Value			string				`xml:"value,attr"`
}

type TimeCodeTrack struct {
	DropFrameFlag			DropFrameFlag
	DurationSMPTE			DurationSMPTE
	StartTimeCodeSMPTE		StartTimeCodeSMPTE
}

type ActualTracks struct {
	Value			string				`xml:"value,attr"`
}

type Field struct {
	Name			string				`xml:"name,attr"`	
	Value			string				`xml:"value,attr"`
	ActualTracks	ActualTracks				
	TimeCodeTrack	TimeCodeTrack	
	ChromaBitDepth	ChromaBitDepth
}

type Info struct {
	Field			[]Field				`xml:"field"`	
}

type Error struct {
	FrameDuration				int				`xml:",attr"`
	SMPTETimecodeDuration		string			`xml:",attr"`
	Description					string			`xml:"description,attr"`
	Endsmptetimecode			string			`xml:"endsmptetimecode,attr"`
	Item						string			`xml:"item,attr"`
	Startsmptetimecode			string			`xml:"startsmptetimecode,attr"`
	Synopsis					string			`xml:"synopsis,attr"`
	Severity					string			`xml:"severity,attr"`
}

type Decodedaudiochecks struct {
	Name			string				`xml:"name,attr"`	
	Error			[]Error				`xml:"error"`	
}

type Decodedvideochecks struct {
	Name			string				`xml:"name,attr"`	
	Error			[]Error				`xml:"error"`	
}

type Customchecks struct {
	Decodedvideochecks			Decodedvideochecks				`xml:"decodedvideochecks"`	
	Decodedaudiochecks			Decodedaudiochecks				`xml:"decodedaudiochecks"`	
}

type Errors struct {
	Customchecks			Customchecks				`xml:"customchecks"`	
}

type Streamnode struct {
	EncodedDuration			string				`xml:",attr"`
	Name					string				`xml:"name,attr"`
	Info					Info				`xml:"info"`
	Errors					Errors				`xml:"errors"`	
	Id						string				`xml:"id,attr"`	
}

type Toplevelinfo struct {
	Error			string				`xml:",attr"`	
	Format			string				`xml:",attr"`
}

type TaskReport struct {	
	XMLName 			xml.Name					`xml:"taskReport"`
	Toplevelinfo		Toplevelinfo				`xml:"toplevelinfo"`
	Streamnode  		[]Streamnode				`xml:"streamnode"`
}

type QCInfos struct {
	TITLE					string			`default:" "`
	PROVIDER				string			`default:" "`
	VALIDATION_TYPE			string			`default:" "`
	VERSION					string			`default:" "`
	DATE					string			`default:" "`
	YEAR					string			`default:" "`
	OPERATOR				string			`default:" "`
	SOURCE_FILENAME			string			`default:" "`
	ASSET					string			`default:" "`
	ASSET_NUM				string			`default:" "`
	ASSET_PASS				string			`default:" "`
	NUM01					string			`default:" "`
	NUM02					string			`default:" "`
	FILE_DATE				string			`default:" "`
	REV						string			`default:" "`
	EXT						string			`default:".mov"`
	FILENAME				string			`default:" "`
	ASSET_LANG				string			`default:" "`
	AUDIO_LANG				string			`default:" "`
	SUB_LANG				string			`default:" "`
	TEXT_PRES				string			`default:" "`
	SUB_PRES				string			`default:" "`
	RUN_TIME				string			`default:" "`
	FILESIZE				int				`default:"0"`
	BITRATE					string			`default:" "`
	DIF_RATE				string			`default:" "`
	SRC_RATE				string			`default:" "`
	RESOLUTION				string			`default:" "`
	RATIO					string			`default:" "`
	PXLRATIO				string			`default:" "`
	CODEC					string			`default:" "`
	RANGE					string			`default:"SDR"`
	COLORSAMPLE				string			`default:" "`
	AUDIOSAMPLE				string			`default:" "`
	AUDIO_BIT				string			`default:" "`
	VID_BIT					string			`default:" "`
	V						string			`default:" "`
	MODIF_FROM_SOURCE		string			`default:"N/A"`
	HEAD_LOGO				HEAD_LOGO
	TAIL_LOGO				TAIL_LOGO
	DIF_TC					DIF_TC
	SRC_TC					SRC_TC
	DIF_CH					[]string		`default:"[\" \", \" \", \" \",\" \", \" \", \" \", \" \", \" \", \" \"]"`
	SRC_CH					[]string		
	GENCOM					GENCOM
	PRIM_ISSUES 			[]PRIM_ISSUES	
	OTT_REJECT				[]OTT_REJECT
	TEXT_PRESENCE			[]TEXT_PRESENCE
	NEEDED_TEXT				[]NEEDED_TEXT
	TEXT_ISSUES				[]TEXT_ISSUES
}

type HEAD_LOGO struct {
	DIF_TC_LOGO			string			`default:" "`
	SRC_TC_LOGO			string			`default:" "`
	LOGO				string			`default:" "`
}	

type TAIL_LOGO struct {
	DIF_TC_TAIL_LOGO			string			`default:" "`
	SRC_TC_TAIL_LOGO			string			`default:" "`
	LOGO						string			`default:" "`
}	


type DIF_TC struct {
	DIF_START_TC			string			`default:" "`
	DIF_PROGIN_TC			string			`default:" "`
	DIF_TITLE_TC			string			`default:" "`
	DIF_MAINCRED_TC			string			`default:" "`
	DIF_ENDCRED_TC			string			`default:" "`
	DIF_PROGOUT_TC			string			`default:" "`
	DIF_TXTLIN_TC			string			`default:" "`
	DIF_TXTLOUT_TC			string			`default:" "`
}	

type SRC_TC struct {
	SRC_START_TC			string			`default:" "`
	SRC_PROGIN_TC			string			`default:" "`
	SRC_TITLE_TC			string			`default:" "`
	SRC_MAINCRED_TC			string			`default:" "`
	SRC_ENDCRED_TC			string			`default:" "`
	SRC_PROGOUT_TC			string			`default:" "`
	SRC_TXTLIN_TC			string			`default:" "`
	SRC_TXTLOUT_TC			string			`default:" "`
}	

type GENCOM struct {
	GENCOM_CODE			[]string		`default:"[\" \", \" \", \" \",\" \", \" \", \" \", \" \"]"`
	GENCOM_VALUE		[]string		`default:"[\" \", \" \", \" \",\" \", \" \", \" \", \" \"]"`
}

type PRIM_ISSUES struct {
	DIF_TC_PRIM				string			`default:" "`
	SRC_TC_PRIM				string			`default:" "`
	DESC_PRIM				string			`default:" "`
	DURATION_PRIM			string			`default:" "`
	ISAUDIO_PRIM			string			`default:" "`
	ISVIDEO_PRIM			string			`default:" "`
	CHAN_PRIM				string			`default:" "`
	SCALE_PRIM				string			`default:" "`
	FIXED_PRIM				string			`default:" "`
}

type OTT_REJECT struct {
	DIF_OTT_REJECT_TC				string			`default:" "`
	SRC_OTT_REJECT_TC				string			`default:" "`
	OTT_REJECT_DESC					string			`default:" "`
	OTT_REJECT_DUR					string			`default:" "`
	OTT_REJECT_ISAUDIO				string			`default:" "`
	OTT_REJECT_ISVIDEO				string			`default:" "`
	OTT_REJECT_CHAN					string			`default:" "`
	OTT_REJECT_SCALE				string			`default:" "`
	OTT_REJECT_FIXE					string			`default:" "`
}

type TEXT_PRESENCE struct {
	DIF_TC_TEXT			string			`default:" "`
	SRC_TC_TEXT			string			`default:" "`
	TEXT				string			`default:" "`
	DUR_TEXT			string			`default:" "`
	TYPE_TEXT			string			`default:" "`
}

type NEEDED_TEXT struct {
	DIF_TC_NEEDEDTEXT			string			`default:" "`
	SRC_TC_NEEDEDTEXT			string			`default:" "`
	NEEDEDTEXT					string			`default:" "`
	NEEDEDTEXT_DUR				string			`default:" "`
	NEEDEDTEXT_TYPE				string			`default:" "`
	NEEDEDTEXT_SCALE			string			`default:" "`
	NEEDEDTEXT_FIXE				string			`default:" "`
}

type TEXT_ISSUES struct {
	DIFF_TC_ISSUETEXT			string			`default:" "`
	SRC_TC_ISSUETEXT			string			`default:" "`
	ISSUETEXT					string			`default:" "`
	ISSUETEXT_DUR				string			`default:" "`
	ISSUETEXT_TYPE				string			`default:" "`
	ISSUETEXT_SCALE				string			`default:" "`
	ISSUETEXT_FIXE				string			`default:" "`
}

//"420127" 420108
func main() {			
	CleanChar := regexp.MustCompile("[^a-zA-Z0-9]+")
	if len(os.Args) < 2 {
		log.Println("Il faut un argument.")
	} else {
		switch os.Args[1] {
			case "v":
				log.Println(Build)
			default:	
				if len(os.Args[1]) > 5 {
					if Build != "" {log.Println("Rapport-QC v"+Build)}
					ex, _ := os.Executable()
					var (
						Tpl_File *os.File
					 	Vers Tpl_Version
						Template *template.Template
					)					
					CatInfos := CatRest{resty.New(), strings.Split(os.Args[1], ","), "http://10.99.139.220:8080/api/9/"}

					if _, err := os.Stat(filepath.Dir(ex)+tmpl); err == nil {						
						Tpl_File, err = os.Open(filepath.Dir(ex)+tmpl)
						if err != nil {
							log.Println(err)
						}
					} else if os.IsNotExist(err) {						
						Tpl_File, err = os.Open("C:/Users/daudels/go/src/rapport-qc/qc_report.tmpl")
						if err != nil {
							log.Println(err)
						}						
					}
					TplByte, err := ioutil.ReadAll(Tpl_File)
						if err != nil {
							log.Fatal("Erreur de lecture du template :", err)
						}
					xml.Unmarshal(TplByte, &Vers)
					if Vers.Meta.Version == MetaVers {
						log.Println("Utilisation du template v"+Vers.Meta.Version)
					} else {
						log.Fatal("La version du template n'est pas bonne.")
					}
					Template = template.Must(template.New("tpl").Parse(string(TplByte)))			

					for _, Clip := range CatInfos.CatDV() {
						var (
							Baton TaskReport
							InfosCatDV CatDV
							tpl bytes.Buffer
							)
						QC := Init(&QCInfos{})
						if err := defaults.Set(QC); err != nil {
							panic(err)
						}
						err := json.Unmarshal(Clip, &InfosCatDV)
						if err != nil {
							log.Println("JSON error:", err)
						}
						HeadQC(QC, InfosCatDV.Data.Fields)			

						err = xml.Unmarshal(BatonRestCall(InfosCatDV.Data.Fields.Baton_taskid), &Baton)	
						if err != nil {
							log.Fatal("error baton :", err)							
						}

						for _, Field := range Baton.Streamnode[0].Info.Field {
							if Field.Name == "MP4::TimeCodeTrack" {
								QC.RUN_TIME = Field.TimeCodeTrack.DurationSMPTE.Value
								QC.DIF_TC.DIF_START_TC = Field.TimeCodeTrack.StartTimeCodeSMPTE.Value
							}
							if Field.Name == "FileSize" {
								Size, _ := strconv.Atoi(Field.Value)
								QC.FILESIZE = ByteToMB(Size)
							}
							if Field.Name == "Container::Bitrate" {			
								QC.BITRATE = Field.Value
							}
						}
						for _, Field := range Baton.Streamnode[1].Info.Field {		
							if Field.Name == "Picture Scanning Type" {				
								if Field.Value != "Progressive" {
									VideoScan = "i"
								}
							}
							if Field.Name == "Frame Rate" {			
								QC.DIF_RATE = Field.Value+VideoScan
							}
							if Field.Name == "Resolution" {			
								QC.RESOLUTION = Field.Value
							}
							if Field.Name == "ProresCodecType" {			
								QC.CODEC = Field.Value
							}
							if Field.Name == "Sample Aspect Ratio" {			
								QC.PXLRATIO = Field.Value
							}
							if Field.Name == "Bits Per Pixel" {			
								QC.VID_BIT = Field.ChromaBitDepth.Value+" bits"
							}
							if Field.Name == "Chroma Format" {				
								QC.COLORSAMPLE = Field.Value
							}
						}	
						for _, Field := range Baton.Streamnode[2].Info.Field {		
							if Field.Name == "Sampling Frequency" {	
								Sample, _ := strconv.Atoi(Field.Value)			
								QC.AUDIOSAMPLE = strconv.Itoa(Sample/1000)+" KHz"
							}
							if Field.Name == "Bits per sample" {			
								QC.AUDIO_BIT = Field.Value+" bits"
							}
							if Field.Name == "Audio Channels" {			
								if Field.Value == "2" {
									QC.DIF_CH[1] = "STEREO LEFT"
									QC.DIF_CH[2] = "STEREO RIGHT"
								} else {
									QC.DIF_CH[1] = "5.1 LEFT"
									QC.DIF_CH[2] = "5.1 RIGHT"
									QC.DIF_CH[3] = "5.1 CENTER"
									QC.DIF_CH[4] = "5.1 CENTER"
									QC.DIF_CH[5] = "5.1 LEFT SURROUND"
									QC.DIF_CH[6] = "5.1 RIGHT SURROUND"
									QC.DIF_CH[7] = "LEFT TOTAL"
									QC.DIF_CH[8] = "RIGHT TOTAL"
								}
							}
						}
						ErrorIDX := 0
						for _, Error := range Baton.Streamnode[1].Errors.Customchecks.Decodedvideochecks.Error {
							QC.PRIM_ISSUES[ErrorIDX].DIF_TC_PRIM 		= Error.Startsmptetimecode
							QC.PRIM_ISSUES[ErrorIDX].DESC_PRIM 			= Error.Item
							QC.PRIM_ISSUES[ErrorIDX].DURATION_PRIM  	= Error.SMPTETimecodeDuration
							QC.PRIM_ISSUES[ErrorIDX].ISVIDEO_PRIM 		= "x"
							QC.PRIM_ISSUES[ErrorIDX].SCALE_PRIM 		= Scale(Error.Severity)
							ErrorIDX++
						}
						
						for _, Node := range Baton.Streamnode {
							if Node.Name == "LPCM Audio" {			 
								for _, Error := range Node.Errors.Customchecks.Decodedaudiochecks.Error {
									QC.PRIM_ISSUES[ErrorIDX].DIF_TC_PRIM 		= Error.Startsmptetimecode
									QC.PRIM_ISSUES[ErrorIDX].DESC_PRIM 			= Error.Item
									QC.PRIM_ISSUES[ErrorIDX].DURATION_PRIM  	= Error.SMPTETimecodeDuration
									QC.PRIM_ISSUES[ErrorIDX].ISAUDIO_PRIM 		= "x"
									QC.PRIM_ISSUES[ErrorIDX].SCALE_PRIM 		= Scale(Error.Severity)
									QC.PRIM_ISSUES[ErrorIDX].CHAN_PRIM 			= Node.Id
									ErrorIDX++
								}
							}
						}
						QC.VERSION = QC.Version()				

						if err := Template.Execute(&tpl, QC); err != nil {
							log.Println(err)
						}
						PDF(tpl.String(), CleanChar.ReplaceAllString(QC.FILENAME[:len(QC.FILENAME)-4],"_"))
					}
				} else {log.Println("ClipID invalide.")}	
			}
		}
}

func BatonRestCall(BatonTaskId string) []byte {
	client := resty.New()
	client.SetDisableWarn(true)
	resp, err := client.R().
		SetBasicAuth("dev_user", "jhg45W&sd_18").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").		
		Get("http://baton-2:8080/Baton/api/tasks/"+BatonTaskId+"/report?type=xml")		
	if err != nil {
		log.Fatal("Problème d'accès à Baton...")
	}
	
	return []byte(resp.String())
}

func PDF(Rapport, Filename string) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
	  log.Fatal(err)
	}	
	
	if _, err = os.Stat("C:/Users/daudels/go/src/template_test/"); err == nil {
		pdfg.OutputFile = "C:/Users/daudels/go/src/template_test/"+Filename+"_qc_report.pdf"
	} else if os.IsNotExist(err) {		
		pdfg.OutputFile = "/data/"+Filename+"_qc_report.pdf"
	}
	
	page := wkhtmltopdf.NewPageReader(strings.NewReader(Rapport))
	page.Zoom.Set(1.05)
	pdfg.AddPage(page)
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Le rapport QC en PDF a été généré :",pdfg.OutputFile)
	}
}

func Init(QCChamps *QCInfos) *QCInfos {
	for i := 0; i < 11; i++ {
		QCChamps.PRIM_ISSUES = append(QCChamps.PRIM_ISSUES, PRIM_ISSUES{" "," "," "," "," "," "," "," "," "})
	}
	for i := 0; i < 5; i++ {
		QCChamps.OTT_REJECT = append(QCChamps.OTT_REJECT, OTT_REJECT{" "," "," "," "," "," "," "," "," "})
	}	
	for i := 0; i < 5; i++ {
		QCChamps.TEXT_PRESENCE = append(QCChamps.TEXT_PRESENCE, TEXT_PRESENCE{" "," "," "," "," "})
	}	
	for i := 0; i < 5; i++ {
		QCChamps.NEEDED_TEXT = append(QCChamps.NEEDED_TEXT, NEEDED_TEXT{" "," "," "," "," "," "," "})
	}	
	for i := 0; i < 5; i++ {
		QCChamps.TEXT_ISSUES = append(QCChamps.TEXT_ISSUES, TEXT_ISSUES{" "," "," "," "," "," "," "})
	}	
	return QCChamps
}

func ByteToMB(b int) int {
	return (b/(1024*1024))
}	

func Scale(BatonScale string) string {
	switch BatonScale {
		case "Fatal":
			return "1"
		case "Warning":
			return "3"
		case "Serious":
			return "2"
		default:
			return "FYI"			
	}
}

func HeadQC(QC *QCInfos, fd Fields) {
	QC.TITLE = fd.Titre
	QC.AUDIO_LANG = lib.Locale(fd.Nom_original)
	QC.PROVIDER = fd.Provider
	QC.VALIDATION_TYPE = fd.Validation_type						
	QC.ASSET = fd.Asset_type
	switch QC.ASSET {
		case "Episodic":
			QC.ASSET_NUM = fd.Saison
			QC.ASSET_PASS = fd.Episode
			QC.NUM01 = "Season #"
			QC.NUM02 = "Episode #"
		case "Trailer":	
			QC.ASSET_NUM = fd.Saison
			QC.ASSET_PASS = fd.Episode
			QC.NUM01 = "Trl #"
			QC.NUM02 = "Pass #"
	} 				
	QC.DATE = fd.Date
	QC.YEAR	= fd.Production_year
	QC.OPERATOR = fd.Qc_operator
	QC.SOURCE_FILENAME = fd.Source_filename
	QC.FILENAME = fd.Nom_original
	QC.TEXT_PRES = fd.Narrative_text_presence_1
	QC.SUB_LANG = fd.Text_language
	QC.SUB_PRES = fd.Burnedin_dialogs_presence	
	QC.GENCOM.GENCOM_VALUE[1] = fd.Commentaires_generaux
	QC.MODIF_FROM_SOURCE = fd.Commentaires_packaging
	QC.RATIO = fd.Aspect_ratio
	QC.SRC_CH = SRCAUDIO(&fd)
}

func (CatInfos CatRest) CatDV() [][]byte {
	CatInfos.Open()
	Clips := CatInfos.Get()
	CatInfos.Delete()

	return Clips
}

func (CatInfos CatRest) Open() {
	_, err :=CatInfos.Client.R().Get(CatInfos.Endpoint+"session?usr=api_admin&pwd=Brute-Fence8-Backboned")
		if err != nil {
			log.Fatal("Erreur de connexion à CatDV.")
		}
}

func (CatInfos CatRest) Get() [][]byte {
	var ClipsBody [][]byte
	for _, Clip := range CatInfos.ClipIDs {
		resp, err := CatInfos.Client.R().Get(CatInfos.Endpoint+"clips/"+Clip)
		if err == nil {
			ClipsBody = append(ClipsBody, resp.Body())
		} else {
			log.Println("Erreur dans la récupération de clips dans CatDV.")
		}
		
	}
	return ClipsBody
}

func (CatInfos CatRest) Delete() {
	CatInfos.Client.R().	
		Delete(CatInfos.Endpoint+"session")
}

func SRCAUDIO(SRC *Fields) []string {
	audioAss := []string{" ",
				SRC.Source_audio_ch_1,
				SRC.Source_audio_ch_2,
				SRC.Source_audio_ch_3,
				SRC.Source_audio_ch_4,
				SRC.Source_audio_ch_5,
				SRC.Source_audio_ch_6,
				SRC.Source_audio_ch_7,
				SRC.Source_audio_ch_8,
				" "," "," "," "," "," ",
				" "," "," "}
	return audioAss
}

func (QC QCInfos) Version() string {
	var mRes = map[string]Resolution{
		"3840x2160": Resolution{
			Nom: "uhd",
			Court: "UHD",
		},
		"1920x1080": Resolution{
			Nom: "1080",
			Court: "HD",
		},
		"1280x720": Resolution{
			Nom: "720",
			Court: "HD",
		},
		"1024x576": Resolution{
			Nom: "pal16",
			Court: "SD",
		},
		"720x480": Resolution{
			Nom: "ntsc16",
			Court: "SD",
		},
		"720x486": Resolution{
			Nom: "ntsc16",
			Court: "SD",
		},
		"720x576": Resolution{
			Nom: "pal16",
			Court: "SD",
		},
		"768x576": Resolution{
			Nom: "pal16",
			Court: "SD",
		},
		"640x480": Resolution{
			Nom: "ntsc16",
			Court: "SD",
		},
		"640x486": Resolution{
			Nom: "ntsc16",
			Court: "SD",
		},
		"853x480": Resolution{
			Nom: "ntsc16",
			Court: "SD",
		},
		"853x486": Resolution{
			Nom: "ntsc16",
			Court: "SD",
		},
		"2048×1080": Resolution{
			Nom: "2k",
			Court: "2K",
		},
		"4096x2160": Resolution{
			Nom: "4k",
			Court: "4K",
		},
	}

	return mRes[QC.RESOLUTION].Court+" "+mRes[QC.RESOLUTION].Nom+" "+QC.RANGE+" "+QC.RATIO+" "+QC.DIF_RATE
}