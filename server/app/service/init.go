package service

import "gf-admin/app/service/guard"

var (
	Guard = guard.NewJwtGuard()
)
