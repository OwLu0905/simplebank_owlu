-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency
  ) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
select *
from accounts
where id = $1
limit 1
;

-- name: GetAccountForUpdate :one
select *
from accounts
where id = $1
limit 1
for no key update
;

-- name: ListAccounts :many
select *
from accounts
WHERE owner = $1
order by id
limit $2
offset $3
;

-- name: UpdateAccount :one
UPDATE accounts 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts 
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
delete from accounts
where id = $1
;
