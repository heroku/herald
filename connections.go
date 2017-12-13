package herald

import "log"
import "os"

import "github.com/garyburd/redigo/redis"

// Redis Stuff.
var REDIS_URL = os.Getenv("REDIS_URL")


type Redis struct {
	URL string
	Connection redis.Conn
}

func NewRedis(url string) Redis {
	// Support default $REDIS_URL, if none was provided.
	if url == "" {
		url = REDIS_URL
	}

	r := Redis{
		URL: url,
	}
	r.Connect()
	return r
}

// Opens (and returns) the Redis connection.
func (r *Redis) Connect() redis.Conn {
	c, err := redis.DialURL(r.URL)
	r.Connection = c

	if err != nil {
		// Fail epically.
		log.Fatal(err)
	}

	return c
}

// Closes the Redis connection.
func (r *Redis) Close() {
	defer r.Connection.Close()
}
