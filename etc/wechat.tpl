事件状态：P{{.Priority}} {{.Status}}
策略名称：{{.Sname}}
endpoint：{{.Endpoint}}
metric：{{.Metric}}
tags：{{.Tags}}
当前值：{{.Value}}
报警说明：{{.Info}}
触发时间：{{.Etime}}
报警详情：{{.Elink}}
{{if .IsUpgrade}}
---
报警已升级!!!
{{end}}