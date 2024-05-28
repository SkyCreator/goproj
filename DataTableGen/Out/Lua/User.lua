-- ------------------------------------------------------------
-- 此文件由工具自动生成，请勿直接修改。
-- ------------------------------------------------------------

DRUser = {
        [1] = {Id = 1,Sex = 1, },
        [2] = {Id = 2,Name = "bbb",Height = 178.65,Sex = 0,AGE = 35, },
        [3] = {Id = 3,Name = "ccc",Height = 192,Sex = true,AGE = 41, },
        [4] = {Id = 4,Name = "ddd",Height = 162,Sex = false,AGE = 82, },
}

local _default = {Name = "aaa",Height = 182,Sex = true,AGE = 18,
}

local _base = {
    __index = function(tbl,key)
        return _default[key]
    end, 
    __newindex = function(tbl,key,value)
        print( "Attempt to modify read-only table DRUser: key:" .. key .. ", value:" .. value)
        rawset(tbl,key,value)
    end,
    __metatable = false,
}

local mt = setmetatable
for k, v in pairs( DRUser ) do
    mt( v, _base )
end