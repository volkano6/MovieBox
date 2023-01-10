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
-- Table structure for table `movie_genres`
--

DROP TABLE IF EXISTS `movie_genres`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movie_genres` (
  `movie_id` int NOT NULL,
  `genre_name` varchar(255) NOT NULL,
  PRIMARY KEY (`movie_id`,`genre_name`),
  KEY `movie_genres_genre_name_fk` (`genre_name`),
  CONSTRAINT `movie_genres_genre_name_fk` FOREIGN KEY (`genre_name`) REFERENCES `genres` (`genre`),
  CONSTRAINT `movie_genres_movie_id_fk` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie_genres`
--

LOCK TABLES `movie_genres` WRITE;
/*!40000 ALTER TABLE `movie_genres` DISABLE KEYS */;
INSERT INTO `movie_genres` VALUES (11,'Action'),(14,'Action'),(17,'Action'),(19,'Action'),(22,'Action'),(23,'Action'),(25,'Action'),(38,'Action'),(39,'Action'),(40,'Action'),(48,'Action'),(51,'Action'),(53,'Action'),(64,'Action'),(66,'Action'),(67,'Action'),(68,'Action'),(69,'Action'),(72,'Action'),(77,'Action'),(79,'Action'),(81,'Action'),(14,'Adventure'),(17,'Adventure'),(19,'Adventure'),(22,'Adventure'),(25,'Adventure'),(30,'Adventure'),(32,'Adventure'),(38,'Adventure'),(40,'Adventure'),(48,'Adventure'),(52,'Adventure'),(56,'Adventure'),(67,'Adventure'),(68,'Adventure'),(69,'Adventure'),(70,'Adventure'),(72,'Adventure'),(75,'Adventure'),(79,'Adventure'),(81,'Adventure'),(103,'Adventure'),(32,'Animation'),(52,'Animation'),(55,'Animation'),(65,'Animation'),(67,'Animation'),(70,'Animation'),(75,'Animation'),(79,'Animation'),(16,'Biography'),(24,'Biography'),(27,'Biography'),(44,'Biography'),(47,'Biography'),(63,'Biography'),(66,'Biography'),(28,'Comedy'),(35,'Comedy'),(44,'Comedy'),(56,'Comedy'),(60,'Comedy'),(61,'Comedy'),(73,'Comedy'),(87,'Comedy'),(92,'Comedy'),(105,'Comedy'),(106,'Comedy'),(10,'Crime'),(11,'Crime'),(12,'Crime'),(13,'Crime'),(15,'Crime'),(24,'Crime'),(31,'Crime'),(34,'Crime'),(36,'Crime'),(37,'Crime'),(42,'Crime'),(46,'Crime'),(50,'Crime'),(51,'Crime'),(64,'Crime'),(80,'Crime'),(86,'Crime'),(88,'Crime'),(95,'Crime'),(96,'Crime'),(106,'Crime'),(9,'Drama'),(10,'Drama'),(11,'Drama'),(12,'Drama'),(13,'Drama'),(14,'Drama'),(15,'Drama'),(16,'Drama'),(18,'Drama'),(19,'Drama'),(20,'Drama'),(22,'Drama'),(24,'Drama'),(26,'Drama'),(27,'Drama'),(28,'Drama'),(29,'Drama'),(30,'Drama'),(31,'Drama'),(33,'Drama'),(34,'Drama'),(35,'Drama'),(36,'Drama'),(37,'Drama'),(39,'Drama'),(40,'Drama'),(41,'Drama'),(42,'Drama'),(43,'Drama'),(44,'Drama'),(45,'Drama'),(46,'Drama'),(47,'Drama'),(48,'Drama'),(49,'Drama'),(51,'Drama'),(52,'Drama'),(54,'Drama'),(55,'Drama'),(59,'Drama'),(60,'Drama'),(61,'Drama'),(62,'Drama'),(63,'Drama'),(64,'Drama'),(65,'Drama'),(66,'Drama'),(68,'Drama'),(71,'Drama'),(73,'Drama'),(74,'Drama'),(76,'Drama'),(77,'Drama'),(80,'Drama'),(82,'Drama'),(83,'Drama'),(85,'Drama'),(86,'Drama'),(88,'Drama'),(89,'Drama'),(91,'Drama'),(92,'Drama'),(93,'Drama'),(94,'Drama'),(95,'Drama'),(96,'Drama'),(97,'Drama'),(99,'Drama'),(100,'Drama'),(101,'Drama'),(102,'Drama'),(103,'Drama'),(104,'Drama'),(107,'Drama'),(108,'Drama'),(109,'Drama'),(32,'Family'),(41,'Family'),(60,'Family'),(70,'Family'),(74,'Family'),(75,'Family'),(102,'Family'),(25,'Fantasy'),(34,'Fantasy'),(38,'Fantasy'),(41,'Fantasy'),(65,'Fantasy'),(94,'Fantasy'),(91,'Film-Noir'),(16,'History'),(27,'History'),(63,'History'),(58,'Horror'),(82,'Horror'),(84,'Horror'),(94,'Horror'),(43,'Music'),(47,'Music'),(95,'Music'),(85,'Musical'),(36,'Mystery'),(39,'Mystery'),(45,'Mystery'),(50,'Mystery'),(58,'Mystery'),(76,'Mystery'),(77,'Mystery'),(78,'Mystery'),(83,'Mystery'),(86,'Mystery'),(88,'Mystery'),(90,'Mystery'),(100,'Mystery'),(20,'Romance'),(35,'Romance'),(54,'Romance'),(59,'Romance'),(61,'Romance'),(104,'Romance'),(105,'Romance'),(109,'Romance'),(17,'Sci-Fi'),(23,'Sci-Fi'),(30,'Sci-Fi'),(45,'Sci-Fi'),(53,'Sci-Fi'),(56,'Sci-Fi'),(69,'Sci-Fi'),(84,'Sci-Fi'),(104,'Sci-Fi'),(28,'Thriller'),(37,'Thriller'),(42,'Thriller'),(46,'Thriller'),(50,'Thriller'),(58,'Thriller'),(76,'Thriller'),(78,'Thriller'),(90,'Thriller'),(93,'Thriller'),(96,'Thriller'),(33,'War'),(55,'War'),(59,'War'),(83,'War'),(89,'War'),(92,'War'),(93,'War'),(100,'War'),(103,'War'),(21,'Western'),(57,'Western'),(71,'Western');
/*!40000 ALTER TABLE `movie_genres` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-01-10 23:20:27
