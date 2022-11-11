package main
import "fmt"

func main(){

	// slice itu sama kaya array namun bisa berubah
	// dalam slice ada 3
	// pointer, length, dan capacity
	// length banyaknya element di slice
	// capacity panjang slicenya (length tidak boleh lebih dari capacity)

	// contoh
	// kalo emang gatau berapa jumlah element di array bisa pake ...
	var months = [12]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	slice := months[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// slice2 := months[2:4]
	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "cobaAppend") // karena ketika kita ubah capacitynya masih ada maka ikut terubah
	// jika capacitynya sudah penuh maka akan membuat array baru dan tidak akan mengubah array origin

	fmt.Println(slice3)
	slice3[0] = "cobaUbah"

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// making newslice from scracth
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Mujahid"
	newSlice[1] = "Ansori"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int {1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}