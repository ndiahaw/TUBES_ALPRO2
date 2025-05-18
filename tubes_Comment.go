package main

import "fmt"

const NMAX int = 100

type comment struct {
	longcomment    int
	text, sentimen string
}
type tabComment [NMAX]comment

var kataPositif = []string{"bagus", "indah", "cantik", "ganteng"}
var kataNegatif = []string{"jelek", "buruk", "menyebalkan", "benci"}
var tab tabComment
var nData int = 0
var kataKunci string

// prosedur ini tu buat input/tambahin komentar
func inputComment(A *tabComment, nData *int) {
	var text string
	fmt.Println("")
	fmt.Println("========================================================")
	fmt.Println("||		    TAMBAH KOMENTAR                   ||")
	fmt.Println("--------------------------------------------------------")
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
		*nData++
		fmt.Scan(&text)

	}
	fmt.Println("")
	fmt.Println("Komentar yang di masukan :", *nData)
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
	for i = 0; i < nData; i++ {
		if A[i].sentimen == "" {
			fmt.Printf("| %d.  | %-19s | %-19s |\n", i+1, A[i].text, "Belum Tersedia")
		} else {
			fmt.Printf("| %d.  | %-19s | %-19s |\n", i+1, A[i].text, A[i].sentimen)
		}
	}
	fmt.Println("---------------------------------------------------")
	fmt.Println()
}

// fungsi ini buat ubah komentar
func ubahKomentar() {

	fmt.Println("Silahkan isi Index dari komentar yang ingin diubah :")

}

// fungsi ini buat hapus komentar
func hapusKomentar() {

	fmt.Println("Silahkan isi Index dari komentar yang ingin dihapus :")

}

// prosedur ini tu buat output
func analisisSentimen(A *tabComment, nData int) {
	var i, j int
	var x, y bool
	i = 0
	for i < nData {
		for j = 0; j < len(kataPositif); j++ {
			if cariKataKunci((*A)[i], kataPositif[j]) {
				x = true
			}
		}
		for j = 0; j < len(kataNegatif); j++ {
			if cariKataKunci((*A)[i], kataNegatif[j]) {
				y = true
			}
		}
		if x == true {
			A[i].sentimen = "Positif"
		} else if y == true {
			A[i].sentimen = "Negatif"
		} else {
			A[i].sentimen = "Netral"
		}
		i++
	}
	fmt.Println("")
	fmt.Println("Berikut adalah hasil analisis sentimen dari komentar yang telah diinputkan :")
	outputComment(*A, nData)

}

func menuCariKataKunci(A tabComment, nData int) {
	var kataKunci string
	var i int
	fmt.Scan(&kataKunci)
	for i = 0; i < nData; i++ {
		if cariKataKunci(A[i], kataKunci) == true {
			fmt.Println(A[i].text)
		}
	}
	if cariKataKunci(A[i], kataKunci) == true {
		fmt.Println("Komentar tidak ditemukan")
	}

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
func urutkanPanjangKomentar(A *tabComment, nData int) {
	var i, idx, pass int
	var temp comment
	pass = 1
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
}

// fungsi ini buat ngurutin berdasarkan sentimen komentarnya menggunakan insertion sort
func urutkanSentimenKomentar(A *tabComment, nData int) {

}

func urutkanKomentar(A *tabComment, nData int) {
	// buat array comment
	var pilihan int
	urutkanPanjangKomentar(&tab, nData)
	fmt.Println("")
	fmt.Println("1. Berdasarkan Panjang Komentar")
	fmt.Println("2. Berdasarkan Sentimen Komentar")
	if pilihan == 1 {
		urutkanPanjangKomentar(&tab, nData)
	} else if pilihan == 2 {
		urutkanSentimenKomentar(&tab, nData)
	} else {

	}
	fmt.Println("")
	fmt.Println("Silahkan Pilih Tipe Pengurutan : ")
	fmt.Scan(&pilihan)
}

func statistikSentimen() {

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
		fmt.Println("5. Cari Komentar (Sequential)                     |")
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
			ubahKomentar()
		case 3:
			hapusKomentar()
		case 4:
			analisisSentimen(&tab, nData)
		case 5:
			menuCariKataKunci(tab, nData)
		case 6:
			urutkanKomentar(&tab, nData)
		case 7:
			statistikSentimen()
		case 8:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
		}
		if pilihan < 0 || pilihan > 8 {
			fmt.Println("Fitur tidak tersedia, silakan pilih kembali!")
			fmt.Println("")
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
