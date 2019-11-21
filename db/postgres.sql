
CREATE TABLE PlantsList
(
  id   SERIAL PRIMARY KEY,
);

CREATE TABLE PlantsState
(
    id            SERIAL PRIMARY KEY,
    moistureLevel       FLOAT,
    serverTime    timestamp,
    plantid      INTEGER,
    FOREIGN KEY (plantid) REFERENCES PlantsList(id)
);

