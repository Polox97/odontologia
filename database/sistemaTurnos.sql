-- MySQL dump 10.13  Distrib 8.0.16, for macos10.14 (x86_64)
--
-- Host: 127.0.0.1    Database: sistemaTurnos
-- ------------------------------------------------------
-- Server version	5.7.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `pacientes`
--

DROP TABLE IF EXISTS `pacientes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `pacientes`(
    `id` int NOT NULL primary key auto_increment,
    `dni` text NOT NULL,
    `nombre` text NOT NULL,
    `apellido` text NOT NULL,
    `domicilio` text NOT NULL,
    `fecha_alta` TIMESTAMP DEFAULT current_timestamp(),
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `pacientes` WRITE;
/*!40000 ALTER TABLE `phone_books` DISABLE KEYS */;
INSERT INTO `pacientes` VALUES (1,'123456789','Luis','Dias','calle 80'),(2,'987654321','Brandon','Salgado','calle 90'),(3,'543216789','Maria','Monrroy','calle 100'),(4,'678954321','Andrea','Ceron','calle 200');
/*!40000 ALTER TABLE `phone_books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pacientes`
--

DROP TABLE IF EXISTS `pacientes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE dentistas(
    `id` int NOT NULL primary key auto_increment,
    `matricula` text null,
    `nombre` text null,
    `apellido` text null
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `pacientes` WRITE;
/*!40000 ALTER TABLE `phone_books` DISABLE KEYS */;
INSERT INTO `pacientes` VALUES (1,'1234','Juan','Pulido'),(2,'5678','Pepito','Perez'),(3,'4321','Josefa','Ruiz'),(4,'8765','Leonel','Messi');
/*!40000 ALTER TABLE `phone_books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `turnos`
--

DROP TABLE IF EXISTS `turnos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE turnos(
    `id` int NOT NULL primary key auto_increment,
    `paciente_id` int NOT NULL
    `dentista_id` int NOT NULL
    `fecha_hora` DATETIME NOT NULL,
    `descripcion` text NOT NULL
    CONSTRAINT turnos_pacienteId_fk FOREIGN KEY (paciente_id) REFERENCES pacientes(id),
    CONSTRAINT turnos_dentistaId_fk FOREIGN KEY (dentista_id) REFERENCES dentistas(id),
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `pacientes` WRITE;
/*!40000 ALTER TABLE `phone_books` DISABLE KEYS */;
INSERT INTO `pacientes` VALUES (1,2,1,'2022-09-23 10:30:27.000', 'extracción'),(2,2,2,'2022-09-23 9:30:27.000', 'limpieza'),(3,3,3,'2022-09-23 11:30:27.000', 'limpieza'),(4,3,2,'2022-09-23 09:30:27.000', 'extracción'),(5,4,2,'2022-09-23 09:50:27.000', 'extracción');
/*!40000 ALTER TABLE `phone_books` ENABLE KEYS */;
UNLOCK TABLES;

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-06-11 10:36:24