package plants

import "database/sql"

type Store struct {
	Db *sql.DB
}

type Plant struct {
	Id int64 `json:"id"`
}

type State struct {
	Id            int64  `json:"id"`
	MoistureLevel string `json:"moistureLevel"`
	ServerTime    string `json:"serverTime"`
}

type SendData struct {
	Id            int64  `json:"id"`
	MoistureLevel string `json:"moistureLevel"`
	ServerTime    string `json:"serverTime"`
}

type Response struct {
	State State `json:"state`
}
