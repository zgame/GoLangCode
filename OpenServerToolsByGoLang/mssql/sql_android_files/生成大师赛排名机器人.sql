-- 生成机器人,@Num:指定生成数量，一半名字使用游客***，一半的名字读取AndroidUserName，

DECLARE @Accounts NVARCHAR(32)
DECLARE @Num INT
DECLARE @FaceId INT
DECLARE @ServerID INT
DECLARE @VipLevel INT
DECLARE @FaceUrl NVARCHAR(256)
DECLARE @FaceHttp NVARCHAR(256)
DECLARE @NameCount INT
DECLARE @AndroidCount INT
DECLARE @NickName NVARCHAR(64)

SET @Num = %s
SET @FaceId = 1
SET @VipLevel = 0
-- 服务器ID 根据具体服务器ID进行修改
SET @ServerID = %s
-- 这个地址根据版本修改
SET @FaceHttp=N'http://dl.yx.soonyo.com/tb/robot/'

--修改机器人名字
-- 查询ID
SELECT id=identity(int,1,1), A.Name INTO #regtemp11 FROM AndroidUserName(NOLOCK) A
SELECT @NameCount=COUNT(*) FROM #regtemp11
SELECT @AndroidCount=COUNT(*) FROM AccountsInfo(NOLOCK)

-- WHILE @Num < 1000
BEGIN
	-- 生成账号
	SET @Accounts = N'JQL_CHECK_DP' + cast((@ServerID*1000+@Num) AS NVARCHAR(32))
	SET @AndroidCount=@AndroidCount+1
-- 	SET @Num = @Num + 1
	IF @Num <= 10
	BEGIN
		SET @VipLevel = Rand()*2+6
	END
	ELSE IF @Num <= 100
	BEGIN
		SET @VipLevel = Rand()*3+4
	END
	ELSE IF @Num <= 300
	BEGIN
		SET @VipLevel = Rand()*3+3
	END
	ELSE IF @Num <= 500
	BEGIN
		SET @VipLevel = Rand()*2+3
	END
	ELSE IF @Num <= 1000
	BEGIN
		SET @VipLevel = Rand()*3+1
	END
	-- 设置头像
	SET @FaceId = cast( floor(RAND()*6) AS INT) + 1
	SET @FaceUrl = @FaceHttp+CAST(CAST(RAND()*200 AS INT)+1 AS NVARCHAR)+N'.jpg'
	SELECT @NickName=Name FROM #regtemp11 WHERE id=(@AndroidCount-1)%s@NameCount+1
	-- 注册用户
	INSERT dbo.AccountsInfo (Accounts,NickName,RegAccounts,LogonPass,InsurePass,IsAndroid,RegisterIP,LastLogonIP,VipLev,FaceID,FaceUrl)
	VALUES (@Accounts,@NickName,@Accounts,'96E79218965EB72C92A549DD5A330112','96E79218965EB72C92A549DD5A330112',1,'127.0.0.1','127.0.0.1',@VipLevel,@FaceId,@FaceUrl)


	-- GameID
	DECLARE @UserID INT
	UPDATE dbo.AccountsInfo SET @UserID=UserID, GameID=dbo.GSP_FUN_GetGameID(@UserID, RAND(), RAND()) WHERE Accounts=@Accounts

	-- 将机器人插入记录表
	--INSERT INTO dbo.AndroidManager(UserID, ServerID, ServiceGender, MinPlayDraw, MaxPlayDraw, MinTakeScore, MaxTakeScore, MinReposeTime, MaxReposeTime, ServiceTime) 
	--VALUES(@UserID, @ServerID, 7, 1, 100, 800000, 500000000,30,360,16777215 )
	INSERT dbo.UserArenaInfo(UserID) VALUES(@UserID)
END

DROP TABLE #regtemp11

---------------------------------------------------------------------------------------------------------------------


