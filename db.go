package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ipfs/go-datastore"
	levelds "github.com/ipfs/go-ds-leveldb"
	"github.com/mitchellh/go-homedir"
	ldbopts "github.com/syndtr/goleveldb/leveldb/opt"
)

var DB datastore.Batching

func setupLevelDs(path string, readonly bool) (datastore.Batching, error) {
	if _, err := os.ReadDir(path); err != nil {
		if os.IsNotExist(err) {
			//mkdir
			err = os.MkdirAll(path, 0777)
			if err != nil {
				return nil, err
			}
		}
	}

	db, err := levelds.NewDatastore(path, &levelds.Options{
		Compression: ldbopts.NoCompression,
		NoSync:      false,
		Strict:      ldbopts.StrictAll,
		ReadOnly:    readonly,
	})
	if err != nil {
		fmt.Printf("NewDatastore: %s\n", err)
		return nil, err
	}

	return db, err
}

func DataStores() {
	repodir, err := homedir.Expand(repoPath)
	if err != nil {
		return
	}

	idb, err := setupLevelDs(path.Join(repodir, "datastore"), false)
	if err != nil {
		log.Printf("setup infodb: err %s", err)
		return
	}
	DB = idb

	log.Printf("DB: %+v", DB)
}

// func saveMInfo(addr string, info ltypes.MinerInfo) error {
// 	key := datastore.NewKey(addr)

// 	// isHas, err := MinerInfoDB.Has(key)
// 	// if err != nil {
// 	// 	log.Infof("minfo: has %s", err)
// 	// 	return err
// 	// }

// 	// if !isHas {
// 	in, err := json.Marshal(info)
// 	if err != nil {
// 		return err
// 	}

// 	err = MinerInfoDB.Put(key, in)
// 	if err != nil {
// 		log.Infof("minfo: begin %s", err)
// 		return err
// 	}

// 	log.Infof("write minfo for addr: %s val %v", key.String(), info)
// 	// }

// 	return nil
// }

// func readMInfos() ([]ltypes.MinerInfo, error) {
// 	res, err := MinerInfoDB.Query(query.Query{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer res.Close()

// 	minfos := []ltypes.MinerInfo{}

// 	var errs error

// 	for {
// 		res, ok := res.NextSync()
// 		if !ok {
// 			break
// 		}

// 		if res.Error != nil {
// 			return nil, res.Error
// 		}

// 		minfo := &ltypes.MinerInfo{}
// 		err := json.Unmarshal(res.Value, minfo)
// 		if err != nil {
// 			errs = multierr.Append(errs, xerrors.Errorf("decoding state for key '%s': %w", res.Key, err))
// 			continue
// 		}

// 		minfos = append(minfos, *minfo)
// 	}

// 	log.Infof("read minfos ok, len %d", len(minfos))

// 	return minfos, nil
// }

// func readmInfo(maddr string) (*ltypes.MinerInfo, error) {
// 	key := datastore.NewKey(maddr)
// 	isHas, err := MinerInfoDB.Has(key)
// 	if err != nil {
// 		log.Infof("minfo: has %s", err)
// 		return nil, xerrors.Errorf("has err for %s err %s", key.String(), err.Error())
// 	}

// 	if !isHas {
// 		return nil, xerrors.Errorf("minfo not exist: %s", key.String())
// 	}

// 	res, err := MinerInfoDB.Get(key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	minfo := &ltypes.MinerInfo{}
// 	err = json.Unmarshal(res, minfo)
// 	if err != nil {
// 		return nil, xerrors.Errorf("unmarsal err %s", err.Error())
// 	}

// 	log.Infof("read minfo(%s) ok", key.String())

// 	return minfo, nil
// }
