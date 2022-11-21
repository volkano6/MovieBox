DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS actor_filmography;
DROP TABLE IF EXISTS movie_genres;
DROP TABLE IF EXISTS user_ratings;
DROP TABLE IF EXISTS user_watchlist;
DROP TABLE IF EXISTS user_watched;
DROP TABLE IF EXISTS user_comments;

CREATE TABLE users (
    id INT,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    displayname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    dateofbirth DATE NOT NULL,
    avatar VARCHAR(255),
    bio VARCHAR(1000),
    location VARCHAR(255),
    social_twitter VARCHAR(64),
    social_instagram VARCHAR(64),
    type TINYINT,
    PRIMARY KEY (id)
);

CREATE TABLE movies (
    id INT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(1000) NOT NULL,
    release_date DATE NOT NULL,
    poster VARCHAR(255),
    rating FLOAT(24),
    length INT,
    trailer VARCHAR(255)
);

CREATE TABLE genres (
    genre VARCHAR(255) PRIMARY KEY
);

CREATE TABLE actors (
    id INT PRIMARY KEY,
    fullname VARCHAR(255),
    dateofbirth DATE,
    gender CHAR(1),
    bio VARCHAR(1000),
    `location` VARCHAR(255)
);

CREATE TABLE movie_genres (
    movie_id INT,
    genre_name VARCHAR(255),
    PRIMARY KEY (movie_id, genre_name),
    CONSTRAINT movie_genres_movie_id_fk FOREIGN KEY movie_genres_movie_id_fk (movie_id) REFERENCES movies(id),
    CONSTRAINT movie_genres_genre_name_fk FOREIGN KEY movie_genres_genre_name_fk (genre_name) REFERENCES genres(genre)
);

CREATE TABLE actor_filmography (
    movie_id INT,
    actor_id INT,
    PRIMARY KEY (movie_id, actor_id),
    CONSTRAINT actor_filmography_movie_id_fk FOREIGN KEY actor_filmography_movie_id_fk (movie_id) REFERENCES movies(id),
    CONSTRAINT actor_filmography_actor_id_fk FOREIGN KEY actor_filmography_actor_id_fk (actor_id) REFERENCES actors(id)
);

CREATE TABLE user_ratings (
    movie_id INT,
    user_id INT,
    rating INT,
    PRIMARY KEY (movie_id, user_id),
    CONSTRAINT user_ratings_movie_id_fk FOREIGN KEY user_ratings_movie_id_fk (movie_id) REFERENCES movies(id),
    CONSTRAINT user_ratings_user_id_fk FOREIGN KEY user_ratings_user_id_fk (user_id) REFERENCES users(id)
);

CREATE TABLE user_watchlist (
    movie_id INT,
    user_id INT,
    added_date DATE,
    PRIMARY KEY (movie_id, user_id),
    CONSTRAINT user_watchlist_movie_id_fk FOREIGN KEY user_watchlist_movie_id_fk (movie_id) REFERENCES movies(id),
    CONSTRAINT user_watchlist_user_id_fk FOREIGN KEY user_watchlist_user_id_fk (user_id) REFERENCES users(id)
);

CREATE TABLE user_watched (
    movie_id INT,
    user_id INT,
    watch_date DATE,
    PRIMARY KEY (movie_id, user_id),
    CONSTRAINT user_watched_movie_id_fk FOREIGN KEY user_watched_movie_id_fk (movie_id) REFERENCES movies(id),
    CONSTRAINT user_watched_user_id_fk FOREIGN KEY user_watched_user_id_fk (user_id) REFERENCES users(id)
);

CREATE TABLE user_comments (
    movie_id INT,
    user_id INT,
    comment VARCHAR(1000),
    PRIMARY KEY (movie_id, user_id),
    CONSTRAINT user_comments_movie_id_fk FOREIGN KEY user_comments_movie_id_fk (movie_id) REFERENCES movies(id),
    CONSTRAINT user_comments_user_id_fk FOREIGN KEY user_comments_user_id_fk (user_id) REFERENCES users(id)
);