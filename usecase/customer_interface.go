package usecase

import (
	"context"
	"fmt"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/model/entities"
)

func (c *Usecase) FindById(ctx context.Context, id string) (entities.TblCustomer, error) {
	customer, err := c.repo.GetCustomer(ctx, id)
	if err != nil {
		return entities.TblCustomer{}, fmt.Errorf("failed to find customer by id! :%v", err)
	}

	return customer, nil
}

func (c *Usecase) RegisterCustomer(ctx context.Context, arg entities.TblCustomer) (entities.TblCustomer, error) {
	customer, err := c.repo.CreateCustomers(ctx, arg)
	if err != nil {
		return entities.TblCustomer{}, fmt.Errorf("register failed! :%v", err)
	}

	return customer, nil
}

func (c *Usecase) EditCustomer(ctx context.Context, arg entities.TblCustomer) error {
	err := c.repo.UpdateCustomer(ctx, arg)
	if err != nil {
		return fmt.Errorf("update failed! :%v", err)
	}
	return nil

}

func (c *Usecase) FetchingAllCustomers(ctx context.Context) ([]entities.TblCustomer, error) {
	customers, err := c.repo.ListCustomers(ctx)
	if err != nil {
		return []entities.TblCustomer{}, fmt.Errorf("failed to retrieve Customers data!. error: %v", err)
	}

	return customers, nil

}

func (c *Usecase) DeleteCustomer(ctx context.Context, id string) error {
	err := c.repo.DeleteCustomer(ctx, id)
	if err != nil {
		return fmt.Errorf("the customer deletion process failed!. :%v", err)
	}
	return nil
}
