package contracts

import (
	"fmt"
	"log"
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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//InsertNews adalah blabla
func (news News) InsertNews() {
	db := Connect()
	defer db.Close()
	//query := fmt.Sprintf(`INSERT INTO News(author, body, created) (` + news.Author + `, ` + news.Body + `,` + news.Created + ` )`)
	stmt, err := db.Prepare("INSERT News SET author=?,body=?,created=?")
	failOnError(err, "Failed to Insert News Statment")
	res, err := stmt.Exec(news.Author, news.Body, time.Now())
	failOnError(err, "Failed to Insert News Execution")
	fmt.Println("result :", res.RowsAffected)
	fmt.Println("news :", news)
}
