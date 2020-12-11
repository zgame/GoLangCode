--Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
module('Proto_User_pb')

BASEUSER = protobuf.Descriptor();
BASEUSER_USER_ID_FIELD = protobuf.FieldDescriptor();
BASEUSER_OPEN_ID_FIELD = protobuf.FieldDescriptor();
BASEUSER_NICK_NAME_FIELD = protobuf.FieldDescriptor();
BASEUSER_LEVEL_FIELD = protobuf.FieldDescriptor();
BASEUSER_EXP_FIELD = protobuf.FieldDescriptor();
BASEUSER_ROOM_ID_FIELD = protobuf.FieldDescriptor();
BASEUSER_CHAIR_ID_FIELD = protobuf.FieldDescriptor();

BASEUSER_USER_ID_FIELD.name = "user_id"
BASEUSER_USER_ID_FIELD.full_name = ".Proto.BaseUser.user_id"
BASEUSER_USER_ID_FIELD.number = 1
BASEUSER_USER_ID_FIELD.index = 0
BASEUSER_USER_ID_FIELD.label = 1
BASEUSER_USER_ID_FIELD.has_default_value = false
BASEUSER_USER_ID_FIELD.default_value = 0
BASEUSER_USER_ID_FIELD.type = 13
BASEUSER_USER_ID_FIELD.cpp_type = 3

BASEUSER_OPEN_ID_FIELD.name = "open_id"
BASEUSER_OPEN_ID_FIELD.full_name = ".Proto.BaseUser.open_id"
BASEUSER_OPEN_ID_FIELD.number = 2
BASEUSER_OPEN_ID_FIELD.index = 1
BASEUSER_OPEN_ID_FIELD.label = 1
BASEUSER_OPEN_ID_FIELD.has_default_value = false
BASEUSER_OPEN_ID_FIELD.default_value = 0
BASEUSER_OPEN_ID_FIELD.type = 13
BASEUSER_OPEN_ID_FIELD.cpp_type = 3

BASEUSER_NICK_NAME_FIELD.name = "nick_name"
BASEUSER_NICK_NAME_FIELD.full_name = ".Proto.BaseUser.nick_name"
BASEUSER_NICK_NAME_FIELD.number = 3
BASEUSER_NICK_NAME_FIELD.index = 2
BASEUSER_NICK_NAME_FIELD.label = 1
BASEUSER_NICK_NAME_FIELD.has_default_value = false
BASEUSER_NICK_NAME_FIELD.default_value = ""
BASEUSER_NICK_NAME_FIELD.type = 9
BASEUSER_NICK_NAME_FIELD.cpp_type = 9

BASEUSER_LEVEL_FIELD.name = "level"
BASEUSER_LEVEL_FIELD.full_name = ".Proto.BaseUser.level"
BASEUSER_LEVEL_FIELD.number = 4
BASEUSER_LEVEL_FIELD.index = 3
BASEUSER_LEVEL_FIELD.label = 1
BASEUSER_LEVEL_FIELD.has_default_value = false
BASEUSER_LEVEL_FIELD.default_value = 0
BASEUSER_LEVEL_FIELD.type = 13
BASEUSER_LEVEL_FIELD.cpp_type = 3

BASEUSER_EXP_FIELD.name = "exp"
BASEUSER_EXP_FIELD.full_name = ".Proto.BaseUser.exp"
BASEUSER_EXP_FIELD.number = 5
BASEUSER_EXP_FIELD.index = 4
BASEUSER_EXP_FIELD.label = 1
BASEUSER_EXP_FIELD.has_default_value = false
BASEUSER_EXP_FIELD.default_value = 0
BASEUSER_EXP_FIELD.type = 13
BASEUSER_EXP_FIELD.cpp_type = 3

BASEUSER_ROOM_ID_FIELD.name = "room_id"
BASEUSER_ROOM_ID_FIELD.full_name = ".Proto.BaseUser.room_id"
BASEUSER_ROOM_ID_FIELD.number = 6
BASEUSER_ROOM_ID_FIELD.index = 5
BASEUSER_ROOM_ID_FIELD.label = 1
BASEUSER_ROOM_ID_FIELD.has_default_value = false
BASEUSER_ROOM_ID_FIELD.default_value = 0
BASEUSER_ROOM_ID_FIELD.type = 13
BASEUSER_ROOM_ID_FIELD.cpp_type = 3

BASEUSER_CHAIR_ID_FIELD.name = "chair_id"
BASEUSER_CHAIR_ID_FIELD.full_name = ".Proto.BaseUser.chair_id"
BASEUSER_CHAIR_ID_FIELD.number = 7
BASEUSER_CHAIR_ID_FIELD.index = 6
BASEUSER_CHAIR_ID_FIELD.label = 1
BASEUSER_CHAIR_ID_FIELD.has_default_value = false
BASEUSER_CHAIR_ID_FIELD.default_value = 0
BASEUSER_CHAIR_ID_FIELD.type = 13
BASEUSER_CHAIR_ID_FIELD.cpp_type = 3

BASEUSER.name = "BaseUser"
BASEUSER.full_name = ".Proto.BaseUser"
BASEUSER.nested_types = {}
BASEUSER.enum_types = {}
BASEUSER.fields = {BASEUSER_USER_ID_FIELD, BASEUSER_OPEN_ID_FIELD, BASEUSER_NICK_NAME_FIELD, BASEUSER_LEVEL_FIELD, BASEUSER_EXP_FIELD, BASEUSER_ROOM_ID_FIELD, BASEUSER_CHAIR_ID_FIELD}
BASEUSER.is_extendable = false
BASEUSER.extensions = {}

BaseUser = protobuf.Message(BASEUSER)

