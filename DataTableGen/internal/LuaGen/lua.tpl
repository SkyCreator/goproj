-- ------------------------------------------------------------
-- 此文件由工具自动生成，请勿直接修改。
-- ------------------------------------------------------------

DR{{.TableName}} = {
     {{- range $index, $_ := .TableRowData}}
        [{{.CellId}}] = { {{- range .CellData}} {{- if eq .CellType "STRING"}}{{.CellName}} = "{{.CellValue}}",{{- else}}{{.CellName}} = {{.CellValue}},{{- end}}{{- end}} },
    {{- end}}
}

local _default = {
    {{- range .SpecialRowData.CellData}}
    {{- if eq .CellType "STRING"}}{{.CellName}} = "{{.CellValue}}",{{- else}}{{.CellName}} = {{.CellValue}},{{- end}}
    {{- end}}
}

local _base = {
    __index = function(tbl,key)
        return _default[key]
    end, 
    __newindex = function(tbl,key,value)
        print( "Attempt to modify read-only table DR{{.TableName}}: key:" .. key .. ", value:" .. value)
        rawset(tbl,key,value)
    end,
    __metatable = false,
}

local mt = setmetatable
for k, v in pairs( DR{{.TableName}} ) do
    mt( v, _base )
end