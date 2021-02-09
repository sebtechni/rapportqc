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

	//"encoding/json"	
	// "io/ioutil"	
	//"os/exec"
	//"net/http"
	// "reflect"
	//"rapport-qc/lib"
	// "encoding/xml"
	// "time"	
	// "fmt"

)

var (    
	Build    string = ""
)

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
	MODIF_FROM_SOURCE		string			`default:" "`
	HEAD_LOGO				HEAD_LOGO
	TAIL_LOGO				TAIL_LOGO
	DIF_TC					DIF_TC
	SRC_TC					SRC_TC
	DIF_CH					[]string		`default:"[\" \", \" \", \" \",\" \", \" \", \" \", \" \", \" \", \" \"]"`
	SRC_CH					[]string		`default:"[\" \", \" \", \" \",\" \", \" \", \" \", \" \", \" \", \" \", \" \", \" \",\" \", \" \", \" \", \" \", \" \", \" \", \" \"]"`
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
					CatInfos := CatRest{Client: resty.New(), ClipIDs: strings.Split(os.Args[1], ","), Endpoint: "http://10.99.139.220:8080/api/9/"}

					CatInfos.Open()
					Clips := CatInfos.Get()
					CatInfos.Delete()

					for _, Clip := range Clips {
						var Baton TaskReport
						QC := &QCInfos{}
						if err := defaults.Set(QC); err != nil {
							panic(err)
						}
						QC = Init(QC)

						var InfosCatDV CatDV
						err := json.Unmarshal(Clip, &InfosCatDV)
						if err != nil {
							log.Println("JSON error:", err)
						}
						
						QC.TITLE = InfosCatDV.Data.Fields.Titre
						QC.AUDIO_LANG = lib.Locale(InfosCatDV.Data.Fields.Nom_original)
						QC.PROVIDER = InfosCatDV.Data.Fields.Provider
						QC.VALIDATION_TYPE = InfosCatDV.Data.Fields.Validation_type
						QC.ASSET_NUM = InfosCatDV.Data.Fields.Saison
						QC.ASSET_PASS = InfosCatDV.Data.Fields.Episode
						QC.ASSET = InfosCatDV.Data.Fields.Asset_type
						QC.DATE = InfosCatDV.Data.Fields.Date
						QC.YEAR	= InfosCatDV.Data.Fields.Production_year
						QC.OPERATOR = InfosCatDV.Data.Fields.Qc_operator
						QC.SOURCE_FILENAME = InfosCatDV.Data.Fields.Source_filename
						QC.FILENAME = InfosCatDV.Data.Fields.Nom_original
						QC.TEXT_PRES = InfosCatDV.Data.Fields.Narrative_text_presence_1
						QC.SUB_LANG = InfosCatDV.Data.Fields.Text_language
						QC.SUB_PRES = InfosCatDV.Data.Fields.Burnedin_dialogs_presence	
						QC.GENCOM.GENCOM_VALUE[1] = InfosCatDV.Data.Fields.Commentaires_generaux
						QC.MODIF_FROM_SOURCE = InfosCatDV.Data.Fields.Commentaires_packaging
						QC.RATIO = InfosCatDV.Data.Fields.Aspect_ratio


						xml.Unmarshal(BatonRestCall(InfosCatDV.Data.Fields.Baton_taskid), &Baton)	
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
							if Field.Name == "Frame Rate" {			
								QC.DIF_RATE = Field.Value
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


						//QC.RUN_TIME = Baton.
						//log.Println(Baton)
						var Template *template.Template
						ex, _ := os.Executable()
						if _, err := os.Stat(filepath.Dir(ex)+"/qc_report.tmpl"); err == nil {
							Template = template.Must(template.ParseFiles(filepath.Dir(ex)+"/qc_report.tmpl"))
						} else if os.IsNotExist(err) {
							Template = template.Must(template.ParseFiles("C:/Users/daudels/go/src/rapport-qc/qc_report.tmpl"))
						}
						
						var tpl bytes.Buffer
						if err := Template.Execute(&tpl, QC); err != nil {
							log.Println(err)
						}
						RapportFinal := tpl.String()

						PDF(RapportFinal, CleanChar.ReplaceAllString(QC.FILENAME[:len(QC.FILENAME)-4],"_"))			
						//lib.WriteHTML()

						//log.Println(TEST)
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
		
		//fmt.Println(resp.String())
		//fmt.Println(reflect.TypeOf(resp.String()))
	return []byte(resp.String())
}

func PDF(HTMLRapport, Filename string) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
	  log.Fatal(err)
	}	
	
	if _, err = os.Stat("C:/Users/daudels/go/src/template_test/"); err == nil {
		pdfg.OutputFile = "C:/Users/daudels/go/src/template_test/"+Filename+"_qc_report.pdf"
	} else if os.IsNotExist(err) {		
		pdfg.OutputFile = "/data/"+Filename+"_qc_report.pdf"
	}
	
	page := wkhtmltopdf.NewPageReader(strings.NewReader(HTMLRapport))
	page.Zoom.Set(1.05)
	pdfg.AddPage(page)
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Le rapport QC en PDF a été généré : ",pdfg.OutputFile)
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

