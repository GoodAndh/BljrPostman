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

func Index(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("index.html")
		if err != nil {
			fmt.Println(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		filename := "data.json"
		_, err := os.Stat(filename) //pengecekan apakah ada file dgn nama ....
		if os.IsNotExist(err) {     // jika tdk ada akan dibuat
			file, err := os.Create(filename)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
		}

		metode := "s="
		apikey := "ed7fa361"
		Masukan := &Bebas{ //menampung user
			Data: r.Form.Get("cari_judul"),
		} // minta data ke server dari user
		baseUrl := "http://www.omdbapi.com/" + "?apikey=" + apikey + "&" + metode + Masukan.Data
		respone, _ := http.Get(baseUrl)
		bodis, _ := ioutil.ReadAll(respone.Body)
		Search := &Search{} // menampung hasil json
		json.Unmarshal(bodis, Search)

		// fmt.Println(respone)

		//mengubah ke tipe json
		jsonData, _ := json.MarshalIndent(Search, "", "  ")

		//menulis file json
		Nulisjson, _ := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
		Nulisjson.Write(jsonData)
		defer Nulisjson.Close() //masalah yg diatasi yaitu akan mengcopy paste tidak menulis isi file

		// jika ada error
		var message error
		if Search.Response == "False" {
			message = errors.New("Judul tidak ditemukan!")
		}
		if message != nil {
			data := map[string]interface{}{
				"error": message,
			}
			temp, err := template.ParseFiles("index.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, data)

		} else { // jika tdk ada error atau bebas.respone== True
			data := map[string]interface{}{}
			// tempatid := []string{}
			// Searchid := &Searchid{}
			Nulisjsons, _ := os.OpenFile("byid.json", os.O_WRONLY|os.O_TRUNC, 0644) //mwmbuka file byid.json
			defer Nulisjsons.Close()
			var jsonDslice []interface{}

			for i, v := range Search.Pencari {
				data1 := map[string]interface{}{
					"Urutan": strconv.Itoa(i + 1),
					"Title":  v.Judul,
					"Year":   v.Year,
					"Imdbid": v.ImdbID,
					"Type":   v.Type,
					"Poster": v.Poster,
					"klikme": "http://www.omdbapi.com/" + "?apikey=" + apikey + "&i=" + v.ImdbID,
				}

				data[strconv.Itoa(i+1)] = data1

				klikmeURL := data1["klikme"].(string)
				// fmt.Print("=======")
				// fmt.Println(klikmeURL)
				jsonDslice = append(jsonDslice, klikmeURL)

			}
			// fmt.Println(data["1"])
			// for i, v := range data {
			// 	fmt.Println("=== ISI JSONDSLICE ===")
			// 	fmt.Print(i, ". ")
			// 	fmt.Println(v)

			// 	fmt.Println()
			// 	fmt.Println()

			// }

			// for _,v:=range data{
			// 	fmt.Println("======")
			// 	fmt.Println(v)
			// 	if v==io.EOF{
			// 		fmt.Println()
			// 	}
			// 	fmt.Println("======")

			// }
			// json, err := json.MarshalIndent(data, "", "  ")
			// if err != nil {
			// 	panic(err)
			// }

			// fmt.Println(string(json))

			temp, err := template.ParseFiles("index.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, data)

		}

	}
}
