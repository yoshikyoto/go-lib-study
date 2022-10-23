// + wireinject
// ↑wireの設定ファイルを通常のビルド対象から外す
package main

import "github.com/google/wire"

func InitializeService() Service {
	wire.Build(NewService, NewRepositoryImpl)
	// 戻り値は動作に影響しないので何でも良い
	return Service{}
}
