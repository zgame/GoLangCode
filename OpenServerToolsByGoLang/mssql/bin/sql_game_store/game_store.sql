SET NOCOUNT ON
	DECLARE @ServerName NVARCHAR(64)
	DECLARE @wKindID INT
	DECLARE @RoomCellScore INT
	DECLARE @wServerID INT
	DECLARE @wAndroidCount INT
	DECLARE @stockCode INT

	SET @wServerID = '%s'
	SET @wAndroidCount = '%s'
	SET @stockCode = (CASE WHEN @wKindID/100=4 THEN 5 WHEN @wKindID=3 AND @RoomCellScore=1 THEN 1 WHEN @wKindID=3 AND @RoomCellScore=10 THEN 2 WHEN @wKindID=3 AND @RoomCellScore=100 THEN 3 WHEN @wKindID=3 AND @RoomCellScore=1000 THEN 4 ELSE 1 END)

	BEGIN
	SELECT @ServerName=ISNULL(ServerName,N''),@wKindID = GameID FROM GameRoomInfo WHERE ServerID=@wServerID
	SELECT @RoomCellScore=CellScore FROM GameRoomInfo(NOLOCK) WHERE ServerID=@wServerID;
	SET @RoomCellScore=ISNULL(@RoomCellScore,0)
	--
	INSERT INTO dbo.GameStoreInfo(ServerID, KindID, ServerName, GameStore,ScoreLineUp,ScoreLineDown,LineUpOffSet,PutMoneyMax,PerRevenue)
	SELECT @wServerID, @wKindID,@ServerName,stockValue,stockMaxValue,stockMinValue,stockMaxpeak,drawMaxScore,taxrate
	FROM dbo.RoomStockConfig(NOLOCK) WHERE stockCode= @stockCode;

	update dbo.GameStoreInfo set AndroidCount = @wAndroidCount where ServerID = @wServerID
	Select * from dbo.GameStoreInfo where ServerID = @wServerID

	END