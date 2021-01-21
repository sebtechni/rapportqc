package lib

import (
	// "strings"
	// "encoding/xml"
	//  "io/ioutil"		
	// "fmt"
	// "xlsx"	
	// "path"
	// "regexp"
	"os"
)



func WriteHTML() {

// TEST = "UNPROVIDER"

HtmlCode := `<html xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns="http://www.w3.org/TR/REC-html40"><head><meta http-equiv="Content-Type" content="text/html; charset=windows-1252">

<meta name="ProgId" content="Excel.Sheet">
<meta name="Generator" content="Microsoft Excel 15">
<!--[if !mso]>
<style>
v\:* {behavior:url(#default#VML);}
o\:* {behavior:url(#default#VML);}
x\:* {behavior:url(#default#VML);}
.shape {behavior:url(#default#VML);}
</style>
<![endif]-->
<style id="qc_report_14491_Styles">
<!--table
	{mso-displayed-decimal-separator:"\.";
	mso-displayed-thousand-separator:"\,";}
.font514491
	{color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;}
.font614491
	{color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;}
.xl8414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl8514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl8614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl8714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl8814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl8914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl9014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:none;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl9114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl9214491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:middle;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl9314491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:none;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl9414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl9514491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:middle;
	background:#404040;
	mso-pattern:black none;
	white-space:normal;}
.xl9614491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl9714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	border-top:none;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl9814491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	border-top:none;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl9914491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"Short Date";
	text-align:general;
	vertical-align:bottom;
	border-top:none;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl10014491
	{padding-left:35px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	border-top:none;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl10114491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:7.5pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	border-top:none;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl10214491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:7.5pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:bottom;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl10314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	background:white;
	mso-pattern:#BFBFBF none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl10414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:middle;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl10514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl10614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl10714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:bottom;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl10814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:bottom;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl10914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl11014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl11114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl11214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:general;
	vertical-align:bottom;
	border-top:none;
	border-right:.5pt solid #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl11314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl11414491
	{padding:0px;
	border: 1px solid lightgray!important;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl11514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl11614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:9.0pt;
	font-weight:700;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:top;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl11714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:9.0pt;
	font-weight:700;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:top;
	border:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl11814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl11914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl12014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:bottom;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl12114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl12214491
	{padding:0px;
	mso-ignore:padding;
	border: 1px solid lightgray!important;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl12314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl12414491
	{padding:0px;
	border: 1px solid lightgray!important;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl12514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl12614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	border: 1px solid lightgray!important;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl126144912
	{padding-left:5px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}	
.xl12714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl12814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl12914491
	{padding:0px;
	mso-ignore:padding;
	border: 1px solid lightgray!important;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl13014491
	{padding:0px;
	border: 1px solid lightgray!important;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl13114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl13214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl13314491
	{padding:0px;
	mso-ignore:padding;
	border: 1px solid lightgray!important;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl13414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl13514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl13614491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl13714491
	{padding:0px;
	border: 1px solid lightgray!important;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl13814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl13914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #BFBFBF;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl14014491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl14114491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt solid #404040;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl14214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	border: 1px solid lightgray!important;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl14314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl14414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl14514491
	{padding:0px;
	border: 1px solid lightgray!important;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl14614491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl14714491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl14814491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl14914491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl15014491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl15114491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl15214491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl15314491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt solid #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl15414491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt solid #404040;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl15514491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt solid #404040;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl15614491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl15714491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl15814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl15914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl16014491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl16114491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl16214491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl16314491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl16414491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl16514491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl16614491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl16714491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl16814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl16914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl17014491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl17114491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl17214491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl17314491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl17414491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl17514491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl17614491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl17714491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl17814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl17914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt solid #404040;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl18014491
	{padding:0px;
	border: 1px solid lightgray!important;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl18114491
	{padding:0px;
	border: 1px solid lightgray!important;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl18214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl18314491
	{padding:0px;
	mso-ignore:padding;
	border: 1px solid lightgray!important;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl18414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl18514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl18614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl18714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl18814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl18914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl19014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl19114491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl19214491
	{padding:0px;
	mso-ignore:padding;
	border: 1px solid lightgray!important;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl19314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl19414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl19514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl19614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl19714491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:none;
	border-left:.5pt solid #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl19814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:none;
	border-left:.5pt solid #404040;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl19914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:.5pt solid #404040;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl20014491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl20114491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl20214491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl20314491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl20414491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl20514491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl20614491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl20714491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl20814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl20914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl21014491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl21114491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl21214491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl21314491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl21414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt solid #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl21514491
	{padding:0px;
	border: 1px solid lightgray!important;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl215144912
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}	
.xl21614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl21714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl21814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:.5pt solid #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl21914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:none;
	border-left:.5pt solid #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl22014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt solid #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl22114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl22214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl22314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl22414491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:.5pt solid #404040;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl22514491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl22614491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:none;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl22714491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl22814491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl22914491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl23014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl23114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl23214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl23314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl23414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl23514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl23614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl23714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl23814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl23914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl24014491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl24114491
	{padding:0px;
	border: 1px solid lightgray!important;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl24214491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:9.0pt;
	font-weight:700;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl24314491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl24414491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl24514491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl24614491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl24714491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl24814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl24914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl25014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:#BFBFBF;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl25114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl25214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl25314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl25414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:text-top;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:normal;}
.xl25514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:normal;}
.xl25614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:normal;}
.xl25714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:normal;}
.xl25814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:normal;}
.xl25914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:normal;}
.xl26014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:normal;}
.xl26114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:normal;}
.xl26214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:normal;}
.xl26314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl26414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl26514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl26614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl26714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl26814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl26914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:#BFBFBF thin-diag-stripe;
	white-space:nowrap;}
.xl27014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:#BFBFBF thin-diag-stripe;
	white-space:nowrap;}
.xl27114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:#BFBFBF thin-diag-stripe;
	white-space:nowrap;}
.xl27214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:#BFBFBF thin-diag-stripe;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl27314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:#BFBFBF thin-diag-stripe;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl27414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:#BFBFBF thin-diag-stripe;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl27514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl27614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl27714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl27814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl27914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl28014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl28114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl28214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl28314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl28414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl28514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl28614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl28714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl28814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl28914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl29014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl29114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl29214491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid windowtext;
	border-left:.5pt solid windowtext;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl29314491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid windowtext;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl29414491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid windowtext;
	border-bottom:.5pt solid windowtext;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl29514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl29614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl29714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl29814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl29914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl30014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt solid #404040;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl30114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl30214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl30314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl30414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl30514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl30614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl30714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl30814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl30914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl31014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl31114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl31214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl31314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl31414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl31514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl31614491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl31714491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl31814491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	border-top:.5pt solid #404040;
	border-right:.5pt solid windowtext;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl31914491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid windowtext;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl32014491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl32114491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl32214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl32314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl32414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl32514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl32614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl32714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl32814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:none;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl32914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid windowtext;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl33014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl33114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl33214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:#BFBFBF thin-diag-stripe;
	white-space:nowrap;}
.xl33314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\;\;\;";
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:#BFBFBF thin-diag-stripe;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl33414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl33514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl33614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl33714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl33814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl33914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:#BFBFBF thin-diag-stripe;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl34014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl34114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl34214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl34314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl34414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:7.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl34514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl34614491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid windowtext;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl34714491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid windowtext;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid windowtext;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl34814491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid windowtext;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl34914491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:10.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid windowtext;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl35014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl35114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl35214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl35314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl35414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl35514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"00\\\:00\\\:00\\\:00";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl35614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl35714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl35814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl35914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl36014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl36114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl36214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl36314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl36414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl36514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl36614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl36714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl36814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:black none;
	white-space:nowrap;}
.xl36914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:black none;
	white-space:nowrap;}
.xl37014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl37114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl37214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl37314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl37414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl37514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl37614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl37714491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl37814491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl37914491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl38014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl38114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl38214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl38314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl38414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl38514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl38614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl38714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl38814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl38914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl39014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl39114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl39214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl39314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl39414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl39514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl39614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl39714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl39814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl39914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl40014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl40114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl40214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl40314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl40414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl40514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl40614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl40714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl40814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl40914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl41014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl41114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl41214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl41314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl41414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl41514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:00;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl41614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:00;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl41714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:00;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl41814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl41914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl42014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl42114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:00;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl42214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:00;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl42314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:00;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl42414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl42514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl42614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl42714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl42814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:italic;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:normal;}
.xl42914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl43014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl43114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:none;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl43214491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:.5pt solid #404040;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl43314491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl43414491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:none;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl43514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"\@";
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl43614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl43714491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl43814491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	white-space:nowrap;}
.xl43914491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl44014491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl44114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl44214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl44314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl44414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl44514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl44614491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl44714491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl44814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl44914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"Short Date";
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl45014491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl45114491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl45214491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:"Short Date";
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt hairline #BFBFBF;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl45314491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl45414491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:8.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:left;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt hairline #BFBFBF;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	white-space:nowrap;}
.xl45514491
	{padding:0px;
	mso-ignore:padding;
	color:yellow;
	font-size:12.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:"Lucida Handwriting", cursive;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl45614491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:12.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl45714491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:12.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:none;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl45814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl45914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl46014491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt hairline #BFBFBF;
	border-right:.5pt solid #404040;
	border-bottom:.5pt solid #404040;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl46114491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:.5pt solid #404040;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl46214491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:none;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl46314491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	border-top:none;
	border-right:none;
	border-bottom:none;
	border-left:.5pt solid #404040;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl46414491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl46514491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	border-top:none;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:.5pt solid #404040;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl46614491
	{padding:0px;
	mso-ignore:padding;
	color:black;
	font-size:11.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:bottom;
	border-top:none;
	border-right:none;
	border-bottom:.5pt solid #404040;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl46714491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:.5pt solid #404040;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl46814491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:none;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl46914491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:9.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:.5pt solid #404040;
	border-right:.5pt solid #404040;
	border-bottom:.5pt hairline #BFBFBF;
	border-left:none;
	background:white;
	mso-pattern:black none;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl47014491
	{padding-right:5px;
	mso-ignore:padding;
	color:white;
	font-size:12.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl47114491
	{padding:0px;
	mso-ignore:padding;
	color:white;
	font-size:12.0pt;
	font-weight:400;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:right;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:none;
	border-left:none;
	background:#404040;
	mso-pattern:black none;
	white-space:nowrap;}
.xl47214491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:none;
	border-bottom:none;
	border-left:.5pt solid #404040;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl47314491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
.xl47414491
	{padding:0px;
	mso-ignore:padding;
	color:#404040;
	font-size:11.0pt;
	font-weight:700;
	font-style:normal;
	text-decoration:none;
	font-family:Calibri, sans-serif;
	mso-font-charset:0;
	mso-number-format:General;
	text-align:center;
	vertical-align:middle;
	border-top:none;
	border-right:.5pt solid #404040;
	border-bottom:none;
	border-left:none;
	mso-background-source:auto;
	mso-pattern:auto;
	mso-protection:unlocked visible;
	white-space:nowrap;}
-->
</style>
</head>

<body>

<div id="qc_report_14491" align="center" x:publishsource="Excel">

<table border="0" cellpadding="0" cellspacing="0" width="756" class="xl8514491" style="border-collapse:collapse;table-layout:fixed;width:578pt">
 <colgroup><col class="xl8814491" width="21" style="mso-width-source:userset;mso-width-alt:
 739;width:16pt">
 <col class="xl8514491" width="17" span="42" style="mso-width-source:userset;
 mso-width-alt:597;width:13pt">
 <col class="xl8514491" width="21" style="mso-width-source:userset;mso-width-alt:
 739;width:16pt">
 </colgroup><tbody><tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" width="21" style="height:14.4pt;width:16pt">&nbsp;</td>
  <td colspan="14" rowspan="5" height="98" width="238" style="border-bottom:.5pt solid #404040;
  height:73.8pt;width:182pt" align="left" valign="top">
  <span style="mso-ignore:vglayout;
  position:absolute;z-index:1;margin-left:14px;margin-top:24px;width:238px;
  height:59px"><img width="238" height="59" src="./template_qc_report_files/difuze_logo.png" v:shapes="Picture_x0020_1"></span><!--[endif]--><span style="mso-ignore:vglayout2">
  <table cellpadding="0" cellspacing="0">
   <tbody><tr>
    <td colspan="14" rowspan="5" height="98" class="xl46114491" width="238" style="border-bottom:.5pt solid #404040;height:73.8pt;width:190pt">&nbsp;</td>
   </tr>
  </tbody></table>
  </span></td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8914491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl9014491" width="17" style="width:13pt">&nbsp;</td>
  <td class="xl8514491" width="21" style="width:16pt">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9214491">&nbsp;</td>
  <td class="xl9214491">&nbsp;</td>
  <td class="xl9214491">&nbsp;</td>
  <td class="xl9214491">&nbsp;</td>
  <td class="xl9314491">&nbsp;</td>
  <td colspan="9" class="xl46714491" style="border-right:.5pt solid #404040;
  border-left:none">Domestic QC</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9414491">&nbsp;</td>
  <td class="xl9514491" width="17" style="width:13pt">&nbsp;</td>
  <td colspan="5" class="xl47014491" style="border-right:.5pt solid #404040">Provider
  : </td>
  <td colspan="9" class="xl47214491" style="border-right:.5pt solid #404040;
  border-left:none">`+TEST+`</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td class="xl9114491">&nbsp;</td>
  <td class="xl9614491">&nbsp;</td>
  <td colspan="10" class="xl45514491" style="font-size:12.0pt;color:yellow;
  font-weight:700;text-decoration:none;text-underline-style:none;text-line-through:
  none;font-family:&quot;Lucida Handwriting&quot;, cursive;border:none;background:#404040;
  mso-pattern:black none"></td>
  <td colspan="7" class="xl45614491" style="border-right:.5pt solid #404040">Validation
  Type :</td>
  <td colspan="9" class="xl45814491" style="border-right:.5pt solid #404040;
  border-left:none">$%VALIDATION_TYPE%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9814491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9914491">&nbsp;</td>
  <td class="xl10014491">&nbsp;</td>
  <td class="xl10114491">&nbsp;</td>
  <td class="xl10014491">&nbsp;</td>
  <td class="xl10014491">File Verification Report 5.0.0.0</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491"></td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491"></td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl9714491">&nbsp;</td>
  <td class="xl10214491"></td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl10314491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="5" class="xl26314491" style="border-right:.5pt hairline #BFBFBF">Title:</td>
  <td colspan="23" class="xl44314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%TITLE%$</td>
  <td colspan="5" class="xl44614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Version:<span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="9" class="xl45214491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%VERSION%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="5" class="xl26314491" style="border-right:.5pt hairline #BFBFBF">Localized
  Title:</td>
  <td colspan="23" class="xl44314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="5" class="xl44614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Date:</td>
  <td colspan="9" class="xl44914491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%DATE%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="5" class="xl26314491" style="border-right:.5pt hairline #BFBFBF">Director:</td>
  <td colspan="9" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="5" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Production Year:</td>
  <td colspan="9" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%YEAR%$</td>
  <td colspan="5" class="xl43914491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">QC Operators:</td>
  <td colspan="9" class="xl44014491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%OPERATOR%$<span style="mso-spacerun:yes">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span></td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10414491" style="height:14.4pt">&nbsp;</td>
  <td colspan="5" class="xl26314491" style="border-right:.5pt hairline #BFBFBF">PO
  #:</td>
  <td colspan="9" class="xl27714491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">N/A</td>
  <td colspan="5" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Coordinator:</td>
  <td colspan="9" class="xl21614491" style="border-right:.5pt hairline #BFBFBF">N/A</td>
  <td colspan="5" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Facility:</td>
  <td colspan="9" class="xl43614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Montréal</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="7" class="xl42414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Source Filename:</td>
  <td colspan="35" class="xl42714491" width="595" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:455pt">$%SOURCE_FILENAME%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td colspan="5" class="xl42914491">&nbsp;</td>
  <td class="xl10614491" style="border-top:none">&nbsp;</td>
  <td class="xl10614491" style="border-top:none">&nbsp;</td>
  <td class="xl10614491" style="border-top:none">&nbsp;</td>
  <td class="xl10614491" style="border-top:none">&nbsp;</td>
  <td class="xl10614491" style="border-top:none">&nbsp;</td>
  <td class="xl10614491" style="border-top:none">&nbsp;</td>
  <td class="xl10614491" style="border-top:none">&nbsp;</td>
  <td class="xl10614491" style="border-top:none">&nbsp;</td>
  <td class="xl10614491" style="border-top:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" class="xl43214491" style="border-right:.5pt solid #404040">PRIMARY
  ASSET:</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl40614491" style="border-right:.5pt hairline #BFBFBF">Primary
  Asset:</td>
  <td colspan="7" class="xl40914491" width="119" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:91pt">$%ASSET%$</td>
  <td colspan="7" class="xl41214491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Trl #</td>
  <td colspan="7" class="xl41514491" width="119" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:91pt">$%ASSET_NUM%$</td>
  <td colspan="7" class="xl41814491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Pass #</td>
  <td colspan="7" class="xl42114491" width="119" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:91pt">$%ASSET_PASS%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl38914491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"></td>
  <td colspan="18" class="xl39114491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:none;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;background:white;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="3" class="xl39414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">File Date:</td>
  <td colspan="4" class="xl39714491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%FILE_DATE%$</td>
  <td colspan="3" class="xl40014491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Revision:<span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="3" class="xl40314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%REV%$</td>
  <td colspan="2" class="xl39414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Ext:</td>
  <td colspan="2" class="xl40314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%EXT%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl10414491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl38014491" style="border-right:.5pt hairline #BFBFBF">Notes:</td>
  <td colspan="7" class="xl38314491" width="119" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:91pt">&nbsp;</td>
  <td colspan="7" class="xl38614491" width="119" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:91pt;font-size:8.0pt;color:black;font-weight:400;
  text-decoration:none;text-underline-style:none;text-line-through:none;
  font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;border-right:
  none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="7" class="xl38614491" width="119" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:91pt;font-size:8.0pt;color:black;font-weight:400;
  text-decoration:none;text-underline-style:none;text-line-through:none;
  font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;border-right:
  none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="7" class="xl38614491" width="119" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:91pt;font-size:8.0pt;color:black;font-weight:400;
  text-decoration:none;text-underline-style:none;text-line-through:none;
  font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;border-right:
  none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="7" class="xl38614491" width="119" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:91pt;font-size:8.0pt;color:black;font-weight:400;
  text-decoration:none;text-underline-style:none;text-line-through:none;
  font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;border-right:
  none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl10514491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl26314491" style="border-left:none">Difuze&#7475;&#7484;
  Filename:</td>
  <td colspan="35" class="xl126144912" width="595" style="border-right:.5pt hairline #BFBFBF;
  width:455pt">$%FILENAME%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl12014491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl10714491" style="border-top:none">&nbsp;</td>
  <td class="xl12014491">&nbsp;</td>
  <td class="xl12014491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl12014491">&nbsp;</td>
  <td class="xl12014491">&nbsp;</td>
  <td class="xl10814491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" class="xl37714491" style="border-right:.5pt solid #404040">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl37014491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:none;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">Additional
  Asset:</td>
  <td colspan="5" class="xl37214491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  background:white;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="3" class="xl36414491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="4" class="xl36214491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="2" class="xl36414491" style="font-size:8.0pt;color:black;font-weight:
  400;text-decoration:none;text-underline-style:none;text-line-through:none;
  font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;border-right:
  none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="2" class="xl36614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:700;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="2" class="xl36814491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:400;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="2" class="xl36614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:700;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl35614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Filename:</td>
  <td colspan="35" class="xl35914491" width="595" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:455pt;font-size:8.0pt;color:black;font-style:italic;
  font-weight:400;text-decoration:none;text-underline-style:none;text-line-through:
  none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  background:white;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl37014491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:none;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">Additional
  Asset:</td>
  <td colspan="5" class="xl37214491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  background:white;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="3" class="xl36414491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="4" class="xl36214491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="2" class="xl36414491" style="font-size:8.0pt;color:black;font-weight:
  400;text-decoration:none;text-underline-style:none;text-line-through:none;
  font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;border-right:
  none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="2" class="xl36614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:700;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="2" class="xl36814491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:400;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="2" class="xl36614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:700;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl35614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Filename:</td>
  <td colspan="35" class="xl35914491" width="595" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:455pt;font-size:8.0pt;color:black;font-style:italic;
  font-weight:400;text-decoration:none;text-underline-style:none;text-line-through:
  none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  background:white;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl37014491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:none;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">Additional
  Asset:</td>
  <td colspan="5" class="xl37214491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  background:white;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="3" class="xl36414491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="4" class="xl36214491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="2" class="xl36414491" style="font-size:8.0pt;color:black;font-weight:
  400;text-decoration:none;text-underline-style:none;text-line-through:none;
  font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;border-right:
  none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="2" class="xl36614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:700;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="2" class="xl36814491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:400;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="2" class="xl36614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:700;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl35614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Filename:</td>
  <td colspan="35" class="xl35914491" width="595" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:455pt;font-size:8.0pt;color:black;font-style:italic;
  font-weight:400;text-decoration:none;text-underline-style:none;text-line-through:
  none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  background:white;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl37014491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:none;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">Additional
  Asset:</td>
  <td colspan="5" class="xl37214491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  background:white;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="5" class="xl37414491" width="85" style="width:65pt;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="3" class="xl36414491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="4" class="xl36214491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:700;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="2" class="xl36414491" style="font-size:8.0pt;color:black;font-weight:
  400;text-decoration:none;text-underline-style:none;text-line-through:none;
  font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;border-right:
  none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="2" class="xl36614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:700;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="2" class="xl36814491" style="border-left:none;font-size:8.0pt;
  color:black;font-weight:400;text-decoration:none;text-underline-style:none;
  text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe"><span style="mso-spacerun:yes">&nbsp;</span></td>
  <td colspan="2" class="xl36614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:700;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="7" class="xl35614491" style="border-right:.5pt hairline #BFBFBF;
  font-size:8.0pt;color:black;font-weight:400;text-decoration:none;text-underline-style:
  none;text-line-through:none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  mso-background-source:auto;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Filename:</td>
  <td colspan="35" class="xl35914491" width="595" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:455pt;font-size:8.0pt;color:black;font-style:italic;
  font-weight:400;text-decoration:none;text-underline-style:none;text-line-through:
  none;font-family:Calibri, sans-serif;border-top:.5pt hairline #BFBFBF;
  border-right:none;border-bottom:.5pt hairline #BFBFBF;border-left:.5pt hairline #BFBFBF;
  background:white;mso-pattern:#BFBFBF thin-diag-stripe">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10914491" style="height:14.4pt">&nbsp;</td>
  <td colspan="14" class="xl34614491" style="border-left:none">PRIMARY ASSET</td>
  <td colspan="14" class="xl31914491" style="border-right:.5pt solid black;
  border-left:none">DIFUZE&#7475;&#7484; TIMECODES</td>
  <td colspan="14" class="xl34714491" style="border-right:.5pt solid #404040;
  border-left:none">SOURCE TIMECODES</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="7" class="xl32214491">Asset Language:</td>
  <td colspan="7" class="xl34414491" style="border-left:none">$%ASSET_LANG%$</td>
  <td colspan="7" class="xl35014491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Starting TC:</td>
  <td colspan="7" class="xl35314491" style="border-left:none">$%DIF_START_TC%$</td>
  <td colspan="7" class="xl35414491" style="border-left:none">Starting TC:</td>
  <td colspan="7" class="xl35514491" style="border-left:none">$%SRC_START_TC%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="7" class="xl26514491">Audio Language:</td>
  <td colspan="7" class="xl34414491" style="border-left:none">$%AUDIO_LANG%$</td>
  <td colspan="7" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Program In:</td>
  <td colspan="7" class="xl34514491" style="border-left:none">$%DIF_PROGIN_TC%$</td>
  <td colspan="7" class="xl33814491" style="border-left:none">Program In:</td>
  <td colspan="7" class="xl33714491" style="border-left:none">$%SRC_PROGIN_TC%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="7" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Text and Sub Language:</td>
  <td colspan="7" class="xl34414491" style="border-left:none">$%SUB_LANG%$</td>
  <td colspan="7" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Main Title:</td>
  <td colspan="7" class="xl18814491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%DIF_TITLE_TC%$</td>
  <td colspan="7" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Main Title:</td>
  <td colspan="7" class="xl18814491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%SRC_TITLE_TC%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="7" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Text Presence:</td>
  <td colspan="7" class="xl34114491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%TEXT_PRES%$</td>
  <td colspan="7" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Main Credits:</td>
  <td colspan="7" class="xl18814491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%DIF_MAINCRED_TC%$</td>
  <td colspan="7" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Main Credits:</td>
  <td colspan="7" class="xl18814491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%SRC_MAINCRED_TC%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="7" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Subtitle Presence:</td>
  <td colspan="7" class="xl34014491" style="border-left:none">$%SUB_PRES%$</td>
  <td colspan="7" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">End Credits:</td>
  <td colspan="7" class="xl18814491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%DIF_ENDCRED_TC%$</td>
  <td colspan="7" class="xl33814491" style="border-left:none">End Credits:</td>
  <td colspan="7" class="xl33714491" style="border-left:none">$%SRC_ENDCRED_TC%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="7" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Running Time:</td>
  <td colspan="7" class="xl33714491" style="border-left:none">$%RUN_TIME%$</td>
  <td colspan="7" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Program Out:</td>
  <td colspan="7" class="xl33714491" style="border-left:none">$%DIF_PROGOUT_TC%$</td>
  <td colspan="7" class="xl33814491" style="border-left:none">Program Out:</td>
  <td colspan="7" class="xl33714491" style="border-left:none">$%SRC_PROGOUT_TC%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="7" class="xl33214491">&nbsp;</td>
  <td colspan="7" class="xl33914491" style="border-left:none">&nbsp;</td>
  <td colspan="7" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Textless In:</td>
  <td colspan="7" class="xl33714491" style="border-left:none">$%DIF_TXTLIN_TC%$</td>
  <td colspan="7" class="xl33814491" style="border-left:none">Textless In:</td>
  <td colspan="7" class="xl33714491" style="border-left:none">$%SRC_TXTLIN_TC%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="7" class="xl33214491">&nbsp;</td>
  <td colspan="7" class="xl33314491" style="border-left:none">&nbsp;</td>
  <td colspan="7" class="xl33414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Textless Out:</td>
  <td colspan="7" class="xl33714491" style="border-left:none">$%DIF_TXTLOUT_TC%$</td>
  <td colspan="7" class="xl33814491" style="border-left:none">Textless Out:</td>
  <td colspan="7" class="xl33714491" style="border-left:none">$%SRC_TXTLOUT_TC%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td colspan="2" class="xl31514491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl11114491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="12" class="xl31614491" style="border-right:.5pt solid black">VIDEO
  &amp; AUDIO SPECIFICATIONS</td>
  <td colspan="12" class="xl29214491" style="border-right:.5pt solid black;
  border-left:none;font-size:10.0pt;color:white;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt solid #404040;border-right:none;border-bottom:.5pt solid windowtext;
  border-left:.5pt solid windowtext;background:#404040;mso-pattern:black none">DIFUZE&#7475;&#7484;
  MEZZ AUDIO CONFIG.</td>
  <td colspan="18" class="xl31914491" style="border-right:.5pt solid #404040;
  border-left:none">SOURCE MEZZ AUDIO CONFIGURATION</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl32214491">File Size:</td>
  <td colspan="3" class="xl31014491" style="border-left:none">$%FILESIZE%$</td>
  <td colspan="3" class="xl31014491" style="border-left:none">MB</td>
  <td colspan="5" class="xl32414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt solid windowtext;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;background:white;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Ch.1:</td>
  <td colspan="7" class="xl32714491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt solid windowtext;border-right:none;border-bottom:none;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe">$%DIF_CH1%$</td>
  <td colspan="2" class="xl33014491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Ch.1</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH1%$</td>
  <td colspan="2" class="xl31014491" style="border-left:none">Ch.18</td>
  <td colspan="7" class="xl30914491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26514491">Bit Rate:</td>
  <td colspan="3" class="xl18614491" style="border-left:none">$%BITRATE%$</td>
  <td colspan="3" class="xl18614491" style="border-left:none">Mbit/s</td>
  <td colspan="5" class="xl29514491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;background:white;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Ch.2:</td>
  <td colspan="7" class="xl31214491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:none;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe">$%DIF_CH2%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.2</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH2%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.19</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26414491" style="border-right:.5pt hairline #BFBFBF">Mezz
  Frame Rate:</td>
  <td colspan="6" class="xl215144912" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%DIF_RATE%$</td>
  <td colspan="5" class="xl295144912" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;padding-left:11px;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;background:white;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Ch.3:</td>
  <td colspan="7" class="xl301144912" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe">$%DIF_CH3%$</td>
  <td colspan="2" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Ch.3</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH3%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.20</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26414491" style="border-right:.5pt hairline #BFBFBF">Source
  Frame Rate:</td>
  <td colspan="6" class="xl215144912" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%SRC_RATE%$</td>
  <td colspan="5" class="xl29514491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;background:white;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Ch.4:</td>
  <td colspan="7" class="xl30114491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe">$%DIF_CH4%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.4</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH4%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.21</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26414491" style="border-right:.5pt hairline #BFBFBF">Resolution:</td>
  <td colspan="6" class="xl215144912" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%RESOLUTION%$</td>
  <td colspan="5" class="xl29514491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;background:white;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Ch.5:</td>
  <td colspan="7" class="xl30114491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe">$%DIF_CH5%$</td>
  <td colspan="2" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Ch.5</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH5%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.22</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Aspect Ratio:</td>
  <td colspan="6" class="xl30414491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%RATIO%$</td>
  <td colspan="5" class="xl29514491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;background:white;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Ch.6:</td>
  <td colspan="7" class="xl30114491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe">$%DIF_CH6%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.6</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH6%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.23</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Pixel Aspect Ratio:</td>
  <td colspan="6" class="xl215144912" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%PXLRATIO%$</td>
  <td colspan="5" class="xl29514491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;background:white;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Ch.7:</td>
  <td colspan="7" class="xl30114491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe">$%DIF_CH7%$</td>
  <td colspan="2" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Ch.7</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH7%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.24</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Video Levels:</td>
  <td colspan="6" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Standard</td>
  <td colspan="5" class="xl29514491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;background:white;mso-pattern:#BFBFBF thin-diag-stripe">Difuze&#7475;&#7484;
  Ch.8:</td>
  <td colspan="7" class="xl29814491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:.5pt hairline #BFBFBF;border-right:none;border-bottom:.5pt solid #404040;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe">$%DIF_CH8%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.8</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH8%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.25</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF">Tape
  Tracking #:</td>
  <td colspan="6" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">N/A</td>
  <td colspan="12" class="xl29214491" style="border-right:.5pt solid black">DYNAMIC
  RANGE METADATA</td>
  <td colspan="2" class="xl26714491" style="border-right:.5pt hairline #BFBFBF">Ch.9</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH9%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.26</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Codec:</td>
  <td colspan="6" class="xl215144912" style="border-left:none">$%CODEC%$</td>
  <td colspan="6" class="xl28614491" style="border-right:.5pt hairline #BFBFBF">Dynamic
  Range:</td>
  <td colspan="6" class="xl28914491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%RANGE%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.10</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH10%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.27</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Picture Structure:</td>
  <td colspan="6" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">I-Frames</td>
  <td colspan="6" class="xl28014491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:none;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;background:white;mso-pattern:#BFBFBF thin-diag-stripe"></td>
  <td colspan="6" class="xl28314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:none;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="2" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Ch.11</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH11%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.28</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Color Sampling:</td>
  <td colspan="6" class="xl27714491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%COLORSAMPLE%$</td>
  <td colspan="6" class="xl28014491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:none;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe"></td>
  <td colspan="6" class="xl28314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;font-size:8.0pt;color:black;font-weight:400;text-decoration:
  none;text-underline-style:none;text-line-through:none;font-family:Calibri, sans-serif;
  border-top:none;border-right:none;border-bottom:.5pt hairline #BFBFBF;
  border-left:.5pt hairline #BFBFBF;mso-background-source:auto;mso-pattern:
  #BFBFBF thin-diag-stripe">&nbsp;</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.12</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH12%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.29</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Video Bit Depth:</td>
  <td colspan="6" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%VID_BIT%$</td>
  <td colspan="6" class="xl26914491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="6" class="xl27214491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="2" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Ch.13</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH13%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.30</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Audio Sample Rate:</td>
  <td colspan="6" class="xl215144912" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%AUDIOSAMPLE%$</td>
  <td colspan="6" class="xl26914491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="6" class="xl27214491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.14</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH14%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.31</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Audio Bit Depth:</td>
  <td colspan="6" class="xl215144912" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%AUDIO_BIT%$</td>
  <td colspan="6" class="xl26914491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="6" class="xl27214491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="2" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Ch.15</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH15%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.32</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Audio Levels:</td>
  <td colspan="6" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Standard</td>
  <td colspan="6" class="xl26914491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="6" class="xl27214491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.16</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH16%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.33</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" class="xl26314491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Closed Captioning:</td>
  <td colspan="6" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">No</td>
  <td colspan="6" class="xl26914491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="6" class="xl27214491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="2" class="xl26614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Ch.17</td>
  <td colspan="7" class="xl18714491" style="border-left:none">$%SRC_CH17%$</td>
  <td colspan="2" class="xl18614491" style="border-left:none">Ch.34</td>
  <td colspan="7" class="xl18714491" style="border-left:none">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="30" style="mso-height-source:userset;height:22.5pt">
  <td height="30" class="xl11014491" style="height:22.5pt">&nbsp;</td>
  <td colspan="5" rowspan="4" class="xl25414491" width="85" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:65pt">Modification from source:</td>
  <td colspan="2" class="xl25114491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%V%$:</td>
  <td colspan="35" class="xl25314491" width="595" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:455pt">$%MODIF_FROM_SOURCE%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="2" class="xl25114491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="35" class="xl25314491" width="595" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:455pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="2" class="xl25114491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="35" class="xl25314491" width="595" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:455pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="2" class="xl25114491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">&nbsp;</td>
  <td colspan="35" class="xl25314491" width="595" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:455pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="page-break-before:always;height:120px">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" rowspan="2" class="xl19114491">PRIMARY ASSET<span style="mso-spacerun:yes">&nbsp; </span>NOTES</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8714491" colspan="10">Comments/Problems Severity Code Scale</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8714491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="2" rowspan="2" class="xl24214491">FYI</td>
  <td colspan="20" rowspan="2" class="xl24314491" width="340" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:260pt">For your information:
  Additional information that does not affect the quality, but may have some
  importance to the client.</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="2" rowspan="2" class="xl24214491">1</td>
  <td colspan="20" rowspan="2" class="xl24314491" width="340" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:260pt">Minor: Noticeable, but not a
  rejection.</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td colspan="6" class="xl13614491">Sector Key</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td colspan="2" class="xl25014491">UL</td>
  <td colspan="2" class="xl25014491" style="border-left:none">UC</td>
  <td colspan="2" class="xl25014491" style="border-left:none">UR</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="2" rowspan="2" class="xl24214491">2</td>
  <td colspan="20" rowspan="2" class="xl24314491" width="340" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:260pt">Conditional to client's
  approval: Not enough to be rejected for platform/broadcast deliveries, but
  recommended to fix if possible.</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td colspan="2" class="xl25014491">ML</td>
  <td colspan="2" class="xl25014491" style="border-left:none">MC</td>
  <td colspan="2" class="xl25014491" style="border-left:none">MR</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td colspan="2" class="xl25014491">LL</td>
  <td colspan="2" class="xl25014491" style="border-left:none">LC</td>
  <td colspan="2" class="xl25014491" style="border-left:none">LR</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl8814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="2" rowspan="2" class="xl24214491">3</td>
  <td colspan="20" rowspan="2" class="xl24314491" width="340" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:260pt">Rejected: Major issues that
  should be fixed and would likely be rejected for platform/broadcast
  deliveries.</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>

 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl10414491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" class="xl24914491">Primary Asset General Comments (Please
  itemize specific comments / problems in next section)</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" class="xl12614491" width="714" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:546pt"><font class="font614491">$%GENCOM_CODE_01%$:</font><font class="font514491"> $%GENCOM_VALUE_01%$</font></td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" class="xl12614491" width="714" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:546pt"><font class="font614491">$%GENCOM_CODE_02%$:</font><font class="font514491"> $%GENCOM_VALUE_02%$</font></td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" class="xl12614491" width="714" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:546pt"><font class="font614491">$%GENCOM_CODE_03%$:</font><font class="font514491"> $%GENCOM_VALUE_03%$</font></td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" class="xl12614491" width="714" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:546pt"><font class="font614491">$%GENCOM_CODE_04%$:</font><font class="font514491"> $%GENCOM_VALUE_04%$</font></td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" class="xl12614491" width="714" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:546pt"><font class="font614491">$%GENCOM_CODE_05%$:</font><font class="font514491"> $%GENCOM_VALUE_05%$</font></td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11014491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" class="xl12614491" width="714" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:546pt"><font class="font614491">$%GENCOM_CODE_06%$:</font><font class="font514491"> $%GENCOM_VALUE_06%$</font></td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr class="xl8514491" height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="8" class="xl14614491">&nbsp;</td>
  <td colspan="34" class="xl14814491" style="border-right:.5pt solid #404040">Primary
  Asset Head Logos</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11214491" style="height:15.0pt">&nbsp;</td>
  <td colspan="4" rowspan="2" class="xl15014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:52pt">Difuze&#7475;&#7484; Timecode</td>
  <td colspan="4" rowspan="2" class="xl13814491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:52pt">Source Timecode</td>
  <td colspan="34" rowspan="2" class="xl16214491" width="578" style="border-right:.5pt solid #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:442pt">Description</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11214491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl10514491" style="height:15.0pt">&nbsp;</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%DIF_TC_LOGO%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_LOGO%$</td>
  <td colspan="34" class="xl12614491" width="578" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:442pt">$%LOGO%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr class="xl8514491" height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td colspan="8" class="xl14614491">&nbsp;</td>
  <td colspan="34" class="xl14814491" style="border-right:.5pt solid #404040">Primary
  Asset Issues</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11214491" style="height:15.0pt">&nbsp;</td>
  <td colspan="4" rowspan="2" class="xl15014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:52pt">Difuze&#7475;&#7484; Timecode</td>
  <td colspan="4" rowspan="2" class="xl15614491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:52pt">Source Timecode</td>
  <td colspan="21" rowspan="2" class="xl16214491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:273pt">Description</td>
  <td colspan="3" rowspan="2" class="xl16814491" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF">Duration</td>
  <td rowspan="2" class="xl17214491" style="border-bottom:.5pt hairline #BFBFBF">A</td>
  <td rowspan="2" class="xl17214491" style="border-bottom:.5pt hairline #BFBFBF">V</td>
  <td colspan="4" rowspan="2" class="xl17414491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:52pt">Chan / Sector</td>
  <td colspan="2" rowspan="2" class="xl17614491" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF">Scale</td>
  <td colspan="2" rowspan="2" class="xl13814491" width="34" style="border-right:.5pt solid #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:26pt">Fixed issue</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11214491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="36" style="mso-height-source:userset;height:27.0pt">
  <td height="36" class="xl11314491" style="height:27.0pt">1</td>
  <td colspan="4" class="xl24114491" width="68" style="border-left:none;width:52pt">$%DIF_TC_PRIM_01%$</td>
  <td colspan="4" class="xl13014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%SRC_TC_PRIM_01%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt"><font class="font514491">$%DESC_PRIM_01%$</font></td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DURATION_PRIM_01%$</td>
  <td class="xl11414491" width="17" style="border-left:none;width:13pt">$%ISAUDIO_PRIM_01%$</td>
  <td class="xl12114491" width="17" style="width:13pt">$%ISVIDEO_PRIM_01%$</td>
  <td colspan="4" class="xl18314491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%CHAN_PRIM_01%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%SCALE_PRIM_01%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%FIXED_PRIM_01%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">2</td>
  <td colspan="4" class="xl24114491" width="68" style="border-left:none;width:52pt">$%DIF_TC_PRIM_02%$</td>
  <td colspan="4" class="xl13014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%SRC_TC_PRIM_02%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt"><font class="font514491">$%DESC_PRIM_02%$</font></td>
  <td colspan="3" class="xl18114491" width="51" style="width:39pt">$%DURATION_PRIM_02%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISAUDIO_PRIM_02%$</td>
  <td class="xl12214491" width="17" style="border-left:none;width:13pt">$%ISVIDEO_PRIM_02%$</td>
  <td colspan="4" class="xl12914491" width="68" style="border-left:none;width:52pt">$%CHAN_PRIM_02%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%SCALE_PRIM_02%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%FIXED_PRIM_02%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11314491" style="height:15.0pt;border-top:none">3</td>
  <td colspan="4" class="xl24114491" width="68" style="border-left:none;width:52pt">$%DIF_TC_PRIM_03%$</td>
  <td colspan="4" class="xl13014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%SRC_TC_PRIM_03%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%DESC_PRIM_03%$</td>
  <td colspan="3" class="xl18214491" width="51" style="border-right:.5pt hairline #BFBFBF;
  width:39pt">$%DURATION_PRIM_03%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISAUDIO_PRIM_03%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISVIDEO_PRIM_03%$</td>
  <td colspan="4" class="xl18314491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%CHAN_PRIM_03%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%SCALE_PRIM_03%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%FIXED_PRIM_03%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11314491" style="height:15.0pt;border-top:none">4</td>
  <td colspan="4" class="xl24114491" width="68" style="border-left:none;width:52pt">$%DIF_TC_PRIM_04%$</td>
  <td colspan="4" class="xl13014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%SRC_TC_PRIM_04%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%DESC_PRIM_04%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DURATION_PRIM_04%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISAUDIO_PRIM_04%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISVIDEO_PRIM_04%$</td>
  <td colspan="4" class="xl18314491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%CHAN_PRIM_04%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%SCALE_PRIM_04%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%FIXED_PRIM_04%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11314491" style="height:15.0pt;border-top:none">5</td>
  <td colspan="4" class="xl24114491" width="68" style="border-left:none;width:52pt">$%DIF_TC_PRIM_05%$</td>
  <td colspan="4" class="xl13014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%SRC_TC_PRIM_05%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%DESC_PRIM_05%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DURATION_PRIM_05%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISAUDIO_PRIM_05%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISVIDEO_PRIM_05%$</td>
  <td colspan="4" class="xl18314491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%CHAN_PRIM_05%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%SCALE_PRIM_05%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%FIXED_PRIM_05%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11314491" style="height:15.0pt;border-top:none">6</td>
  <td colspan="4" class="xl24114491" width="68" style="border-left:none;width:52pt">$%DIF_TC_PRIM_06%$</td>
  <td colspan="4" class="xl13014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%SRC_TC_PRIM_06%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%DESC_PRIM_06%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DURATION_PRIM_06%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISAUDIO_PRIM_06%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISVIDEO_PRIM_06%$</td>
  <td colspan="4" class="xl18314491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%CHAN_PRIM_06%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%SCALE_PRIM_06%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%FIXED_PRIM_06%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11314491" style="height:15.0pt;border-top:none">7</td>
  <td colspan="4" class="xl24114491" width="68" style="border-left:none;width:52pt">$%DIF_TC_PRIM_07%$</td>
  <td colspan="4" class="xl13014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%SRC_TC_PRIM_07%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%DESC_PRIM_07%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DURATION_PRIM_07%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISAUDIO_PRIM_07%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISVIDEO_PRIM_07%$</td>
  <td colspan="4" class="xl18314491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%CHAN_PRIM_07%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%SCALE_PRIM_07%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%FIXED_PRIM_07%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11314491" style="height:15.0pt;border-top:none">8</td>
  <td colspan="4" class="xl24114491" width="68" style="border-left:none;width:52pt">$%DIF_TC_PRIM_08%$</td>
  <td colspan="4" class="xl13014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%SRC_TC_PRIM_08%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%DESC_PRIM_08%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DURATION_PRIM_08%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISAUDIO_PRIM_08%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISVIDEO_PRIM_08%$</td>
  <td colspan="4" class="xl18314491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%CHAN_PRIM_08%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%SCALE_PRIM_08%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%FIXED_PRIM_08%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">9</td>
  <td colspan="4" class="xl24114491" width="68" style="border-left:none;width:52pt">$%DIF_TC_PRIM_09%$</td>
  <td colspan="4" class="xl13014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%SRC_TC_PRIM_09%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%DESC_PRIM_09%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DURATION_PRIM_09%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISAUDIO_PRIM_09%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISVIDEO_PRIM_09%$</td>
  <td colspan="4" class="xl18314491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%CHAN_PRIM_09%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%SCALE_PRIM_09%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%FIXED_PRIM_09%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">10</td>
  <td colspan="4" class="xl24114491" width="68" style="border-left:none;width:52pt">$%DIF_TC_PRIM_10%$</td>
  <td colspan="4" class="xl13014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%SRC_TC_PRIM_10%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%DESC_PRIM_10%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DURATION_PRIM_10%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISAUDIO_PRIM_10%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%ISVIDEO_PRIM_10%$</td>
  <td colspan="4" class="xl18314491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%CHAN_PRIM_10%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%SCALE_PRIM_10%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%FIXED_PRIM_10%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr class="xl8514491" height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="8" class="xl14614491">&nbsp;</td>
  <td colspan="34" class="xl14814491" style="border-right:.5pt solid #404040">Primary
  Asset Tail Logos</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11514491" style="height:15.0pt">&nbsp;</td>
  <td colspan="4" rowspan="2" class="xl15014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:52pt">Difuze&#7475;&#7484; Timecode</td>
  <td colspan="4" rowspan="2" class="xl15614491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:52pt">Source Timecode</td>
  <td colspan="34" rowspan="2" class="xl16214491" width="578" style="border-right:.5pt solid #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:442pt">Description</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="4" class="xl13714491" width="68" style="width:52pt">$%DIF_TC_TAIL_LOGO%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_TAIL_LOGO%$</td>
  <td colspan="34" class="xl12614491" width="578" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:442pt">$%TAIL_LOGO%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl12314491">&nbsp;</td>
  <td class="xl12314491">&nbsp;</td>
  <td class="xl12314491">&nbsp;</td>
  <td class="xl12314491">&nbsp;</td>
  <td class="xl12314491">&nbsp;</td>
  <td class="xl12314491">&nbsp;</td>
  <td class="xl12314491">&nbsp;</td>
  <td class="xl12314491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8614491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr class="xl8514491" height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="8" class="xl14614491">&nbsp;</td>
  <td colspan="34" class="xl14814491" style="border-right:.5pt solid #404040">Primary
  Asset OTT Specifics Rejections</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11514491" style="height:15.0pt">&nbsp;</td>
  <td colspan="4" rowspan="2" class="xl15014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:52pt">Difuze&#7475;&#7484; Timecode</td>
  <td colspan="4" rowspan="2" class="xl15614491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:52pt">Source Timecode</td>
  <td colspan="21" rowspan="2" class="xl16214491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:273pt">Description</td>
  <td colspan="3" rowspan="2" class="xl16814491" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF">Duration</td>
  <td rowspan="2" class="xl17214491" style="border-bottom:.5pt hairline #BFBFBF">A</td>
  <td rowspan="2" class="xl17214491" style="border-bottom:.5pt hairline #BFBFBF">V</td>
  <td colspan="4" rowspan="2" class="xl17414491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:52pt">Chan / Sector</td>
  <td colspan="2" rowspan="2" class="xl17614491" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF">Scale</td>
  <td colspan="2" rowspan="2" class="xl13814491" width="34" style="border-right:.5pt solid #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:26pt">Fixed issue</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11314491" style="height:15.0pt">1</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_OTT_REJECT_TC_01%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_OTT_REJECT_TC_01%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%OTT_REJECT_DESC_01%$</td>
  <td colspan="3" class="xl13314491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%OTT_REJECT_DUR_01%$</td>
  <td class="xl11414491" width="17" style="border-left:none;width:13pt">$%OTT_REJECT_ISAUDIO_01%$</td>
  <td class="xl12114491" width="17" style="width:13pt">$%OTT_REJECT_ISVIDEO_01%$</td>
  <td colspan="4" class="xl14214491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">$%OTT_REJECT_CHAN_01%$</td>
  <td colspan="2" class="xl14514491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%OTT_REJECT_SCALE_01%$</td>
  <td colspan="2" class="xl14514491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%OTT_REJECT_FIXE_01%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">2</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_OTT_REJECT_TC_02%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_OTT_REJECT_TC_02%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%OTT_REJECT_DESC_02%$</td>
  <td colspan="3" class="xl12214491" width="51" style="border-left:none;width:39pt">$%OTT_REJECT_DUR_02%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%OTT_REJECT_ISAUDIO_02%$</td>
  <td class="xl12214491" width="17" style="border-left:none;width:13pt">$%OTT_REJECT_ISVIDEO_02%$</td>
  <td colspan="4" class="xl12914491" width="68" style="border-left:none;width:52pt">$%OTT_REJECT_CHAN_02%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%OTT_REJECT_SCALE_02%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%OTT_REJECT_FIXE_02%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">3</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_OTT_REJECT_TC_03%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_OTT_REJECT_TC_03%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%OTT_REJECT_DESC_03%$</td>
  <td colspan="3" class="xl12214491" width="51" style="border-left:none;width:39pt">$%OTT_REJECT_DUR_03%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%OTT_REJECT_ISAUDIO_03%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%OTT_REJECT_ISVIDEO_03%$</td>
  <td colspan="4" class="xl12914491" width="68" style="border-left:none;width:52pt">$%OTT_REJECT_CHAN_03%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%OTT_REJECT_SCALE_03%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%OTT_REJECT_FIXE_03%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">4</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_OTT_REJECT_TC_04%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_OTT_REJECT_TC_04%$</td>
  <td colspan="21" class="xl12614491" width="357" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:273pt">$%OTT_REJECT_DESC_04%$</td>
  <td colspan="3" class="xl12214491" width="51" style="border-left:none;width:39pt">$%OTT_REJECT_DUR_04%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%OTT_REJECT_ISAUDIO_04%$</td>
  <td class="xl12214491" width="17" style="border-top:none;border-left:none;
  width:13pt">$%OTT_REJECT_ISVIDEO_04%$</td>
  <td colspan="4" class="xl12914491" width="68" style="border-left:none;width:52pt">$%OTT_REJECT_CHAN_04%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%OTT_REJECT_SCALE_04%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%OTT_REJECT_FIXE_04%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:245px">
  <td height="19" class="xl8814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="page-break-before:always;height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="42" rowspan="2" class="xl22414491" style="border-right:.5pt solid #404040;
  border-bottom:.5pt solid #404040">PRIMARY ASSET TEXT PRESENCE</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl12314491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl12314491" style="height:14.4pt">&nbsp;</td>
  <td colspan="6" rowspan="2" class="xl23014491" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF">Text Type Key</td>
  <td class="xl11614491" style="border-left:none">t</td>
  <td colspan="8" class="xl23614491" style="border-right:.5pt hairline #BFBFBF">On-screen
  text</td>
  <td class="xl11714491" style="border-left:none">p</td>
  <td colspan="8" class="xl23614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Partial subtitle</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl12314491" style="height:14.4pt">&nbsp;</td>
  <td class="xl11714491" style="border-top:none;border-left:none">s</td>
  <td colspan="8" class="xl23614491" style="border-left:none">Subtitles narratives</td>
  <td class="xl11714491" style="border-top:none">f</td>
  <td colspan="8" class="xl23614491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">Full subtitle</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl12314491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl12314491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
  <td class="xl8414491">&nbsp;</td>
 </tr>
 <tr class="xl8514491" height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="8" class="xl14614491">&nbsp;</td>
  <td colspan="34" class="xl14814491" style="border-right:.5pt solid #404040">Primary
  Asset Text Presence</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11814491" style="height:15.0pt">&nbsp;</td>
  <td colspan="4" rowspan="2" class="xl21814491" width="68" style="border-right:.5pt hairline #BFBFBF;
  width:52pt">Difuze&#7475;&#7484; Timecode</td>
  <td colspan="4" rowspan="2" class="xl15614491" width="68" style="border-right:.5pt hairline #BFBFBF;
  width:52pt">Source Timecode</td>
  <td colspan="24" rowspan="2" class="xl15614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:312pt">Description</td>
  <td colspan="3" rowspan="3" class="xl20314491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:39pt">Duration</td>
  <td colspan="7" rowspan="3" class="xl20314491" width="119" style="border-right:.5pt solid #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:91pt">Text type</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11814491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11914491" style="height:14.4pt">&nbsp;</td>
  <td colspan="4" class="xl22014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">&nbsp;</td>
  <td colspan="4" class="xl22314491" width="68" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:52pt">&nbsp;</td>
  <td colspan="24" class="xl19414491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">Examples of text presence :</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11314491" style="height:15.0pt;border-top:none">1</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_TC_TEXT_01%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_TEXT_01%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%TEXT_01%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DUR_TEXT_01%$</td>
  <td colspan="7" class="xl21514491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%TYPE_TEXT_01%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11314491" style="height:15.0pt;border-top:none">2</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_TC_TEXT_02%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_TEXT_02%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%TEXT_02%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DUR_TEXT_02%$</td>
  <td colspan="7" class="xl21514491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%TYPE_TEXT_02%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">3</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_TC_TEXT_03%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_TEXT_03%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%TEXT_03%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DUR_TEXT_03%$</td>
  <td colspan="7" class="xl21514491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%TYPE_TEXT_03%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">4</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_TC_TEXT_04%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_TEXT_04%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%TEXT_04%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%DUR_TEXT_04%$</td>
  <td colspan="7" class="xl21514491" style="border-right:.5pt hairline #BFBFBF;
  border-left:none">$%TYPE_TEXT_04%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr class="xl8514491" height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="8" class="xl14614491">&nbsp;</td>
  <td colspan="34" class="xl14814491" style="border-right:.5pt solid #404040">Primary
  Asset Needed Text and/or Subtitle</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11514491" style="height:15.0pt">&nbsp;</td>
  <td colspan="4" rowspan="2" class="xl15014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  width:52pt">Difuze&#7475;&#7484; Timecode</td>
  <td colspan="4" rowspan="2" class="xl15614491" width="68" style="border-right:.5pt hairline #BFBFBF;
  width:52pt">Source Timecode</td>
  <td colspan="24" rowspan="2" class="xl15614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:312pt">Description</td>
  <td colspan="3" rowspan="3" class="xl20314491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:39pt">Duration</td>
  <td colspan="3" rowspan="3" class="xl20314491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:39pt">Text type</td>
  <td colspan="2" rowspan="3" class="xl20714491" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF">Scale</td>
  <td colspan="2" rowspan="3" class="xl15614491" width="34" style="border-right:.5pt solid #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:26pt">Fixed issue</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11914491" style="height:15.0pt">&nbsp;</td>
  <td colspan="4" class="xl21414491" width="68" style="border-left:none;width:52pt">&nbsp;</td>
  <td colspan="4" class="xl19314491" width="68" style="border-left:none;width:52pt">&nbsp;</td>
  <td colspan="24" class="xl19414491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">Examples of needed text:</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">1</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_TC_NEEDEDTEXT_01%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_NEEDEDTEXT_01%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%NEEDEDTEXT_01%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%NEEDEDTEXT_DUR_01%$</td>
  <td colspan="3" class="xl19214491" width="51" style="border-left:none;width:39pt">$%NEEDEDTEXT_TYPE_01%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%NEEDEDTEXT_SCALE_01%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%NEEDEDTEXT_FIXE_01%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">2</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_TC_NEEDEDTEXT_02%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_NEEDEDTEXT_02%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%NEEDEDTEXT_02%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">%NEEDEDTEXT_DUR_02%$</td>
  <td colspan="3" class="xl19214491" width="51" style="border-left:none;width:39pt">$%NEEDEDTEXT_TYPE_02%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%NEEDEDTEXT_SCALE_02%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%NEEDEDTEXT_FIXE_02%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">3</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_TC_NEEDEDTEXT_03%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_NEEDEDTEXT_03%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%NEEDEDTEXT_03%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">%NEEDEDTEXT_DUR_03%$</td>
  <td colspan="3" class="xl19214491" width="51" style="border-left:none;width:39pt">$%NEEDEDTEXT_TYPE_03%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%NEEDEDTEXT_SCALE_03%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%NEEDEDTEXT_FIXE_03%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">4</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIF_TC_NEEDEDTEXT_04%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_NEEDEDTEXT_04%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%NEEDEDTEXT_04%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%NEEDEDTEXT_DUR_04%$</td>
  <td colspan="3" class="xl19214491" width="51" style="border-left:none;width:39pt">$%NEEDEDTEXT_TYPE_04%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%NEEDEDTEXT_SCALE_04%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%NEEDEDTEXT_FIXE_04%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr class="xl8514491" height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td colspan="8" class="xl14614491">&nbsp;</td>
  <td colspan="34" class="xl14814491" style="border-right:.5pt solid #404040">Primary
  Asset Text and/or Subtitles issues</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11514491" style="height:15.0pt">&nbsp;</td>
  <td colspan="4" rowspan="2" class="xl15014491" width="68" style="border-right:.5pt hairline #BFBFBF;
  width:52pt">Difuze&#7475;&#7484; Timecode</td>
  <td colspan="4" rowspan="2" class="xl15614491" width="68" style="border-right:.5pt hairline #BFBFBF;
  width:52pt">Source Timecode</td>
  <td colspan="24" rowspan="2" class="xl15614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:312pt">Description</td>
  <td colspan="3" rowspan="3" class="xl20314491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:39pt">Duration</td>
  <td colspan="3" rowspan="3" class="xl20314491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:39pt">Text type</td>
  <td colspan="2" rowspan="3" class="xl20714491" style="border-right:.5pt hairline #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF">Scale</td>
  <td colspan="2" rowspan="3" class="xl15614491" width="34" style="border-right:.5pt solid #BFBFBF;
  border-bottom:.5pt hairline #BFBFBF;width:26pt">Fixed issue</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="20" style="mso-height-source:userset;height:15.0pt">
  <td height="20" class="xl11914491" style="height:15.0pt">&nbsp;</td>
  <td colspan="4" class="xl21414491" width="68" style="border-left:none;width:52pt">&nbsp;</td>
  <td colspan="4" class="xl19314491" width="68" style="border-left:none;width:52pt">&nbsp;</td>
  <td colspan="24" class="xl19414491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">Examples of text issue:</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">1</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIFF_TC_ISSUETEXT_01%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_ISSUETEXT_01%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%ISSUETEXT_01%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%ISSUETEXT_DUR_01%$</td>
  <td colspan="3" class="xl19214491" width="51" style="border-left:none;width:39pt">$%ISSUETEXT_TYPE_01%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%ISSUETEXT_SCALE_01%$</td>
  <td colspan="2" class="xl18014491" width="34" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:26pt">$%ISSUETEXT_FIXE_01%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">2</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIFF_TC_ISSUETEXT_02%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_ISSUETEXT_02%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%ISSUETEXT_02%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%ISSUETEXT_DUR_02%$</td>
  <td colspan="3" class="xl19214491" width="51" style="border-left:none;width:39pt">$%ISSUETEXT_TYPE_02%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%ISSUETEXT_SCALE_02%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%ISSUETEXT_FIXE_02%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">3</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIFF_TC_ISSUETEXT_03%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_ISSUETEXT_03%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%ISSUETEXT_03%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%ISSUETEXT_DUR_03%$</td>
  <td colspan="3" class="xl19214491" width="51" style="border-left:none;width:39pt">$%ISSUETEXT_TYPE_03%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%ISSUETEXT_SCALE_03%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%ISSUETEXT_FIXE_03%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11314491" style="height:14.4pt;border-top:none">4</td>
  <td colspan="4" class="xl13714491" width="68" style="border-left:none;width:52pt">$%DIFF_TC_ISSUETEXT_04%$</td>
  <td colspan="4" class="xl12414491" width="68" style="width:52pt">$%SRC_TC_ISSUETEXT_04%$</td>
  <td colspan="24" class="xl12614491" width="408" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:312pt">$%ISSUETEXT_04%$</td>
  <td colspan="3" class="xl18014491" width="51" style="border-right:.5pt hairline #BFBFBF;
  border-left:none;width:39pt">$%ISSUETEXT_DUR_04%$</td>
  <td colspan="3" class="xl19214491" width="51" style="border-left:none;width:39pt">$%ISSUETEXT_TYPE_04%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%ISSUETEXT_SCALE_04%$</td>
  <td colspan="2" class="xl12214491" width="34" style="border-left:none;width:26pt">$%ISSUETEXT_FIXE_04%$</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <tr height="19" style="height:14.4pt">
  <td height="19" class="xl11514491" style="height:14.4pt">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
  <td class="xl8514491">&nbsp;</td>
 </tr>
 <!--[if supportMisalignedColumns]-->
 <tr height="0" style="display:none">
  <td width="21" style="width:16pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="17" style="width:13pt"></td>
  <td width="21" style="width:16pt"></td>
 </tr>
 <!--[endif]-->
</tbody></table>

</div>

</body></html>`



	f, err := os.Create("C:/Users/daudels/Desktop/test_html/test.html")
	check(err)
	defer f.Close()

	f.WriteString(HtmlCode)
}
 
func check(e error) {
    if e != nil {
        panic(e)
    }
}