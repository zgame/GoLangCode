---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by Administrator.
--- DateTime: 2018/10/12 15:56
---




local user_pb = require('user_pb')

print("protocol")

local person= user_pb.Friend()
person.name = "zsw"
person.adress = "zsw adress"

local data = person:SerializeToString()

local msg = user_pb.Friend()
msg:ParseFromString(data)
print(msg)