


function SandRockRoom:ResourcePointUpdate()
    -- 判断生命周期， 到期的给删除掉
    for i, v in ipairs(self.resourcePoint) do

    end
    -- 开始刷新新东西
    local areaList = CSV_resourceGenerate.GetAllKeys()
    for _, areaName in ipairs(areaList) do
        local count = CSV_resourceGenerate.GetValue(areaName, 'Count')
        local list = ZString.Split(count, ',')
        local num = ZRandom.GetRandom(list[1], list[2])

        local number_now = #self.resourcePoint[areaName]        -- 已经包含多少个点
        local number_max = CSV_resourceGenerate.GetValue(areaName, 'Points')


        if num > number_now then
            for i = 1, i <= num - number_now do
                -- 生成一个point
                local element ={}
                local point = 1     -- 位置
                local pointList = {}
                for i=1,i<=number_now do
                    table.insert(pointList,i)           -- 全数组
                end
                for s,k in ipairs(self.resourcePoint[areaName]) do
                    table.remove(pointList,k.areaPoint)         -- 把已经占用的删掉
                end


                local type = tonumber(CSV_resourceGenerate.GetValue(areaName,"Resource"))    -- 以后改为多个权重

                element.areaPoint = point
                element.resourceType = type
                table.insert(self.resourcePoint[areaName],element)

            end
        end

    end
end