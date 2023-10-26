-- +goose Up
CREATE TABLE IF NOT EXISTS prefectures (
    prefecture_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
