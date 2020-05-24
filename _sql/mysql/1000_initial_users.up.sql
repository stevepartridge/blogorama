-- -------------
--
-- users
--
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `api_key` char(36) NOT NULL,
  `name` char(128) NOT NULL,
  `active` tinyint(1) unsigned NOT NULL default '0',
  `updated_by` int(11) unsigned NOT NULL default '0',
  `updated_at` timestamp(6) default CURRENT_TIMESTAMP(6) on update CURRENT_TIMESTAMP(6),
  `created_by` int(11) unsigned NOT NULL default '0',
  `created_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `api_key` (`api_key`),
  KEY `name` (`name`),
  KEY `active` (`active`),
  KEY `updated_by` (`updated_by`),
  KEY `updated_at` (`updated_at`),
  KEY `created_by` (`created_by`),
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `users_history`;
CREATE TABLE `users_history` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `api_key` char(36) NOT NULL,
  `name` char(128) NOT NULL,
  `active` tinyint(1) unsigned NOT NULL default '0',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0',
  `updated_at` datetime(6) NOT NULL,
  `created_by` int(11) unsigned NOT NULL default '0',
  `created_at` datetime(6) NOT NULL,
  `history_method` enum('update','delete') NOT NULL DEFAULT 'update',
  `history_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `api_key` (`api_key`),
  KEY `name` (`name`),
  KEY `active` (`active`),
  KEY `updated_by` (`updated_by`),
  KEY `updated_at` (`updated_at`),
  KEY `created_by` (`created_by`),
  KEY `created_at` (`created_at`),
  KEY (`history_method`),
  KEY (`history_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- +migration BeginStatement
DROP TRIGGER IF EXISTS `users_history_trigger_update`;
CREATE TRIGGER `users_history_trigger_update` BEFORE UPDATE ON `users` FOR EACH ROW BEGIN
    INSERT INTO `users_history` SET
    `user_id`         = OLD.id,
    `api_key`            = old.api_key,
    `name`            = old.name,
    `active`          = old.active,
    `updated_by`      = old.updated_by,
    `updated_at`      = OLD.updated_at,
    `created_by`      = OLD.created_by,
    `created_at`      = OLD.created_at,
    `history_method`  = 'update',
    `history_at` = NOW(6);
END;;
-- +migration EndStatement

-- +migration BeginStatement
DROP TRIGGER IF EXISTS `users_history_trigger_delete`;
CREATE TRIGGER `users_history_trigger_delete` BEFORE DELETE ON `users` FOR EACH ROW BEGIN
    INSERT INTO `users_history` SET
    `user_id`         = OLD.id,
    `api_key`            = old.api_key,
    `name`            = old.name,
    `active`          = old.active,
    `updated_by`      = old.updated_by,
    `updated_at`      = OLD.updated_at,
    `created_by`      = OLD.created_by,
    `created_at`      = OLD.created_at,
    `history_method`  = 'delete',
    `history_at` = NOW(6);
END;;
-- +migration EndStatement