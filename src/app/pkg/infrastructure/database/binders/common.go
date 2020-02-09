package binders

import "gonference/pkg/utils"

type Scanner = func(dest ...interface{}) error

func scan(scanner Scanner, dest ...interface{}) {
	utils.Check(scanner(dest...))
}
