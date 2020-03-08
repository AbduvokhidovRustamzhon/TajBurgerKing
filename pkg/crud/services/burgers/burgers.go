package burgers

import (
	"context"
	"crud/pkg/crud/errors"
	"crud/pkg/crud/models"
	"crud/pkg/crud/services"
	errors2 "errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BurgersSvc struct {
	pool *pgxpool.Pool // dependency
}

func (service *BurgersSvc) InitDB() error {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return errors.ApiError("can't init db: ", err)
	}
	_, err = conn.Query(context.Background(), services.BurgersDDL)
	if err != nil {
		return errors.ApiError("can't init db: ", err)
	}
	return nil
}
func NewBurgersSvc(pool *pgxpool.Pool) *BurgersSvc {
	if pool == nil {
		panic(errors2.New("pool can't be nil")) // <- be accurate
	}
	return &BurgersSvc{pool: pool}
}

func (service *BurgersSvc) BurgersList() (list []models.Burger, err error) {
	list = make([]models.Burger, 0)
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return nil, errors.ApiError("can't execute pool: ", err)
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), services.GetBurgers)
	if err != nil {
		return nil, errors.ApiError("can't query: execute pool", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Burger{}
		err := rows.Scan(&item.Id, &item.Name, &item.Price, &item.Description)
		if err != nil {
			return nil, errors.ApiError("can't scan row: ", err)
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (service *BurgersSvc) Save(model models.Burger) (err error) {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return errors.ApiError("can't execute pool: ", err)
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), services.SaveBurger, model.Name, model.Price, model.Description)
	if err != nil {
		return errors.ApiError("can't save burger: ", err)
	}
	return nil
}
func (service *BurgersSvc) RemoveById(id int64) (err error) {

	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return errors.ApiError("can't execute pool: ", err)
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), services.RemoveBurger, id)
	if err != nil {
		return errors.ApiError("can't remove burger: ", err)
	}
	return nil
}
