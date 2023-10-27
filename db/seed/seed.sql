-- +goose Up
-- Inserting data into the 'animes' table
INSERT INTO animes (name, img_url) VALUES
('Anime Title 1', 'http://example.com/path/to/image1.jpg'),
('Anime Title 2', 'http://example.com/path/to/image2.jpg'),
('Anime Title 3', 'http://example.com/path/to/image3.jpg');

-- Inserting data into the 'prefectures' table
INSERT INTO prefectures (name) VALUES
('Tokyo'),
('Osaka'),
('Kyoto');

-- Inserting data into the 'spots' table
INSERT INTO spots (name, anime_img_url, real_img_url, address, anime_id, prefecture_id) VALUES
('Spot Name 1', 'http://example.com/path/to/anime/image1.jpg', 'http://example.com/path/to/real/image1.jpg', '123 Anime Street, Tokyo', 1, 1),
('Spot Name 2', 'http://example.com/path/to/anime/image2.jpg', 'http://example.com/path/to/real/image2.jpg', '456 Anime Way, Osaka', 2, 2),
('Spot Name 3', 'http://example.com/path/to/anime/image3.jpg', 'http://example.com/path/to/real/image3.jpg', '789 Anime Avenue, Kyoto', 3, 3);

-- Inserting data into 'genders' table
INSERT INTO genders (gender) VALUES
('Male'),
('Female'),
('Other');

-- Inserting data into 'countries' table
INSERT INTO countries (name, img_url) VALUES
('Japan', 'url_to_image_japan'),
('United States', 'url_to_image_usa'),
('Canada', 'url_to_image_canada');

-- Assuming that you want to insert data into 'guides' with references to the above tables, you might do something like:
-- Please note, the gender_id and country_id should exist in 'genders' and 'countries' tables respectively.
INSERT INTO guides (name, gender_id, country_id, age, comment, stance, favorite_character) VALUES
('Guide One', 1, 1, 25, 'Comment One', 'Stance One', 'Character One'),
('Guide Two', 2, 2, 30, 'Comment Two', 'Stance Two', 'Character Two'),
('Guide Three', 3, 3, 35, 'Comment Three', 'Stance Three', 'Character Three');
