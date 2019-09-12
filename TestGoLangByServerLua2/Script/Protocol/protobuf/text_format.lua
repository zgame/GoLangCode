--
--------------------------------------------------------------------------------
--  FILE:  text_format.lua
--  DESCRIPTION:  protoc-gen-lua
--      Google's Protocol Buffers project, ported to lua.
--      https://code.google.com/p/protoc-gen-lua/
--
--      Copyright (c) 2010 , 林卓毅 (Zhuoyi Lin) netsnail@gmail.com
--      All rights reserved.
--
--      Use, modification and distribution are subject to the "New BSD License"
--      as listed at <url: http://www.opensource.org/licenses/bsd-license.php >.
--  COMPANY:  NetEase
--  CREATED:  2010年08月05日 15时14分13秒 CST
--------------------------------------------------------------------------------
--
local string = string
local math = math
local print = print
local getmetatable = getmetatable
local table = table
local ipairs = ipairs
local tostring = tostring

local descriptor = require "descriptor"

module "text_format"

function format(buffer)
    local len = string.len( buffer )	
    for i = 1, len, 16 do		
        local text = ""	
        for j = i, math.min( i + 16 - 1, len ) do	
            text = string.format( "%s  %02x", text, string.byte( buffer, j ) )			
        end			
        print( text )	
    end
end

local FieldDescriptor = descriptor.FieldDescriptor

msg_format_indent = function(write, msg, indent)       -- 内部调用， 用来打印msg
    for field, value in msg:ListFields() do
        local print_field = function(field_value)
            local name = field.name
            write(string.rep(" ", indent))
            if field.type == FieldDescriptor.TYPE_MESSAGE then      -- 输出message
                local extensions = getmetatable(msg)._extensions_by_name
                if extensions[field.full_name] then
                    write("[" .. name .. "] {\n")
                else
                    write(name .. " {\n")
                end
                msg_format_indent(write, field_value, indent + 4)
                write(string.rep(" ", indent))
                write("}\n")
            else
                write(string.format("%s: %s\n", name, tostring(field_value)))
            end
        end
        if field.label == FieldDescriptor.LABEL_REPEATED then
            for _, k in ipairs(value) do
                print_field(k)
            end
        else
            print_field(value)
        end
    end
end

function msg_format(msg)            -- protobuf.lua调用， 用来__tostring
    local out = {}
    local write = function(value)
        out[#out + 1] = value       --在表的长度#out后面写入value
    end
    msg_format_indent(write, msg, 0)
    return table.concat(out)
end

