package models

import (
	"database/sql"
	"fmt"
	"go-postgres-crud/config"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// Buku schema dari tabel Buku
// kita coba dengan null string
type Buku struct {
	ID            int64             `json:"id"`
	Judul_buku    config.NullString `json:"judul_buku"`
	Penulis       string            `json:"penulis"`
	Tgl_publikasi string            `json:"tgl_publikasi"`
}

// ambil satu buku
func AmbilSemuaBuku() ([]Buku, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var bukus []Buku

	// kita buat sleect query
	sqlStatement := `SELECT * FROM buku`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()

	// kita iterasi mengambil datanya
	for rows.Next() {
		var buku Buku

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&buku.ID, &buku.Judul_buku, &buku.Penulis, &buku.Tgl_publikasi)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		// masukkan kedalam slice bukus
		bukus = append(bukus, buku)

	}

	// return empty buku atau jika error
	return bukus, err
}

// mengambil satu buku
func AmbilSatuBuku(id int64) (Buku, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var buku Buku

	// buat sql query
	sqlStatement := `SELECT * FROM buku WHERE id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&buku.ID, &buku.Judul_buku, &buku.Penulis, &buku.Tgl_publikasi)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return buku, nil
	case nil:
		return buku, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return buku, err
}

// delete user in the DB
func HapusBuku(id int64) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// buat sql query
	sqlStatement := `DELETE FROM buku WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected
}
