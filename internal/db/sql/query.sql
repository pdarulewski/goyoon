-- name: IncrementScore :exec
UPDATE scores
SET
    score = score + ?,
    played_rounds = played_rounds + 1
WHERE datetime = ?
    AND name = ?;
