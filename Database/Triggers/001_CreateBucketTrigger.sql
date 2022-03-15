CREATE
OR REPLACE FUNCTION ticket_insert_trigger_fnc()
    RETURNS trigger AS

$$

BEGIN

UPDATE "User"
SET "TicketCount" = "TicketCount" + 1
WHERE NEW."UserId" = "Id"
    RETURNING *;

END;

$$
LANGUAGE 'plpgsql';

-- --------------------------------------------------------

CREATE TRIGGER ticket_insert_trigger

    AFTER INSERT

    ON "Ticket"

    FOR EACH ROW
    EXECUTE PROCEDURE ticket_insert_trigger_fnc();