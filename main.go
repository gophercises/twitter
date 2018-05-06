package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gophercises/twitter/twitter"
)

func main() {
	var (
		keyFile    string
		usersFile  string
		tweetID    string
		numWinners int
	)
	flag.StringVar(&keyFile, "key", ".keys.json", "The file where you store your consumer key and secret for the Twitter API.")
	flag.StringVar(&usersFile, "users", "users.csv", "The file where users who have retweeted the tweet are stored. This will be created if it does not exist.")
	flag.StringVar(&tweetID, "tweet", "991053593250758658", "The ID of the Tweet you wish to find retweeters of.")
	flag.IntVar(&numWinners, "winners", 0, "The number of winners to pick for the contest.")
	flag.Parse()

	key, secret, err := keys(keyFile)
	if err != nil {
		panic(err)
	}
	client, err := twitter.New(key, secret)
	if err != nil {
		panic(err)
	}
	newUsernames, err := client.Retweeters(tweetID)
	if err != nil {
		panic(err)
	}
	existUsernames := existing(usersFile)
	allUsernames := merge(newUsernames, existUsernames)
	err = writeUsers(usersFile, allUsernames)
	if err != nil {
		panic(err)
	}

	if numWinners == 0 {
		return
	}
	existUsernames = existing(usersFile)
	winners := pickWinners(existUsernames, numWinners)
	fmt.Println("The winners are:")
	for _, username := range winners {
		fmt.Printf("\t%s\n", username)
	}
}

func keys(keyFile string) (key, secret string, err error) {
	var keys struct {
		Key    string `json:"consumer_key"`
		Secret string `json:"consumer_secret"`
	}
	f, err := os.Open(keyFile)
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	dec.Decode(&keys)
	return keys.Key, keys.Secret, nil
}

func existing(usersFile string) []string {
	f, err := os.Open(usersFile)
	if err != nil {
		return []string{}
	}
	defer f.Close()
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	users := make([]string, 0, len(lines))
	for _, line := range lines {
		users = append(users, line[0])
	}
	return users
}

func merge(a, b []string) []string {
	uniq := make(map[string]struct{}, 0)
	for _, user := range a {
		uniq[user] = struct{}{}
	}
	for _, user := range b {
		uniq[user] = struct{}{}
	}
	ret := make([]string, 0, len(uniq))
	for user := range uniq {
		ret = append(ret, user)
	}
	return ret
}

func writeUsers(usersFile string, users []string) error {
	f, err := os.OpenFile(usersFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	for _, user := range users {
		if err := w.Write([]string{user}); err != nil {
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func pickWinners(users []string, numWinners int) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(users))
	winners := perm[:numWinners]
	ret := make([]string, 0, numWinners)
	for _, idx := range winners {
		ret = append(ret, users[idx])
	}
	return ret
}
