# Tinder REST API client

- [x] Recommendations
- [x] Profile
- [x] User by ID
- [x] Like
- [x] Pass

# Usage
```
client := tinderapi.New("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")

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