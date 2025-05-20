// fungsi ini buat ngurutin berdasarkan sentimen komentarnya menggunakan insertion sort
func urutkanSentimenKomentar(A *tabComment, nData int) {
	var pass, i int
	var temp comment
	pass = 2

	analisisSentimen(*&A, nData)

	fmt.Println("Komentar sebelum diurutkan :")
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
	fmt.Println("Komentar setelah diurutkan :")
	outputComment(*A, nData)
}
