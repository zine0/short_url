-- name: GetUrl :one
select * from url
where key = $1 limit 1;

-- name: AddUrl :exec
insert into url (key, url)
values ($1,$2);