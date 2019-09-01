package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func routers() *chi.Mux {
	router.Get("/", ping)

	router.Post("/articles", CreateArticle)
	router.Get("/articles", GetAllArticle)
	router.Get("/articles/{id}", GetArticleByID)

	return router
}

//-------------------------API STARTING POINT-------------------------//
func ping(w http.ResponseWriter, r *http.Request) {
	rbody, isDBAvailable := dbAvailablityResponse(db)

	if !isDBAvailable {
		rbody.Message = "Server Unavailable"
	}

	respondwithJSON(w, rbody.Status, rbody)
}

//----------------------------API ENDPOINT----------------------------//
//All API endpoint request will start to check server availability before
//proceed to user's request input

//returns all article data
func GetAllArticle(w http.ResponseWriter, r *http.Request) {
	errors := []error{}
	data := []Article{}
	rbody, isDBAvailable := dbAvailablityResponse(db)

	if !isDBAvailable {
		respondwithJSON(w, rbody.Status, rbody)
		return
	}

	rows, err := db.Query("Select id, title, content, author From articles")
	catch(err)

	defer rows.Close()

	rbody = ResponseBody{
		Status:  http.StatusNotFound,
		Message: "No article found",
	}

	for rows.Next() {
		post := Article{}

		er := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)

		if er != nil {
			errors = append(errors, er)
		}
		data = append(data, post)

		rbody = ResponseBody{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    data,
		}
	}

	respondwithJSON(w, http.StatusOK, rbody)
}

//create a new post
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	rbody, isDBAvailable := dbAvailablityResponse(db)

	if !isDBAvailable {
		respondwithJSON(w, rbody.Status, rbody)
		return
	}

	var article Article
	json.NewDecoder(r.Body).Decode(&article)

	rbody, isPostValid := articleValidityResponse(article)
	if isPostValid {
		query, err := db.Prepare("Insert articles SET title=?, content=?, author=?")
		catch(err)

		_, er := query.Exec(article.Title, article.Content, article.Author)
		catch(er)
		defer query.Close()

		post := Article{}
		rows := db.QueryRow("Select id, title, content, author From articles Where id=(SELECT max(id) FROM articles)")
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author)

		//fmt.Println("db last id: ", post.ID)

		rbody = ResponseBody{
			Status:  http.StatusCreated,
			Message: "Success",
			Data: []Article{Article{
				ID: post.ID,
			}},
		}
	}

	respondwithJSON(w, rbody.Status, rbody)
}

func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	rbody, isDBAvailable := dbAvailablityResponse(db)

	if !isDBAvailable {
		respondwithJSON(w, rbody.Status, rbody)
		return
	}

	article := Article{}
	id := chi.URLParam(r, "id")

	row := db.QueryRow("Select id, title, content, author From articles where id=?", id)
	err := row.Scan(&article.ID, &article.Title, &article.Content, &article.Author)

	if err != nil {
		rbody.Status = http.StatusNotFound
		rbody.Message = "No article found in selected id"
		respondwithJSON(w, rbody.Status, rbody)
		return
	}

	rbody = ResponseBody{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    []Article{article},
	}

	respondwithJSON(w, rbody.Status, rbody)
}
