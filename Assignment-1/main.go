package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Id         string
	Name       string
	Address    string
	Occupation string
	Reason     string
}

var friends []Friend = []Friend{
	{
		Id:         "C1",
		Name:       "Muhammad Arief",
		Address:    "Jalan Beton 3",
		Occupation: "Back End Engineer",
		Reason:     "Suka sama golang",
	},
	{
		Id:         "C2",
		Name:       "Siti Nurhaliza",
		Address:    "Jalan Cipinang Melayu",
		Occupation: "Data Science",
		Reason:     "Penasaran sama golang",
	},
	{
		Id:         "C3",
		Name:       "Agus Senja",
		Address:    "Jalan Tebet",
		Occupation: "Web Program",
		Reason:     "Memantabkan ilmu golang sebelumnya",
	},
	{
		Id:         "C4",
		Name:       "Sri Ayu",
		Address:    "Jalan Slamet Riyadi",
		Occupation: "Front End Engineer",
		Reason:     "Ingin mendalami golang",
	},
	{
		Id:         "C5",
		Name:       "Ahmad Ropi",
		Address:    "Jalan Taman Aster",
		Occupation: "Mahasiswa",
		Reason:     "Tertarik dengan golang",
	},
}

func searchFriendById(id string) {
	for _, v := range friends {
		if v.Id == id {
			fmt.Printf("id: %s\n", v.Id)
			fmt.Printf("nama: %s\n", v.Name)
			fmt.Printf("alamat: %s\n", v.Address)
			fmt.Printf("pekerjaan: %s\n", v.Occupation)
			fmt.Printf("alasan pilih kelas Golang: %s\n", v.Reason)
			return
		}

	}

	fmt.Printf("Teman id %s tidak ditemukan\n", id)

}

func main() {
	// Mohon berikan validasi ketika user tidak memberikan argument apapun
	// Mohon diberikan warning message ketika user tidak memberikan argument apapun.
	// Jangan sampai terkena panic error.

	var arguments []string = os.Args

	if len(arguments) < 2 {
		fmt.Println("Tolong beri argument, menjadi:\n go run main.go <no id>")
		return
	}

	searchFriendById(arguments[1])

}
