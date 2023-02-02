package migrate_controller

import (
	"github.com/lowl11/lazylog/layers"
	"wkey-core/src/definition"
)

/*
	_categories запуск миграции категорий с Prosklad
*/
func (controller *Controller) _categories() error {
	logger := definition.Logger

	if err := controller.migrate.Categories(); err != nil {
		logger.Error(err, "Call migrating categories error", layers.Service)
		return err
	}

	return nil
}

/*
	_products запуск миграции продуктов с Prosklad
*/
func (controller *Controller) _products() error {
	logger := definition.Logger

	if err := controller.migrate.Products(); err != nil {
		logger.Error(err, "Call migrating products error", layers.Service)
		return err
	}

	return nil
}
