-------------------------------------------------------
--- 压缩，解压缩
-------------------------------------------------------

local zip = require("zip")                -- 位运算

-- 压缩
function ZipCompression(str)
    return zip.encode(str)
end

-- 解压缩
function ZipUnCompression(str)
    return zip.decode(str)
end
