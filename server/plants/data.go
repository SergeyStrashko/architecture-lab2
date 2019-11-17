package plants

import (
	"database/sql"
)

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

// GET request
func (s *Store) getData(plant *Plant) (*Response, error) {
	// Init response
	var response Response
	response.State.Id = plant.Id
	row := s.Db.QueryRow(`SELECT id FROM PlantsList WHERE id=$1`, plant.Id)
	err := row.Scan(&response.State)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return &response, nil
	}

	rows, err := s.Db.Query(`
		SELECT 
			id, moistureLevel, servertime 
		FROM plantsState 
			WHERE id=$1`, response.State.Id)
	if err != nil {
		return nil, err
	}
	// Proccesing states
	defer rows.Close()
	var res []*State
	for rows.Next() {
		var state State
		if err := rows.Scan(&state.Id, &state.MoistureLevel, &state.ServerTime); err != nil {
			return nil, err
		}
		res = append(res, &state)
	}
	if res == nil {
		res = make([]*State, 0)
	}
	// Sending response
	return &response, nil
}

// Post request
func (s *Store) setData(sdata *SendData) error {
	var id int64
	row := s.Db.QueryRow(`SELECT id FROM PlantsList WHERE id = $1`, sdata.Id)
	err := row.Scan(&id)
	// If it doesnt, so create it
	if err != sql.ErrNoRows && err != nil {
		return err
	} else if err == sql.ErrNoRows {
		_, err := s.Db.Exec(`INSERT INTO PlantsList ("id") VALUES ($1)`, sdata.Id)
		if err != nil {
			return err
		}
		row = s.Db.QueryRow(`SELECT id FROM PlantsList WHERE id = $1`, sdata.Id)
	}
	// Insert Data linked to table
	_, err = s.Db.Exec(`
  INSERT INTO PlantsState
   ("moistureLevel", "serverTime", "plantid")
  VALUES
   ($1, CURRENT_TIMESTAMP, $3)`,
		sdata.MoistureLevel, sdata.ServerTime, id)
	return err
}
