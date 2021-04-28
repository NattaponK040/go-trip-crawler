package repository

import (
	"context"
	"go-crawler/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type MgRepository struct {
	ctx            context.Context
	wnTravel       *mongo.Collection
	wnTravelReview *mongo.Collection
	wnHotel        *mongo.Collection
	wnHotelReview  *mongo.Collection
	wnShop         *mongo.Collection
	wnShopReview   *mongo.Collection

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
		wnTravel:       db.Database(mcf.DBname).Collection(mcf.WonnaiTravelData),
		wnTravelReview: db.Database(mcf.DBname).Collection(mcf.WonnaiTravelReview),
		wnHotel:        db.Database(mcf.DBname).Collection(mcf.WonnaiHotelData),
		wnHotelReview:  db.Database(mcf.DBname).Collection(mcf.WonnaiHotelReview),
		wnShop:         db.Database(mcf.DBname).Collection(mcf.WonnaiShopData),
		wnShopReview:   db.Database(mcf.DBname).Collection(mcf.WonnaiShopReview),

		tripTravel:       db.Database(mcf.DBname).Collection(mcf.TripadvisorTravelData),
		tripTravelReview: db.Database(mcf.DBname).Collection(mcf.TripadvisorTravelReview),
		tripHotel:        db.Database(mcf.DBname).Collection(mcf.TripadvisorHotelData),
		tripHotelReview:  db.Database(mcf.DBname).Collection(mcf.TripadvisorHotelReview),
		tripShop:         db.Database(mcf.DBname).Collection(mcf.TripadvisorShopData),
		tripShopReview:   db.Database(mcf.DBname).Collection(mcf.TripadvisorShopReview),

		db: db,
	}
}
