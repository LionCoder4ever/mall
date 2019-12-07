CREATE TABLE `account`(
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'account id',
  `name` varchar(32) NOT NULL COMMENT 'account name ',
  `avatar` varchar(255) NOT NULL COMMENT 'account avatar',
  `created_at` timestamp  COMMENT 'account create time',
  `updated_at` timestamp  COMMENT 'account info update time',
  `deleted_at` timestamp COMMENT 'account delete time',
  -- privacy
  `real_name` varchar(32)  COMMENT 'account real name',
  `id_card`   varchar(32)  COMMENT 'account identity card',
  `phone` varchar(16)  COMMENT 'account phone number',
  `regist_ip` varchar(32)  COMMENT 'account regist ip',
  `hand_img`  varchar(255) COMMENT 'accountaccount hand info',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user account'