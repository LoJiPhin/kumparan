package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//News Adalah struct yang digunakan untuk menyimpan Berita baru
type News struct {
	//ID adalah identity pada struct News
	ID int `json:"id"`
	//Author adalah variable yang menyimpan data nama penulis berita
	Author string `json:"author"`
	//Body adalah menyimpan isi berita
	Body string `json:"body"`
	//Created diisi secara otomatis oleh system bedasarkan waktu pembuatan berita
	Created string `json:"Created"`
}

//CreateNews adalah
func (news News) InsertNews() {
	db := Connect()
	defer db.Close()
	//query := fmt.Sprintf(`INSERT INTO News(author, body, created) (` + news.Author + `, ` + news.Body + `,` + news.Created + ` )`)
	stmt, err := db.Prepare("INSERT News SET author=?,body=?,created=?")
	if err != nil {
		panic(err.Error())
	}
	res, err := stmt.Exec(news.Author, news.Body, time.Now())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("result :", res.RowsAffected)
	fmt.Println("news :", news)
}
