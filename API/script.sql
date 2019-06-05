-- phpMyAdmin SQL Dump
-- version 4.7.2
-- https://www.phpmyadmin.net/
--
-- Hôte : localhost:3306
-- Généré le :  mer. 05 juin 2019 à 10:33
-- Version du serveur :  5.6.35
-- Version de PHP :  7.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Base de données :  `saconnect`
--

-- --------------------------------------------------------

--
-- Structure de la table `dayItem`
--

CREATE TABLE `dayItem` (
  `id` int(11) NOT NULL,
  `idItem` int(11) NOT NULL,
  `day` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Structure de la table `item`
--

CREATE TABLE `item` (
  `id` int(11) NOT NULL,
  `isInBag` tinyint(1) NOT NULL,
  `name` varchar(255) NOT NULL,
  `idTag` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `item`
--

INSERT INTO `item` (`id`, `isInBag`, `name`, `idTag`) VALUES
(1, 1, 'cahier maths', '12345');

-- --------------------------------------------------------

--
-- Structure de la table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `weight` double NOT NULL,
  `size` double NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `users`
--

INSERT INTO `users` (`id`, `name`, `weight`, `size`) VALUES
(1, 'test1', 67, 3);

-- --------------------------------------------------------

--
-- Structure de la table `weight`
--

CREATE TABLE `weight` (
  `id` int(11) NOT NULL,
  `valueLeft` double NOT NULL,
  `dateTime` datetime NOT NULL,
  `valueRight` double NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Déchargement des données de la table `weight`
--

INSERT INTO `weight` (`id`, `valueLeft`, `dateTime`, `valueRight`) VALUES
(1, 67, '2019-06-04 10:34:00', 45),
(2, 67, '0000-00-00 00:00:00', 45),
(3, 67, '0000-00-00 00:00:00', 45),
(4, 67, '0000-00-00 00:00:00', 45),
(5, 67, '2019-06-04 10:47:18', 45),
(6, 67, '2019-06-04 10:47:58', 45),
(7, 234, '2019-06-04 11:10:49', 342),
(8, 276, '2019-06-04 11:21:37', 342),
(9, 276, '2019-06-05 08:38:42', 342);

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `dayItem`
--
ALTER TABLE `dayItem`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `item`
--
ALTER TABLE `item`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `weight`
--
ALTER TABLE `weight`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `dayItem`
--
ALTER TABLE `dayItem`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT pour la table `item`
--
ALTER TABLE `item`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT pour la table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT pour la table `weight`
--
ALTER TABLE `weight`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;