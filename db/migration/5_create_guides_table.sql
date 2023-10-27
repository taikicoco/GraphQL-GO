-- +goose Up
CREATE TABLE guides (
    guide_id INT AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    gender_id INT NOT NULL,
    country_id INT NOT NULL,
    age INT NOT NULL,
    comment TEXT,
    stance TEXT,
    favorite_character VARCHAR(255),
    PRIMARY KEY (guide_id),
    FOREIGN KEY (gender_id) REFERENCES genders(gender_id),
    FOREIGN KEY (country_id) REFERENCES countries(country_id)
);
