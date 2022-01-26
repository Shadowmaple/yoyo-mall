CREATE DATABASE IF NOT EXISTS `yoyo-mall`;

USE `yoyo-mall`;


CREATE TABLE IF NOT EXISTS `user` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `nickname` VARCHAR(255) NOT NULL DEFAULT '',
    `wechat_unique_id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '微信用户唯一id',
    `avatar` VARCHAR(255) NOT NULL DEFAULT '',
    `gender` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '0男，1女',
    `username` VARCHAR(255) NOT NULL DEFAULT '',
    `password` VARCHAR(255) NOT NULL DEFAULT '',
    `state` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '0正常，1失效',
    `role` TINYINT NOT NULL DEFAULT 0 COMMENT '0普通用户，1管理员，2商家',
    `create_time` DATETIME NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_wechat_id` (`wechat_unique_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 分类
-- 数据量小，不建索引

CREATE TABLE IF NOT EXISTS `category` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL DEFAULT '',
    `order` INT NOT NULL DEFAULT 0 COMMENT '排列序号，0最大',
    `image` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '类目图片',
    `create_time` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `delete_time` DATETIME DEFAULT NULL,

    `parent_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父类目id，0为根类目',

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 商品

CREATE TABLE IF NOT EXISTS `product` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '商品标题',
    `book_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '书名',
    `author` VARCHAR(255) NOT NULL DEFAULT '',
    `publisher` VARCHAR(255) NOT NULL DEFAULT '',
    `price` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '原价',
    `cur_price` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '优惠价',
    `stock` INT NOT NULL DEFAULT '' COMMENT '库存',
    `detail` TEXT NOT NULL DEFAULT '' COMMENT '详情信息',
    `images` TEXT NOT NULL DEFAULT '' COMMENT '图片，分号分割',
    `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态，0正常，1下架',
    `publish_time` DATETIME DEFAULT NULL COMMENT '出版时间',
    `create_time` DATETIME NOT NULL,
    `update_time` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `delete_time` DATETIME DEFAULT NULL,

    `cid` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '一级类目',
    `cid2` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '二级类目',

    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_cid` (`cid`, `cid2`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 收藏

CREATE TABLE IF NOT EXISTS `collection` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `create_time` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `delete_time` DATETIME DEFAULT NULL,

    `user_id` INT UNSIGNED NOT NULL DEFAULT 0,
    `product_id` INT UNSIGNED NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 购物车

CREATE TABLE IF NOT EXISTS `cart` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `num` SMALLINT NOT NULL DEFAULT 0 COMMENT '商品数量',
    `create_time` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `delete_time` DATETIME DEFAULT NULL,

    `user_id` INT UNSIGNED NOT NULL DEFAULT 0,
    `product_id` INT UNSIGNED NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 订单

CREATE TABLE IF NOT EXISTS `order` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态：0->待付款，1->待发货，2->待收货，3->待评价，4->交易完成，5->交易取消，6->退货中，7->交易关闭',
    `payment` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '实付金额',
    `freight` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '运费',
    `total_fee` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '应付总金额',
    `coupon` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '优惠金额',
    `receive_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '收件人',
    `receive_tel` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '收件人联系电话',
    `receive_addr` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '收件人地址',
    `refund` TEXT NOT NULL DEFAULT '' COMMENT '退货退款内容（暂时占位）',
    `order_code` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '订单编号',
    `create_time` DATETIME NOT NULL,
    `pay_time` DATETIME NOT NULL COMMENT '付款时间',
    `deliver_time` DATETIME NOT NULL COMMENT '发货时间',
    `confirm_time` DATETIME NOT NULL COMMENT '签收时间',

    `user_id` INT UNSIGNED NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 订单-商品

CREATE TABLE IF NOT EXISTS `order_product` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `num` INT NOT NULL DEFAULT 0,
    `price` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '单价-原价',
    `cur_price` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '单价-优惠价',
    `total_fee` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '总价',
    `image` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '封面图片',
    `create_time` DATETIME NOT NULL,

    `order_id` INT UNSIGNED NOT NULL DEFAULT 0,
    `product_id` INT UNSIGNED NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 物流

CREATE TABLE IF NOT EXISTS `logistics` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态：0->待发货，1->已发货待收货，2->已完成，3->已取消，4->退货中，5->退货完成',
    `content` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '物流信息',
    `create_time` DATETIME NOT NULL,

    `order_id` INT UNSIGNED NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 评价

CREATE TABLE IF NOT EXISTS `evaluation` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `content` VARCHAR(255) NOT NULL DEFAULT '',
    `score` TINYINT NOT NULL DEFAULT 0 COMMENT '评分',
    `rank` TINYINT NOT NULL DEFAULT 0 COMMENT '0好评，1一般，2差评',
    `is_anoymous` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否匿名评价',
    `pictures` TEXT NOT NULL DEFAULT '' COMMENT '评价图片，分号分割',
    `create_time` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `delete_time` DATETIME DEFAULT NULL,

    `order_id` INT UNSIGNED NOT NULL DEFAULT 0,
    `product_id` INT UNSIGNED NOT NULL DEFAULT 0,
    `user_id` INT UNSIGNED NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 评论

CREATE TABLE IF NOT EXISTS `comment` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `content` VARCHAR(255) NOT NULL DEFAULT '',
    `is_anoymous` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否匿名',
    `create_time` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `delete_time` DATETIME DEFAULT NULL,

    `user_id` INT UNSIGNED NOT NULL DEFAULT 0,
    `evaluation_id` INT UNSIGNED NOT NULL DEFAULT 0,
    `parent_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父评论id，默认为0，暂时占位',
    `reply_user_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '向谁回复，暂时占位'

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 点赞

CREATE TABLE IF NOT EXISTS `like` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `kind` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '0评价点赞，1评论点赞',
    `create_time` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,

    `user_id` INT UNSIGNED NOT NULL DEFAULT 0,
    `comment_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '评价/评论id',

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 优惠券

CREATE TABLE IF NOT EXISTS `coupon` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `discount` SMALLINT NOT NULL DEFAULT 0 COMMENT '折扣金额',
    `threshold` INT NOT NULL DEFAULT 0 COMMENT '满减门槛',
    `kind` TINYINT NOT NULL DEFAULT 0 COMMENT '优惠券种类，默认为0',
    `is_public` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否公共可领取，0不可，1可',
    `code` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '兑换码',
    `title` VARCHAR(255) NOT NULL DEFAULT '',
    `remain` INT NOT NULL DEFAULT 0 COMMENT '剩余数',
    `begin_time` DATETIME NOT NULL COMMENT '生效开始时间',
    `end_time` DATETIME NOT NULL COMMENT '生效截止时间',
    `grab_begin_time` DATETIME NOT NULL COMMENT '领取开始时间',
    `grab_end_time` DATETIME NOT NULL COMMENT '领取截止时间',
    `code_begin_time` DATETIME NOT NULL COMMENT '兑换开始时间',
    `code_end_time` DATETIME NOT NULL COMMENT '兑换截止时间',
    `create_time` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `delete_time` DATETIME DEFAULT NULL,

    `cid` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '一级类目',
    `cid2` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '二级类目',

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 用户-优惠券

CREATE TABLE IF NOT EXISTS `user_coupon` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `status` TINYINT NOT NULL COMMENT '使用状态：0未使用，1已使用',
    `access` TINYINT NOT NULL COMMENT '获取方式：0领取，1兑换码',
    `create_time` DATETIME NOT NULL COMMENT '获取时间',

    `user_id` INT UNSIGNED NOT NULL DEFAULT 0,
    `coupon_id` INT UNSIGNED NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 地址

CREATE TABLE IF NOT EXISTS `address` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '收货人',
    `tel` VARCHAR(255) NOT NULL DEFAULT '',
    `detail` VARCHAR(255) NOT NULL DEFAULT '',
    `is_default` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否是默认地址',
    `create_time` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `delete_time` DATETIME DEFAULT NULL,

    `user_id` INT UNSIGNED NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 用户反馈

CREATE TABLE IF NOT EXISTS `feedback` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `kind` TINYINT NOT NULL DEFAULT 0 COMMENT '反馈类型：0产品建议，1功能异常，2违规举报，3交易投诉',
    `content` VARCHAR(255) NOT NULL DEFAULT '',
    `pictures` TEXT NOT NULL DEFAULT '' COMMENT '图片，分号分割',
    `has_read` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否已读',
    `create_time` DATETIME NOT NULL,
    `read_time` DATETIME DEFAULT NULL,

    `user_id` INT UNSIGNED NOT NULL DEFAULT 0,

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 消息记录
-- 还没考虑好怎么设计

-- CREATE TABLE IF NOT EXISTS `message` (
--     `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
--     `content` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '消息内容',
--     `has_read` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否已读',
--     `kind` TINYINT NOT NULL DEFAULT 0 COMMENT '消息类型：0订单物流，1评论，2点赞',
--     `create_time` DATETIME NOT NULL,
--     `read_time` DATETIME DEFAULT NULL,

--     `pub_user_id` INT UNSIGNED NOT NULL DEFAULT 0,
--     `sub_user_id` INT UNSIGNED NOT NULL DEFAULT 0,
--     `order_id` INT UNSIGNED NOT NULL DEFAULT 0,
--     `evaluation_id` INT UNSIGNED NOT NULL DEFAULT 0,
--     `comment_id` INT UNSIGNED NOT NULL DEFAULT 0,

--     PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
