// vi:filetype=sql

BEGIN TRANSACTION;
  CREATE TABLE sessions(
    id int,
    data string,
    created_on time,
    updated_on time,
    expires_on time);
COMMIT;


