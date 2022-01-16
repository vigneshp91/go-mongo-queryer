package mongodb

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"

	"log"
	"sync"

	pgx "gopkg.in/jackc/pgx.v2"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

//Pool of database connection
var ConnPool *pgx.ConnPool
var once sync.Once

// Connect connects to the Mongo Cluster using the given URI
func Connect(uri string) {
	var err error
	// cfg := log.NewConfig(config.AppName)
	// cfg.SetRemoteConfig(config.LogRemoteURL, config.LogToken, config.LogUname)
	// cfg.SetLevelStr(config.LogLevel)
	// cfg.SetReference("db")
	// cfg.SetFilePathSizeStr(config.LogFilePathSize)
	// l := log.New(cfg)
	client, err = getClient(uri)
	if err != nil {
		log.Print("Error getting client in DB init()")
	}
	log.Print("Client created in DB init()")
}

// Connection is a DB Connection Object
type Connection struct {
	mongo.Client
}

// NewConnection returns a New Mongo Connection
// using the package Client variable
func NewConnection() *Connection {
	return &Connection{
		Client: *client,
	}
}

// IsConnected Returns the status of connection by performing a Ping
func (c *Connection) IsConnected() bool {
	log.Print("Pinging Mongo Cluster: %#v", c)

	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Print("Error in Pinging Mongo Cluster: %#v", err)
		return false
	}

	return true
}

// Databaser interface for accessing DB of a give name
type Databaser interface {
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
	IsConnected() bool
}

// DatabaseQueryer contains useful methods applicable for a database
// 1. Collection Helper Functions
// 2. Client Handle
type DatabaseQueryer interface {
	Name() string
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
	ListCollectionNames(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) ([]string, error)
	ListCollections(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) (*mongo.Cursor, error)
	Client() *mongo.Client
}

// getClient returns a Mongo Client connection
func getClient(uri string) (*mongo.Client, error) {
	// cfg := log.NewConfig(config.AppName)
	// cfg.SetRemoteConfig(config.LogRemoteURL, config.LogToken, config.LogUname)
	// cfg.SetLevelStr(config.LogLevel)
	// cfg.SetReference("db")
	// cfg.SetFilePathSizeStr(config.LogFilePathSize)
	// l := log.New(cfg)

	//log.Print("Establishing Connection to MongoDB URI = %s", uri)
	var clientOptions *options.ClientOptions
	clientOptions = options.Client().ApplyURI(uri)

	// if config.Env != "dev" {
	// 	file, err := helpers.DownloadFile(config.PemFile)
	// 	if err != nil {
	// 		log.Print("Cannot download file", err)
	// 		return nil, err
	// 	}

	// 	tlsConfig, err := getCustomTLSConfig(file.Name())
	// 	if err != nil {
	// 		log.Print("Cannot get custom tls config", err)
	// 		return nil, err
	// 	}

	// 	clientOptions = options.Client().ApplyURI(uri).SetTLSConfig(tlsConfig).SetRetryWrites(false)
	// } else {
	// 	clientOptions = options.Client().ApplyURI(uri)
	// }
	if clientOptions == nil {
		return nil, errors.New("Unable to connect to mongo")
	}

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Print("Cannot create New Client for Mongo", err)
		return nil, err
	}

	log.Print("Mongo Client Created Successfully!!")

	err = client.Connect(context.Background())
	if err != nil {
		log.Print("Cannot connect to Mongo Client", err)
		return nil, err
	}

	log.Print("Mongo Client Connected to Cluster Successfully!!")

	return client, nil
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("failed parsing pem file")
	}

	return tlsConfig, nil
}
