CREATE TABLE `registrations` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `job_id` int(11) unsigned NOT NULL,
  `api_key_id` char(36) NOT NULL,
  `services` json NOT NULL,  
  `requester_id` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `job_id` (`job_id`)
)