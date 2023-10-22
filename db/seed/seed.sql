-- +goose Up
INSERT INTO animes (name, img_url) VALUES
('Anime Title 1', 'http://example.com/path/to/image1.jpg'),
('Anime Title 2', 'http://example.com/path/to/image2.jpg'),
('Anime Title 3', 'http://example.com/path/to/image3.jpg');
