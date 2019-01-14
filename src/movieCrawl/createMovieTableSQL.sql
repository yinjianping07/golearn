CREATE TABLE `movie_info`(
  `id` INT(10) UNSIGNED NOT null auto_increment,
  `movie_id` int(11) unsigned not null comment '电影id',
  `movie_name` varchar(100) comment '电影名',
  `movie_picture` varchar(200) comment '电影图片',
  `movie_director` varchar(50) comment '电影导演',
  `movie_writer` varchar(50) comment '电影编剧',
  `movie_country` varchar(50) comment '电影产地',
  `movie_main_character` varchar(50) comment '电影主演',
  `movie_type` varchar(50) comment '电影类型',
  `movie_on_time` timestamp default '0000-00-00 00:00:00' comment '电影上映时间',
  `movie_span` varchar(20) comment '电影时长',
  `movie_grade` varchar(5) comment '电影评分',
  `remark` varchar(500) default '' comment '电影评论',
  `_create_time` timestamp not null default '0000-00-00 00:00:00' comment '创建时间',
  `_modify_time` timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
  `_status` tinyint(1) default '1',
  primary key (`id`),
  key `idx_movie_id` (`movie_id`),
  key `idx_create_time` (`_create_time`),
  key `idx_modify_time` (`_modify_time`)
)engine = InnoDB AUTO_INCREMENT = 20 DEFAULT CHARSET = utf8 comment '电影信息表'