CREATE TABLE IF NOT EXISTS Bucket (
    Id          uuid            NOT NULL            DEFAULT gen_random_uuid()  PRIMARY KEY,
    Name        varchar(255)    NOT NULL            UNIQUE,
    Active      Boolean         DEFAULT TRUE,
    CreatedAt   Date            Default now(),
    UpdatedAt   Date            Default null,
    DeletedAt   Date            Default null
)
