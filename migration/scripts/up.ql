// vi:filetype=sql

BEGIN TRANSACTION;
  CREATE TABLE IF NOT EXISTS sessions(
    key string,
    data blob,
    created_on time,
    updated_on time,
    expires_on time);
  CREATE TABLE IF NOT EXISTS tasks(
    uuid string,
    done bool,
    user_id int,
    project_id int,
    created_at time,
    updated_at time);
  CREATE TABLE IF NOT EXISTS users(
    username string,
    password blob,
    email string,
    created_at time,
    updated_at time);
  CREATE TABLE IF NOT EXISTS projects(
    user_id int,
    name string,
    created_at time,
    updated_at time);
  
  CREATE UNIQUE INDEX UQE_users on users (username,email);
  CREATE UNIQUE INDEX UQE_tasks on tasks (uuid);
  CREATE UNIQUE INDEX UQE_sessions on sessions (key);
COMMIT;


