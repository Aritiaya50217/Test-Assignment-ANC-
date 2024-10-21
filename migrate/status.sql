CREATE TABLE `statuses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `status` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


INSERT INTO statuses (id,status,created_at,updated_at)
VALUES (1,"placed_order","2024-10-15","2024-10-15") ,
(2,"paid","2024-10-15","2024-10-15"),
(3,"out of shipping","2024-10-15","2024-10-15"),
(4,"completed","2024-10-15","2024-10-15")