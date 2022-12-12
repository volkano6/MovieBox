DROP TABLE IF EXISTS actor_filmography;
DROP TABLE IF EXISTS movie_genres;
DROP TABLE IF EXISTS user_ratings;
DROP TABLE IF EXISTS user_watchlists;
DROP TABLE IF EXISTS user_watched;
DROP TABLE IF EXISTS user_comments;
DROP TABLE IF EXISTS user_favorites;
DROP TABLE IF EXISTS watchlist_movies;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS watchlists;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    displayname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    dateofbirth DATE NOT NULL,
    avatar VARCHAR(255),
    bio VARCHAR(1000),
    location VARCHAR(255),
    social_twitter VARCHAR(64),
    social_instagram VARCHAR(64),
    type TINYINT DEFAULT 0
);

CREATE TABLE movies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(1000) NOT NULL,
    release_date VARCHAR(50) NOT NULL,
    poster VARCHAR(255) NOT NULL,
    rating FLOAT(24),
    length INT NOT NULL
);

CREATE TABLE genres (
    genre VARCHAR(255) PRIMARY KEY
);

CREATE TABLE actors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    dateofbirth DATE NOT NULL,
    gender CHAR(1) NOT NULL,
    bio VARCHAR(1000) NOT NULL,
    `location` VARCHAR(255) NOT NULL
);

CREATE TABLE watchlists (
    id INT AUTO_INCREMENT PRIMARY KEY,
    watchlist_name VARCHAR(255) NOT NULL
);

CREATE TABLE watchlist_movies (
    watchlist_id INT,
    movie_id INT,
    PRIMARY KEY (watchlist_id, movie_id),
    CONSTRAINT watchlist_movies_watchlist_id_fk FOREIGN KEY watchlist_movies_watchlist_id_fk (watchlist_id) REFERENCES watchlists(id),
    CONSTRAINT watchlist_movies_movie_id_fk FOREIGN KEY watchlist_movies_movie_id_fk (movie_id) REFERENCES movies(id)
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
    rating INT NOT NULL,
    PRIMARY KEY (movie_id, user_id),
    CONSTRAINT user_ratings_movie_id_fk FOREIGN KEY user_ratings_movie_id_fk (movie_id) REFERENCES movies(id),
    CONSTRAINT user_ratings_user_id_fk FOREIGN KEY user_ratings_user_id_fk (user_id) REFERENCES users(id)
);

CREATE TABLE user_watchlist (
    watchlist_id INT,
    user_id INT,
    PRIMARY KEY (watchlist_id, user_id),
    CONSTRAINT user_watchlist_watchlist_id_fk FOREIGN KEY user_watchlist_watchlist_id_fk (watchlist_id) REFERENCES watchlists(id),
    CONSTRAINT user_watchlist_user_id_fk FOREIGN KEY user_watchlist_user_id_fk (user_id) REFERENCES users(id)
);

CREATE TABLE user_watched (
    movie_id INT,
    user_id INT,
    watched_date DATE NOT NULL,
    PRIMARY KEY (movie_id, user_id),
    CONSTRAINT user_watched_movie_id_fk FOREIGN KEY user_watched_movie_id_fk (movie_id) REFERENCES movies(id),
    CONSTRAINT user_watched_user_id_fk FOREIGN KEY user_watched_user_id_fk (user_id) REFERENCES users(id)
);

CREATE TABLE user_favorites (
    movie_id INT,
    user_id INT,
    favorited_date DATE NOT NULL,
    PRIMARY KEY (movie_id, user_id),
    CONSTRAINT user_favorited_movie_id_fk FOREIGN KEY user_watched_movie_id_fk (movie_id) REFERENCES movies(id),
    CONSTRAINT user_favoriteed_user_id_fk FOREIGN KEY user_watched_user_id_fk (user_id) REFERENCES users(id)
);

CREATE TABLE user_comments (
    movie_id INT,
    user_id INT,
    comment VARCHAR(1000) NOT NULL,
    comment_date DATE NOT NULL,
    PRIMARY KEY (movie_id, user_id),
    CONSTRAINT user_comments_movie_id_fk FOREIGN KEY user_comments_movie_id_fk (movie_id) REFERENCES movies(id),
    CONSTRAINT user_comments_user_id_fk FOREIGN KEY user_comments_user_id_fk (user_id) REFERENCES users(id)
);