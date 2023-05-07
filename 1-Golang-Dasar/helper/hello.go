package helper

var application = "Belajar Golang"
var Version = "Versi 1.1"

func SayHello(name string) string {
	return "Halo " + name
}

// Tidak bisa diakses ke luar karena nama function huruf kecil
func sayGoodBye(name string) string {
	return "Halo " + name
}