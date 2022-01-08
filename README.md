# JapaneseNewsAPI
A web API written in Go that digests Japanese news for English speakers. Site runs on Cloud Run, and database uses MongoDB Atlas.

## How it works
This API only has 3 endpoints as of now:
* **/** : This is a GET method. It's like a homepage, just responds with how the user can use the API.
* **/getnews/1999-12-31** : This is a GET method. It retrieves a news JSON object for that date which has the headlines in English and Japanese, links to said headlines, as well as what was trending in Japan at that time in both English and Japanese. An abridged JSON response is shown [here](https://github.com/travis-james/JapaneseNewsAPI/blob/main/sample.png).
* **/updatenews** : This is a POST method. It will retrieve headlines and links from Asahi's and NHK's RSS feeds, and what's trending on Twitter in Japan, and translate that information to English using Google's Translate API. If the date already exists in the database, no update will happen.

You can demo it here:

~~https://jpnewsapi-ifbwemlbfq-uc.a.run.app/~~ I took my app offline now that I've found a job.

Rate limiting only allows 1 request per second, so if you get a "rate limit exceeded" error, that's why.

**Note:** Since the database is updated manually not every date will be in the database.

## Why did I make this?
I was going to make this just a local utility for myself, but thought I'd challege myself and try to deploy it as an app. Seemed like a good opportunity to use various APIs (Twitter, Google Translate), and other technologies (MongoDB).

## What's next for this project?
I'd like to restructure my date/news objects so that I can query by keyword and return headlines with that keyword.
Would also like to add users at some point who can post articles, and have a favorites list.
