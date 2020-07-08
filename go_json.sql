-- phpMyAdmin SQL Dump
-- version 4.9.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 08, 2020 at 04:17 PM
-- Server version: 10.4.10-MariaDB
-- PHP Version: 7.3.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_json`
--

-- --------------------------------------------------------

--
-- Table structure for table `clients`
--

CREATE TABLE `clients` (
  `id` varchar(25) NOT NULL,
  `isActive` varchar(5) NOT NULL,
  `age` int(3) NOT NULL,
  `name` varchar(25) NOT NULL,
  `gender` varchar(6) NOT NULL,
  `company` varchar(25) NOT NULL,
  `email` varchar(25) NOT NULL,
  `phone` varchar(12) NOT NULL,
  `street` varchar(25) NOT NULL,
  `city` varchar(25) NOT NULL,
  `state` varchar(10) NOT NULL,
  `zip` varchar(5) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `clients`
--

INSERT INTO `clients` (`id`, `isActive`, `age`, `name`, `gender`, `company`, `email`, `phone`, `street`, `city`, `state`, `zip`) VALUES
('59761c23b30d971669fb42ff', 'true', 36, 'Dunlap Hubbard', 'male', 'CEDWARD', 'dunlaphubbard@cedward.com', '+1 (890) 543', '6649 N Blue Gum St', 'New Orleans', 'LA', '70116');

-- --------------------------------------------------------

--
-- Table structure for table `product`
--

CREATE TABLE `product` (
  `id` varchar(4) NOT NULL,
  `type` varchar(15) NOT NULL,
  `name` varchar(15) NOT NULL,
  `img_url` varchar(50) NOT NULL,
  `img_wth` int(5) NOT NULL,
  `img_hgt` int(5) NOT NULL,
  `thm_url` varchar(50) NOT NULL,
  `thm_wth` int(5) NOT NULL,
  `thm_hgt` int(5) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`id`, `type`, `name`, `img_url`, `img_wth`, `img_hgt`, `thm_url`, `thm_wth`, `thm_hgt`) VALUES
('0001', 'donut', 'Cake', 'images/0001.jpg', 0, 0, 'images/thumbnails/0001.jpg', 0, 0),
('0001', 'donut', 'Cake', 'images/0001.jpg', 0, 0, 'images/thumbnails/0001.jpg', 0, 0),
('0001', 'donut', 'Cake', 'images/0001.jpg', 200, 200, 'images/thumbnails/0001.jpg', 32, 32);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
