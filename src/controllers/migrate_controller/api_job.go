package migrate_controller

import (
	"github.com/lowl11/lazylog/layers"
	"wkey-core/src/definition"
)

/*
	CategoriesJob обертка для _categories
*/
func (controller *Controller) CategoriesJob() error {
	definition.Logger.Info("Categories migrate job started...", layers.Job)
	return controller._categories()
}

/*
	ProductsJob обертка для _products
*/
func (controller *Controller) ProductsJob() error {
	definition.Logger.Info("Products migrate Job started...", layers.Job)
	return controller._products()
}
