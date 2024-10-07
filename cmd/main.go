package main

import (
	"fmt"
	"livecode/internal/pkg/storage"
	"log"
)

func main() {
	st, err := storage.NewStorage()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	st.Set("key2", "222")
	rlt2 := st.Get("key2")
	fmt.Println(*rlt2)
}
