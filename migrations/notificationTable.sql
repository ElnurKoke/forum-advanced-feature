CREATE TABLE IF NOT EXISTS notification (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	post_id INT,
	author TEXT,
	reactauthor TEXT,
	message TEXT,
	activity INT DEFAULT 1,
	created_at DATE
);