package singleton

type singleton struct {
	Name string
}

var createdSingleton *singleton

// New Create Singleton
func New(name string) *singleton {

	if createdSingleton != nil {
		return createdSingleton
	}

	singleton := singleton{
		Name: name,
	}

	createdSingleton = &singleton

	return createdSingleton
}
