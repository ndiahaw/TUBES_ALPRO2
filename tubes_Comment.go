package main

import "fmt"

const NMAX int = 100

type comment struct {
	longcomment, nilaiSentimen int
	text, sentimen             string
}
type tabComment [NMAX]comment

var komentarPositif = []string{"bagus", "indah", "cantik", "ganteng", "modern", "keren", "anjay", "wow", "ramah", "enak", "elegan"}
var komentarNegatif = []string{"jelek", "buruk", "menyebalkan", "benci", "burik", "gajelas", "anjir", "gila", "kurang", "najis", "murahan"}
var tab tabComment
var IDKomen int
var nData int = 1
var kataKunci string

// prosedur ini tu buat input/tambahin komentar
func inputComment(A *tabComment, nData *int) {
	var text string
	fmt.Println("")
	fmt.Println("========================================================")
	fmt.Println("||		    TAMBAH KOMENTAR                   ||")
	fmt.Println("========================================================")
	fmt.Println("PETUNJUK:")
	fmt.Println("Untuk spasi saat penulisan komentar, gunakan tanda '_' |")
	fmt.Println("Jika ingin berhenti, Silahkan isi 0 di akhir           |")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("| 0. Back                                              |")
	fmt.Println("========================================================")
	fmt.Println("")
	fmt.Println("Masukkan Komentar : ")
	fmt.Scan(&text)
	for text != "0" && *nData < NMAX {
		A[*nData].text = text
		A[*nData].longcomment = len(A[*nData].text)
		A[*nData].sentimen = ""
		fmt.Scan(&text)
		*nData++
	}
	fmt.Println()
	fmt.Println("Komentar berhasil ditambahkan!")
	fmt.Println("Komentar yang di masukan :", *nData-1)
	outputComment(*A, *nData)
}

func outputComment(A tabComment, nData int) {
	var i int
	fmt.Println("")
	fmt.Println("===================================================")
	fmt.Println("||		  DATA KOMENTAR                  ||")
	fmt.Println("===================================================")
	fmt.Printf("| %s. | %-19s | %-19s |\n", "No", "Komentar", "Sentimen")
	fmt.Println("---------------------------------------------------")
	for i = 1; i < nData; i++ {
		if A[i].sentimen == "" {
			fmt.Printf("| %d.  | %-19s | %-19s |\n", i, A[i].text, "Belum Tersedia")
		} else {
			fmt.Printf("| %d.  | %-19s | %-19s |\n", i, A[i].text, A[i].sentimen)
		}
	}
	fmt.Println("---------------------------------------------------")
	fmt.Println()
}

// fungsi ini buat ubah komentar
func ubahKomentar(A *tabComment, nData int) {
	var idx int
	var komentarBaru string
	outputComment(*A, nData)
	fmt.Println("")
	fmt.Println("=======================================================")
	fmt.Println("||		    UBAH KOMENTAR                    ||")
	fmt.Println("=======================================================")
	fmt.Println("Index komentar yang telah terisi :(", nData-1, "komentar)      |")
	fmt.Println("-------------------------------------------------------")
	fmt.Println("0. Back                                               |")
	fmt.Println("=======================================================")
	fmt.Println()
	fmt.Println("Silakan masukkan index komentar yang ingin diubah :   ")
	fmt.Scan(&idx)
	fmt.Println("")
	fmt.Println("========================================================")
	fmt.Println("||		    TAMBAH KOMENTAR                   ||")
	fmt.Println("========================================================")
	fmt.Println("Untuk spasi saat penulisan komentar, gunakan tanda '_' |")
	fmt.Println("--------------------------------------------------------")
	if idx > nData || idx <= 0 {
		fmt.Println("Index tidak valid!")
	} else {
		fmt.Println("Komentar lama :", (*A)[idx].text)
		fmt.Println()
		fmt.Println("Masukkan Komentar :")
		fmt.Scan(&komentarBaru)
		(*A)[idx].text = komentarBaru
		(*A)[idx].longcomment = len(komentarBaru)
		(*A)[idx].sentimen = ""
		fmt.Println()
		fmt.Println("Komentar berhasil diubah.")
	}
	outputComment(tab, nData)
}

// fungsi ini buat hapus komentar
func hapusKomentar(A *tabComment, nData *int) {
	var IDKomen int

	outputComment(*A, *nData)

	// tampilan
	fmt.Println("==========================================================")
	fmt.Println("||		       HAPUS KOMENTAR                   ||")
	fmt.Println("==========================================================")
	fmt.Println("PETUNJUK :                                               |")
	fmt.Println("Mohon masukkan Index komentar yang ingin di hapus!       |")
	fmt.Println("Index komentar merupakan No. urutan pada daftar komentar |")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("0. Back                                                  |")
	fmt.Println("==========================================================")
	fmt.Println("ID komentar :")
	fmt.Scan(&IDKomen)
	if IDKomen <= 0 || IDKomen >= *nData {
		fmt.Println("Komentar tidak ditemukan")
	} else {
		for IDKomen <= *nData-2 {
			A[IDKomen] = A[IDKomen+1]
			IDKomen++
		}
		*nData--
		fmt.Println()
		fmt.Println("Komentar anda telah dihapus!")
		outputComment(*A, *nData)
	}

}

// codenya masih rusakkkk
// prosedur ini tu buat output
func analisisSentimen(A *tabComment, nData int) {
	var i, j int
	var x, y bool
	i = 0
	outputComment(*A, nData)
	for i < nData {
		x = false
		y = false
		for j = 0; j < len(komentarPositif); j++ {
			if cariKataKunci((*A)[i], komentarPositif[j]) {
				x = true
			}
		}
		for j = 0; j < len(komentarNegatif); j++ {
			if cariKataKunci((*A)[i], komentarNegatif[j]) {
				y = true
			}
		}
		if x == true {
			A[i].sentimen = "Positif"
			A[i].nilaiSentimen = 2
		} else if y == true {
			A[i].sentimen = "Negatif"
			A[i].nilaiSentimen = 0
		} else {
			A[i].sentimen = "Netral"
			A[i].nilaiSentimen = 1
		}
		i++
	}
	fmt.Println()
	fmt.Println("Berikut adalah hasil analisis sentimen dari komentar yang telah diinputkan :")
	outputComment(*A, nData)

}

func menuCariKataKunci(A tabComment, nData int) {
	var kataKunci string
	var i int
	var found bool
	outputComment(A, nData)
	// tampilan
	fmt.Println("Masukkan kata kunci : ")
	fmt.Scan(&kataKunci)
	for i = 0; i < nData; i++ {
		if cariKataKunci(A[i], kataKunci) == true {
			fmt.Println("Komentar berada di No :", i)
			fmt.Println(A[i].text)
			found = true
		}
	}
	if !found {
		fmt.Println("Komentar tidak ditemukan")
	}
	outputComment(A, nData)
}

// fungsi ini tu buat fitur nyari kata kunci
func cariKataKunci(A comment, kataKunci string) bool {
	var j int
	var subString string

	subString = ""
	for j = 0; j < A.longcomment; j++ {
		if A.text[j] != '_' && j != A.longcomment-1 {
			subString += string(A.text[j])
			if subString == kataKunci {
				return true
			}
		} else {
			if A.text[j] != '_' {
				subString += string(A.text[j])
			}
			if subString == kataKunci {
				return true
			}
			subString = ""
		}
	}
	return false
}

// fungsi ini buat ngurutin komentar berdasarkan panjang teksnya menggunakan selection sort
func urutkanPanjangKomentar(A *tabComment, nData, pilihan int) {
	var i, idx, pass int
	var temp comment
	pass = 2

	if pilihan == 1 {
		fmt.Println("Komentar Ascending sebelum diurutkan :")
		outputComment(*A, nData)
		for pass < nData {
			idx = pass - 1
			i = pass
			temp = A[pass]
			for i < nData {
				if A[i].longcomment < A[idx].longcomment {
					idx = i
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
		fmt.Println("Komentar Ascending setelah diurutkan : ")
	} else if pilihan == 2 {
		fmt.Println("Komentar Descending sebelum diurutkan :")
		outputComment(*A, nData)
		for pass < nData {
			idx = pass - 1
			i = pass
			temp = A[pass]
			for i < nData {
				if A[i].longcomment > A[idx].longcomment {
					idx = i
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
		fmt.Println("Komentar Desscending setelah diurutkan : ")
	}
}

// fungsi ini buat ngurutin berdasarkan sentimen komentarnya menggunakan insertion sort
func urutkanSentimenKomentar(A *tabComment, nData int) {
	var pass, i int
	var temp comment
	pass = 2
	analisisSentimen(*&A, nData)
	fmt.Println("Komentar Sentimen sebelum diurutkan :")
	outputComment(*A, nData)

	for pass < nData {
		i = pass
		temp = A[pass]
		for i > 1 && temp.nilaiSentimen > A[i-1].nilaiSentimen {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
	fmt.Println("Komentar Sentimen setelah diurutkan :")
}

func urutkanKomentar(A *tabComment, nData int) {
	// buat array comment
	var pilihan int
	urutkanPanjangKomentar(&tab, nData, pilihan)
	fmt.Println()
	fmt.Println("1. Berdasarkan Panjang Komentar")
	fmt.Println("2. Berdasarkan Sentimen Komentar")
	fmt.Println("Silahkan Pilih Tipe Pengurutan : ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilih urutan: ")
		fmt.Scan(&pilihan)
		urutkanPanjangKomentar(A, nData, pilihan)
	} else if pilihan == 2 {
		urutkanSentimenKomentar(A, nData)
	}
	outputComment(*A, nData)
}

func statistikSentimen(A *tabComment, nData int) {
	var positif, negatif, netral, i int
	for i = 1; i < nData; i++ {
		if A[i].sentimen == "Positif" {
			positif = positif + 1
		} else if A[i].sentimen == "Negatif" {
			negatif = negatif + 1
		} else if A[i].sentimen == "Netral" {
			netral = netral + 1
		}
	}
	fmt.Println()
	fmt.Println("Statistik Sentimen Komentar : ")
	fmt.Println("Positif:", positif, "komentar")
	fmt.Println("Netral :", netral, "komentar")
	fmt.Println("Negatif:", negatif, "komentar")
	outputComment(tab, nData)

}

//fungsi ini buat tampilan aplikasinya
func tampilanMenu(pilihan int) {
	for pilihan != 8 {
		fmt.Println("===================================================")
		fmt.Println("||		     MENU	              	 ||")
		fmt.Println("===================================================")
		fmt.Println("1. Tambah Komentar                                |")
		fmt.Println("2. Ubah Komentar                                  |")
		fmt.Println("3. Hapus Komentar                                 |")
		fmt.Println("4. Analisis Sentimen Komentar                     |")
		fmt.Println("5. Cari Komentar                                  |")
		fmt.Println("6. Urutkan Komentar (Panjang / Sentimen)          |")
		fmt.Println("7. Statistik Sentimen                             |")
		fmt.Println("--------------------------------------------------|")
		fmt.Println("8. Keluar                                         |")
		fmt.Println("==================================================|")
		fmt.Println()
		fmt.Println("Pilih Fitur : ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			inputComment(&tab, &nData)
		case 2:
			ubahKomentar(&tab, nData)
		case 3:
			hapusKomentar(&tab, &nData)
		case 4:
			analisisSentimen(&tab, nData)
		case 5:
			menuCariKataKunci(tab, nData)
		case 6:
			urutkanKomentar(&tab, nData)
		case 7:
			statistikSentimen(&tab, nData)
		case 8:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
		}
		if pilihan < 0 || pilihan > 8 {
			fmt.Println("Fitur tidak tersedia, silakan pilih kembali!")
		}
	}
}
func main() {
	var pilihan int = 0
	fmt.Println()
	fmt.Println("		Selamat Datang! 			")
	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("||	APLIKASI ANALISIS SENTIMEN KOMENTAR	 ||")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	tampilanMenu(pilihan)
}
