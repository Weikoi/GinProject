CREATE TABLE IF NOT EXISTS `books_info`(
    `book_id` INT(10) AUTO_INCREMENT primary key ,
    `book_name` VARCHAR(40) NOT NULL,
    `press_name` VARCHAR(40) NOT NULL comment '出版社名称',
    `pub_time` date NOT NULL comment '出版日期',
    `author` VARCHAR(40) NOT NULL comment '作者名字',
    `total_num` INT(10) default 0 comment '馆存数量',
    `on_shelf_num` INT(10) default 0 comment '在架数量',
    `desc` text null comment '图书简介',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `users_info`(
    `user_id` INT(10) AUTO_INCREMENT primary key ,
    `user_name` VARCHAR(40) NOT NULL,
    `password` VARCHAR(40) NOT NULL,
    `auth` INT(1) default 1 COMMENT '用户权限 0:admin 1:normal 2:paid -1:forbidden',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `records`(
    `serial_no` INT(10) AUTO_INCREMENT primary key,
    `status`  INT(1) default 0 COMMENT '借阅状态：0:借阅中， 1:已归还',
    `user_id` INT(10) not NULL ,
    `book_id` INT(10) not NULL ,
    `borrow_time` timestamp not NULL,
    `return_time` timestamp NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
