-- -------------
--
-- posts
--
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(36) NOT NULL,
  `title` char(128) NOT NULL,
  `content` TEXT,
  `private` tinyint(1) unsigned NOT NULL default '0',
  `updated_by` int(11) unsigned NOT NULL default '0',
  `updated_at` timestamp(6) default CURRENT_TIMESTAMP(6) on update CURRENT_TIMESTAMP(6),
  `created_by` int(11) unsigned NOT NULL default '0',
  `created_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`),
  KEY `title` (`title`),
  KEY `private` (`private`),
  KEY `updated_by` (`updated_by`),
  KEY `updated_at` (`updated_at`),
  KEY `created_by` (`created_by`),
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `posts_history`;
CREATE TABLE `posts_history` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `post_id` int(11) unsigned NOT NULL,
  `uuid` char(36) NOT NULL,
  `title` char(128) NOT NULL,
  `content` TEXT,
  `private` tinyint(1) unsigned NOT NULL default '0',
  `updated_by` int(11) unsigned NOT NULL DEFAULT '0',
  `updated_at` datetime(6) NOT NULL,
  `created_by` int(11) unsigned NOT NULL default '0',
  `created_at` datetime(6) NOT NULL,
  `history_method` enum('update','delete') NOT NULL DEFAULT 'update',
  `history_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `post_id` (`post_id`),
  KEY `uuid` (`uuid`),
  KEY `title` (`title`),
  KEY `private` (`private`),
  KEY `updated_by` (`updated_by`),
  KEY `updated_at` (`updated_at`),
  KEY `created_by` (`created_by`),
  KEY `created_at` (`created_at`),
  KEY (`history_method`),
  KEY (`history_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- +migration BeginStatement
DROP TRIGGER IF EXISTS `posts_history_trigger_update`;
CREATE TRIGGER `posts_history_trigger_update` BEFORE UPDATE ON `posts` FOR EACH ROW BEGIN
    INSERT INTO `posts_history` SET
    `post_id`         = OLD.id,
    `uuid`            = old.uuid,
    `title`           = old.title,
    `content`         = old.content,
    `private`         = old.private,
    `updated_by`      = old.updated_by,
    `updated_at`      = OLD.updated_at,
    `created_by`      = OLD.created_by,
    `created_at`      = OLD.created_at,
    `history_method`  = 'update',
    `history_at` = NOW(6);
END;;
-- +migration EndStatement

-- +migration BeginStatement
DROP TRIGGER IF EXISTS `posts_history_trigger_delete`;
CREATE TRIGGER `posts_history_trigger_delete` BEFORE DELETE ON `posts` FOR EACH ROW BEGIN
    INSERT INTO `posts_history` SET
    `post_id`         = OLD.id,
    `uuid`            = old.uuid,
    `title`           = old.title,
    `content`         = old.content,
    `private`         = old.private,
    `updated_by`      = old.updated_by,
    `updated_at`      = OLD.updated_at,
    `created_by`      = OLD.created_by,
    `created_at`      = OLD.created_at,
    `history_method` = 'delete',
    `history_at` = NOW(6);
END;;
-- +migration EndStatement