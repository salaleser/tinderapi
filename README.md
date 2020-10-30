# Tinder REST API client

- [x] Get recommendations
- [x] Get matches
- [x] Get profile info
- [x] Get user info
- [x] Get match info
- [ ] Get message info
- [x] Like user
- [x] Pass user
- [x] Send message

# Usage
```go
client, _ := tinderapi.New("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")

page, _ := client.GetRecommendations()

for i, v := range page.Data.Results {
	fmt.Printf(
		"%d: %s (%d) [%s]\n",
		i+1,
		v.User.Name,
		time.Now().Year()-v.User.BirthDate.Year(),
		v.User.ID,
	)

	like, _ := client.Like(v.User.ID)
	fmt.Printf("%v\n\n", like)

	time.Sleep(time.Millisecond * 100)
}
```