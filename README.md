
```mysql
DROP TABLE IF EXISTS `bs_article`;
CREATE TABLE `bs_article` (
    `article_id` bigint(11) NOT NULL COMMENT '文章ID',
    `author_id` bigint(8) unsigned NOT NULL DEFAULT '0' COMMENT '文章作者id',
    `article_url` varchar(100) DEFAULT NULL COMMENT '文章跳转链接',
    `article_show` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '文章是否显示，0为否，1为是，默认为1',
    `article_sort` tinyint(3) unsigned NOT NULL DEFAULT '255' COMMENT '文章排序',
    `article_title` varchar(100) DEFAULT NULL COMMENT '文章标题',
    `article_content` text COMMENT '内容',
    `article_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章发布时间',
    `article_pic` varchar(255) NOT NULL DEFAULT '' COMMENT '文章主图',
    `created_at` datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at` datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    PRIMARY KEY (`article_id`),
    KEY `author_id` (`author_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='文章表';

```

```mysql
DROP TABLE IF EXISTS `bs_article_category`;
CREATE TABLE `bs_article_category` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `article_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章id',
    `category_id` int(11) unsigned NOT NULL COMMENT '文章标签id',
    `created_at` datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at` datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    PRIMARY KEY (`id`),
    KEY `article_id` (`article_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='文章标签关联表';
```

```mysql
DROP TABLE IF EXISTS `bs_category`;
CREATE TABLE `bs_category` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `name`  varchar(100) DEFAULT NULL COMMENT '标签名称',
    `level` tinyint DEFAULT 1 COMMENT '标签层级',
    `parent_id` int(11) DEFAULT 0 COMMENT '标签父级id',
    `status` int(11) unsigned NOT NULL COMMENT '标签状态',
    `created_at` datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at` datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='文章标签表';
```

```mysql
DROP TABLE IF EXISTS `bs_author`;
create table bs_author
(
    id          int  comment '用户uid' primary key,
    name      varchar(20)                                 not null comment '作者名称',
    mobile       varchar(13)                                 not null comment '作者手机号',
    country_code smallint unsigned default '86'              not null comment '手机号国家区号',
    status        tinyint unsigned  default '1'               not null comment '用户状态: 1启用 2封存3注销',
    created_at   datetime          default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at   datetime          default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint un_account unique (name) comment '登录账号',
    constraint un_phone unique (mobile) comment '手机号'
)
    comment '作者表' charset = utf8;
```

```mysql
DROP TABLE IF EXISTS `bs_comment`;
create table bs_comment
(
    `id`          bigint(8)  comment '评论id' primary key,
    `article_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章id',
    `commenter` varchar(100) DEFAULT '' COMMENT '评论者名称',
    `content` text COMMENT '评论内容',
    `parent_id` bigint(8) DEFAULT 0 COMMENT '回复评论id',
    `created_at`   datetime          default CURRENT_TIMESTAMP not null comment '创建时间',
    `updated_at`   datetime          default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间'
)
    comment '评论表' charset = utf8;
```