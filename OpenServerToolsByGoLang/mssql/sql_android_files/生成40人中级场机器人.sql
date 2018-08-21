---------------------------------------------------------------------------------------------------------------------
-- 生成机器人,@Num:指定生成数量，一半名字使用游客***，一半的名字读取AndroidUserName，

DECLARE @Accounts NVARCHAR(32)
DECLARE @Num INT
DECLARE @FaceId INT
DECLARE @ServerID INT
DECLARE @VipLevel INT
DECLARE @NameCount INT
DECLARE @AndroidCount INT
DECLARE @NickName NVARCHAR(64)

SET @Num = %s
SET @FaceId = 1
SET @VipLevel = 0

--以下为动态修改内容
--服务器ID
SET @ServerID = %s

--读取机器人名字
SELECT id=identity(int,1,1), Name INTO #regtemp11 FROM AndroidUserName(NOLOCK)
--取名字数量
SELECT @NameCount=COUNT(*) FROM #regtemp11
--取机器人数量
SELECT @AndroidCount=COUNT(*) FROM AccountsInfo(NOLOCK) WHERE IsAndroid=1


-- WHILE @Num < 100
BEGIN
	-- 生成账号
	SET @Accounts = N'JQL_CHECK_X' + cast((@ServerID*100+@Num) AS NVARCHAR(32))
	SET @AndroidCount=@AndroidCount+1
-- 	SET @Num = @Num + 1
	--VIP1-10个, VIP2-10个,VIP3-30个,VIP4-30个,VIP5-20个
	IF @Num <= 10
		SET @VipLevel = 1
	ELSE IF @Num <= 10+10
		SET @VipLevel = 2
	ELSE IF @Num <= 20+30
		SET @VipLevel = 3
	ELSE IF @Num <= 50+30
		SET @VipLevel = 4
	ELSE
		SET @VipLevel = 5

	-- 设置头像
	SET @FaceId = cast( floor(RAND()*6) AS INT) + 1
	SELECT @NickName=Name FROM #regtemp11 WHERE id=(@AndroidCount-1)%s@NameCount+1

	-- 注册用户
	INSERT dbo.AccountsInfo (Accounts,NickName,RegAccounts,LogonPass,InsurePass,IsAndroid,RegisterIP,LastLogonIP,VipLev,FaceID)
	VALUES (@Accounts,@NickName,@Accounts,'96E79218965EB72C92A549DD5A330112','96E79218965EB72C92A549DD5A330112',1,'127.0.0.1','127.0.0.1',@VipLevel,@FaceId)

	-- GameID
	DECLARE @UserID INT
	UPDATE dbo.AccountsInfo SET @UserID=UserID, GameID=dbo.GSP_FUN_GetGameID(@UserID, RAND(), RAND()) WHERE Accounts=@Accounts

	-- 将机器人插入记录表
	INSERT INTO dbo.AndroidManager(UserID, ServerID, ServiceGender, MinPlayDraw, MaxPlayDraw, MinTakeScore, MaxTakeScore, MinReposeTime, MaxReposeTime, ServiceTime) 
	VALUES(@UserID, @ServerID, 7, 1, 100, 800000, 500000000,30,360,16777215 )

END

 
DROP TABLE #regtemp11

---------------------------------------------------------------------------------------------------------------------


