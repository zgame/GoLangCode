-------------------------------------------------------
--- 加密，解密， 验证
-------------------------------------------------------

Crypto = require("crypto")                -- 位运算

-- MD5 进行验证
function Crypto.MD5(str)
    return Crypto.md5(str)
end


-- Base64加密
function Crypto.BASE64EncodeStr(str)
    return Crypto.base64_encode(str)
end


-- Base64解密
function Crypto.BASE64DecodeStr(str)
    return Crypto.base64_decode(str)
end