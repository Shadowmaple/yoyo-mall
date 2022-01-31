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

INSERT INTO `user` (`nickname`, `username`, `password`, `role`, `create_time`) VALUES('admin', 'admin', '', 1, now());

-- 分类
-- 数据量小，不建索引

CREATE TABLE IF NOT EXISTS `category` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL DEFAULT '',
    `image` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '类目图片',

    `parent_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父类目id，0为根类目',

    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


INSERT INTO `category` (`name`, `parent_id`, `image`) VALUES
('文学', 0, ''), ('人文社科', 0, ''), ('经管', 0, ''), ('艺术', 0, ''), ('科技', 0, ''), ('教育', 0, ''), ('生活', 0, ''), ('成功/励志', 0, ''), ('童书', 0, ''),
('小说', 1, 'https://img9.doubanio.com/view/subject/s/public/s29350294.jpg'),
('文学', 1, 'https://img9.doubanio.com/view/subject/m/public/s3552626.jpg'),
('青春文学', 1, 'https://img1.doubanio.com/view/subject/m/public/s29766608.jpg'),
('传记', 1, 'https://img1.doubanio.com/view/subject/m/public/s33699918.jpg'),
('散文', 1, 'https://img9.doubanio.com/view/subject/m/public/s25616816.jpg'),
('动漫/幽默', 1, 'https://img1.doubanio.com/view/subject/m/public/s2206907.jpg'),
('纪实', 1, 'https://img9.doubanio.com/view/subject/m/public/s29111585.jpg'),
('古诗词', 1, 'https://img1.doubanio.com/view/subject/m/public/s32318099.jpg'),
('历史', 2, 'https://img9.doubanio.com/view/subject/m/public/s1800355.jpg'),
('哲学/宗教', 2, 'https://img3.doubanio.com/view/subject/m/public/s29897220.jpg'),
('文化', 2, 'https://img9.doubanio.com/view/subject/s/public/s29347294.jpg'),
('社会科学', 2, 'https://img1.doubanio.com/view/subject/m/public/s33841867.jpg'),
('心理学', 2, 'https://img2.doubanio.com/view/subject/m/public/s28338983.jpg'),
('法律', 2, 'https://img2.doubanio.com/view/subject/m/public/s33655741.jpg'),
('政治/军事', 2, 'https://img2.doubanio.com/view/subject/m/public/s27269441.jpg'),
('经济', 3, 'https://img9.doubanio.com/view/subject/m/public/s27780875.jpg'),
('管理', 3, 'https://img9.doubanio.com/view/subject/m/public/s29113625.jpg'),
('投资理财', 3, 'https://img2.doubanio.com/view/subject/m/public/s3354143.jpg'),
('摄影', 4, 'https://img9.doubanio.com/view/subject/s/public/s21942845.jpg'),
('绘画', 4, 'https://img1.doubanio.com/view/subject/m/public/s33664998.jpg'),
('书法篆刻', 4, 'https://img9.doubanio.com/view/subject/m/public/s33764806.jpg'),
('音乐', 4, 'https://img2.doubanio.com/view/subject/m/public/s28845543.jpg'),
('舞蹈', 4, 'https://img9.doubanio.com/view/subject/m/public/s33552965.jpg'),
('科普', 5, 'https://img9.doubanio.com/view/subject/m/public/s9111416.jpg'),
('计算机', 5, 'https://img1.doubanio.com/view/subject/m/public/s32513229.jpg'),
('建筑', 5, 'https://img9.doubanio.com/view/subject/m/public/s29667115.jpg'),
('医学', 5, 'https://img2.doubanio.com/view/subject/m/public/s28117212.jpg'),
('农林', 5, 'https://img9.doubanio.com/view/subject/m/public/s6824945.jpg'),
('自然科学', 5, 'https://img9.doubanio.com/view/subject/m/public/s33533954.jpg'),
('工业', 5, 'https://img3.doubanio.com/view/subject/m/public/s33942780.jpg'),
('中小学教辅', 6, 'https://img2.doubanio.com/view/subject/m/public/s29109031.jpg'),
('考试', 6, 'https://img9.doubanio.com/view/subject/m/public/s33322276.jpg'),
('外语', 6, 'https://img9.doubanio.com/view/subject/m/public/s29731505.jpg'),
('教材', 6, 'https://img2.doubanio.com/view/subject/m/public/s28340131.jpg'),
('工具书', 6, 'https://img9.doubanio.com/view/subject/m/public/s27287585.jpg'),
('运动', 7, 'https://img3.doubanio.com/view/subject/m/public/s33450470.jpg'),
('保健', 7, 'https://img9.doubanio.com/view/subject/m/public/s33744385.jpg'),
('旅游', 7, 'https://img9.doubanio.com/view/subject/m/public/s22702375.jpg'),
('两性', 7, 'https://img9.doubanio.com/view/subject/m/public/s28045305.jpg'),
('亲子/家教', 7, 'https://img9.doubanio.com/view/subject/m/public/s29624136.jpg'),
('美妆', 7, 'https://img3.doubanio.com/view/subject/m/public/s28302750.jpg'),
('手工', 7, 'https://img1.doubanio.com/view/subject/m/public/s28876288.jpg'),
('美食', 7, 'https://img2.doubanio.com/view/subject/m/public/s29140511.jpg'),
('心灵与修养', 8, 'https://img1.doubanio.com/view/subject/m/public/s29237648.jpg'),
('人际交往', 8, 'https://img1.doubanio.com/view/subject/m/public/s33901897.jpg'),
('成功/激励', 8, 'https://img2.doubanio.com/view/subject/m/public/s33963701.jpg'),
('人生哲学', 8, 'https://img2.doubanio.com/view/subject/m/public/s29962521.jpg'),
('口才/演讲/辩论', 8, 'https://img1.doubanio.com/view/subject/m/public/s9346487.jpg'),
('性格与习惯', 8, 'https://img9.doubanio.com/view/subject/m/public/s29714854.jpg'),
('童书', 9, 'https://img1.doubanio.com/view/subject/m/public/s33539207.jpg');


-- 商品

CREATE TABLE IF NOT EXISTS `product` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '商品标题',
    `book_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '书名',
    `author` VARCHAR(255) NOT NULL DEFAULT '',
    `publisher` VARCHAR(255) NOT NULL DEFAULT '',
    `price` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '原价',
    `cur_price` DECIMAL(10, 2) NOT NULL DEFAULT 0 COMMENT '优惠价',
    `stock` INT NOT NULL DEFAULT 0 COMMENT '库存',
    `detail` TEXT COMMENT '详情信息',
    `images` TEXT COMMENT '图片，分号分割',
    `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态，0正常，1下架',
    `publish_time` DATETIME DEFAULT NULL COMMENT '出版时间',
    `create_time` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `delete_time` DATETIME DEFAULT NULL,

    `cid` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '一级类目',
    `cid2` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '二级类目',

    PRIMARY KEY (`id`),
    KEY `idx_cid` (`cid`, `cid2`)
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
    `refund` TEXT COMMENT '退货退款内容（暂时占位）',
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
    `level` TINYINT NOT NULL DEFAULT 0 COMMENT '0好评，1一般，2差评',
    `is_anoymous` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否匿名评价',
    `pictures` TEXT COMMENT '评价图片，分号分割',
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
    `reply_user_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '向谁回复，暂时占位',

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
    `province` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '省',
    `city` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '市',
    `district` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '区县',
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
    `pictures` TEXT COMMENT '图片，分号分割',
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
