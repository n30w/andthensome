package reddit

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/theckman/yacspin"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var (
	spinner, _ = yacspin.New(
		yacspin.Config{
			Frequency:       100 * time.Millisecond,
			CharSet:         yacspin.CharSets[43],
			Suffix:          " retrieving posts and comments",
			SuffixAutoColon: true,
			Message:         "", // Set this to the page "after" setting from struct
			StopCharacter:   "✓",
			StopColors:      []string{"fgGreen"},
		},
	)
	ResultsPerRedditRequest = 50
)

// GrabSaved reads all cached posts on the Reddit account.
// This can be used to mass refresh an entire SQL database.
func GrabSaved(postsTable, commentsTable *Table[Row[id, text]], key *Key) {

	var mySavedPosts []*reddit.Post
	var mySavedComments []*reddit.Comment
	var response *reddit.Response

	// Last position of for loops
	lastPos1 := 0
	lastPos2 := 0

	totalRequests := 1

	if PleasePopulateIDs {
		ResultsPerRedditRequest = 100
		totalRequests = 10
	}

	ctx := context.Background()
	opts := &reddit.ListUserOverviewOptions{
		ListOptions: reddit.ListOptions{
			Limit:  ResultsPerRedditRequest,
			After:  "",
			Before: "",
		},
		Sort: "new",
		Time: "all",
	}

	// Establish connection to Reddit API
	httpClient := &http.Client{Timeout: time.Second * 30}
	client, err := reddit.NewClient(*key.NewKey(), reddit.WithHTTPClient(httpClient))

	if err != nil {
		log.Println(Warn.Sprint("Login failed :("))
	} else {
		log.Println(Information.Sprint("Contacting Reddit API..."))
	}

	_ = spinner.Start()

	for i := 0; i < totalRequests; i++ {
		mySavedPosts, mySavedComments, response, err = client.User.Saved(ctx, opts)
		if err != nil {
			log.Fatal(Warn.Sprint(err))
		}

		// TODO go routine optimization can occur here
		for _, post := range mySavedPosts {
			postsTable.Rows[lastPos1] = &Row[id, text]{
				0,
				text(post.Title),
				text(post.Permalink),
				text(post.SubredditName),
				text(post.URL),
			}
			lastPos1++
		}

		// TODO go routine optimization can occur here
		for _, comment := range mySavedComments {
			commentsTable.Rows[lastPos2] = &Row[id, text]{
				0,
				text(comment.Author),
				text(comment.Body),
				text(comment.Permalink),
				text(comment.SubredditName),
			}
			lastPos2++
		}

		spinner.Message(opts.ListOptions.After)

		opts.ListOptions.After = response.After
		time.Sleep(1 * time.Second) // Its recommend to hit Reddit with only 1 request/sec
	}

	// TODO rework ID numbering system, causes synchronization issues
	// with the planetscale database.
	if PleasePopulateIDs {
		populateIDs(postsTable, lastPos1)
		populateIDs(commentsTable, lastPos2)
	}

	_ = spinner.Stop()

	log.Print(Result.Sprint("Saved posts and comments retrieved"))
	// log.Print(Result.Sprintf("Comments: %x", commentsTable.Rows))
	if err != nil {
		log.Fatal(Warn.Sprint(err))
	}
}

// populateIDs populates IDs given a new request to Reddit.
// This is used to either refresh a database from the beginning,
// or assign new IDs to additional saved posts to be put into the database.
func populateIDs(t DBTable, lastPosition int) {
	for i, row := range t.Rows {
		if row == nil {
			break
		}
		row.Col1 = id(lastPosition - i)
	}
}