CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT UNIQUE,
	username TEXT UNIQUE,
	password TEXT,
	session_token TEXT DEFAULT NULL,
	expiresAt DATETIME DEFAULT NULL
);