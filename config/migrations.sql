CREATE TABLE `user`
(
    `id`                 bigint(10) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name`               varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `mobile`             varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '0' COMMENT '手机号',
    `mobile_verified_at` timestamp NULL DEFAULT NULL,
    `password`           varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
    `remember_token`     varchar(100) COLLATE utf8mb4_unicode_ci          DEFAULT NULL COMMENT '找回token',
    `created_at`         datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`         datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `mobile_unique` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

CREATE TABLE `user_token`
(
    `token`      varchar(191) NOT NULL COMMENT '令牌',
    `user_id`    varchar(191) NOT NULL COMMENT '用户ID',
    `expires_at` varchar(191) NOT NULL COMMENT '过期时间',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`token`),
    KEY          `user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户授权token';


CREATE TABLE `user_company`
(
    `id`               bigint(10) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id`          bigint(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '用户ID',
    `name`             varchar(50) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '公司名称',
    `brief`            text COLLATE utf8mb4_unicode_ci         NOT NULL COMMENT '公司简介',
    `location`         varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公司所在地',
    `address`          varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公司详细地址',
    `business_license` varchar(300) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '营业执照',
    `credit_code`      varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '企业统一征信代码',
    `created_at`       datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`       datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY                `user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户公司信息';

CREATE TABLE `sms_captcha`
(
    `id`         bigint(10) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `mobile`     varchar(191) NOT NULL DEFAULT '' COMMENT '手机号',
    `captcha`    varchar(11)  NOT NULL DEFAULT '' COMMENT '验证码',
    `invalid`    tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否失效 0未失效 1已失效',
    `expires_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '到期时间',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY          `mobile_index` (`mobile`),
    KEY          `captcha_index` (`captcha`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='短信验证码';

CREATE TABLE `product`
(
    `id`         bigint(10) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name`       varchar(50)    NOT NULL DEFAULT '' COMMENT '商品名称',
    `type`       tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型:1单件商品,2多规格商品,3定制商品',
    `price`      decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '价格',
    `status`     tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态:1正常,2删除',
    `created_at` datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品';

CREATE TABLE `product_specification`
(
    `id`         bigint(10) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `product_id` bigint(10) NOT NULL DEFAULT '0' COMMENT '商品ID',
    `name`       varchar(50) NOT NULL DEFAULT '' COMMENT '规格名称(材质名称或工艺名称)',
    `spec_id`    bigint(10) NOT NULL DEFAULT '0' COMMENT '父ID',
    `type`       tinyint(1) NOT NULL COMMENT '类型:1材质,2工艺3其他',
    `setting`    json        NOT NULL COMMENT '配置项:材质和工艺的扩展配置',
    `created_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY          `product_id_index` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品规格';
