CREATE TABLE `patterns` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



INSERT INTO patterns (id,name ,created_at,updated_at)
VALUES (1,"polka dots","2024-10-15","2024-10-15") ,
(2,"drawing","2024-10-15","2024-10-15") ,
(3,"calligraphy","2024-10-15","2024-10-15") ;