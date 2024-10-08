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

	st.Set("key1", "11")
	st.Set("key2", "two")
	st.Set("key3", "")
	rlt1 := st.GetKind("key1")
	rlt2 := st.GetKind("key2")
	rlt3 := st.Get("key2")
	fmt.Println(rlt1, rlt2, *rlt3)
}
