-------------------------------------------------------
--- 压缩，解压缩
-------------------------------------------------------

Zip = require("zip")                -- 位运算

-- 压缩
function Zip.Compression(str)
    return Zip.encode(str)
end

-- 解压缩
function Zip.UnCompression(str)
    return Zip.decode(str)
end
