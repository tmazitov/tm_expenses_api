\echo ************* Creating user, database *************

CREATE USER :"db_user" WITH PASSWORD :'db_pass';

CREATE DATABASE :"db_name" OWNER :"db_user";

\connect :"db_name"

CREATE SCHEMA :"db_schema" AUTHORIZATION :"db_user";

GRANT USAGE ON SCHEMA :"db_schema" TO :"db_user";
GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA :"db_schema" TO :"db_user";

ALTER DEFAULT PRIVILEGES IN SCHEMA :"db_schema"
    GRANT SELECT, INSERT, UPDATE ON TABLES TO :"db_user";

ALTER DATABASE :"db_name" SET search_path TO :"db_schema", public;
