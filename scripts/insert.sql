INSERT INTO `users` (
    `id`, `name`, `username`, `email`, `birthday`, `password`, `gender`, `photo_url`, `time`, `active`) 
    VALUES ('0000ffff-0000-0000-0000-000000000000', 'Admin', 'admin', 'test@mail.com', '0000-00-00', 'admin', '', '', 0, 1);
    
INSERT INTO `tags` (`id`, `name`) VALUES (1, 'Action');
INSERT INTO `tags` (`id`, `name`) VALUES (2, 'Adventure');
INSERT INTO `tags` (`id`, `name`) VALUES (3, 'Casual');
INSERT INTO `tags` (`id`, `name`) VALUES (4, 'Indie');
INSERT INTO `tags` (`id`, `name`) VALUES (5, 'Massively Multiplayer');
INSERT INTO `tags` (`id`, `name`) VALUES (6, 'Racing');
INSERT INTO `tags` (`id`, `name`) VALUES (7, 'RPG');
INSERT INTO `tags` (`id`, `name`) VALUES (8, 'Simulation');
INSERT INTO `tags` (`id`, `name`) VALUES (9, 'Sports');
INSERT INTO `tags` (`id`, `name`) VALUES (10, 'Strategy');
INSERT INTO `tags` (`id`, `name`) VALUES (11, 'Free to Play');
INSERT INTO `tags` (`id`, `name`) VALUES (12, 'Early Access');
INSERT INTO `tags` (`id`, `name`) VALUES (13, 'Singleplayer');

INSERT INTO `games` (
    `id`, `name`, `desc`, `image`, `genre`, `release_date`) 
    VALUES ('0000ffff-0000-0000-0000-000000000000', 'Game', 'Game description', 'https://steamcdn-a.akamaihd.net/steam/apps/570/header.jpg?t=1586176520', 'Action', '2013-07-09');
INSERT INTO `games` (
    `id`, `name`, `desc`, `image`, `genre`, `release_date`) 
    VALUES ('00000000-0000-0000-0000-000000000001', 'Game 2', 'Game description 2', 'https://steamcdn-a.akamaihd.net/steam/apps/570/header.jpg?t=1586176520', 'Action', '2013-07-09');  
INSERT INTO `games` (
    `id`, `name`, `desc`, `image`, `genre`, `release_date`) 
    VALUES ('00000000-0000-0000-0000-000000000002', 'Game 3', 'Game description 3', 'https://steamcdn-a.akamaihd.net/steam/apps/570/header.jpg?t=1586176520', 'Action', '2013-07-09');
INSERT INTO `games` (
    `id`, `name`, `desc`, `image`, `genre`, `release_date`) 
    VALUES ('00000000-0000-0000-0000-000000000003', 'Game 4', 'Game description 4', 'https://steamcdn-a.akamaihd.net/steam/apps/570/header.jpg?t=1586176520', 'Action', '2013-07-09');
INSERT INTO `games` (
    `id`, `name`, `desc`, `image`, `genre`, `release_date`) 
VALUES ('00000000-0000-0000-0000-000000000004', 'Game 5', 'Game description 5', 'https://steamcdn-a.akamaihd.net/steam/apps/570/header.jpg?t=1586176520', 'Action', '2013-07-09');

INSERT INTO `links` (
    `id`, `name`, `desc`, `url`, `game_id`) 
    VALUES ('0000ffff-0000-0000-0000-000000000000', 'Link', 'Link description', 'https://steamcdn-a.akamaihd.net/steam/apps/570/header.jpg?t=1586176520', '0000ffff-0000-0000-0000-000000000000');  
INSERT INTO `links` (
    `id`, `name`, `desc`, `url`, `game_id`) 
    VALUES ('00000000-0000-0000-0000-000000000001', 'Link 2', 'Link description 2', 'https://steamcdn-a.akamaihd.net/steam/apps/570/header.jpg?t=1586176520', '0000ffff-0000-0000-0000-000000000000');
INSERT INTO `links` (
    `id`, `name`, `desc`, `url`, `game_id`) 
    VALUES ('00000000-0000-0000-0000-000000000002', 'Link 3', 'Link description 3', 'https://steamcdn-a.akamaihd.net/steam/apps/570/header.jpg?t=1586176520', '0000ffff-0000-0000-0000-000000000000');
INSERT INTO `links` (
    `id`, `name`, `desc`, `url`, `game_id`) 
    VALUES ('00000000-0000-0000-0000-000000000003', 'Link 4', 'Link description 4', 'https://steamcdn-a.akamaihd.net/steam/apps/570/header.jpg?t=1586176520', '00000000-0000-0000-0000-000000000001');
INSERT INTO `links` (
    `id`, `name`, `desc`, `url`, `game_id`) 
    VALUES ('00000000-0000-0000-0000-000000000004', 'Link 5', 'Link description 5', 'https://steamcdn-a.akamaihd.net/steam/apps/570/header.jpg?t=1586176520', '00000000-0000-0000-0000-000000000001');
