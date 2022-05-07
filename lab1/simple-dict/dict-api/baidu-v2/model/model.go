package model

type Symbol struct {
	PhEn  string `json:"ph_en"`
	PhAm  string `json:"ph_am"`
	Parts []struct {
		Part  string   `json:"part"`
		Means []string `json:"means"`
	} `json:"parts"`
	PhOther string `json:"ph_other"`
}

type SimpleMeans struct {
	Symbols   []Symbol `json:"symbols"`
	WordName  string   `json:"word_name"`
	From      string   `json:"from"`
	WordMeans []string `json:"word_means"`
	Exchange  struct {
		WordPl []string `json:"word_pl"`
	}
}

type Answer struct {
	TransResult any `json:"trans_result"`
	DictResult  struct {
		Edict       any         `json:"edict"`
		From        string      `json:"from"`
		BaikeImgURL string      `json:"baike_img_url"`
		Means       SimpleMeans `json:"simple_means"`
		Collins     any         `json:"collins"`
		Lang        string      `json:"lang"`
		Oxford      any         `json:"oxford"`
	} `json:"dict_result"`
	LijuResult any `json:"liju_result"`
	Logid      any `json:"logid"`
}
