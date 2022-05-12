CREATE TABLE blobs.files (
    id int NOT NULL AUTO_INCREMENT,
    file_name VARCHAR(30),
    file_size INT,
    date_of_insert VARCHAR(30),
    bin_data LONGBLOB,
    PRIMARY KEY(id)
);