CREATE TABLE IF NOT EXISTS artist
(
    id          INTEGER NOT NULL,
    artist_name varchar(150),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS album
(
    id                  INTEGER NOT NULL,
    album_name          varchar(150),
    artist              INTEGER REFERENCES artist (id),
    cover_url           varchar(500),
    cover_thumbnail_url varchar(500),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS song
(
    id                  INTEGER NOT NULL,
    song_name           varchar(150),
    artist              INTEGER REFERENCES artist (id),
    album               INTEGER REFERENCES album (id),
    song_lq_url         varchar(500),
    song_hq_url         varchar(500),
    instrumental_lq_url varchar(500),
    instrumental_hq_url varchar(500),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tag
(
    id       INTEGER NOT NULL,
    tag_name varchar(150),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tag_song
(
    id   INTEGER NOT NULL,
    tag  INTEGER REFERENCES tag (id),
    song INTEGER REFERENCES song (id),
    PRIMARY KEY (id)
);

