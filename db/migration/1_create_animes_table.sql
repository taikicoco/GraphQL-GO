-- +goose Up
CREATE TABLE IF NOT EXISTS animes (
    anime_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    img_url TEXT
);
