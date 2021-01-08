----------------------------------------------------------------
--- 资源刷新， 每天刷新一次， 到生命周期就删掉  ， 不足上限数量就补充，但是不一定非要补充，是随机的
----------------------------------------------------------------


-- 获取没有被占用的资源点列表
local function _getEmpty(areaName,resourcePoint)
    local number_max = CSV_resourceGenerate.GetValue(areaName, 'Points')
    local pointList = {}            --没有占用的列表
    for j=1,number_max do
        table.insert(pointList,j)           -- 生成全数组
    end
    for pointIndex, point in pairs(resourcePoint[areaName]) do
        table.remove(pointList, pointIndex)         -- 把已经占用的删掉
    end
    -- 剩下的就是没有被占用的了
    local ran = ZRandom.GetRandom(1,#pointList)
    return pointList[ran]
end


-- 资源点刷新
function SandRockRoom:ResourcePointUpdate()
    -- 判断生命周期， 到期的给删除掉
    for areaName, pointList in pairs(self.resourcePoint) do
        --local temp = {}     -- 不包含过期的
        for index, point in pairs(pointList)do
            if point.live <= 1 then
                pointList[index] = nil              -- 删掉生命周期已经到了的点
                --table.remove(pointList,index)       -- 删掉生命周期已经到了的点
            else
                point.live = point.live - 1
                --table.insert(temp,point)
            end
        end
        --self.resourcePoint[areaName] = temp
    end
    -- 开始刷新新东西
    local areaList = CSV_resourceGenerate.GetAllKeys()
    for _, areaName in ipairs(areaList) do
        if areaName == "ResourceArea_Herb_1" then
            if self.resourcePoint[areaName] == nil then
                self.resourcePoint[areaName] = {}           -- 初始化生成点列表
            end

            local count = CSV_resourceGenerate.GetValue(areaName, 'Count')
            local list = ZString.Split(count, ',')
            local num = ZRandom.GetRandom(tonumber(list[1]), tonumber(list[2]))
            --print("随机获取本次更新资源数量num ："..num)
            local number_now = #self.resourcePoint[areaName]        -- 已经包含多少个点
            --print("number_now"..number_now)
            if num > number_now then
                for i = 1, num - number_now do
                    --print('生成一个point, 下面是point的结构')
                    local resourceType = CSV_resourceGenerate.GetValue(areaName,"Resource")
                    local element ={}
                    local areaPoint = _getEmpty(areaName, self.resourcePoint)
                    element.resourceType = tonumber(resourceType)    -- 以后改为多个权重
                    element.live = CSV_resourceType.GetValue(resourceType,"LifeCycle")
                    --print("保存到房间的资源列表里面")
                    self.resourcePoint[areaName][areaPoint] = element
                    --table.insert(self.resourcePoint[areaName],element)
                    --printTable(self.resourcePoint[areaName])
                end
            end
        end
    end

    --printTable(self.resourcePoint)
end


-----------------------------------采集------------------------------------------
function SandRockRoom:GetResource(userId, areaName, pointIndex, resourceType)
    if self.resourcePoint[areaName] == nil then
        ZLog.Logger("GetResource  areaName error ".. areaName)
        return
    end

    local point = self.resourcePoint[areaName][pointIndex]
    if point == nil then
        ZLog.Logger("GetResource pointIndex error".. pointIndex)
        return
    end
    if point.resourceType ~= resourceType then
        ZLog.Logger("GetResource resourceType error".. resourceType)
    end
    local player = GameServer.GetPlayerByUID(userId)
    -- 采集
    local spCost = CSV_resourceType.GetValue(resourceType,"SpCost")
    local exp = CSV_resourceType.GetValue(resourceType,"Exp")

    -- 获得物品


    -- 销毁采集点
    self.resourcePoint[areaName][pointIndex] = nil


end