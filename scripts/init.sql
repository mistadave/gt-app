CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    birthday DATE DEFAULT '0000-00-00',
    password TEXT NOT NULL,
    gender TEXT DEFAULT '',
    photo_url TEXT DEFAULT '',
    time INTEGER DEFAULT 0,
    active BOOLEAN NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS link_tags (
    link_id UUID,
    tag_id INTEGER,
    PRIMARY KEY (link_id, tag_id),
    FOREIGN KEY (link_id) REFERENCES links (id),
    FOREIGN KEY (tag_id) REFERENCES tags (id)
);

CREATE TABLE IF NOT EXISTS links (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    desc TEXT DEFAULT '',
    url TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    game_id UUID,
    FOREIGN KEY (game_id) REFERENCES games (id)
);

CREATE TABLE IF NOT EXISTS votes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    link_id INTEGER REFERENCES links(id),
    vote_type VARCHAR(10) CHECK(vote_type IN ('up', 'down')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    desc TEXT NOT NULL,
    image TEXT,
    genre TEXT NOT NULL,
    release_date TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Path: scripts/insert.sql