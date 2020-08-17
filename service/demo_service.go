package service

import (
	"time"

	"github.com/my-gin-server/base/applog"
	"github.com/my-gin-server/dao"
	"github.com/my-gin-server/model"
	"github.com/my-gin-server/service/dto"
)

type DemoService struct {
	demoDao *dao.DemoDao
}

func NewDemoService(demoDao *dao.DemoDao) *DemoService {
	return &DemoService{demoDao: demoDao}
}

func (service *DemoService) Create(demoDTO *dto.DemoDTO) (int, error) {
	applog.Info.Printf("This extra info: \"%s\" is not useful at all "+
		"but to explain why we need a dto sometimes.",
		demoDTO.ExtraInfo)
	demo := &model.Demo{
		Name:        demoDTO.Name,
		Description: demoDTO.Description,
		CreateTime:  time.Now().UTC(),
	}
	return service.demoDao.Create(demo)
}

func (service *DemoService) DeleteById(id int) error {
	return service.demoDao.DeleteById(id)
}

func (service *DemoService) DeleteRange(offset int, limit int) (int, error) {
	// Do not implement like this in real project.
	// This is just a example(not that good but simple...) of transaction usage.
	demos, _, err := service.List(offset, limit)
	if err != nil {
		return 0, err
	}

	// Transactions always involve many DAOs,
	// so usually we need a transaction creation at the service level.
	tx, err := service.demoDao.BeginTx()
	if err != nil {
		return 0, err
	}
	// The rollback will be ignored if the tx has been committed later in the function.
	defer tx.Rollback()

	countDeleted := 0
	for _, each := range demos {
		err := service.demoDao.DeleteWithTx(tx, each.Id)
		if err != nil {
			continue
		}
		countDeleted += 1
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return countDeleted, nil
}

func (service *DemoService) List(offset int, limit int) (
	[]*model.Demo, int, error) {
	return service.demoDao.List(offset, limit)
}

func (service *DemoService) QueryById(id int) (*model.Demo, error) {
	return service.demoDao.QueryById(id)
}
