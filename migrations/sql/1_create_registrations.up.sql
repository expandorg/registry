CREATE TABLE `registrations` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL DEFAULT '',
  `service` varchar(255) NOT NULL DEFAULT '',
  `job_id` int(11) unsigned NOT NULL,
  `requester_id` int(11) unsigned NOT NULL,
  `active` tinyint(4) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `job_service` (`job_id`,`service`),
  KEY `job_id` (`job_id`)
)