CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT '0',
  `amount` int(11) DEFAULT '0',
  `status_id` int(11) DEFAULT '0',
  `user_id` int(11) DEFAULT '0',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO orders (id,product_id ,amount ,status_id ,user_id ,created_at,updated_at)
VALUES 
(1,1,2,1,1,"2024-10-17","2024-10-17"),
(2,2,2,1,1,"2024-10-17","2024-10-17") ,
(3,1,2,1,2,"2024-10-17","2024-10-17");