local hello2 = require("hello2")
require("class_test")


print("-----------Lua-----------")


function Zsw2(inss,ss)
    if (ss == nil) then
        ss = 0
    end

    return inss*100 + ss,ss
end



num,sss = double(20,2,"zsw")
print(num)
print(sss)
print("lua call go:  ",num,sss)

print('---------------------')
print("go call lua:  ",Zsw2(9))
print("call require file")
hello2.zsw()
print(hello2.ss)



print('---------------------')
counter = Counter:new(100)
for i=1,10 do
    print(counter:incr(i))
end
print(counter:get())