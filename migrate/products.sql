CREATE TABLE `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `style` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `gender_id` int(11) DEFAULT '0',
  `size_id` int(11) DEFAULT '0',
  `color_id` int(11) DEFAULT '0',
  `pattern_id` int(11) DEFAULT '0',
  `figure_id` int(11) DEFAULT '0',
  `price`  int(11) DEFAULT '0',
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



INSERT INTO products (id,`style`,gender_id ,size_id,color_id,pattern_id,figure_id,price,created_at,updated_at)
VALUES
(1,"T-Shirt",1,1,1,1,1,400,"2024-10-15","2024-10-15"),
(2,"T-Shirt",1,2,1,2,2,400,"2024-10-15","2024-10-15"),
(3,"T-Shirt",1,3,1,3,3,420,"2024-10-15","2024-10-15"),
(4,"T-Shirt",1,4,1,1,1,430,"2024-10-15","2024-10-15"),
(5,"T-Shirt",1,5,1,2,2,450,"2024-10-15","2024-10-15"),
(6,"T-Shirt",2,1,2,3,3,290,"2024-10-15","2024-10-15"),
(7,"T-Shirt",2,2,1,1,1,290,"2024-10-15","2024-10-15"),
(8,"T-Shirt",2,3,1,2,2,290,"2024-10-15","2024-10-15"),
(9,"T-Shirt",2,4,1,3,3,320,"2024-10-15","2024-10-15"),
(10,"T-Shirt",2,5,1,1,1,320,"2024-10-15","2024-10-15")
;
