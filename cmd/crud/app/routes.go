package app

// за разделение handler'ов по адресам -> routing
func (receiver *server) InitRoutes() {
	mux := receiver.router.(*exactMux)
	// panic, если происходят конфликты
	// Handle - добавляет Handler (неудобно)
	// HandleFunc

	// стандартный mux:
	// - если адрес начинается со "/" - то под действие обработчика попадает всё, что начинается со "/"
	// https://dropmefiles.com/k0P8d
	mux.GET("/", receiver.handleBurgersList())
	mux.POST("/", receiver.handleBurgersList())
	//mux.POST("/burgers/save", receiver.handleBurgersSave())
	//mux.POST("/burgers/remove", receiver.handleBurgersRemove())
	// - но если есть более "специфичный", то используется он
	mux.POST("/burgers/save", receiver.handleBurgersSave())

	mux.POST("/burgers/remove", receiver.handleBurgersRemove())
	mux.GET("/favicon.ico", receiver.handleFavicon())

}
