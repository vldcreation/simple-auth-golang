SELECT
    u."id",
    u."fullname",
    u."username",
    u."email",
    u."password"
FROM
    edufund.users u
where 
    u.username = $1 
or 
    u.email = $2
LIMIT 1;
