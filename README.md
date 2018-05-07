# Exercise #16: Twitter Contest CLI

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/twitter)

## Exercise details

The Twitter Contest CLI exercise can be summed up in roughly 5 steps:

1.  Create a Twitter App
2.  Authenticate with the Twitter API
3.  Use the Twitter API to determine which users have retweeted a specific tweet.
4.  Open (or create) a file containing a list of unique usernames who have retweeted the tweet, and update it with new users retrieved from your API call in step (2).
5.  If a flag or some other option is set to pick a winner, read the full list of unique usernames and pick a winner and then print out that winner's username. You can also add support for picking multiple winners.

Below I'll give some guidance to help you out with each step, along with an explanation of why each step is helpful or useful.

### Create a Twitter App

To create a Twitter App, head over to <http://apps.twitter.com/>, sign in, then create a new app.

After creating your app, grab the authentication URLs from the bottom of the page:

|                         |                                             |
|-------------------------|---------------------------------------------|
| App-only authentication | https://api.twitter.com/oauth2/token        |
| Request token URL       | https://api.twitter.com/oauth/request_token |
| Authorize URL           | https://api.twitter.com/oauth/authorize     |
| Access token URL        | https://api.twitter.com/oauth/access_token  |

Finally, head to the "Keys and Access Tokens" tab and get your Consumer Key and Secret.

|              |                          |
|--------------|--------------------------|
| Consumer Key | `Lx.... your key here`   |
| Consumer Key | `ED... your secret here` |

If you plan on using a third party library that needs an Access Token and Secret then you also need to scroll to the bottom of the "Keys and Access Tokens" tab and generate one of these. This will give the app access to your personal Twitter account (*I think...*). We won't need this for the the Gophercises videos because we will be using application-only authentication, but if you are using a third party Twitter library like [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda) I'm pretty certain they expect this information.


### Authenticate with the Twitter API

Before you can determine who has retweeted a specific tweet, you need to authenticate with the Twitter API. There are basically two authentication flows supported by Twitter that will work here:

1.  [Using OAuth](https://developer.twitter.com/en/docs/basics/authentication/overview/using-oauth)
2.  [Application-only Authentication](https://developer.twitter.com/en/docs/basics/authentication/overview/application-only)

I suggest using application-only authentication (2) because it suits our needs best, and won't require you to give the application access to your Twitter account. That said, if you are using a third party library like [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda) it will probably be easier to just generate an access token on your app settings page and (*I think*) this uses the OAuth flow. I'm not certain though, because the application-only authentication flow uses some parts of OAuth towards the end of the flow.

Regardless of which you choose, follow the Twitter docs or your third party library's instructions until you can successfully make API calls. If you get stuck, check out the video on this section.


### Use the Twitter API to determine retweeters

The docs for getting a single tweet's retweets are here: <https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/get-statuses-retweets-id>

In the API tweets are called "statuses", which can be confusing but is something we have to deal with.

It is also important to note that the API limits this endpoint to the 100 most recent retweets. That means that if our CLI was being used for a very popular person on Twitter that we might need to hit this API endpoint fairly frequently to keep up with all the people retweeting, and we can't just run the CLI after the contest ends.

To account for this, we will be creating a file to store all unique users who have retweeted and appending this information to it in the next step. For now just try to get a list of the usernames returned from the API - you can ignore pretty much all other info returned by the API endpoint.

### Persist unique usernames of retweeters

Open (or create if it doesn't exist) a file to store all of the usernames of people who have retweeted the tweet.

Read the list of usernames from the file, then merge those with the users returned from the API call in the last step.

Finally, update the file with the full list of users. Be sure to include both those who were previously in the file and the new ones provided by the API call.

This step is necessary because the API only returns the most recent 100 retweets, as explained in the previous section.


### Pick a winner

By default your program should just connect to the Twitter API and update the list of users who have retweeted the tweet, but eventually you will want to pick a winner.

Add a flag or some other way for users to specify that they want to pick a winner in this step. When this value is set, pick a winner from the list of unique retweeters.

Whether you update the list first then pick a winner or just pick a winner from the file is up to you. I will be updating the list one last time then picking a winner, but it is a minor difference and you may not prefer to do this.

If you want to take these a step further, add a way for users of your program to also specify how many winners they want and pick exactly that many winners without any duplicates.


## Bonus

A few bonus ideas are:

1. Turn your code into a nice API client that can be used as a separate package, and then add support for a few more Twitter API endpoints.
2. Use a timer and update your program to automatically poll the Twitter API every `X` minutes getting an updated list of users, then after a set period of time (`Y` minutes/hours/days) have it pick a winner automatically. This woudl enable you to just run the program once and walk away without needing to manually run it every few minutes.

