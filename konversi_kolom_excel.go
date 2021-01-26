//fungsi konversi nilai integer ke index column excel
func konversi_int2idx_column(angka int) string {
	int_to_huruf := make(map[int] string)
	int_to_huruf[0] = "Z"
	int_to_huruf[1] = "A"
	int_to_huruf[2] = "B"
	int_to_huruf[3] = "C"
	int_to_huruf[4] = "D"
	int_to_huruf[5] = "E"
	int_to_huruf[6] = "F"
	int_to_huruf[7] = "G"
	int_to_huruf[8] = "H"
	int_to_huruf[9] = "I"
	int_to_huruf[10] = "J"
	int_to_huruf[11] = "K"
	int_to_huruf[12] = "L"
	int_to_huruf[13] = "M"
	int_to_huruf[14] = "N"
	int_to_huruf[15] = "O"
	int_to_huruf[16] = "P"
	int_to_huruf[17] = "Q"
	int_to_huruf[18] = "R"
	int_to_huruf[19] = "S"
	int_to_huruf[20] = "T"
	int_to_huruf[21] = "U"
	int_to_huruf[22] = "V"
	int_to_huruf[23] = "W"
	int_to_huruf[24] = "X"
	int_to_huruf[25] = "Y"
	int_to_huruf[26] = "Z"
	
	angka_pengganti := angka
	var array_huruf []string
	nilai_kembalian := ""
	
	if angka <= 26 {
		return int_to_huruf[angka]
	} else {
		for angka_pengganti > 26 {
			huruf_baru := angka_pengganti%26
			array_huruf = append(array_huruf, int_to_huruf[huruf_baru])
			if huruf_baru == 0 {
				angka_pengganti -= 26
			} else {
				angka_pengganti -= huruf_baru
			}
			angka_pengganti /= 26
			fmt.Println(angka_pengganti)
		}
		array_huruf = append(array_huruf, int_to_huruf[angka_pengganti])
		for j := range array_huruf {
			nilai_kembalian += array_huruf[len(array_huruf)-j-1]
		}
		return nilai_kembalian
	}
}
