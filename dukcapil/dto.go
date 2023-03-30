package dukcapil

type Response struct {
	Message string `json:"message"`
}

type Content struct {
	PlaceBirth    string `json:"TMPT_LHR,omitempty"`
	MariageStatus string `json:"STATUS_KAWIN,omitempty"`
	MotherName    string `json:"LAMA_LGKP_IBU,omitempty"`
	Name          string `json:"NAMA_LGKP,omitempty"`
	DateOfBirth   string `json:"TGL_LHR,omitempty"`
	Gender        string `json:"JENIS_KLMIN,omitempty"`
	ResponseCode  string `json:"RESPONSE_CODE,omitempty"`
	ResponseDesc  string `json:"RESPONSE_DESC,omitempty"`
	Response      string `json:"RESPONSE,omitempty"`
}

type DukcapilResponse struct {
	Content          []Content   `json:"content"`
	LastPage         bool        `json:"lastPage"`
	NumberOfElements int         `json:"numberOfElements"`
	Sort             interface{} `json:"sort"`
	TotalElements    int         `json:"totalElements"`
	FirstPage        bool        `json:"firstPage"`
	Number           int         `json:"number"`
	Size             int         `json:"size"`
	QuotaLimiter     int         `json:"quotaLimiter"`
}
