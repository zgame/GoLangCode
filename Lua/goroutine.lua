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

function deal(ok,v)
    if not ok then
        print("channel closed")
        --                    exit = true
    else
        print("received:", v)

        quit:send("返回给你")
    end
end


function receivezz()
--    local exit = false
----    while not exit do
        channel.select(
            {"|<-", ch, deal},
--            {"|<-", quit, function(ok, v)
--                print("quit"..v)
----                exit = true
--            end},
            {"default"}
        )
----    end


--    local idx, recv, ok = channel.select(
--        {"|<-", ch},
--        {"default"}
--    )
--    if not ok then
----        print("closed")
--    elseif idx == 1 then -- received from ch1
--        print(recv)
--        quit:send(recv.."收到")
----    elseif idx == 2 then -- received from ch2
----        print(recv)
-- end


--    local ok, v = ch:receive()
--    print(v)
--    quit:send(v.."收到")

end


function sendzz(myName)
--    ch:send(1)
    ch:send(myName)
--    ch:send(true)
--ch:send(1.2)
--local tt = {}
--tt[1] = 1
--tt["22"] = "sdf"
--tt[3] = {}
--tt[3]["3"] = 3
--ch:send(tt)
--    quit:send(false)
--    quit:send(true)
    local ok, v = quit:receive()
    print("发送后等待收到消息:",v)
end
