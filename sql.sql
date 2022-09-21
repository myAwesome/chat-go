CREATE DATABASE IF NOT EXISTS `chat_v1_1663684348`;
CREATE TABLE `chat_v1_1663684348`.`tbl_room` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  
  `name` VARCHAR(255) DEFAULT NULL,
  `is_direct` BOOLEAN DEFAULT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `chat_v1_1663684348`.`tbl_participant` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  
  `name` VARCHAR(255) DEFAULT NULL,
  `email` VARCHAR(255) DEFAULT NULL,
  `password` VARCHAR(255) DEFAULT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `chat_v1_1663684348`.`tbl_message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `text` VARCHAR(255) DEFAULT NULL,
  `created` DATETIME DEFAULT NULL,
  `author` INT(11) DEFAULT NULL,
  `room` INT(11) DEFAULT NULL,
  
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `chat_v1_1663684348`.`tbl_room_has_participants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  
  `room` INT(11) DEFAULT NULL,
  `participant` INT(11) DEFAULT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




                        
ALTER TABLE `chat_v1_1663684348`.`tbl_message` ADD INDEX `fk_tbl_message_author_idx` (`author` ASC);
ALTER TABLE `chat_v1_1663684348`.`tbl_message` ADD CONSTRAINT `fk_tbl_message_author` FOREIGN KEY (`author`)
REFERENCES `chat_v1_1663684348`.`tbl_participant` (`id`) ON DELETE NO ACTION   ON UPDATE NO ACTION;
  
ALTER TABLE `chat_v1_1663684348`.`tbl_message` ADD INDEX `fk_tbl_message_room_idx` (`room` ASC);
ALTER TABLE `chat_v1_1663684348`.`tbl_message` ADD CONSTRAINT `fk_tbl_message_room` FOREIGN KEY (`room`)
REFERENCES `chat_v1_1663684348`.`tbl_room` (`id`) ON DELETE NO ACTION   ON UPDATE NO ACTION;
        
ALTER TABLE `chat_v1_1663684348`.`tbl_room_has_participants` ADD INDEX `fk_tbl_room_has_participants_room_idx` (`room` ASC);
ALTER TABLE `chat_v1_1663684348`.`tbl_room_has_participants` ADD CONSTRAINT `fk_tbl_room_has_participants_room` FOREIGN KEY (`room`)
REFERENCES `chat_v1_1663684348`.`tbl_room` (`id`) ON DELETE NO ACTION   ON UPDATE NO ACTION;
  
ALTER TABLE `chat_v1_1663684348`.`tbl_room_has_participants` ADD INDEX `fk_tbl_room_has_participants_participant_idx` (`participant` ASC);
ALTER TABLE `chat_v1_1663684348`.`tbl_room_has_participants` ADD CONSTRAINT `fk_tbl_room_has_participants_participant` FOREIGN KEY (`participant`)
REFERENCES `chat_v1_1663684348`.`tbl_participant` (`id`) ON DELETE NO ACTION   ON UPDATE NO ACTION;
  