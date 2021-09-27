DROP TABLE IF EXISTS `attributes`;
DROP TABLE IF EXISTS `entities`;

CREATE TABLE `entities`(
    `id` int(8) NOT NULL AUTO_INCREMENT comment '',
    `body` varchar(1024) NOT NULL DEFAULT '' comment '',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '',
    `deleted_at` datetime comment '',
    PRIMARY KEY (`id`)
) comment='';

CREATE TABLE `attributes`(
    `id` int(8) NOT NULL AUTO_INCREMENT comment '',
    `body` varchar(1024) NOT NULL DEFAULT '' comment '',
    `entity_id` int(8) NOT NULL comment '',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '',
    `deleted_at` datetime comment '',
    PRIMARY KEY (`id`),
    FOREIGN KEY (entity_id) REFERENCES entities(id)
) comment='';

