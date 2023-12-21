CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `user_id` bigint(20) NOT NULL,
                        `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `email` varchar(64) COLLATE utf8mb4_general_ci,
                        `gender` tinyint(4) NOT NULL DEFAULT '0',
                        `creat_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_username` (`username`) USING BTREE,
                        UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



drop table if exists `community`;
create table `community` (
    `id` int(11) not null auto_increment,
    `community_id` int(10) unsigned not null,
    `community_name` varchar(128) collate utf8mb4_general_ci not null,
    `introduction` varchar(256) collate utf8mb4_general_ci not null,
    `creat_time` timestamp not null default current_timestamp,
    `update_time` timestamp not null default current_timestamp on update current_timestamp,
    primary key(`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_general_ci;


INSERT INTO `community` VALUES ('1', '1', 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');
INSERT INTO `community` VALUES ('2', '2', 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00');
INSERT INTO `community` VALUES ('3', '3', 'CS:GO', 'Rush B。。。', '2018-08-07 08:30:00', '2018-08-07 08:30:00');
INSERT INTO `community` VALUES ('4', '4', 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');