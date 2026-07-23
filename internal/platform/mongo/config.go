package mongo

import (
	"blog-management-system/internal/config"
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func Config(cfg *config.Config) *mongo.Client {
	opts := options.Client().
		ApplyURI(cfg.DatabaseURI).
		SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal().Err(err).Msg("connect to MongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal().Err(err).Msg("ping MongoDB")
	}

	log.Info().Msg("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}
