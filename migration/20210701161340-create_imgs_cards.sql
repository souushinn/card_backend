-- +migrate Up
CREATE TABLE IF NOT EXISTS cards (
    no int(11) NOT NULL AUTO_INCREMENT,
    word VARCHAR(128)
     );

CREATE TABLE IF NOT EXISTS images (
     id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
     img BLOB NOT NULL,
     img_name VARCHAR(255) NOT NULL,
     card_no int(11) NOT NULL
)ENGINE = InnoDB;

-- +migrate Down
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS images;