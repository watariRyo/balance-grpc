package repository

type AllRepository struct {
	DBConnection 					DBConnection
	DBTransaction					DBTransaction
	RedisClient 					RedisClient
	UserRepository 					UserRepository
	UserTagRepository 				UserTagRepository
	GroupRepository 				GroupRepository
	IncomeAndExpenditureRepository 	IncomeAndExpenditureRepository
}