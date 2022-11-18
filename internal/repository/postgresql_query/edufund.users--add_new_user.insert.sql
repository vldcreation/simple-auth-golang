insert into
    edufund.users(fullname, username, email, "password")
values
    ($1,$2,$3,$4)
RETURNING id;