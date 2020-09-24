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

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `admin`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `account` varchar(50) NOT NULL,
  `password` char(32) NOT NULL,
  `email_address` char(32) NOT NULL,
  `phone_number` char(32) NOT NULL,
  `descript` varchar(255) DEFAULT '',
  `headico` varchar(200) DEFAULT '',
  `type` varchar(20) NOT NULL DEFAULT 'u',
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `username` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `user`(created_at,updated_at,deleted_at,account,password,email_address,phone_number,descript,headico,type) VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Admin@HUST.builder.com', '21232f297a57a5a743894a0e4a801fc3','129843719@qq.com','18821218478', 'admin of HUST.builder.com', '/public/adminlit/dist/img/user2-160x160.jpg', 'admin');
INSERT INTO `user`(created_at,updated_at,deleted_at,account,password,email_address,phone_number,descript,headico,type) VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'asdasda@HUST.builder.com', '21232f297a57a5a743894a0e4a801fc3','1221412@qq.com','18124121411', 'user of HUST.builder.com', '/public/adminlit/dist/img/user2-160x160.jpg', 'user');
INSERT INTO `user`(created_at,updated_at,deleted_at,account,password,email_address,phone_number,descript,headico,type) VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Admin@WH-zhijianju.supervisor.com', '21232f297a57a5a743894a0e4a801fc3','3251354123@qq.com','13252523232', 'admin of WH-zhijianju.supervisor.com', '/public/adminlit/dist/img/user2-160x160.jpg','admin');
INSERT INTO `user`(created_at,updated_at,deleted_at,account,password,email_address,phone_number,descript,headico,type) VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Admin@zhongjian-1-ju.constructor.com', '21232f297a57a5a743894a0e4a801fc3','1213112312@qq.com','18124128478', 'admin of zhongjian-1-ju.constructor.com', '/public/adminlit/dist/img/user2-160x160.jpg','admin');
INSERT INTO `user`(created_at,updated_at,deleted_at,account,password,email_address,phone_number,descript,headico,type) VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Admin@zhongjian-2-ju.constructor.com', '21232f297a57a5a743894a0e4a801fc3','241123412@qq.com','18124112418', 'admin of zhongjian-2-ju.constructor.com', '/public/adminlit/dist/img/user2-160x160.jpg','admin');


-- ----------------------------
-- Table structure for `Company`
-- ----------------------------


DROP TABLE IF EXISTS `company`;
CREATE TABLE `company` (
                        `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                        `created_at` timestamp NULL DEFAULT NULL,
                        `updated_at` timestamp NULL DEFAULT NULL,
                        `deleted_at` timestamp NULL DEFAULT NULL,
                        `company_name` varchar(50) NOT NULL,
                        `company_owner_account` varchar(50) NOT NULL,
                        `company_build_year` varchar(50) NOT NULL,
                        `email_address` char(50) NOT NULL,
                        `phone_number` char(50) NOT NULL,
                        `descript` varchar(255) DEFAULT '',
                        `headico` varchar(200) DEFAULT '',
                        `company_invest`    varchar(50) DEFAULT '',
                        `company_size`     varchar(50) DEFAULT '',
                        `company_location`  varchar(50) DEFAULT '',
                        `company_website`   varchar(50) DEFAULT '',
                        PRIMARY KEY (`id`),
                        KEY `idx_user_deleted_at` (`deleted_at`),
                        KEY `companyname` (`company_name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


INSERT INTO `company`(created_at,updated_at,deleted_at,company_name,company_owner_account,company_build_year,email_address,phone_number,descript,headico,company_invest,company_size,company_location,company_website) VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'SanFrancisco Power','Admin@HUST.builder.com', '1978','SanPowerOfficial@gmail.com','+86-219031213', 'Best power supplier in SanFancisco', '/public/adminlit/dist/img/user2-160x160.jpg', '50 million','2000-5000 employees','SanFrancisco Long Street 182','www.SanFanciscoPower.com');
INSERT INTO `company`(created_at,updated_at,deleted_at,company_name,company_owner_account,company_build_year,email_address,phone_number,descript,headico,company_invest,company_size,company_location,company_website) VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Trustworth Energy Company','Admin@HUST.builder.com', '2007','TrustworthEnergyOffcial@gmail.com','+86-32156243', 'Energy and power are the future', '/public/adminlit/dist/img/user2-160x160.jpg', '12 million','100-1000 employees','Seatle Dann Street 78','www.TrustworthEngergy.com');
INSERT INTO `company`(created_at,updated_at,deleted_at,company_name,company_owner_account,company_build_year,email_address,phone_number,descript,headico,company_invest,company_size,company_location,company_website) VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Santa Claus Power','Admin@WH-zhijianju.supervisor.com', '2015','SantaClausPower@gmail.com','+70-2193891213', 'Creative Energy means everything', '/public/adminlit/dist/img/user2-160x160.jpg', '50 million','100-1000 employees','Los Angeles Beach Street 89','www.SantaClausPower.com');
INSERT INTO `company`(created_at,updated_at,deleted_at,company_name,company_owner_account,company_build_year,email_address,phone_number,descript,headico,company_invest,company_size,company_location,company_website) VALUES ('2018-10-22 14:03:48', '2018-11-01 15:01:13', null, 'Believing Energy','Admin@WH-zhijianju.supervisor.com', '1984','Be-Energy@gmail.com','+86-3546336', 'Believe it or not, we can offer you the best and the cleanest power', '/public/adminlit/dist/img/user2-160x160.jpg', '100 million','5000+ employees','New York City King road 17','www.BelievingEnergy.com');

-- ----------------------------
-- Records of company







-- ----------------------------
-- Table structure for `addTable`
-- ----------------------------

DROP TABLE IF EXISTS `table_for_web`;
CREATE TABLE `table_for_web` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tid` varchar(100) NOT NULL DEFAULT '' COMMENT '表单ID',
  `tname` varchar(100) NOT NULL DEFAULT '' COMMENT '表单名称',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tid` (`tid`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of table_for_web
-- ----------------------------