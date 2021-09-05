PRAGMA foreign_keys = false;
CREATE TABLE IF NOT EXISTS "database"
(
    "id"        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "show_name" TEXT,
    "driver"    TEXT,
    "host"      TEXT,
    "port"      TEXT,
    "name"      TEXT,
    "username"  TEXT,
    "password"  TEXT
);
PRAGMA foreign_keys = true;

PRAGMA foreign_keys = false;
CREATE TABLE IF NOT EXISTS "collection"
(
    "id"          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "parent_id"   INTEGER,
    "parent_path" TEXT,
    "label"       TEXT,
    "type"        TEXT,
    "sort_no"     INTEGER,
    "del_flag"    INTEGER
);
PRAGMA foreign_keys = true;