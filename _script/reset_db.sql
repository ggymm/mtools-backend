DELETE
FROM database;

DELETE
FROM sqlite_sequence
WHERE name = 'database';

INSERT INTO "database"
VALUES (1, 'office_platform(本地)', 'mysql', 'localhost', 3306, 'office_platform', 'root', 'root');

INSERT INTO "database"
VALUES (2, 'ninelock(本地)', 'mysql', 'localhost', 3306, 'ninelock', 'root', 'root');

INSERT INTO "database"
VALUES (4, 'develop-platform(本地)', 'mysql', 'localhost', 3306, 'develop-platform', 'root', 'root');

INSERT INTO "database"
VALUES (5, 'outsourcing_new(本地)', 'mysql', 'localhost', 3306, 'outsourcing_new', 'root', 'root');

UPDATE "sqlite_sequence"
SET seq = 5
WHERE name = 'database';

DELETE
FROM collection;

DELETE
FROM sqlite_sequence
WHERE name = 'collection';