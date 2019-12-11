CREATE TABLE `user`(
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'user id',
  `uid` int(11) unsigned NOT NULL DEFAULT  0 COMMENT 'account id',
  `name` varchar(32) NOT NULL COMMENT 'user name ',
  `avatar` varchar(255) NOT NULL COMMENT 'account avatar',
  `gender` enum('male', 'female', 'unknown') NOT NULL DEFAULT 'unknown' COMMENT 'user gender',
  `created_at` timestamp  COMMENT 'account create time',
  `updated_at` timestamp  COMMENT 'account info update time',
  `deleted_at` timestamp COMMENT 'account delete time',
  PRIMARY KEY (`id`)
  KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user'