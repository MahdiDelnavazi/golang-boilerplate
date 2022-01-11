-- create bucket stored procedure
CREATE OR REPLACE PROCEDURE create_bucket(name character varying)
language plpgsql
AS
$$
begin
INSERT INTO "bucket"(name)
VALUES(name);
end;
$$
