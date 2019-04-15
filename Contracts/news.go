package contracts

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
