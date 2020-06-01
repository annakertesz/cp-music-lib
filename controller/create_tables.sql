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
    id           SERIAL NOT NULL,
    album_name   varchar(150),
    album_artist SERIAL REFERENCES artist (id),
    album_cover  INT,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS song
(
    id                  SERIAL NOT NULL,
    song_name           varchar(150),
    song_album          INTEGER REFERENCES album (id),
    song_tag            INTEGER REFERENCES tag (id),
    song_lq_url         varchar(500),
    song_hq_url         varchar(500),
    instrumental_lq_url varchar(500),
    instrumental_hq_url varchar(500),
    PRIMARY KEY (id)
);



CREATE TABLE IF NOT EXISTS tag_song
(
    id       SERIAL NOT NULL,
    map_tag  INTEGER REFERENCES tag (id),
    map_song INTEGER REFERENCES song (id),
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS cp_update
(
    id            SERIAL NOT NULL,
    ud_date       DATE,
    found_songs   INTEGER,
    created_songs INTEGER,
    failed_songs  INTEGER,
    deleted_songs INTEGER,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS failed_song
(
    id            SERIAL NOT NULL,
    box_id        varchar(500),
    error_message varchar(500),
    cp_update     INTEGER REFERENCES cp_update (id),
    PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS cp_user
(
    id            SERIAL NOT NULL,
    username      varchar(150),
    first_name    varchar(150),
    last_name     varchar(150),
    email         varchar(150),
    password_hash varchar(150),
    phone         varchar(500),
    user_status   int,
    PRIMARY KEY (id)
);

