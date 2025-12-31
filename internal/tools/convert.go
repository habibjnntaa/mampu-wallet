package tools

import "strconv"

func FormatRupiah(amount int) string {
	// Konversi angka ke string
	str := strconv.Itoa(amount)
	var result string
	length := len(str)

	for i := 0; i < length; i++ {
		// Sisipkan titik setiap 3 digit dari kanan
		if i > 0 && (length-i)%3 == 0 {
			result += "."
		}
		result += string(str[i])
	}

	return "Rp. " + result
}
