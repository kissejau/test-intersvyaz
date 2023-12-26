DO $$ 
DECLARE
    i INT;
BEGIN
    FOR i IN 1..31 LOOP
        EXECUTE FORMAT('CREATE TABLE metrics_%s (LIKE metrics INCLUDING ALL);', i);
        EXECUTE FORMAT('ALTER TABLE metrics_%s INHERIT metrics;', i);
    END LOOP;
END $$;