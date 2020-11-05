-- 用户表
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`           int(11)        NOT NULL AUTO_INCREMENT COMMENT '主键',

    `unionid`      char(80) COMMENT 'UNIONID',
    `openid`       char(80) COMMENT 'OPENID',
    `session_key`  char(80) COMMENT 'OPENID',
    `nickname`     char(80)       NOT NULL DEFAULT '' COMMENT '用户名',
    `avatar`       varchar(255)   NOT NULL DEFAULT '' COMMENT '头像',
    `sex`          tinyint(1)     NOT NULL DEFAULT 0 COMMENT '性别,0-未知 1-男 2-女',
    `mobile`       char(18)       NOT NULL DEFAULT '' COMMENT '手机',
    `introduction` char(80)       NOT NULL DEFAULT '' COMMENT '简介',
    `balance`      decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '余额',
    `coin`         decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '平台币',
    `password`     varchar(255)   NOT NULL DEFAULT '' COMMENT '密码',
    `ssk`          varchar(255)   NOT NULL DEFAULT '' COMMENT 'sessionKey',

    `created_on`   int(10) unsigned        DEFAULT '0' COMMENT '写入时间',
    `modified_on`  int(10) unsigned        DEFAULT '0' COMMENT '更新时间',
    `deleted_at`   timestamp      NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE (`mobile`),
    KEY (`openid`),
    KEY (`unionid`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='用户表';

-- 首页轮播
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner`
(
    `id`          int(11)      NOT NULL AUTO_INCREMENT,
    `name`        varchar(50)           DEFAULT NULL COMMENT 'Banner名称',
    `image_url`   varchar(255) NOT NULL DEFAULT '' COMMENT '轮播图地址',
    `key_word`    int(11)      NOT NULL COMMENT '执行关键字，根据不同的type含义不同',
    `type`        tinyint(4)   NOT NULL DEFAULT '1' COMMENT '跳转类型，可能导向商品，可能导向专题，可能导向其他。0，无导向；1：导向商品;2:导向专题',
    `created_on`  int(10) unsigned      DEFAULT '0' COMMENT '写入时间',
    `modified_on` int(10) unsigned      DEFAULT '0' COMMENT '更新时间',
    `deleted_at`  timestamp    NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 6
  DEFAULT CHARSET = utf8mb4 COMMENT ='banner表';

INSERT INTO `banner`
VALUES ('1', 'banner1', 'https://www.pexels.com/zh-cn/photo/2265082/', 5, 1, null, null, null);
INSERT INTO `banner`
VALUES ('2', 'banner2', 'https://www.pexels.com/zh-cn/photo/h2o-1295138/', 5, 1, null, null, null);
INSERT INTO `banner`
VALUES ('3', 'banner3', 'https://www.pexels.com/zh-cn/photo/346529/', 6, 1, null, null, null);

-- 商品表
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`
(
    `id`            int(11)       NOT NULL AUTO_INCREMENT COMMENT '主键',

    `title`         char(80)      NOT NULL DEFAULT '' COMMENT '商品名',
    `cover`         varchar(255)  NOT NULL DEFAULT '' COMMENT '商品封面图',
    `carousel`      text COMMENT '商品图集',
    `brand`         char(40)      NOT NULL DEFAULT '' COMMENT '品牌名',
    `series`        char(40)      NOT NULL DEFAULT '' COMMENT '系列名',
    `price`         decimal(6, 2) NOT NULL DEFAULT 0.00 COMMENT '零售价',
    `selling_price` decimal(6, 2) NOT NULL DEFAULT 0.00 COMMENT '售价',
    `cost`          decimal(6, 2) NOT NULL DEFAULT 0.00 COMMENT '成本价',
    `tags`          char(80)      NOT NULL DEFAULT '' COMMENT '标签,多个逗号分隔',
    `sales`         int(8)        NOT NULL DEFAULT 0 COMMENT '销量',
    `inventory`     int(8)        NOT NULL DEFAULT 0 COMMENT '库存',
    `status`        tinyint(1)    NOT NULL DEFAULT 0 COMMENT '状态,0-下架 1-上架',
    `on_sale`       tinyint(1)    NOT NULL DEFAULT 0 COMMENT '状态,0-未折扣 1-折扣中',

    `created_on`    int(10) unsigned       DEFAULT '0' COMMENT '写入时间',
    `modified_on`   int(10) unsigned       DEFAULT '0' COMMENT '更新时间',
    `deleted_at`    timestamp     NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY (`series`),
    KEY (`selling_price`),
    KEY (`cost`),
    KEY (`title`),
    KEY (`tags`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='商品表';


-- 库存表
DROP TABLE IF EXISTS `inventories`;
CREATE TABLE `inventories`
(
    `p_id`        int(11)   NOT NULL DEFAULT 0 COMMENT '商品id',

    `f`           int(8)    NOT NULL DEFAULT 0 COMMENT 'F',
    `xxs`         int(8)    NOT NULL DEFAULT 0 COMMENT 'XXS',
    `xs`          int(8)    NOT NULL DEFAULT 0 COMMENT 'XS',
    `s`           int(8)    NOT NULL DEFAULT 0 COMMENT 'S',
    `m`           int(8)    NOT NULL DEFAULT 0 COMMENT 'M',
    `l`           int(8)    NOT NULL DEFAULT 0 COMMENT 'L',
    `xl`          int(8)    NOT NULL DEFAULT 0 COMMENT 'XL',
    `xxl`         int(8)    NOT NULL DEFAULT 0 COMMENT 'XXL',
    `s35`         int(8)    NOT NULL DEFAULT 0 COMMENT '35',
    `s36`         int(8)    NOT NULL DEFAULT 0 COMMENT '36',
    `s37`         int(8)    NOT NULL DEFAULT 0 COMMENT '37',
    `s38`         int(8)    NOT NULL DEFAULT 0 COMMENT '38',
    `s39`         int(8)    NOT NULL DEFAULT 0 COMMENT '39',
    `s40`         int(8)    NOT NULL DEFAULT 0 COMMENT '40',
    `s405`        int(8)    NOT NULL DEFAULT 0 COMMENT '40.5',
    `s41`         int(8)    NOT NULL DEFAULT 0 COMMENT '41',
    `s415`        int(8)    NOT NULL DEFAULT 0 COMMENT '41.5',
    `s42`         int(8)    NOT NULL DEFAULT 0 COMMENT '42',
    `s425`        int(8)    NOT NULL DEFAULT 0 COMMENT '42.5',
    `s43`         int(8)    NOT NULL DEFAULT 0 COMMENT '43',
    `s435`        int(8)    NOT NULL DEFAULT 0 COMMENT '43.5',
    `s44`         int(8)    NOT NULL DEFAULT 0 COMMENT '44',
    `s445`        int(8)    NOT NULL DEFAULT 0 COMMENT '44.5',
    `s45`         int(8)    NOT NULL DEFAULT 0 COMMENT '45',
    `s46`         int(8)    NOT NULL DEFAULT 0 COMMENT '46',
    `s47`         int(8)    NOT NULL DEFAULT 0 COMMENT '47',


    `created_on`  int(10) unsigned   DEFAULT '0' COMMENT '写入时间',
    `modified_on` int(10) unsigned   DEFAULT '0' COMMENT '更新时间',
    `deleted_at`  timestamp NULL COMMENT '删除时间',

    CONSTRAINT p2i_id_fk FOREIGN KEY (p_id) REFERENCES products (id)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='库存表';


-- 收藏表
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`
(
    `user_id`    int(11) NOT NULL COMMENT '用户id',
    `p_id`       int(11) NOT NULL COMMENT '商品主键',

    `created_on` int(10) unsigned DEFAULT '0' COMMENT '写入时间',
    KEY (`user_id`),
    KEY (`p_id`),
    CONSTRAINT p_id_fk FOREIGN KEY (p_id) REFERENCES products (id)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='收藏表';

-- 收货地址
DROP TABLE IF EXISTS `addresses`;
CREATE TABLE `addresses`
(
    `id`            int(11)    NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`       char(32)   NOT NULL DEFAULT '' COMMENT '用户id',
    `province_name` char(40)   NOT NULL DEFAULT '' COMMENT '国标收货地址第一级地址',
    `city_name`     char(40)   NOT NULL DEFAULT '' COMMENT '国标收货地址第二级地址',
    `county_name`   char(40)   NOT NULL DEFAULT '' COMMENT '国标收货地址第三级地址',
    `detail_info`   char(200)  NOT NULL DEFAULT '' COMMENT '详细收货地址信息',
    `postal_code`   char(6)    NOT NULL DEFAULT '' COMMENT '邮编',
    `national_code` char(6)    NOT NULL DEFAULT '' COMMENT '收货地址国家码',
    `user_name`     char(40)   NOT NULL DEFAULT '' COMMENT '收货人姓名',
    `tel_number`    char(18)   NOT NULL DEFAULT '' COMMENT '收货人手机号码',
    `default`       tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否默认,0-否 1-是',
    PRIMARY KEY (`id`),
    KEY (`user_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='收货地址';


-- 购物车
DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts`
(
    `id`          int(11)    NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`     int(11)    NOT NULL DEFAULT 0 COMMENT '用户id',
    `p_id`        int(11)    NOT NULL DEFAULT 0 COMMENT '商品主键',
    `o_id`        int(11)    NOT NULL DEFAULT 0 COMMENT '订单主键',
    `number`      int(4)     NOT NULL DEFAULT 0 COMMENT '数量',
    `status`      tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态,0-未生成订单 1-已生成订单',
    `created_on`  int(10) unsigned    DEFAULT '0' COMMENT '写入时间',
    `modified_on` int(10) unsigned    DEFAULT '0' COMMENT '更新时间',
    `deleted_at`  timestamp  NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE (`p_id`),
    KEY (`user_id`),
    KEY (`o_id`),
    KEY (`status`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='购物车';


-- 订单表
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`
(
    `id`            int(11)        NOT NULL AUTO_INCREMENT COMMENT '主键',
--
    `user_id`       int(11)        NOT NULL DEFAULT 0 COMMENT '用户id',
-- 收获地址
    `province_name` char(40)       NOT NULL DEFAULT '' COMMENT '省份',
    `city_name`     char(40)       NOT NULL DEFAULT '' COMMENT '城市',
    `county_name`   char(40)       NOT NULL DEFAULT '' COMMENT '区县',
    `detail_info`   char(200)      NOT NULL DEFAULT '' COMMENT '详细地址',
    `postal_code`   char(6)        NOT NULL DEFAULT '' COMMENT '邮编',
    `user_name`     char(40)       NOT NULL DEFAULT '' COMMENT '收货人姓名',
    `tel_number`    char(18)       NOT NULL DEFAULT '' COMMENT '联系电话',
-- 物流信息
    `express_title` char(40)       NOT NULL DEFAULT '' COMMENT '物流公司',
    `express_code`  char(20)       NOT NULL DEFAULT '' COMMENT '物流编号',
    `express_no`    char(40)       NOT NULL DEFAULT '' COMMENT '物流单号',
    `express_time`  int(10) unsigned        DEFAULT '0' COMMENT '发货时间',
-- 订单信息
    `total`         decimal(12, 2) NOT NULL DEFAULT 0 COMMENT '汇总金额',
    `sum_pay`       decimal(12, 2) NOT NULL DEFAULT 0 COMMENT '结算金额',
    `status`        tinyint(1)     NOT NULL DEFAULT 0 COMMENT '状态,0-未结算 1-已结算(待发货) 2-已发货(待收货) 3-已完成 9-异常',
    `created_on`    int(10) unsigned        DEFAULT '0' COMMENT '写入时间',
    `modified_on`   int(10) unsigned        DEFAULT '0' COMMENT '更新时间',
    `pushed_at`     int(10) unsigned        DEFAULT '0' COMMENT ' ',
    `deleted_at`    timestamp      NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY (`user_id`),
    KEY (`status`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='订单表';


-- 微信支付表
DROP TABLE IF EXISTS `wxorders`;
CREATE TABLE `wxorders`
(
    `id`             int(11)        NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`        int(11)        NOT NULL DEFAULT 0 COMMENT '用户id',
    `transaction_id` char(32)       NOT NULL DEFAULT '' COMMENT '微信支付订单号',
    `out_trade_no`   char(32)       NOT NULL DEFAULT '' COMMENT '商户订单号',
    `nonce_str`      char(32)       NOT NULL DEFAULT '' COMMENT '随机字符串',
    `sign`           char(32)       NOT NULL DEFAULT '' COMMENT '签名',
    `body`           char(128)      NOT NULL DEFAULT '' COMMENT '商品描述',
    `detail`         varchar(255)   NOT NULL DEFAULT '' COMMENT '商品详情',
    `attach`         char(127)      NOT NULL DEFAULT '' COMMENT '附加数据',
    `sum_pay`        decimal(12, 2) NOT NULL DEFAULT 0 COMMENT '结算金额',
    `total_fee`      int            NOT NULL DEFAULT 0 COMMENT '标价金额,单位为分',
    `created_on`     int(10) unsigned        DEFAULT '0' COMMENT '写入时间',
    `modified_on`    int(10) unsigned        DEFAULT '0' COMMENT '更新时间',
    `pushed_at`      int(10) unsigned        DEFAULT '0' COMMENT '推送时间',
    `deleted_at`     timestamp      NULL COMMENT '完成时间',
    PRIMARY KEY (`id`),
    KEY (`user_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='微信支付表';


-- 管理员表
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`
(
    `id`          int(11)      NOT NULL AUTO_INCREMENT COMMENT '主键',

    `nickname`    char(80)     NOT NULL DEFAULT '' COMMENT '用户名',
    `avatar`      varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
    `username`    char(18)     NOT NULL DEFAULT '' COMMENT 'Username',
    `password`    varchar(255) NOT NULL DEFAULT '' COMMENT '密码',

    `created_on`  int(10) unsigned      DEFAULT '0' COMMENT '写入时间',
    `modified_on` int(10) unsigned      DEFAULT '0' COMMENT '更新时间',
    `deleted_at`  timestamp    NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='管理员表';
-- 插入管理员
INSERT INTO `admins`(nickname, avatar, username, password)
VALUES ('YFjie', 'xxxxxx', 'admin', '123456')