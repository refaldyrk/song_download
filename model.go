package song_donwload

type Song struct {
	Query   string `json:"query"`
	Blocked bool   `json:"blocked"`
	Items   []Item `json:"items"`
}

type Item struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Views     string `json:"views"`
	Duration  string `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}

type Extractor struct {
	Extractor     string `json:"extractor"`
	VideoId       string `json:"videoId"`
	Title         string `json:"title"`
	LengthSeconds string `json:"lengthSeconds"`
	Jwt           string `json:"jwt"`
	Tasks         []Task `json:"tasks"`
}

type Task struct {
	Bitrate  int    `json:"bitrate"`
	Filesize string `json:"filesize"`
	Hash     string `json:"hash"`
}

type Download struct {
	TaskId           string `json:"taskId"`
	Status           string `json:"status"`
	DownloadProgress int    `json:"download_progress"`
	ConvertProgress  int    `json:"convert_progress"`
	Title            string `json:"title"`
	Ext              string `json:"ext"`
	VideoId          string `json:"videoId"`
	Path             string `json:"path"`
	Quality          int    `json:"quality"`
	ReqIP            string `json:"reqIP"`
	ReqRef           string `json:"reqRef"`
	ReqUserAgent     string `json:"reqUserAgent"`
	Origin           string `json:"origin"`
	Filesize         string `json:"filesize"`
	Download         string `json:"download"`
	PostDloadUrl     string `json:"postDloadUrl"`
	Hash             string `json:"hash"`
}

//func (i Item) Download() string {
//
//}
