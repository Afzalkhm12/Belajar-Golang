package variable

import "fmt"

func Variable() {
    var Nama string = "Afzal"
	var usia string = "21"
	var univ string = "Universitas Nasional"
	var asal string = "Bekasi"

    fmt.Printf("halo %s %s!\n", Nama, usia)
	fmt.Printf("%s\n", univ)
	fmt.Printf("%s\n", asal)
}

