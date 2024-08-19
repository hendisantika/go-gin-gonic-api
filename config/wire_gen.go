package config

// Injectors from injector.go:

func Init() *Initialization {
	gormDB := ConnectToDB()
	userRepositoryImpl := repository.UserRepositoryInit(gormDB)
	userServiceImpl := service.UserServiceInit(userRepositoryImpl)
	userControllerImpl := controller.UserControllerInit(userServiceImpl)
	roleRepositoryImpl := repository.RoleRepositoryInit(gormDB)
	initialization := NewInitialization(userRepositoryImpl, userServiceImpl, userControllerImpl, roleRepositoryImpl)
	return initialization
}
