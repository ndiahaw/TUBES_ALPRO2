package main

import "fmt"

const NMAX int = 10

type comment struct {
	longcomment    int
	text, sentimen string
}
type tabComment [NMAX]comment

var tab tabComment
var nData int
var kataKunci string

// prosedur ini tu buat input/tambahin komentar
func inputComment(A *tabComment, nData *int) {
	var i int
	for i <= NMAX && A[i].text != "0" {
		fmt.Scan(&A[i].text)
		A[i].longcomment = len(A[i].text)
		*nData++
	}
}

// fungsi ini buat ubah komentar
func ubahKomentar() {

}

// fungsi ini buat hapus komentar
func hapusKomentar() {

}

// prosedur ini tu buat output
func analisisSentimen(A tabComment, n int) {

}

// fungsi ini tu buat fitur nyari kata kunci
func cariKataKunci(A *tabComment, nData int, kataKunci string) int {
	var i, j, k, found, longKey int
	var subString string
	found = -1
	longKey = len(kataKunci)

	for i = 0; i < nData; i++ {
		if A[i].text != "0" && found == -1 {
			for j = 0; j <= A[i].longcomment; j++ {
				subString = ""
				for k = j; k <= longKey; k++ {
					subString += string(A[i].text[k])
				}
				fmt.Println(subString)
				if subString == kataKunci {
					found = i
					return found
				}
			}
		} else if found != -1 {
			fmt.Print(i)
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

//fungsi ini buat tampilan aplikasinya
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
