/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : gorm

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2018-11-07 17:06:10
*/

SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for `admin`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`            int(10) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`    timestamp        NULL     DEFAULT NULL,
    `updated_at`    timestamp        NULL     DEFAULT NULL,
    `deleted_at`    timestamp        NULL     DEFAULT NULL,
    `account`       varchar(50)      NOT NULL,
    `password`      char(32)         NOT NULL,
    `email_address` char(32)         NOT NULL,
    `phone_number`  char(32)         NOT NULL,
    `descript`      varchar(255)              DEFAULT '',
    `headico`       varchar(200)              DEFAULT '',
    `type`          varchar(20)      NOT NULL DEFAULT 'u',
    PRIMARY KEY (`id`),
    KEY `idx_user_deleted_at` (`deleted_at`),
    KEY `username` (`account`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `user`(created_at, updated_at, deleted_at, account, password, email_address, phone_number, descript,
                   headico,
                   type)
VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Admin@HUST.builder.com',
        '21232f297a57a5a743894a0e4a801fc3', '129843719@qq.com', '18821218478', 'admin of HUST
.builder.com', '/public/images/user2-160x160.jpg', 'admin');
INSERT INTO `user`(created_at, updated_at, deleted_at, account, password, email_address, phone_number, descript,
                   headico, type)
VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'asdasda@HUST.builder.com',
        '21232f297a57a5a743894a0e4a801fc3', '1221412@qq.com', '18124121411', 'user of HUST.builder.com',
        '/public/images/user2-160x160.jpg', 'user');
INSERT INTO `user`(created_at, updated_at, deleted_at, account, password, email_address, phone_number, descript,
                   headico, type)
VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Admin@WH-zhijianju.supervisor.com',
        '21232f297a57a5a743894a0e4a801fc3', '3251354123@qq.com', '13252523232', 'admin of WH-zhijianju.supervisor.com',
        '/public/images/user2-160x160.jpg', 'admin');
INSERT INTO `user`(created_at, updated_at, deleted_at, account, password, email_address, phone_number, descript,
                   headico, type)
VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Admin@zhongjian-1-ju.constructor.com',
        '21232f297a57a5a743894a0e4a801fc3', '1213112312@qq.com', '18124128478',
        'admin of zhongjian-1-ju.constructor.com', '/public/images/user2-160x160.jpg', 'admin');
INSERT INTO `user`(created_at, updated_at, deleted_at, account, password, email_address, phone_number, descript,
                   headico, type)
VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Admin@zhongjian-2-ju.constructor.com',
        '21232f297a57a5a743894a0e4a801fc3', '241123412@qq.com', '18124112418',
        'admin of zhongjian-2-ju.constructor.com', '/public/images/user2-160x160.jpg', 'admin');


-- ----------------------------
-- Table structure for `Company`
-- ----------------------------


DROP TABLE IF EXISTS `company`;
CREATE TABLE `company`
(
    `id`                    int(10) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`            timestamp        NULL DEFAULT NULL,
    `updated_at`            timestamp        NULL DEFAULT NULL,
    `deleted_at`            timestamp        NULL DEFAULT NULL,
    `company_name`          varchar(50)      NOT NULL,
    `company_owner_account` varchar(50)      NOT NULL,
    `company_build_year`    varchar(50)      NOT NULL,
    `email_address`         char(50)         NOT NULL,
    `phone_number`          char(50)         NOT NULL,
    `descript`              varchar(255)          DEFAULT '',
    `headico`               varchar(200)          DEFAULT '',
    `company_invest`        varchar(50)           DEFAULT '',
    `company_size`          varchar(50)           DEFAULT '',
    `company_location`      varchar(50)           DEFAULT '',
    `company_website`       varchar(50)           DEFAULT '',
    PRIMARY KEY (`id`),
    KEY `idx_user_deleted_at` (`deleted_at`),
    KEY `companyname` (`company_name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;


INSERT INTO `company`(created_at, updated_at, deleted_at, company_name, company_owner_account, company_build_year,
                      email_address, phone_number, descript, headico, company_invest, company_size, company_location,
                      company_website)
VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'SanFrancisco Power', 'Admin@HUST.builder.com', '1978',
        'SanPowerOfficial@gmail.com', '+86-219031213', 'Best power supplier in SanFancisco',
        '/public/adminlit/dist/img/user2-160x160.jpg', '50 million', '2000-5000 employees',
        'SanFrancisco Long Street 182', 'www.SanFanciscoPower.com');
INSERT INTO `company`(created_at, updated_at, deleted_at, company_name, company_owner_account, company_build_year,
                      email_address, phone_number, descript, headico, company_invest, company_size, company_location,
                      company_website)
VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Trustworth Energy Company', 'Admin@HUST.builder.com',
        '2007', 'TrustworthEnergyOffcial@gmail.com', '+86-32156243', 'Energy and power are the future',
        '/public/adminlit/dist/img/user2-160x160.jpg', '12 million', '100-1000 employees', 'Seatle Dann Street 78',
        'www.TrustworthEngergy.com');
INSERT INTO `company`(created_at, updated_at, deleted_at, company_name, company_owner_account, company_build_year,
                      email_address, phone_number, descript, headico, company_invest, company_size, company_location,
                      company_website)
VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Santa Claus Power', 'Admin@WH-zhijianju.supervisor.com',
        '2015', 'SantaClausPower@gmail.com', '+70-2193891213', 'Creative Energy means everything',
        '/public/adminlit/dist/img/user2-160x160.jpg', '50 million', '100-1000 employees',
        'Los Angeles Beach Street 89', 'www.SantaClausPower.com');
INSERT INTO `company`(created_at, updated_at, deleted_at, company_name, company_owner_account, company_build_year,
                      email_address, phone_number, descript, headico, company_invest, company_size, company_location,
                      company_website)
VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Believing Energy', 'Admin@WH-zhijianju.supervisor.com',
        '1984', 'Be-Energy@gmail.com', '+86-3546336',
        'Believe it or not, we can offer you the best and the cleanest power',
        '/public/adminlit/dist/img/user2-160x160.jpg', '100 million', '5000+ employees', 'New York City King road 17',
        'www.BelievingEnergy.com');

-- ----------------------------
-- Records of company


-- -----------------------------
-- Bids-------------------------


DROP TABLE IF EXISTS `bid`;
CREATE TABLE `bid`
(
    `id`                    int(10) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`            timestamp        NULL DEFAULT NULL,
    `updated_at`            timestamp        NULL DEFAULT NULL,
    `deleted_at`            timestamp        NULL DEFAULT NULL,
    `contract_id`           varchar(50)      NOT NULL,
    `contract_name`         varchar(50)      NOT NULL,
    `contract_version`      varchar(50)      NOT NULL,
    `contract_company_name` varchar(50)      NOT NULL,
    `contract_company_owner_account` varchar(50)      NOT NULL,
    `contract_company_owner_sig`  varchar(255)          DEFAULT '',
    `contract_details`      varchar(255)          DEFAULT '',
    `energy_type`           varchar(50)      NOT NULL,
    `energy_price`          varchar(50)      NOT NULL,
    `contract_last_time`    varchar(50)      NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_user_deleted_at` (`deleted_at`),
    KEY `contract_id` (`contract_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;


INSERT INTO `bid`(contract_id, contract_name, contract_version,
                  contract_company_name,contract_company_owner_account, contract_details, energy_type, energy_price,
                  contract_last_time)
VALUES ('0001', 'Water Power Project', '1.0', 'SanFrancisco Power','Admin@HUST.builder.com',
        'The water power project started 3 years ago, aiming to provide clean water power for electricity supply.',
        'water',
        '1.5',
        '12');

INSERT INTO `bid`(contract_id, contract_name, contract_version,
                  contract_company_name,contract_company_owner_account, contract_details, energy_type, energy_price, contract_last_time)
VALUES ('0002', 'Nuclear Supply Project', '1.0', 'SanFrancisco Power','Admin@HUST.builder.com',
        'Nuclear Power is good, nuclear power is strong, nuclear power is necessary',
        'nuclear',
        '1.2',
        '24');

INSERT INTO `bid`(contract_id, contract_name, contract_version,
                  contract_company_name,contract_company_owner_account, contract_details, energy_type, energy_price, contract_last_time)
VALUES ('0003', 'Windy Supply', '1.0', 'Believing Energy','Admin@WH-zhijianju.supervisor.com',
        'Spanish wind turbine manufacturer Siemens Gamesa Renewable Energy (SGRE) has secured a contract from Berkshire Hathaway Energy (BHE) Canada for its 130MW Rattlesnake Ridge wind power project',
        'wind',
        '1.6',
        '12');

INSERT INTO `bid`(contract_id, contract_name, contract_version,
                  contract_company_name,contract_company_owner_account, contract_details, energy_type, energy_price, contract_last_time)
VALUES ('0004', 'Tidal Wave Weighs', '1.0', 'Santa Claus Power','Admin@WH-zhijianju.supervisor.com',
        'A operating on the same principle as wind turbines, the power in sea turbines comes from tidal currents
which turn blades similar to ships propellers, but, unlike wind, the tides are predictable and the power input is constant',
        'tidal',
        '1.6',
        '6');

INSERT INTO `bid`(contract_id, contract_name, contract_version,
                  contract_company_name,contract_company_owner_account, contract_details, energy_type, energy_price, contract_last_time)
VALUES ('0005', 'Sunny Everyday', '1.0', 'Trustworth Energy Company','Admin@HUST.builder.com',
        'Solar power is becoming increasingly popular as both individuals and businesses realize the potential of energy independence through power provided by the sun',
        'solar',
        '1.5',
        '6');