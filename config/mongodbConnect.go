


package mogodbConnect

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func main(){
	var ctx = context.TODO()
	mongoDBClientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	mongoDBClient = mongo.connect(ctx, mongoDBClientOptions);
}