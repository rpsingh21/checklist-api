package service

// NewDataBaseConnection Return new Connection of db
// func NewDataBaseConnection(dbConnection string, logger *zap.SugaredLogger) *mongo.Database {
// 	clientOptions := options.Client().ApplyURI(dbConnection)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		logger.Errorf("Database connection failed: %v", err)
// 	}
// 	if client.Connect(ctx); err != nil {
// 		logger.Errorf("Database connection failed: %v", err)
// 	}
// 	// defer client.Disconnect(*ctx)

// 	db := client.Database("checkList")
// 	return db
// }
