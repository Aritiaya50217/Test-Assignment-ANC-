CREATE TABLE `colors` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


INSERT INTO colors (id,name ,created_at,updated_at)
VALUES (1,"Red","2024-10-15","2024-10-15") ,
(2,"Black","2024-10-15","2024-10-15"),
(3,"Green","2024-10-15","2024-10-15");