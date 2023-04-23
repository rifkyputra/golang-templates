package modularHTTP

import (
	. "ModularHTTPGo/types"
	. "ModularHTTPGo/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Id    int
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetArticleHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "Application/json")

		rows, err := db.Query("select id, title, body from articles")
		if err != nil {
			log.Fatalf(err.Error())
			http.Error(w, http.StatusText(500), 500)
			return
		}
		defer rows.Close()

		articles := []Article{}
		for rows.Next() {
			var article Article
			if err := rows.Scan(&article.Id, &article.Title, &article.Body); err != nil {
				fmt.Println(err.Error())
				http.Error(w, http.StatusText(500), 500)
				return
			}
			articles = append(articles, article)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(BaseResponse{
			Status: 200,
			Alert:  map[string]interface{}{},
			Data:   articles,
		})
	}
}

func PostArticleHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		SetContentJson(w)
		article := Article{}
		if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res, err := db.Exec("INSERT INTO articles (title, body) VALUES ($1, $2)", article.Title, article.Body)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		affected, _ := res.RowsAffected()
		if affected < 1 {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(BaseResponse{
			Status: 200,
			Alert:  map[string]interface{}{},
			Data:   article,
		})
	}
}

func UpdateArticleHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "Application/json")

		article := Article{}
		if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		setStatement := ""

		if len(article.Body) > 1 {
			setStatement = fmt.Sprintf(`%v body = '%v',`, setStatement, article.Body)
		}

		if len(article.Title) > 1 {
			setStatement = fmt.Sprintf(`%v title = '%v',`, setStatement, article.Title)
		}

		if len(setStatement) == 0 {
			http.Error(w, http.StatusText(400), 400)

			return
		}

		setStatement = "SET" + setStatement
		fmt.Println(setStatement)
		res, err := db.Exec(fmt.Sprintf(`UPDATE articles 
		%v updated_at = NOW() 
		WHERE id = %d;
		`, setStatement, article.Id))
		if err != nil {
			fmt.Println(err.Error())

			http.Error(w, http.StatusText(500), 500)
			return
		}

		affected, _ := res.RowsAffected()
		if affected < 1 {
			fmt.Println(err.Error())

			http.Error(w, http.StatusText(500), 500)
			return
		}

		if err != nil {
			fmt.Println(err.Error())

			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(BaseResponse{
			Status: 200,
			Alert:  map[string]interface{}{},
			Data:   article,
		})
	}
}

func DeleteArticleHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "Application/json")

		article := Article{}
		if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		res, err := db.Exec(fmt.Sprintf(`DELETE FROM
		articles WHERE id = %d;`, article.Id))
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
			return
		}

		affected, _ := res.RowsAffected()
		if affected < 1 {
			fmt.Println(err.Error())

			http.Error(w, http.StatusText(500), 500)
			return
		}

		if err != nil {
			fmt.Println(err.Error())

			http.Error(w, http.StatusText(500), 500)
			return
		}

		json.NewEncoder(w).Encode(BaseResponse{
			Status: 200,
			Alert:  map[string]interface{}{},
			Data:   article,
		})
	}
}

func GetArticleByIdHandler(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "Application/json")

		articleId := r.URL.Query().Get("id")

		rows, err := db.Query(fmt.Sprintf("select id, title, body from articles where id=%v", articleId))
		if err != nil {
			log.Fatalf(err.Error())
			http.Error(w, http.StatusText(500), 500)
			return
		}
		defer rows.Close()

		var article Article
		for rows.Next() {
			if err := rows.Scan(&article.Id, &article.Title, &article.Body); err != nil {
				fmt.Println(err.Error())
				http.Error(w, http.StatusText(500), 500)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(BaseResponse{
			Status: 200,
			Alert:  map[string]interface{}{},
			Data:   article,
		})
	}
}
