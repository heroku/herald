package herald

import "log"
import "os"
import "fmt"
import "strings"
import "github.com/deckarep/golang-set"
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

func (r Redis) GetTargets(bp string) []Target {
	targets := mapset.NewSet()
	results := []Target{}

	selector := fmt.Sprintf("%s:%s", bp, "*")
	keys, _ := redis.Strings(r.Connection.Do("KEYS", selector))

	// Add Redis results to set.
	for _, key := range keys {
		targets.Add(strings.Split(key, ":")[1])
	}

	// Convert set results into Target type.
	for _, target := range targets.ToSlice() {
		results = append(results, NewTarget(NewBuildpack(bp), target.(string)))
	}

	return results
}

func (r Redis) GetTargetVersions(bp string, target string) mapset.Set {
	targets := mapset.NewSet()

	selector := fmt.Sprintf("%s:%s:%s", bp, target, "*")
	keys, _ := redis.Strings(r.Connection.Do("KEYS", selector))
	for _, key := range keys {
		targets.Add(strings.Split(key, ":")[2])
	}

	return targets
}