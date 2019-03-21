# 4210161009_TCP_Encapsulation - DOTA 2
```go
type Hero struct {
	ID     int // Hero ID
	XCoord int // Hero Coordinate X
	YCoord int // Hero Coordinate Y
	Mana   int // Hero Mana
	Gold   int // Hero Gold
}
```

# Pengiriman memakai Struct
1. Golang memiliki library bernama gob untuk encode/decode Struct
2. Gob NewEncoder memiliki 1 parameter yang menerima io.Writer (bisa berupa tcp writer)
3. Kemudian Gob melakukan encode yang kemudian akan dikirim melalui socket tcp
4. Server yang menerima paket dari client juga menggunakan library gob untuk decode paket yang dikirim tadi
5. Hasil decode tersebut kemudian dimasukkan ke dalam struct yang ada pada server
6. Setelah server menerima paket dari client, server meng- Encode paket tadi dan kemudian mengirimkan kembali ke client.