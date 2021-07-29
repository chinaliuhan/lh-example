package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	for i, r := range "Hello, 世界" {
		log.Printf("%d\t%q\t%d\n", i, r, r)
	}
	/**
	结果
	main.go:11: 0   'H'     72
	main.go:11: 1   'e'     101
	main.go:11: 2   'l'     108
	main.go:11: 3   'l'     108
	main.go:11: 4   'o'     111
	main.go:11: 5   ','     44
	main.go:11: 6   ' '     32
	main.go:11: 7   '世'    19990
	main.go:11: 10  '界'    30028

	*/
}
