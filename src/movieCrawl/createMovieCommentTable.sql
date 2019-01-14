CREATE TABLE `movie_comment`(
  `id` INT(10) UNSIGNED NOT null auto_increment,
  `movie_name` varchar(100) comment '电影名',
  `user_name` varchar(20) comment '用户名',
  `movie_type` varchar(50) comment '电影类型',
  `comment_star` tinyint(1) comment '用户所给星级',
  `comment_time` timestamp not null default '0000-00-00 00:00:00' comment '评论时间',
  `movie_grade` varchar(5) comment '电影评分',
  `good_comment` varchar(10) comment '好评率',
  `middle_comment` varchar(10) comment '一般率',
  `bad_comment` varchar(10) comment '差评率',
  `short_comment` varchar(500) default '' comment '电影短评',
  `_create_time` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `_modify_time` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
  `_status` tinyint(1) default '1',
  primary key (`id`),
  key `idx_create_time` (`_create_time`),
  key `idx_modify_time` (`_modify_time`)
)engine = InnoDB AUTO_INCREMENT = 20 DEFAULT CHARSET = utf8 comment '电影短评表'