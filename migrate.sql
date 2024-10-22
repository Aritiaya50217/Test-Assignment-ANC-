CREATE DATABASE  IF NOT EXISTS `colors` ;
USE `colors` ;

DROP TABLE IF EXISTS `colors`;

CREATE TABLE `colors` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


LOCK TABLES `colors` WRITE;
INSERT INTO `colors` (id,name ,created_at,updated_at)
VALUES (1,"Red","2024-10-15","2024-10-15") ,
(2,"Black","2024-10-15","2024-10-15"),
(3,"Green","2024-10-15","2024-10-15");
UNLOCK TABLES;


CREATE DATABASE  IF NOT EXISTS `figures` ;
USE `figures` ;

DROP TABLE IF EXISTS `figures`;
CREATE TABLE `figures` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

LOCK TABLES `figures` WRITE;
INSERT INTO `figures` (id,name ,created_at,updated_at)
VALUES (1,"Batman","2024-10-15","2024-10-15") ,
(2,"Iron Man","2024-10-15","2024-10-15") ,
(3,"Spiderman,","2024-10-15","2024-10-15") ;
UNLOCK TABLES;



CREATE DATABASE  IF NOT EXISTS `genders` ;
USE `genders` ;

DROP TABLE IF EXISTS `genders`;
CREATE TABLE `genders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

LOCK TABLES `genders` WRITE;
INSERT INTO `genders` (id,name ,created_at,updated_at)
VALUES (1,"Men","2024-10-15","2024-10-15") ,
(2,"Women","2024-10-15","2024-10-15")  ;
UNLOCK TABLES;



CREATE DATABASE  IF NOT EXISTS `orders` ;
USE `orders` ;
DROP TABLE IF EXISTS `orders`;

CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT '0',
  `amount` int(11) DEFAULT '0',
  `status_id` int(11) DEFAULT '0',
  `user_id` int(11) DEFAULT '0',
  `address` varchar(255) DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

LOCK TABLES `orders` WRITE;
INSERT INTO orders (id,product_id ,amount ,status_id ,user_id ,created_at,updated_at)
VALUES 
(1,1,2,1,1,"address ..","2024-10-17","2024-10-17"),
(2,2,2,1,1,"address ..","2024-10-17","2024-10-17") ,
(3,1,2,1,2,"address ..","2024-10-17","2024-10-17");
UNLOCK TABLES;


CREATE DATABASE  IF NOT EXISTS `patterns` ;
USE `patterns` ;
DROP TABLE IF EXISTS `patterns`;

CREATE TABLE `patterns` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
UNLOCK TABLES;

LOCK TABLES `patterns` WRITE;
INSERT INTO `patterns` (id,name ,created_at,updated_at)
VALUES (1,"polka dots","2024-10-15","2024-10-15") ,
(2,"drawing","2024-10-15","2024-10-15") ,
(3,"calligraphy","2024-10-15","2024-10-15") ;
UNLOCK TABLES;


CREATE DATABASE  IF NOT EXISTS `products` ;
USE `products` ;
DROP TABLE IF EXISTS `products`;

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


LOCK TABLES `products` WRITE;
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

UNLOCK TABLES;


CREATE DATABASE  IF NOT EXISTS `sizes` ;
USE `sizes` ;

DROP TABLE IF EXISTS `sizes`;

CREATE TABLE `sizes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `size` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



LOCK TABLES `sizes` WRITE;
INSERT INTO `sizes` (id,`size` ,created_at,updated_at)
VALUES (1,"XS","2024-10-15","2024-10-15") ,
(2,"S","2024-10-15","2024-10-15"),
(3,"M","2024-10-15","2024-10-15"),
(4,"L","2024-10-15","2024-10-15"),
(5,"XL","2024-10-15","2024-10-15") ;
UNLOCK TABLES;

CREATE DATABASE  IF NOT EXISTS `statuses` ;
USE `statuses` ;

DROP TABLE IF EXISTS `statuses`;

CREATE TABLE `statuses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `status` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

LOCK TABLES `statuses` WRITE;
INSERT INTO statuses (id,status,created_at,updated_at)
VALUES (1,"placed_order","2024-10-15","2024-10-15") ,
(2,"paid","2024-10-15","2024-10-15"),
(3,"out of shipping","2024-10-15","2024-10-15"),
(4,"completed","2024-10-15","2024-10-15")

UNLOCK TABLES;


CREATE DATABASE  IF NOT EXISTS `users` ;
USE `users` ;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `address` varchar(255) COLLATE utf8_unicode_ci NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

LOCK TABLES `users` WRITE;
INSERT INTO  `users` (id,username,address,created_at,updated_at)
VALUES (1,"user one","address ...","2024-10-17","2024-10-17"),
(2,"user two","address ...","2024-10-17","2024-10-17");
UNLOCK TABLES;