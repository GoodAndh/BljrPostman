package jauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func SearchMetod[T any](s *T, fl, keyapi, metode, duser string) error {
	filename := fl
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	}
	baseUrl := "http://www.omdbapi.com/?" + "apikey=" + keyapi + "&" + metode + duser
	fmt.Println(baseUrl)
	response, err := http.Get(baseUrl)
	if err != nil {
		fmt.Println(err)
		panic(err)

	}
	defer response.Body.Close()
	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		fmt.Println(err1)
		panic(err1)

	}
	json.Unmarshal(body, s)
	jsonData, _ := json.MarshalIndent(s, "", "  ")
	Nulisjson, _ := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	Nulisjson.Write(jsonData)

	return nil

}

func Dindex(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		r.ParseForm()
		filename := "data.json"

		keyapi := "ed7fa361"
		InputUser := &Inputuser{
			Data: r.Form.Get("cari_judul"),
		}

		Pencari := &Search{}
		searcing := &Searchid{}
		SearchMetod(Pencari, filename, keyapi, "s=", InputUser.Data)

		var Errmessage error

		if Pencari.Response == "False" || searcing.Response == "False" {
			Errmessage = errors.New("Judul tidak ditemukan!")
		}
		if Errmessage != nil {
			t, _ := template.ParseFiles("index.html")
			t.Execute(w, Errmessage)

		} else {
			Newdata := map[string]interface{}{}
			Newdata1 := map[string]interface{}{}
			for i, v := range Pencari.Pencari {
				Lastdata := map[string]interface{}{
					"Urutan": strconv.Itoa(i + 1),
					"Title":  v.Judul,
					"Year":   v.Year,
					"Imdbid": v.ImdbID,
					"Type":   v.Type,
					"Poster": v.Poster,
				}
				Newdata[strconv.Itoa(i+1)] = Lastdata
				Newdata1[strconv.Itoa(i+1)] = v.ImdbID

			}

			data := &Data{
				"Newdata":  Newdata,
				"Newdata1": Newdata1,
				"Isvalue":  r.URL.Query().Get(InputUser.Data),
			}

			for _, v := range Newdata1 {
				SearchMetod(searcing, "byid.json", keyapi, "t=", v.(string))
			}

			temp, err := template.ParseFiles("index.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, data)

		}
	} else if r.Method == http.MethodGet {

		temp, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)

	}
}
