-------------------------------------------------------
--- 加密，解密， 验证
-------------------------------------------------------


-- MD5 进行验证
function MD5Get(str)
    return luaCallGoGetMD5(str)
end


-- Base64加密
function BASE64EncodeStr(str)
    return luaCallGoBASE64EncodeStr(str)
end


-- Base64解密
function BASE64DecodeStr(str)
    return luaCallGoBASE64DecodeStr(str)
end