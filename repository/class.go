package repository

import (
	"context"

	"github.com/izzanzahrial/go-mongo-ex/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type class struct {
}

func (c *class) CreateClass(ctx context.Context, collection *mongo.Collection, class *entity.Class) error {
	if _, err := collection.InsertOne(ctx, class); err != nil {
		return err
	}

	return nil
}

func (c *class) UpdateClass(ctx context.Context, collection *mongo.Collection, class *entity.Class) error {
	bsonClass, err := bson.Marshal(class)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": class.Id}
	update := bson.M{"$set": bsonClass}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func (c *class) DeleteClass(ctx context.Context, collection *mongo.Collection, class *entity.Class) error {
	filter := bson.M{"_id": class.Id}
	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (c *class) GetClassById(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (*entity.Class, error) {
	var class entity.Class

	if err := collection.FindOne(ctx, id).Decode(&class); err != nil {
		return nil, err
	}

	return &class, nil
}

func (c *class) GetClassByName(ctx context.Context, collection *mongo.Collection, name string) ([]*entity.Class, error) {
	filter := bson.M{"name": name}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var classes []*entity.Class

	if err := cursor.All(ctx, &classes); err != nil {
		return nil, err
	}

	return classes, nil
}

func (c *class) GetClassPeriod(ctx context.Context, collection *mongo.Collection, period int) ([]*entity.Class, error) {
	filter := bson.M{"period": period}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var classes []*entity.Class

	if err := cursor.All(ctx, &classes); err != nil {
		return nil, err
	}

	return classes, nil
}

func (c *class) GetAllClass(ctx context.Context, collection *mongo.Collection) ([]*entity.Class, error) {
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var classes []*entity.Class

	if err := cursor.All(ctx, &classes); err != nil {
		return nil, err
	}

	return classes, nil
}

func (c *class) DeleteAllClass(ctx context.Context, collection *mongo.Collection) error {
	_, err := collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}

	return nil
}
