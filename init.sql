CREATE DATABASE  IF NOT EXISTS `clinic` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `clinic`;
-- MySQL dump 10.13  Distrib 8.0.29, for Win64 (x86_64)
--
-- Host: localhost    Database: clinic
-- ------------------------------------------------------
-- Server version	8.0.29

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
-- Table structure for table `appointment`
--

DROP TABLE IF EXISTS `appointment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `appointment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `time` datetime DEFAULT NULL,
  `id_patient` int DEFAULT NULL,
  `id_dentist` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `id_patient_idx` (`id_patient`),
  KEY `id_dentist_idx` (`id_dentist`),
  CONSTRAINT `id_dentist` FOREIGN KEY (`id_dentist`) REFERENCES `dentist` (`id`),
  CONSTRAINT `id_patient` FOREIGN KEY (`id_patient`) REFERENCES `patient` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `appointment`
--

LOCK TABLES `appointment` WRITE;
/*!40000 ALTER TABLE `appointment` DISABLE KEYS */;
/*!40000 ALTER TABLE `appointment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dentist`
--

DROP TABLE IF EXISTS `dentist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dentist` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `last_name` varchar(45) DEFAULT NULL,
  `registration_id` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dentist`
--

LOCK TABLES `dentist` WRITE;
/*!40000 ALTER TABLE `dentist` DISABLE KEYS */;
INSERT INTO `dentist` VALUES (1,'Yorgo','O\'Hagerty','6124364743'),(2,'Skipton','Illston','3737493006'),(3,'Rubie','Kingswood','1618659294'),(4,'Annamarie','Hassey','7847044862'),(5,'Nev','Goodlake','6017669737'),(6,'Land','Henningham','6357113050'),(7,'Marius','Urry','2297677987'),(8,'Carola','Kinsey','1996045687'),(9,'Bing','Nineham','6960804000'),(10,'Nichol','Elgram','3095443552');
/*!40000 ALTER TABLE `dentist` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `patient`
--

DROP TABLE IF EXISTS `patient`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `patient` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `last_name` varchar(45) DEFAULT NULL,
  `address` varchar(45) DEFAULT NULL,
  `dni` varchar(45) DEFAULT NULL,
  `registration_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `patient`
--

LOCK TABLES `patient` WRITE;
/*!40000 ALTER TABLE `patient` DISABLE KEYS */;
INSERT INTO `patient` VALUES (1,'Rahal','Petera','525 Randy Alley','8836043801','2022-02-24 20:44:13'),(2,'Vi','Pallis','35489 Charing Cross Crossing','6576384203','2022-08-14 18:45:19'),(3,'Maurine','Shimwell','72 Eastwood Terrace','1342529782','2022-06-05 02:34:21'),(4,'Bethina','Plaskitt','572 6th Avenue','1952987342','2022-07-28 21:28:32'),(5,'Wrennie','MacIntosh','3 Buena Vista Lane','6280744213','2022-02-28 22:38:33'),(6,'Arline','Yardy','5101 Portage Junction','6700744466','2022-02-15 08:29:16'),(7,'Rodrique','Scoffins','423 Butterfield Lane','8742566010','2022-04-03 04:30:42'),(8,'Lelah','Fairebrother','0787 Shoshone Way','3145015494','2021-10-08 14:44:31'),(9,'Cristy','Houldcroft','50445 Twin Pines Crossing','0946754985','2022-04-17 14:56:22'),(10,'Ginger','Rivalant','4 Heath Plaza','0244570361','2021-11-09 05:35:47'),(11,'Free','Verrillo','365 Ramsey Crossing','5932846690','2021-09-25 02:24:50'),(12,'Pip','Kitchiner','7116 Dunning Point','2139349105','2021-12-30 17:04:58'),(13,'Idalina','Poulter','3886 Almo Trail','7069199198','2022-01-25 02:08:13'),(14,'Nicolle','Bursnell','60 Mosinee Place','5184654305','2022-01-26 08:25:15'),(15,'Orville','Juett','242 Gale Crossing','2235579418','2021-11-29 18:31:12'),(16,'Bethina','Grew','3 Lien Avenue','5106351332','2022-03-13 12:12:52'),(17,'Cleo','Brownrigg','025 Kings Circle','3335396863','2022-01-18 19:20:32'),(18,'Shela','Guerry','700 Huxley Street','3701471525','2022-06-03 10:28:25'),(19,'Willi','Norcock','098 Schurz Place','7320221187','2021-11-27 16:47:04'),(20,'Lind','Remington','974 Toban Alley','3244997376','2021-10-24 13:00:35'),(21,'Wallace','Ducarne','002 Elka Crossing','2046117204','2022-05-22 23:19:51'),(22,'Barrie','Rickford','7 Caliangt Circle','3475021331','2021-12-28 05:07:55'),(23,'Simonette','Eisikowitch','3 Randy Plaza','5533331042','2022-02-13 14:45:22'),(24,'Jessalin','Witcombe','0 Pleasure Street','9112019062','2022-08-13 01:52:22'),(25,'Coretta','Meegan','933 Garrison Crossing','5299904770','2022-01-17 08:41:17'),(26,'Karly','Achromov','0726 Dakota Terrace','7616134882','2021-12-04 14:16:36'),(27,'Raymund','Minigo','76652 Lighthouse Bay Crossing','6702803761','2022-06-08 10:18:45'),(28,'Christiana','Motte','926 2nd Terrace','6693336485','2022-01-04 06:27:20'),(29,'Emanuel','Sentance','9435 Onsgard Point','9373727435','2022-08-09 13:21:07'),(30,'Carlyn','Klemz','30044 Meadow Valley Pass','5566143601','2022-02-02 23:36:18'),(31,'Tiffany','Edelheid','9997 Westport Lane','6252609188','2021-12-20 12:10:18'),(32,'Hesther','Shercliff','06 Charing Cross Parkway','3036820035','2021-11-13 09:00:21'),(33,'Tamma','Giacoppo','6 Becker Center','2233574613','2021-12-26 03:28:24'),(34,'Claire','Grewcock','6744 American Ash Center','8671075605','2022-06-10 17:46:40'),(35,'Liza','Redferne','179 Northview Crossing','3757198476','2022-07-17 09:15:17'),(36,'Jelene','Cessford','955 Drewry Hill','1650579950','2022-02-13 20:07:36'),(37,'Patric','Ferri','37 Comanche Lane','5060302032','2022-03-23 22:00:14'),(38,'Siana','Scrowston','4 Jana Terrace','2091290041','2022-05-31 08:05:00'),(39,'Sutherland','Yelding','52467 Dunning Hill','5926197244','2022-08-03 16:23:36'),(40,'Sibley','Knevit','1 Farwell Circle','4975028204','2022-01-28 10:43:01'),(41,'Hayward','Eadon','88 Scoville Crossing','2181918031','2022-02-09 20:14:56'),(42,'Imogene','Angood','0 Marquette Way','7988961705','2022-06-09 09:29:21'),(43,'Marshall','Golden','5845 Schiller Lane','6103307376','2021-10-21 20:07:22'),(44,'Bekki','Bromet','5 Kim Alley','4792735556','2021-11-30 13:24:16'),(45,'Nanci','Todarini','42 Bunting Junction','2666171957','2022-01-10 19:58:51'),(46,'Ingelbert','Artindale','3 Mayer Plaza','3515364390','2022-06-01 16:20:07'),(47,'Libbi','Shewsmith','4 Elgar Court','3712973934','2022-01-16 06:22:10'),(48,'Read','Pashba','405 Buhler Circle','6481810353','2021-10-02 02:32:42'),(49,'Simon','Hegg','2 Eagan Court','0683448625','2021-12-26 22:25:37'),(50,'Coretta','Flindall','2406 Dawn Street','1233240765','2022-06-13 19:11:32'),(51,'Briny','Marunchak','33 West Alley','2310692883','2021-12-31 17:34:14'),(52,'Dov','Radclyffe','18084 Porter Alley','7219143958','2022-03-25 02:55:25'),(53,'Ardeen','Bastard','467 Main Street','6155219079','2021-10-23 16:22:53'),(54,'Lothaire','Arbor','95995 Lakewood Circle','8917522416','2022-01-25 13:56:04'),(55,'Teodor','Topes','4503 Utah Way','4986185217','2022-09-04 01:25:12'),(56,'Julia','MacCosto','4978 Sunfield Court','4023571911','2022-01-13 09:32:47'),(57,'Halie','Blenkensop','32 Blaine Drive','7752107409','2021-12-12 12:24:50'),(58,'Nedda','Mucillo','67 Almo Terrace','8972350257','2021-11-17 12:01:26'),(59,'Clara','Giraux','4 Steensland Avenue','0448506998','2022-09-16 07:56:18'),(60,'Muriel','Augustin','33 Kennedy Plaza','6466582272','2021-12-28 03:11:28'),(61,'Montague','Redemile','79 Steensland Plaza','8617091614','2022-08-04 07:03:23'),(62,'Constantino','Lohan','6673 Glendale Terrace','8541212254','2021-10-02 23:28:57'),(63,'Orelle','Brody','90068 Charing Cross Court','5744968121','2022-07-20 14:46:23'),(64,'Ashlee','Seton','194 Mallory Avenue','2767793325','2021-09-28 16:43:38'),(65,'Mirabella','Hanington','99 Shasta Crossing','0366541501','2022-01-04 09:33:46'),(66,'Marji','Phetteplace','2 Birchwood Way','8714120003','2022-02-27 09:55:13'),(67,'Venus','Duffitt','831 Lindbergh Alley','1313911518','2022-07-18 11:02:42'),(68,'Joshia','Deavall','56653 La Follette Hill','9770739618','2022-08-14 05:05:33'),(69,'Leeann','Gonnin','55477 Waywood Terrace','6090126969','2022-06-07 14:21:10'),(70,'Nickolas','Martynov','2 Loomis Street','7437210762','2022-05-15 08:06:46'),(71,'Hastings','Indgs','2 Tony Center','3274820237','2022-01-09 16:54:57'),(72,'Des','Eales','593 Holmberg Drive','4907019963','2022-07-03 10:55:13'),(73,'Stacee','Hartop','9 Cottonwood Terrace','0050193201','2022-08-25 19:36:00'),(74,'Wilton','Imbrey','399 Westerfield Drive','7651312546','2022-05-25 15:55:28'),(75,'Alix','McConaghy','41 Colorado Point','2128086949','2022-02-12 22:52:32'),(76,'Dolly','Kesley','2 Southridge Avenue','8637388515','2021-10-28 16:50:44'),(77,'Lewiss','Elbourn','785 Roxbury Crossing','7389697202','2022-07-04 13:15:36'),(78,'Brady','Paler','1569 Heath Place','4482968420','2021-12-06 10:58:28'),(79,'Crawford','Crocket','419 Lighthouse Bay Point','9167632882','2022-09-12 06:05:19'),(80,'Neddie','Glendza','12 Sundown Parkway','0797610820','2021-12-17 03:55:37'),(81,'Scottie','Larmor','9385 Towne Hill','8173025266','2022-01-31 12:59:45'),(82,'Avie','Muddicliffe','532 Fisk Alley','0595637272','2021-10-26 10:18:40'),(83,'Kenon','Preator','672 Eggendart Pass','4023974358','2022-06-24 20:25:00'),(84,'Dominique','Boost','87664 Alpine Trail','5192129060','2022-01-15 23:50:24'),(85,'Harlie','Tranter','30 Golf View Parkway','6484419477','2022-05-06 05:55:02'),(86,'Gilly','Cremin','42335 Thierer Center','4875180314','2021-09-27 15:19:08'),(87,'Humfried','Clunie','410 East Junction','7116346953','2022-04-04 10:10:06'),(88,'Bebe','Mosdill','588 Colorado Park','5660710956','2022-01-30 21:35:44'),(89,'Derron','Terbeck','12 Browning Court','4882962411','2022-04-02 04:05:40'),(90,'Hi','Ascroft','9 Stuart Avenue','5446069021','2022-05-06 02:35:31'),(91,'Austina','Speek','92176 Lien Terrace','1878087932','2021-10-25 14:19:55'),(92,'Wells','Jancy','96 Oriole Circle','2874109991','2022-08-25 16:20:23'),(93,'Jermaine','MacNeilly','3047 Oakridge Drive','5055090960','2022-01-10 05:20:26'),(94,'Mead','Glynn','4 Northland Trail','3904979264','2022-01-21 05:13:42'),(95,'Pacorro','Dymidowicz','8474 Kedzie Pass','8232867779','2022-08-17 15:24:03'),(96,'Care','Winslet','3824 Eastwood Alley','7731796361','2022-04-10 15:39:49'),(97,'Godard','Maxwell','4386 Harper Alley','7939978713','2022-02-18 08:27:14'),(98,'Anthia','Newlin','79706 Loeprich Place','9938536255','2021-11-29 21:08:10'),(99,'Hollis','Itzhayek','987 Kenwood Street','5463237263','2021-11-15 00:48:04'),(100,'Berry','Puckrin','3 Tennessee Place','9898554223','2021-10-13 15:49:42');
/*!40000 ALTER TABLE `patient` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-23  1:32:04
