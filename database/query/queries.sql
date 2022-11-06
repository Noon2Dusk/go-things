-- name: GetSpaceShipById :one
SELECT * FROM spaceship
WHERE id = ?
LIMIT 1;

-- name: GetSpaceShips :many
SELECT * FROM spaceship;

-- name: InsertSpaceship :exec
INSERT INTO spaceship (
                       name,
                       class,
                       crew,
                       image,
                       value,
                       status,
                       armaments
) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateSpaceship :exec
UPDATE spaceship
SET name = ?, class = ?, crew = ?, image = ?, value = ?, status = ?, armaments = ?
WHERE id = ?;

-- name: DeleteSpaceshipById :exec
DELETE FROM spaceship
WHERE id = ?;