local hello2 = require("Script/hello2")
require("Script/class_test")
require("Script/user")

--print("-----------Lua-----------")


-- go 调用lua---------------------------------------------------------------
function Zsw2(inss,ss)
    if (ss == nil) then
        ss = 0
    end
    --print("go call lua:  ",inss*100 + ss, ss)

    --print(hello2.zsw())

    return inss*100 + ss,ss
end




-- lua调用go---------------------------------------------------------------
--num,sss = double(20,2,"zsw")
--print("lua call go:  ",num,sss)





-- 其他lua文件module调用---------------------------------------------------------------
--print("call require file")
--hello2.zsw()
--print(hello2.ss)



-- 类的调用方法---------------------------------------------------------------

print('--------自增-------------')
counter = Counter:new(1)


helloNum = counter:incr(helloNum)
print(counter:get())