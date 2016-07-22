// vi:filetype=sql

BEGIN TRANSACTION;
  CREATE TABLE sessions(
    key string,
    data blob,
    created_on time,
    updated_on time,
    expires_on time);
COMMIT;


