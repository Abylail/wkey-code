package migrate_controller

/*
	CategoriesJob обертка для _categories
*/
func (controller *Controller) CategoriesJob() error {
	return controller._categories()
}

/*
	ProductsJob обертка для _products
*/
func (controller *Controller) ProductsJob() error {
	return controller._products()
}
