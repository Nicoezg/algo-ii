package fnv

import "fmt"

const (
	uint64Offset uint64 = 0xcbf29ce484222325
	uint64Prime  uint64 = 0x00000100000001b3
)

func Hashing[K comparable](clave K, tam int) uint64 {
	return fnvHash[K](clave) % uint64(tam)
}

func convertiraBytes[K comparable](clave K) []byte {
	//Convierte a bytes la clave ingresada
	return []byte(fmt.Sprintf("%v", clave))
}

func fnvHash[K comparable](clave K) (hash uint64) {
	data := convertiraBytes(clave)
	hash = uint64Offset

	for _, b := range data {
		hash ^= uint64(b)
		hash *= uint64Prime
	}
	return hash
}
