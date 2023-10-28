-- +goose Up
CREATE TABLE  IF NOT EXISTS guide_prefectures (
    guide_id INT,
    prefecture_id INT,
    PRIMARY KEY (guide_id, prefecture_id),
    FOREIGN KEY (guide_id) REFERENCES guides(guide_id),
    FOREIGN KEY (prefecture_id) REFERENCES prefectures(prefecture_id)
);
