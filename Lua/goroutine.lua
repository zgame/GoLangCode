--
-- Created by IntelliJ IDEA.
-- User: Administrator
-- Date: 2018/11/8
-- Time: 11:03
-- To change this template use File | Settings | File Templates.
--

require("dumpTable")

function test()
    print("multi")
end

Quit = nil
Num = 0
function deal(ok,v)
    if ok then
        if v == "1" and Num > 1 then
            print(1,"这块不再返回了")
            quit:send("quit")
        else
            Num = Num + 1
            print("received:", v)
            quit:send("返回给你"..v)
        end
    end
end


function receivezz()
        channel.select(
            {"|<-", ch, deal},
            {"default"}
        )
end


function sendzz(myName)
--    print("sendzz, Quit",Quit)
    if Quit ~= nil then

        return
    end

    ch:send(myName)
    local ok, v = quit:receive()
    print("发送后等待收到消息:",v)
    if v == "quit" then
--        ch:close()
--        quit:close()
        Quit = true
        zClose()        -- 关闭Lstate
    end

end
