select 
    u.id, u.fullname, u.username, u.email 
from 
    edufund.users u
where u.id = $1;