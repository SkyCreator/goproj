-- ------------------------------------------------------------
-- 此文件由工具自动生成，请勿直接修改。
-- ------------------------------------------------------------

DRUserCommodity = {
        [91] = {Id = 91,Code = "SUPREME",Name = "王国商城至尊勋章",Price = 49999,Icon = "aaa.png", },
        [327] = {Id = 327,Code = "19VIPPLUS",Name = "赞助助手",Price = 30, },
        [353] = {Id = 353,Code = "18VIPAM",Price = 15, },
        [361] = {Id = 361,Code = "19ITEMR00G",Name = "",Price = 1, },
        [391] = {Id = 391,Code = "18TOOLSF",Name = "永久王牌指挥", },
        [392] = {Id = 392,Name = "永久干扰精通",Price = 168,Icon = "bbb.png", },
        [405] = {Id = 405,Code = "18TOOLSI",Name = "永久运筹帷幄", },
        [406] = {Id = 406,Code = "18TOOLSG",Name = "永久排兵布阵",Price = 168, },
        [407] = {Id = 407,Code = "18TOOLSE",Name = "永久快速建造", },
        [408] = {Id = 408,Code = "18TOOLSD",Name = "永久末日预言",Price = 168, },
        [409] = {Id = 409,Code = "18TOOLSC",Name = "永久空间传送",Price = 288, },
        [410] = {Id = 410,Code = "18TOOLSA",Name = "永久伐木预估", },
        [416] = {Id = 416,Code = "18VIPC",Name = "永久超级赞助",Price = 2000, },
        [417] = {Id = 417,Code = "18VIPB",Name = "永久赞助",Price = 498, },
        [418] = {Id = 418,Code = "18VIPA",Name = "永久会员",Price = 188,Icon = "ccc.png", },
}

local _default = {Code = "18TOOLSB",Name = "会员月卡",Type = 1,Price = 88,Icon = "",
}

local _base = {
    __index = function(tbl,key)
        return _default[key]
    end, 
    __newindex = function(tbl,key,value)
        print( "Attempt to modify read-only table DRUserCommodity: key:" .. key .. ", value:" .. value)
        rawset(tbl,key,value)
    end,
    __metatable = false,
}

local mt = setmetatable
for k, v in pairs( DRUserCommodity ) do
    mt( v, _base )
end