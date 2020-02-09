package testlib

import "fmt"

var pop map[string]string

func init() {
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallin"
	pop["John Legend"] = "All of me"
}

// GetMusic: Popular music by singer 외부 호출 가능
func GetMusic(singer string) string {
	return pop[singer]
}

// 내부에서만 호출 가능
func getKeys() {
	for _, kv := range pop {
		fmt.Println(kv)
	}
}
