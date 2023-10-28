-- +goose Up
CREATE TABLE  IF NOT EXISTS guide_animes (
    guide_id INT,
    anime_id INT,
    PRIMARY KEY (guide_id, anime_id),
    FOREIGN KEY (guide_id) REFERENCES guides(guide_id),
    FOREIGN KEY (anime_id) REFERENCES animes(anime_id)
);
