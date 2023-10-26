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
