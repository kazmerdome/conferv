package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	mongoDb struct {
		client   *mongo.Client
		database *mongo.Database
	}
	mongoCollection struct {
		collection *mongo.Collection
	}
	mongoCursor struct {
		cursor *mongo.Cursor
	}
)

/*
	NewMongoDB returns DB
*/

func NewMongoDB(uri string, name string, retrywrites bool) Database {
	if uri == "" {
		panic(errors.New("uri is required"))
	}

	if name == "" {
		panic(errors.New("database name is required"))
	}

	fmt.Println("\nconnecting \033[0;36m", name, "\033[0m db...")
	connectionURI := uri + "/" + name

	if retrywrites {
		connectionURI = connectionURI + "?retryWrites=true"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		fmt.Println(err)
		log.Fatal("Mongo connection error!")
	}

	// Check the connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	database := client.Database(name)
	fmt.Printf("\033[1;32mconnected successfully!\033[0m\n")

	return &mongoDb{
		database: database,
		client:   client,
	}
}

func (d *mongoDb) Disconnect() {
	fmt.Println("\ndisconnecting \033[0;36m", d.database.Name(), "\033[0m db...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	d.database.Client().Disconnect(ctx)
	fmt.Printf("\033[1;32mdisconnected successfully!\033[0m\n")
}

func (d *mongoDb) Collection(name string, opts ...*options.CollectionOptions) Collection {
	return &mongoCollection{collection: d.database.Collection(name, opts...)}
}

/*
	Collection implements all available operations.
*/
func (c *mongoCollection) Drop(ctx context.Context) error {
	return c.collection.Drop(ctx)
}

func (c *mongoCollection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return c.collection.Aggregate(ctx, pipeline, opts...)
}

func (c *mongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return c.collection.Find(ctx, filter, opts...)
}

func (c *mongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return c.collection.FindOne(ctx, filter, opts...)
}

func (c *mongoCollection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return c.collection.BulkWrite(ctx, models, opts...)
}

func (c *mongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return c.collection.CountDocuments(ctx, filter, opts...)
}

func (c *mongoCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return c.collection.DeleteOne(ctx, filter, opts...)
}

func (c *mongoCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return c.collection.DeleteMany(ctx, filter, opts...)
}

func (c *mongoCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.collection.UpdateMany(ctx, filter, update, opts...)
}

func (c *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(ctx, filter, update, opts...)
}

func (c *mongoCollection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return c.collection.InsertMany(ctx, documents, opts...)
}

func (c *mongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.collection.InsertOne(ctx, document, opts...)
}

func (c *mongoCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return c.collection.FindOneAndUpdate(ctx, filter, update, opts...)
}

func (c *mongoCollection) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	return c.collection.FindOneAndDelete(ctx, filter, opts...)
}
