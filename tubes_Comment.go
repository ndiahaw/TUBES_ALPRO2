package main

import "fmt"

const NMAX int = 100

type comment struct {
	longcomment    int
	text, sentimen string
}
type tabComment [NMAX]comment

var tab tabComment
var nData int = 0

var kataKunci string

// prosedur ini tu buat input/tambahin komentar
func inputComment(A *tabComment, nData *int) {
	var i int = 0
	var text string
	fmt.Scan(&text)
	for text != "0" && i < NMAX {
		A[i].text = text
		A[i].longcomment = len(A[i].text)
		fmt.Scan(&text)
		*nData++
		i++
	}
}

// fungsi ini buat ubah komentar
func ubahKomentar() {

}

// fungsi ini buat hapus komentar
func hapusKomentar() {

}

// prosedur ini tu buat analisis sentimennya berdasarkan kata kunci positif & negatif yang udah di deklarasi di array sebelumnya
func analisisSentimen(A *tabComment, nData int) {
	var i, j int
	var isPositif, isNegatif bool
	isPositif = false
	isNegatif = false
	for i < nData {
		for j = 0; j < len(kataPositif); j++ {
			if isNegatif && cariKataKunci(*A, nData, kataPositif[j]) != -1 {
				isPositif = true
				A[i].sentimen = "Positif"
			}
			if isPositif && cariKataKunci(*A, nData, kataNegatif[j]) != -1 {
				isNegatif = true
				A[i].sentimen = "Negatif"
			} else {
				A[i].sentimen = "Netral"
			}
		}
		i++
	}
	outputComment(tab, nData)
}

// fungsi ini tu buat fitur nyari kata kunci
func cariKataKunci(A *tabComment, nData int, kataKunci string) int {
	var i, j, found int
	var subString string
	found = -1

	subString = ""

	for i = 0; i < nData; i++ {
		if A[i].text != "0" {
			subString = ""
			for j = 0; j < A[i].longcomment; j++ {
				if A[i].text[j] != '_' && j != A[i].longcomment-1 {
					subString += string(A[i].text[j])
					if subString == kataKunci {
						found = i
						fmt.Print(found)
						return found
					}
				} else if A[i].text[j] != '_' && j == A[i].longcomment-1 {
					subString += string(A[i].text[j])
					if subString == kataKunci {
						found = i
						fmt.Print(found)
						return found
					}
					subString = ""
				} else {
					if subString == kataKunci {
						found = i
						fmt.Print(found)
						return found
					}
					subString = ""
				}
			}
		}
	}
	return found
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

// fungsi ini buat tampilan aplikasinya
func tampilanInput() {

	pilihanTampilanInput()
}

// fungsi ini buat nampilin ke tampilan selanjutnya berdasarkan pilihan
func pilihanTampilanInput() {

}

func main() {
	inputComment(&tab, &nData)
	fmt.Scan(&kataKunci)
	cariKataKunci(&tab, nData, kataKunci)
}
