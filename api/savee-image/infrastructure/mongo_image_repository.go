package image_infrastructure

import (
	"context"
	"log"

	image_domain "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/domain"
	infrastructure_shared "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/shared/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoImageCollection = infrastructure_shared.MongoDatabase.Collection("images")

type MongoImageRepository struct {
	collection *mongo.Collection
}

func NewMongoImageRepository() *MongoImageRepository {
	return &MongoImageRepository{
		collection: MongoImageCollection,
	}
}

func (m *MongoImageRepository) Add(image *image_domain.Image) {
	imageBson := bson.M{
		"id":    image.ID,
		"url":   image.Url,
		"Autor": image.Autor,
		"Alt":   image.Alt,
		"Size":  image.Size,
	}

	_, err := m.collection.InsertOne(context.Background(), imageBson)

	if err != nil {
		log.Println("Error while saving todo", err)
	}
}

func (m *MongoImageRepository) GetAll() []*image_domain.Image {
	cursor, err := m.collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Println("Error while getting all images", err)
		return nil
	}

	allImages := []*image_domain.Image{}
	err = cursor.All(context.Background(), &allImages)

	if err != nil {
		log.Println("Error while getting all images", err)
		return nil
	}

	return allImages
}

func (m *MongoImageRepository) Remove(id int) {
	_, err := m.collection.DeleteOne(context.TODO(), bson.D{{"id", id}})

	if err != nil {
		log.Println("Error while removing image", err)
	}
}

func (m *MongoImageRepository) Update(image *image_domain.Image) {
	filter := bson.D{{"id", image.ID}}
	update := bson.D{
		{
			"$set",
			bson.D{
				{"id", image.ID},
				{"url", image.Url},
				{"autor", image.Autor},
				{"alt", image.Alt},
				{"size", *image.Size},
			},
		},
	}

	_, err := m.collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Println("Error while updating image", err)
	}
}

func (m *MongoImageRepository) Find(id int) *image_domain.Image {
	filter := bson.D{{"id", id}}
	image := &image_domain.Image{}

	err := m.collection.FindOne(context.TODO(), filter).Decode(image)

	if err != nil {
		return nil
	}

	return image
}
