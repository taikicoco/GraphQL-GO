-- +goose Up
CREATE TABLE IF NOT EXISTS spots (
    spot_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    anime_img_url TEXT,
    real_img_url TEXT,
    address VARCHAR(255) NOT NULL, 
    anime_id INT NOT NULL, 
    prefecture_id INT NOT NULL, 
    FOREIGN KEY (anime_id) REFERENCES animes(anime_id),
    FOREIGN KEY (prefecture_id) REFERENCES prefectures(prefecture_id)
);
