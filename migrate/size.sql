CREATE TABLE `sizes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `size` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


INSERT INTO sizes (id,`size` ,created_at,updated_at)
VALUES (1,"XS","2024-10-15","2024-10-15") ,
(2,"S","2024-10-15","2024-10-15"),
(3,"M","2024-10-15","2024-10-15"),
(4,"L","2024-10-15","2024-10-15"),
(5,"XL","2024-10-15","2024-10-15") ;