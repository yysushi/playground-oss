CREATE USER user@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO user@'%';
FLUSH PRIVILEGES;
CREATE DATABASE dbname;

/*
USER db1;

CREATE TABLE `users` (
  `id` varchar(255) NOT NULL,
  `name` text,
  `active` boolean,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
*/
