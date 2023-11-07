# Project Backend Invoice
## Dari design figma dan flow diketahui bahwa :
1. Terdapat invoice indexing : yang berisi tampilan semua nama customer beserta 
invoice id, subject, total items, status date dan status pembayaran

2. Terdapat add invoice : yang berisi invoice detail, customer info dan detail item, jadi ini kayanya terpisah untuk fungsi add di post tapi saling menganerete. diantara itu akan ada yg menjadi promary key nya. 
ex : (post) http//localhost:8060/addinvoice
terdapat

======================
Perintah Penting
======================
1. Instal Gorm jinzhu : go get github.com/jinzhu/gorm
2. Buat database dengan mysql jinzhu : go get github.com/jinzhu/gorm
======================
Berarti nanti kalau mau bikin invoice beserta nama customer jadi kaya gini :
http://localhost/8060/addinvoice/id