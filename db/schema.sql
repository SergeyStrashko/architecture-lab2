
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

-- Insert demo data.
INSERT INTO PlantsList (id) VALUES (1);
INSERT INTO PlantsState 
  (moistureLevel, serverTime, plantid) 
VALUES 
  ( 0.1, CURRENT_TIMESTAMP, 
    (SELECT id FROM PlantsList WHERE id = 1)
  );

INSERT INTO PlantsList (id) VALUES (2);
INSERT INTO PlantsState 
  (moistureLevel, serverTime, plantid) 
VALUES 
  ( 0.3, CURRENT_TIMESTAMP, 
    (SELECT id FROM PlantsList WHERE id = 2)
  );
