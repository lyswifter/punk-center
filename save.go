package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ipfs/go-datastore"
)

func saveInfo(info ResultPunk) error {
	key := datastore.NewKey(fmt.Sprintf("%s-%s-%s", info.IP, info.Wal, info.TokenID))
	// ishas, err := DB.Has(key)
	// if err != nil {
	// 	log.Printf("entry: has %s", err)
	// 	return err
	// }

	// if !ishas {
	in, err := json.Marshal(info)
	if err != nil {
		return err
	}

	err = DB.Put(key, in)
	if err != nil {
		log.Fatalf("entry: begin %s", err)
		return err
	}

	log.Printf("write info for: %s ok", key.String())
	// }

	return nil
}
