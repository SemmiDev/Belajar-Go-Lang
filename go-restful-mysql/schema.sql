CREATE TABLE students (
      `id`              int(11)     NOT NULL AUTO_INCREMENT,
      `first_name`      varchar(50) NOT NULL,
      `last_name`       varchar(50) NOT NULL,
      `identifier`      varchar(10) NOT NULL UNIQUE,
      `email`           varchar(50) NOT NULL UNIQUE,
      `phone_number`    varchar(14) NOT NULL UNIQUE,
      PRIMARY KEY (`id`)
);

INSERT INTO students(id, first_name, last_name, identifier, email, phone_number) VALUES (1, 'sam','dev','2003113941', 'sam1@gmail.com', '12');
INSERT INTO students(id, first_name, last_name, identifier, email, phone_number) VALUES (2, 'sam2','dev2','2003113942', 'sam2@gmail.com', '13');
INSERT INTO students(id, first_name, last_name, identifier, email, phone_number) VALUES (3, 'sam3','dev3','2003113943', 'sam3@gmail.com', '14');
INSERT INTO students(id, first_name, last_name, identifier, email, phone_number) VALUES (4, 'sam4','dev4','2003113944', 'sam4@gmail.com', '15');
INSERT INTO students(id, first_name, last_name, identifier, email, phone_number) VALUES (5, 'sam5','dev5','2003113945', 'sam5@gmail.com', '16');
INSERT INTO students(id, first_name, last_name, identifier, email, phone_number) VALUES (6, 'sam6','dev6','2003113946', 'sam6@gmail.com', '17');