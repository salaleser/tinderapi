# Tinder REST API client

- [x] Get recommendations
- [x] Get matches
- [x] Get profile info
- [x] Set profile info
- [x] Get user info
- [x] Get match info
- [ ] Get message info
- [x] Like user
- [x] Pass user
- [x] Send message

# Autolike example
```go
package main

import (
    "github.com/salaleser/tinderapi/cmd/tinderapi"
    "os"
    "time"
)


func main() {
    c := tinderapi.NewClient()

    c.Login(os.Getenv("TINDER_TOKEN"))
    
    for i := 100; i > 0; i-- {
        x, _ := c.GetRecommendations()

        for _, r := range x.Data.Results {
            like, _ := c.Like(r.User.ID)
            if like.LikesRemaining == 0 {
                i = 0
            }
        }

        time.Sleep(time.Millisecond * 50)
    }
}
```
