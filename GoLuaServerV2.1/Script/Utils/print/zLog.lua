
ZLog = {}

-- 输入日志带打印控制台
function ZLog.Logger(str)
    luaCallGoPrintLogger(str,true)
end

-- 输入日志不带打印控制台
function ZLog.Record(str)
    luaCallGoPrintLogger(str,false)
end



-- 打印输出table
require("Script/Utils/print/dumpTable")           -- printTable(tableName)
local MySerpent = require("Script/Utils/print/serpent")      -- 使用方法 ： MySerpent.block(tableName)        也是打印table


function ZLog.printTable1(table)
    printTable(table)
end

function ZLog.printTable2(table)
    MySerpent.block(table)
end


function ZLog.printTime(tag)
    print(tag .. " :" .. ZTime.GetOsTimeMillisecond())
end