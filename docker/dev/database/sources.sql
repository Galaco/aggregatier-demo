-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: database
-- Generation Time: Feb 04, 2019 at 01:12 PM
-- Server version: 5.6.43
-- PHP Version: 7.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `aggregatier`
--

-- --------------------------------------------------------

--
-- Table structure for table `games`
--

CREATE TABLE `games` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `short_name` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `games`
--

INSERT INTO `games` (`id`, `name`, `short_name`) VALUES
(1, 'Smite', 'smite'),
(2, 'League of Legends', 'LoL'),
(3, 'Defence of the Ancients 2', 'DotA 2');

-- --------------------------------------------------------

--
-- Table structure for table `heroes`
--

CREATE TABLE `heroes` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL,
  `game_id` int(11) NOT NULL,
  `icon_url` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tier`
--

CREATE TABLE `tier` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL,
  `game_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tier`
--

INSERT INTO `tier` (`id`, `name`, `game_id`) VALUES
(1, 'S+', 1),
(2, 'S', 1),
(3, 'A+', 1),
(4, 'A', 1),
(5, 'A-', 1),
(6, 'B+', 1),
(7, 'B', 1),
(8, 'B-', 1),
(9, 'C+', 1),
(10, 'C', 1),
(11, 'C-', 1),
(12, 'D', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `games`
--
ALTER TABLE `games`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `heroes`
--
ALTER TABLE `heroes`
  ADD PRIMARY KEY (`id`),
  ADD KEY `game_id` (`game_id`);

--
-- Indexes for table `tier`
--
ALTER TABLE `tier`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_game_id` (`game_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `games`
--
ALTER TABLE `games`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `heroes`
--
ALTER TABLE `heroes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `tier`
--
ALTER TABLE `tier`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `heroes`
--
ALTER TABLE `heroes`
  ADD CONSTRAINT `heroes_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `games` (`id`);

--
-- Constraints for table `tier`
--
ALTER TABLE `tier`
  ADD CONSTRAINT `fk_game_id` FOREIGN KEY (`game_id`) REFERENCES `games` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
