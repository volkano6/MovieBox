-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: 129.159.122.34    Database: project
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `displayname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `dateofbirth` date NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `bio` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `social_twitter` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `social_instagram` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Ali','1234567890','Ali','alii@hotmail.com','2000-04-11','','','','',''),(4,'Buket','1234567890','Buket','buket@hotmail.com','2000-04-11','','','','',''),(6,'','1234567890','','bukket@hotmail.com','2000-04-11','','','','',''),(8,'batuhan','123123','batuhan','vbatuhan@hotmail.com','2000-04-03','','','','',''),(11,'Bukett','1234567890','Bukett','bukekhjtdsfsd@hotmail.com','2000-04-11','','','','',''),(12,'Bukettr','1234567890','Bukettr','bukekhjtdsfsrd@hotmail.com','2000-04-11','','','','',''),(13,'Bukettrr','1234567890','Bukettrr','bukekhjtdsfsrrd@hotmail.com','2000-04-11','','','','',''),(14,'Mehmet','1234567890','Mehmet','mehmet@hotmail.com','2000-04-11','','','','',''),(15,'Arda Serdar','$2a$14$/1sF4zSwHmh06OXBVIuiEefYkJVZFOga2knLf3XzQoRfNwx28ynvO','Arda Serdar','ardap@gmail.com','2002-10-28','','','','',''),(16,'test2','$2a$14$CIWolt5W6jt7cFyOfILn7.zQsPAN/iTq6WB7z61YHPIGE5gUtVzM.','test2','asdf','2020-11-12','','','','',''),(19,'Mehmetd','$2a$14$R3rygMZCrbljyukMNgHIYuHskg6dVJ.yFyW.bSLfY4O/d.MXAFEv.','Mehmetd','mehmewt@hotmail.com','2000-04-11','','','','',''),(20,'Mehdmetd','$2a$14$.dvvLGDpd3PTM5qDHF9.9eDmZEwbLhfZUVozhSMG2ej8Y8mFO5BeO','Mehdmetd','mehmedwt@hotmail.com','2000-04-11','','','','',''),(21,'MehCVXVdmetd','$2a$14$2ECg3GYLz8jPs3zlJVXRUupbjVjP3..QPAiGydDIiV7SPvrnGrXSK','MehCVXVdmetd','mehmedSXCDVwt@hotmail.com','2000-04-11','','','','',''),(22,'pektezol','$2a$14$vDMGOEITU1eMzQCtrnNj8u3cYP7xKeB1xyQcntDWBepRN3SjYIEg2','Arda Serdar Pektezol','ardapektezol@outlook.com','2002-10-28','','helo','Istanbul, TR','','ardaserdarp'),(23,'sfdgd','$2a$14$DLH7ke6gjUsqFpPWXuZ3BOX6sg.oRy4szgWvAcDtgbzlLQXpaJkbO','sfdgd','volkanodfhztoklu@hotmail.com','2000-04-11','','','','',''),(24,'dfgdf','$2a$14$KVER1mPtNnwX66GkjXTbZujrbTcHmDPzgCeFV/ePbKI/IcrHsvlJe','dfgdf','fdgdf@fdgfd','2000-04-11','','','','',''),(26,'kaan','$2a$14$qZkksKQZhcaM2z265nhGhee9eQdgD8oRfyJM1FlRCmpcI3QHW85EO','kaan','kaan@kaan','2000-04-19','https://avatars.githubusercontent.com/u/86519099?v=4','wklhfwıqopeyfiıoqywefıioyqwewueyrtıprwe','İstanbul, Sarıyer','volka_ztoklu','volka_ztoklu'),(28,'kagan','$2a$14$ijyY0wMpneGbk1pcgyVhvO/g2eq7YPviJvufLgeqEuOxKaKr3v10S','kagan','kagan@kagan','2000-04-18','','','','',''),(29,'sdfds','$2a$14$RdaBU2semNL1ofKaYZcneedrteaKYSZsIVR1gc9P6z0QPgpoJu7P6','sdfds','fasdf@dsfsdf','2000-01-01','','','','',''),(30,'dsfasd','$2a$14$37MoHOsg5EQs/MnjdPZK2eMUR26CUTWFjRHJw22Nc5QsyRSZHmGNm','dsfasd','volkanoztoklu@hotmail.comhdgsdk','2000-01-01','','','','',''),(31,'irem','$2a$14$8Erp/ACIBKUAk7fqKh/jhe1TaG8F31k2EJMRcMV7aXs/cU5q15u3O','irem','irem@irem','2000-02-02','','','','',''),(34,'alperen','$2a$14$mSanVvMCqxee0bLnHk0z6uqO1aFYgybXf9VSxOMt4D2X5kkPhvQDK','alperen','alperen@alperen','2000-01-01','','','','',''),(36,'aliasd','$2a$14$l7dJLYMWD/7tumwuH5nVqevsEePlTcq8dFh7SJUNIQfoSp627xcjq','aliasd','volkanoztoklu@hotmail.com','2000-02-02','','','','',''),(38,'volkan','$2a$14$VKF/mq9LUp8YlqiXp3/Um.u6/jAQtjqn3NDp3xOs7ZBWvic4vi3WC','volkan','volkanoztoklu@hotmaidsfl.com','2000-03-02','','','','',''),(39,'aslıenver','$2a$14$AhyPc.LR5kIVF4pFup7Z3O4nXk5eVCy60RS8JwnffcT9P7TC.wHai','aslıenver','aslie@gmail.com','2000-01-01','','','','',''),(41,'serkan','$2a$14$CuDmNOd7LBPr8tzenht1q.aMjOIFHb8y904latx7/R3flsWK/oqJq','serkan','serkan@gmail.com','1985-10-10','','','','',''),(51,'serkan2','$2a$14$8rEq8rEzBI6aUBTLWSg6Cer3Ng6p2knmcilGdYotejZ5UDzPs9yIC','serkan2','serkan2@gmail.com','2000-12-02','','','','',''),(52,'deneme','$2a$14$pTOg2RK/OLg04u7Sac1mpect5Lb/WsalbDo9mBsOL3If3y8nDH0Sm','ilker','deneme@example.com','2000-01-01','','','','','');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-01-10 23:20:38
