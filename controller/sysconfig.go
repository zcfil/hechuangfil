package controller

import "hechuangfil/models"

type SysConfig struct {
	*models.SysConfigModels
}

func NewSysConfig(u models.SysConfigModels) (*SysConfig){
	return &SysConfig{
		&u,
	}
}