CREATE TABLE IF NOT EXISTS artist
(
    id          SERIAL NOT NULL,
    artist_name varchar(150),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tag
(
    id       SERIAL NOT NULL,
    tag_name varchar(150),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS album
(
    id                  SERIAL NOT NULL,
    album_name          varchar(150),
    album_artist              SERIAL REFERENCES artist (id),
    cover_url           varchar(500),
    cover_thumbnail_url varchar(500),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS song
(
    id                  SERIAL NOT NULL,
    song_name           varchar(150),
    song_artist              INTEGER REFERENCES artist (id),
    song_album               INTEGER REFERENCES album (id),
    song_tag              INTEGER REFERENCES tag (id),
    song_lq_url         varchar(500),
    song_hq_url         varchar(500),
    instrumental_lq_url varchar(500),
    instrumental_hq_url varchar(500),
    PRIMARY KEY (id)
);



CREATE TABLE IF NOT EXISTS tag_song
(
    id   SERIAL NOT NULL,
    map_tag  INTEGER REFERENCES tag (id),
    map_song INTEGER REFERENCES song (id),
    PRIMARY KEY (id)
);

