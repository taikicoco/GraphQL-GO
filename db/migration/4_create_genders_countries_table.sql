-- +goose Up
CREATE TABLE IF NOT EXISTS genders (
    gender_id INT AUTO_INCREMENT,
    gender VARCHAR(255) NOT NULL,
    PRIMARY KEY (gender_id)
);

CREATE TABLE IF NOT EXISTS countries (
    country_id INT AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    img_url TEXT,
    PRIMARY KEY (country_id)
);
