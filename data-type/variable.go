package main
import "fmt"

func main(){
	var name string
	var secondName = "Ansori" // tidak harus diberi tipe datanya karena langsung diinisialisasi
	// kita juga tidak diwajibkan memberi var
	lastName := "Majid"

	// tapi semua di atas masih mutable (masih bisa diinisiasi ulang)
	// atau bisa juga kaya gini
	var(
		namaPertama = "Mujahid"
		namaKedua = "Ansori"
		namaKetiga = "Majid"
	)

	name = "Mujahid"
	fmt.Println("my name is", name, secondName, lastName)
	fmt.Println("my name is", namaPertama, namaKedua, namaKetiga)
}