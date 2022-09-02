-- add children for project, foldersformds, foldersforfiles
-- 以下四个字段用来表示当前节点下一层的文件节点，存储形式都是 <id>-<is_folder> 形式
-- id 表示文件的 id，is-folder 表示是不是文件
-- 这里不需要标志是文件和文档，在文件夹下的只能是文件，文档夹下的只能是文档
ALTER TABLE `projects` ADD `file_children` TEXT DEFAULT NULL;

ALTER TABLE `projects` ADD `doc_children` TEXT DEFAULT NULL;

ALTER TABLE `foldersforfiles` ADD `children` TEXT DEFAULT NULL;

ALTER TABLE `foldersformds` ADD `children` TEXT DEFAULT NULL;

-- add father_id for doc, file, folder
-- 以下四个字段用来标志当前节点的父节点
-- father_id 为 0 表示父节点是项目
ALTER TABLE `docs` ADD `father_id` int(11) DEFAULT 0;

ALTER TABLE `files` ADD `father_id` int(11) DEFAULT 0;

ALTER TABLE `foldersformds` ADD `father_id` int(11) DEFAULT 0;

ALTER TABLE `foldersforfiles` ADD `father_id` int(11) DEFAULT 0;

-- add index for user2projects index
ALTER TABLE `user2projects` ADD UNIQUE INDEX(`user_id`,`project_id`);

-- add last_edit_time for doc
ALTER TABLE `docs` ADD `last_edit_time` varchar(30) DEFAULT NULL;

-- add project soft delete
-- 使用 gorm 提供好的软删除
ALTER TABLE `projects` ADD `deleted_at` datetime DEFAULT NULL;

-- add creator_id for projects
ALTER TABLE `projects` ADD COLUMN `creator_id` int(11);

-- user2file2 -> attentions
RENAME TABLE `user2files` TO `attentions`;
ALTER TABLE `attentions` ADD COLUMN `time_day` varchar(20) DEFAULT NULL;
ALTER TABLE `attentions` ADD COLUMN `time_hm` varchar(20) DEFAULT NULL;

-- add team_id for applys
ALTER TABLE `applys` ADD COLUMN `team_id` int(11) DEFAULT NULL;

-- add key for user2projects
ALTER TABLE `user2projects` ADD KEY  `user_id` (`user_id`);

-- add key for attention
ALTER TABLE `attentions` ADD UNIQUE KEY `idx_file_kind_user` (`user_id`, `file_id`, `file_kind`) ;

-- -----------------------------
-- Table structure for comments
-- -----------------------------
-- 项目评论表，从原comments表迁移出数据，分别为file doc status生成共三个comments表
DROP TABLE IF EXISTS `comments_files`;
CREATE TABLE `comments_files` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `kind` int(11) DEFAULT NULL COMMENT '表中kind用来区分是否二级评论',
    `content` text,
    `time` varchar(50) DEFAULT NULL,
    `creator` int(11) DEFAULT NULL,
    `target_id` int(11) DEFAULT NULL,
    `re` tinyint(1) DEFAULT NULL COMMENT '标志是否删除，0-未删除 1-删除 删除时只要将 re 置为 1',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `comments_docs`;
CREATE TABLE `comments_docs` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `kind` int(11) DEFAULT NULL,
    `content` text,
    `time` varchar(50) DEFAULT NULL,
    `creator` int(11) DEFAULT NULL,
    `target_id` int(11) DEFAULT NULL,
    `re` tinyint(1) DEFAULT NULL COMMENT '标志是否删除，0-未删除 1-删除 删除时只要将 re 置为 1',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `comments_status`;
CREATE TABLE `comments_status` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `kind` int(11) DEFAULT NULL,
    `content` text,
    `time` varchar(50) DEFAULT NULL,
    `creator` int(11) DEFAULT NULL,
    `target_id` int(11) DEFAULT NULL,
    `re` tinyint(1) DEFAULT NULL COMMENT '标志是否删除，0-未删除 1-删除 删除时只要将 re 置为 1',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
