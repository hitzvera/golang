package main
import "fmt"

func main(){
	// di go cuma ada satu doang yaitu for
	// engga ada while dan do while

	for i := 1; i <= 10; i++ {
		fmt.Print(i)
	}

	slices := []string{"Mujahid", "ansori", "Majid"}

	for i := 0; i<len(slices); i++ {
		fmt.Print(slices[i])
	}

	fmt.Println()
	// atau bisa juga seperti ini
	for index, item := range slices {
		fmt.Println("index - ",index, " value = ", item)
	} // jika kita tidak akan menggunakan index, bisa diberikan ' _ '

	// bisa juga diberikan kepada map
	person := map[string]string{
		"name":"Mujahid",
		"title":"Programmer",
	}

	for key, value := range person{
		fmt.Println("name ",value, "key ",key)
	}

	// break = perulangan keseluruhannya dihentikan
	// continue = berhenti di current perulangan, lanjut ke perulagnan selanjutnya

	// program anti nomor 5
	breakSlice := []int {2,3,1,12,3,5,5,3,4,32,3}
	for i := 0;i < len(breakSlice); i++ {
		fmt.Print(breakSlice[i])
		if(i == 5){
			break
		}
	}
	fmt.Println()

	// program menampilkan item ganjil
	for _, item := range breakSlice {
		if(item % 2 == 0){
			continue
		}
		fmt.Print(item)
	}


}