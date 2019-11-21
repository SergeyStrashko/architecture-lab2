package plants

import "database/sql"

type Store struct {
	Db *sql.DB
}

type Plant struct {
	Id int64 `json:"id"`
}

type State struct {
	Id            int64   `json:"id"`
	MoistureLevel float64 `json:"moistureLevel"`
	ServerTime    string  `json:"soilDataTimestamp"`
}

type SendData struct {
	Id            int64   `json:"id"`
	MoistureLevel float64 `json:"moistureLevel"`
	ServerTime    string  `json:"soilDataTimestamp"`
}

type Response struct {
	State []*State `json:"res"`
}
