-------------------------------------------------------
--- 加密，解密， 验证
-------------------------------------------------------

local crypto = require("crypto")                -- 位运算

-- MD5 进行验证
function MD5Get(str)
    return crypto.md5(str)
end


-- Base64加密
function BASE64EncodeStr(str)
    return crypto.base64_encode(str)
end


-- Base64解密
function BASE64DecodeStr(str)
    return crypto.base64_decode(str)
end