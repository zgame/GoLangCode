--
-- Created by IntelliJ IDEA.
-- User: Administrator
-- Date: 2018/11/8
-- Time: 11:03
-- To change this template use File | Settings | File Templates.
--

require("dumpTable")

function receivezz()
--    local exit = false
----    while not exit do
--        channel.select(
--            {"|<-", ch, function(ok, v)
--                if not ok then
--                    print("channel closed")
----                    exit = true
--                else
--                    print("received:", v)
--
----                    quit:send("返回给你")
--                end
--            end},
--            {"|<-", quit, function(ok, v)
--                print("quit"..v)
----                exit = true
--            end},
--            {"default", function()
----                print("default action")
--            end}
--        )
----    end


--    local idx, recv, ok = channel.select(
--        {"|<-", ch},
--        {"|<-", quit},
--        {"default", function()
----            print("default action")
--        end}
--    )
--    if not ok then
----        print("closed")
--    elseif idx == 1 then -- received from ch1
--        print(recv)
--    elseif idx == 2 then -- received from ch2
--        print(recv)
-- end


    local ok, v = ch:receive()
    print(v)
    quit:send("收到")

end


function sendzz()
--    ch:send(1)
    ch:send("string")
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
    print(v)
end
