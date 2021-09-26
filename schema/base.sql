DROP TABLE IF EXISTS `models`;
DROP TABLE IF EXISTS `parents`;

CREATE TABLE `parents`(
    `id` int(8) NOT NULL AUTO_INCREMENT comment '',
    `body` varchar(1024) NOT NULL DEFAULT '' comment '',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '',
    `deleted_at` datetime comment '',
    PRIMARY KEY (`id`)
) comment='';

CREATE TABLE `models`(
    `id` int(8) NOT NULL AUTO_INCREMENT comment '',
    `body` varchar(1024) NOT NULL DEFAULT '' comment '',
    `parent_id` int(8) NOT NULL comment '',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '',
    `deleted_at` datetime comment '',
    PRIMARY KEY (`id`),
    FOREIGN KEY (parent_id) REFERENCES parents(id)
) comment='';

