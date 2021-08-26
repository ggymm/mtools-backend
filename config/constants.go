package config

const QueryTableListSQL = `SELECT DISTINCT TABLE_NAME, TABLE_COMMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA = ?`
