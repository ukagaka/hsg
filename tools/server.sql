

CREATE TABLE `user` (
    `uid` bigint(20) NOT NULL COMMENT '玩家唯一ID',
    `openId` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT '' COMMENT '玩家账号',
    `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
    `serverId` int(11) DEFAULT '0' COMMENT '服务器ID',
    `name` varchar(50) DEFAULT NULL COMMENT '玩家昵称',
    `level` int(11) DEFAULT '1' COMMENT '玩家等级',
    `avatar` varchar(200) DEFAULT NULL COMMENT '玩家头像链接',
    `sex` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别',
    `vip` int(11) NOT NULL DEFAULT '0' COMMENT 'VIP等级',
    `medal` int(11) NOT NULL DEFAULT '-1' COMMENT '称号',
    `jobId` int(11) NOT NULL DEFAULT '0' COMMENT '职业ID',
    `updateTime` datetime DEFAULT NULL COMMENT '最后更新时间',
    `createTime` datetime DEFAULT NULL COMMENT '创角时间',
    PRIMARY KEY (`userId`) USING BTREE,
    KEY `idx_openId` (`openId`) USING BTREE COMMENT 'openId索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='玩家信息表';