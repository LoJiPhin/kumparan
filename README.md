# kumparan

Sebelum menjalankan project ini saya asumsikan, project nya sudah berada di gopath nya dan sudah terinstall docker.
silahkan jalankan command dibawah ini untuk menjalankan servicenya
docker-compose -f "docker-compose.yml" up -d --build

pastikan semua container sudah up, jika terdapat container yang belum up, bisa lakukan restart container tersebut.

program terdiri dari dua bagian
producer dan consumer
1. untuk melakukan input bisa akses method POST ke http://localhost:3002/news
dan kirim kan data 
{
  "author" : "Dilan",
  "body" : "Milea"
}
2. untuk view data bisa dengan method GET ke http://localhost:3002/news



