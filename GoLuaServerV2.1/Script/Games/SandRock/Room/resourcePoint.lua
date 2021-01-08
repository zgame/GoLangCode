


function SandRockRoom:ResourcePointUpdate()
    -- 判断生命周期， 到期的给删除掉
    for i, v in ipairs(self.resourcePoint) do

    end
    -- 开始刷新新东西
    local areaList = CSV_resourceGenerate.GetAllKeys()
    for _, areaName in ipairs(areaList) do
        if areaName == "ResourceArea_Herb_1" then
            if self.resourcePoint[areaName] == nil then
                self.resourcePoint[areaName] = {}
            end

            local count = CSV_resourceGenerate.GetValue(areaName, 'Count')
            local list = ZString.Split(count, ',')
            local num = ZRandom.GetRandom(tonumber(list[1]), tonumber(list[2]))

            --print("随机获取本次更新资源数量num ："..num)
            local number_now = #self.resourcePoint[areaName]        -- 已经包含多少个点
            local number_max = CSV_resourceGenerate.GetValue(areaName, 'Points')
            --print("number_now"..number_now)

            if num > number_now then
                for i = 1, num - number_now do
                    --print('生成一个point')
                    local pointList = {}            --没有占用的列表
                    for i=1,number_max do
                        table.insert(pointList,i)           -- 全数组
                    end
                    for s,k in ipairs(self.resourcePoint[areaName]) do
                        table.remove(pointList,k.areaPoint)         -- 把已经占用的删掉
                    end
                    local ran = ZRandom.GetRandom(1,#pointList)

                    local type = tonumber(CSV_resourceGenerate.GetValue(areaName,"Resource"))    -- 以后改为多个权重

                    --print("保存到房间的资源列表里面")
                    local element ={}
                    element.areaPoint = pointList[ran]
                    element.resourceType = type
                    table.insert(self.resourcePoint[areaName],element)

                    --printTable(self.resourcePoint[areaName])
                end
            end
        end
    end
end