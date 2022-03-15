CREATE OR REPLACE FUNCTION newticket(UserId uuid, Subject varchar, Message varchar, Image varchar, Like1 boolean)
    RETURNS TABLE
            (
                Subject1 varchar,
                Message1 varchar,
                Image1 varchar,
                Like2 boolean
            )

AS
$$
BEGIN

    RETURN QUERY
        INSERT INTO "Ticket" ("UserId", "Subject", "Message", "Image", "Like")
            VALUES (UserId, Subject, Message, Image, Like1)
            RETURNING "Subject","Message","Image","Like";

END
$$
    LANGUAGE 'plpgsql';

