/*
Navicat MySQL Data Transfer

Source Server         : 207
Source Server Version : 50621
Source Host           : 192.168.0.207:3307
Source Database       : game_by

Target Server Type    : MYSQL
Target Server Version : 50621
File Encoding         : 65001

Date: 2019-06-24 15:03:46
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `game_state`
-- ----------------------------
DROP TABLE IF EXISTS `game_state`;
CREATE TABLE `game_state` (
  `zkey` varchar(40) NOT NULL DEFAULT '',
  `server_ip` varchar(20) DEFAULT NULL,
  `game_id` int(255) DEFAULT NULL,
  `table_id` int(255) DEFAULT NULL,
  `seat_array` int(255) DEFAULT NULL,
  `pool_all` bigint(255) DEFAULT NULL,
  `jackpot` bigint(255) DEFAULT NULL,
  `reward_rate` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`zkey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of game_state
-- ----------------------------
INSERT INTO `game_state` VALUES ('192.168.0.218:8123_100_1', '192.168.0.218:8123', '100', '1', '0', '0', '0', '0.00%');
INSERT INTO `game_state` VALUES ('192.168.0.218:8123_101_1', '192.168.0.218:8123', '101', '1', '0', '1950', '260', '53.19%');
INSERT INTO `game_state` VALUES ('192.168.0.218:8123_102_1', '192.168.0.218:8123', '102', '1', '0', '0', '0', '0.00%');
INSERT INTO `game_state` VALUES ('192.168.0.218:8123_103_1', '192.168.0.218:8123', '103', '1', '0', '112500', '15000', '58.82%');
INSERT INTO `game_state` VALUES ('192.168.101.109:8123_100_1', '192.168.101.109:8123', '100', '1', '1', '0', '0', '0.00%');
INSERT INTO `game_state` VALUES ('192.168.101.109:8123_101_1', '192.168.101.109:8123', '101', '1', '0', '-29400', '2880', '124.75%');
INSERT INTO `game_state` VALUES ('192.168.101.109:8123_102_1', '192.168.101.109:8123', '102', '1', '1', '0', '0', '0%');
INSERT INTO `game_state` VALUES ('192.168.101.109:8123_103_1', '192.168.101.109:8123', '103', '1', '0', '-20000', '0', '0.00%');
