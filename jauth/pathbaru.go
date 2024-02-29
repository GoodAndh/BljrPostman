package jauth


type Hasil struct {
	Urutan string `json:"Urutan"`
	Judul  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type Search struct {
	Pencari     []Hasil `json:"Search"`
	TotalResult string  `json:"totalResults"`
	Response    string  `json:"Response"`
}


type Searchid struct {
	Title    string `json:"Title"`
	Released string `json:"Released"`
	Poster   string `json:"Poster"`
	Director string `json:"Director"`
	Genre    string `json:"Genre"`
	Actors   string `json:"Actors"`
	Imdbid   string `json:"imdbID"`
	Response string `json:"Response"`
}

type Bebas struct {
	Data     string //Keluaran Data output atau return
}