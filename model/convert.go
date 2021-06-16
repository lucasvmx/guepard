package model

type ConvertStatus struct {
	ConversionStatus string `json:"c_status"`
	DownloadLink     string `json:"dlink"`
	Quality          string `json:"fquality"`
	FormatType       string `json:"ftype"`
	Mess             string `json:"mess"`
	Status           string `json:"status"`
	Title            string `json:"title"`
	Video            string `json:"vid"`
}

type ConvertQuery struct {
	Vid string
	K   string
}

var (
	convertPath = "/api/ajaxConvert/convert"
)

func GetConversionPath() string {
	return convertPath
}
