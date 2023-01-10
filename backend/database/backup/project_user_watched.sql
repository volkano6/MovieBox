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
-- Table structure for table `user_watched`
--

DROP TABLE IF EXISTS `user_watched`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_watched` (
  `movie_id` int NOT NULL,
  `user_id` int NOT NULL,
  `watched_date` datetime NOT NULL,
  PRIMARY KEY (`movie_id`,`user_id`),
  KEY `user_watched_user_id_fk` (`user_id`),
  CONSTRAINT `user_watched_movie_id_fk` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`),
  CONSTRAINT `user_watched_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_watched`
--

LOCK TABLES `user_watched` WRITE;
/*!40000 ALTER TABLE `user_watched` DISABLE KEYS */;
INSERT INTO `user_watched` VALUES (16,22,'2023-01-04 15:33:17'),(17,22,'2023-01-04 15:34:59'),(18,22,'2023-01-04 15:32:48'),(18,26,'2023-01-04 13:25:48'),(22,22,'2023-01-04 15:32:59'),(22,26,'2023-01-04 09:25:11'),(22,39,'2023-01-05 10:05:58'),(25,6,'2000-01-02 00:00:00'),(25,26,'2000-02-02 00:00:00'),(25,34,'2022-12-12 00:00:00'),(26,26,'2022-12-12 23:38:57'),(48,52,'2023-01-05 17:11:29'),(55,26,'2022-12-12 23:38:57'),(88,26,'2022-12-16 19:36:20');
/*!40000 ALTER TABLE `user_watched` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-01-10 23:20:49
