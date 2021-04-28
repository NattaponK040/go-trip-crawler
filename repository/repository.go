package repository

import (
	"context"
	"go-crawler/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type MgRepository struct {
	ctx            context.Context

	tripTravel       *mongo.Collection
	tripTravelReview *mongo.Collection
	tripHotel        *mongo.Collection
	tripHotelReview  *mongo.Collection
	tripShop         *mongo.Collection
	tripShopReview   *mongo.Collection

	db *mongo.Client
}

func NewMongoRepository(db *mongo.Client, ctx context.Context, mcf *config.Mongodb) *MgRepository {
	return &MgRepository{
		ctx:            ctx,
		
		tripTravel:       db.Database(mcf.DBname).Collection(mcf.TripadvisorTravelData),
		tripTravelReview: db.Database(mcf.DBname).Collection(mcf.TripadvisorTravelReview),
		tripHotel:        db.Database(mcf.DBname).Collection(mcf.TripadvisorHotelData),
		tripHotelReview:  db.Database(mcf.DBname).Collection(mcf.TripadvisorHotelReview),
		tripShop:         db.Database(mcf.DBname).Collection(mcf.TripadvisorShopData),
		tripShopReview:   db.Database(mcf.DBname).Collection(mcf.TripadvisorShopReview),

		db: db,
	}
}
