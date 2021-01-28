package main

import (
	"github.com/resty"	
	"log"
	"encoding/xml"	
	"html/template"
	"defaults"
	"bytes"
	// "fmt"
	"go-wkhtmltopdf"
	// "encoding/xml"
	// "time"	
	// "regexp"		
	"strings"
	// "xlsx"	
	// "os"
	//"strconv"
	// "path"
	//"encoding/json"	
	// "io/ioutil"	
	//"os/exec"
	//"net/http"
	//"reflect"
	//"rapport-qc/lib"


)

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
	Value			int8				`xml:"value,attr"`
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
//--------------------------------------------------------------------------------------------
//var TEST string = "LLL"

type QCInfos struct {
	TITLE					string			`default:" "`
	PROVIDER				string			`default:"AYTRE!"`
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
	EXT						string			`default:" "`
	FILENAME				string			`default:" "`
	ASSET_LANG				string			`default:" "`
	AUDIO_LANG				string			`default:" "`
	SUB_LANG				string			`default:" "`
	TEXT_PRES				string			`default:" "`
	SUB_PRES				string			`default:" "`
	RUN_TIME				string			`default:" "`
	FILESIZE				string			`default:" "`
	BITRATE					string			`default:" "`
	DIF_RATE				string			`default:" "`
	SRC_RATE				string			`default:" "`
	RESOLUTION				string			`default:" "`
	RATIO					string			`default:" "`
	PXLRATIO				string			`default:" "`
	CODEC					string			`default:" "`
	RANGE					string			`default:" "`
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


func main() {
//00000176e33b7de348c2b333000a0063008a0046
//0000016e3d423cb80b853502000a0063008a0046
		var Baton TaskReport
		QC := &QCInfos{}
		if err := defaults.Set(QC); err != nil {
			panic(err)
		}
		log.Println(QC)

		QC = Init(QC)


		xml.Unmarshal(RestCall("0000016e3d423cb80b853502000a0063008a0046"), &Baton)
	
		log.Println(Baton)

		
		Template := template.Must(template.ParseFiles("C:/Users/daudels/go/src/template_test/qc_report_for_test.tmpl"))
		//fmt.Println(Template.Execute(os.Stdout, QC))
		
		var tpl bytes.Buffer
		if err := Template.Execute(&tpl, QC); err != nil {
			log.Println(err)
		}
		RapportFinal := tpl.String()

		PDF(RapportFinal)
		log.Println("PDF is OUT!")
		//lib.WriteHTML()

		//log.Println(TEST)
}

func RestCall(BatonTaskId string) []byte {
	client := resty.New()
	resp, _ := client.R().
		SetBasicAuth("dev_user", "jhg45W&sd_18").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").		
		Get("http://baton-2:8080/Baton/api/tasks/"+BatonTaskId+"/report?type=xml")
		
		//fmt.Println(resp.String())
		//fmt.Println(reflect.TypeOf(resp.String()))
	return []byte(resp.String())

}

func PDF(HTMLRapport string) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
	  log.Fatal(err)
	}

	pdfg.OutputFile = "C:/Users/daudels/go/src/template_test/PDF_JEUDI.pdf"
	
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(HTMLRapport)))

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
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