SET NOCOUNT ON    
SET IDENTITY_INSERT dbo.GameRoomInfo ON







INSERT INTO [dbo].[GameRoomInfo] ([ServerID], [ServerName], [KindID], [NodeID], [SortID], [GameID], [TableCount], [ServerType], [ServerPort], [DataBaseName], [DataBaseAddr], [CellScore], [RevenueRatio], [ServiceScore], [RestrictScore], [MinTableScore], [MinEnterScore], [MaxEnterScore], [MinEnterMember], [MaxEnterMember], [MaxPlayer], [ServerRule], [DistributeRule], [MinDistributeUser], [MaxDistributeUser], [DistributeTimeSpace], [DistributeDrawCount], [DistributeStartDelay], [AttachUserRight], [ServiceMachine], [CustomRule], [Nullity], [ServerNote], [CreateDateTime], [ModifyDateTime], [MinEnterCannonLev], [MaxEnterCannonLev], [MinEnterVip], [MaxEnterVip], [ServiceAddr]) VALUES ('%s', N'%s', '0', '0', '0', '3', '150', '1', '%s', N'DataBaseBY', N'172.27.248.55', '500', '0', '0', '0', '0', '600000', '0', '0', '0', '600', '2192', '0', '0', '0', '0', '0', '0', '0', N'%s', N'00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000', '0', N'', '2016-09-23 04:17:01.360', '2016-09-23 04:17:01.360', '23', '0', '0', '0', N'');



SET IDENTITY_INSERT dbo.GameRoomInfo OFF
select @@identity