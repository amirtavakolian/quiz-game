-- +migrate Up
CREATE TABLE `profiles` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `fullname` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci NULL,
  `avatar` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci NULL,
  `bio` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_persian_ci NULL,
  `player_id` BIGINT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uq_profiles_player_id` (`player_id`),   -- 🔥 اضافه شد
  CONSTRAINT `fk_profiles_player` FOREIGN KEY (`player_id`) REFERENCES `players` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_persian_ci
  ROW_FORMAT = DYNAMIC;

-- +migrate Down
DROP TABLE IF EXISTS `profiles`;
