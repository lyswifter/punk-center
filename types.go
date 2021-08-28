package main

type ResultPunk struct {
	IP        string `json:"ip"`
	Wal       string `json:"wal"`
	TokenID   string `json:"tokenID"`
	WalPubkey string `json:"walPubKey"`
	WalPriKey string `json:"walPriKey"`
}

type StatuePunk struct {
	IP     string `json:"ip"`
	Statue string `json:"status"`
}
