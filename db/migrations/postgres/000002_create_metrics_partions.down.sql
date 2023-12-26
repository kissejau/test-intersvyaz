DO $$ 
DECLARE
    i INT;
BEGIN
    FOR i IN 1..31 LOOP
        EXECUTE FORMAT('DROP TABLE IF EXISTS metrics_%s;', i);
    END LOOP;
END $$;