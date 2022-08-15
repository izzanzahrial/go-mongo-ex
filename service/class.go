package service

import (
	"context"
	"strconv"

	"github.com/izzanzahrial/go-mongo-ex/db"
	"github.com/izzanzahrial/go-mongo-ex/entity"
	"github.com/izzanzahrial/go-mongo-ex/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type class struct {
	collectionName string
	repository     repository.Class
	db.Connection
}

func (c *class) CreateClass(ctx context.Context, class *entity.Class) error {
	collection := c.DB().Collection(c.collectionName)

	if err := c.repository.CreateClass(ctx, collection, class); err != nil {
		return err
	}

	return nil
}

func (c *class) UpdateClass(ctx context.Context, class *entity.Class) error {
	collection := c.DB().Collection(c.collectionName)

	if err := c.repository.UpdateClass(ctx, collection, class); err != nil {
		return err
	}

	return nil
}

func (c *class) DeleteClass(ctx context.Context, class *entity.Class) error {
	collection := c.DB().Collection(c.collectionName)

	if err := c.repository.DeleteClass(ctx, collection, class); err != nil {
		return err
	}

	return nil
}

func (c *class) GetClassById(ctx context.Context, id string) (*entity.Class, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := c.DB().Collection(c.collectionName)

	class, err := c.repository.GetClassById(ctx, collection, objectID)
	if err != nil {
		return nil, err
	}

	return class, nil
}

func (c *class) GetClassByName(ctx context.Context, name string) ([]*entity.Class, error) {
	collection := c.DB().Collection(c.collectionName)

	classes, err := c.repository.GetClassByName(ctx, collection, name)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (c *class) GetClassPeriod(ctx context.Context, period string) ([]*entity.Class, error) {
	intPeriod, err := strconv.Atoi(period)
	if err != nil {
		return nil, err
	}

	collection := c.DB().Collection(c.collectionName)

	classes, err := c.repository.GetClassPeriod(ctx, collection, intPeriod)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (c *class) GetAllClass(ctx context.Context) ([]*entity.Class, error) {
	collection := c.DB().Collection(c.collectionName)

	classes, err := c.repository.GetAllClass(ctx, collection)
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (c *class) DeleteAllClass(ctx context.Context) error {
	collection := c.DB().Collection(c.collectionName)

	if err := c.repository.DeleteAllClass(ctx, collection); err != nil {
		return err
	}

	return nil
}
