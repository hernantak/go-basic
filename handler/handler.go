package handler

import (
	"go-web/handler/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_template, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error On This Pages", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "Learning Golang Web",
	// 	"content": 86,
	// }

	// data := entity.Product{
	// 	ID:    1,
	// 	Nama:  "Mobilio",
	// 	Price: 220000,
	// 	Stock: 3,
	// }

	data := []entity.Product{
		{
			ID:    1,
			Nama:  "Mobilio",
			Price: 220000,
			Stock: 2,
		},
		{
			ID:    2,
			Nama:  "Xpander",
			Price: 270000,
			Stock: 8,
		},
		{
			ID:    3,
			Nama:  "Pajero",
			Price: 5000000,
			Stock: 11,
		},
	}

	err = _template.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error With Execute, Please Contact Developer", http.StatusInternalServerError)
		return
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	id_num, err := strconv.Atoi(id)

	if err != nil || id_num < 0 {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"content": id_num,
	}

	_template, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error On This Pages", http.StatusInternalServerError)
		return
	}

	err = _template.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error With Execute, Please Contact Developer", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini GET"))
	case "POST":
		w.Write([]byte("Ini POST"))
	default:
		http.Error(w, "[ERROR] Cannot Get/Post", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_template, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "[Error] While processing GET", http.StatusInternalServerError)
			return
		}

		err = _template.Execute(w, nil)
		if err != nil {
			log.Println(w, "[Error] While Execute GET", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "[Error] Status Bad Request", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "[Error] While Post Data", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		messsage := r.Form.Get("messsage")

		data := map[string]interface{}{
			"name":    name,
			"message": messsage,
		}

		_template, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "[Error] While Execute Data", http.StatusInternalServerError)
			return
		}

		err = _template.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "[Error] Cannot Execute Data", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "[Error] Cannot POST Data", http.StatusBadRequest)
}
