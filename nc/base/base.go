package base

//Base struct
type Base struct {
	Account string `json:"account" xml:"account,attr"`
	Billtype string `json:"billtype" xml:"billtype,attr"`
	Filename string `json:"filename" xml:"filename,attr"`
	Groupcode string `json:"groupcode" xml:"groupcode,attr"`
	Isexchange string `json:"isexchange" xml:"isexchange,attr"`
	Replace string `json:"replace" xml:"replace,attr"`
	Roottag string `json:"roottag" xml:"roottag,attr"`
	Sender string `json:"sender" xml:"sender,attr"`
}
