# JapaneseNewsAPI
A web API written in Go that digests Japanese news for English speakers.

## How it works
This API only has 3 endpoints as of now:
* **/** : This is like a homepage, just responds with how the user can use the API.
* **/getnews/1999-12-31** : This is a GET method. It retrieves a news JSON object for that date which has the headlines in English and Japanese, links to said headlines, as well as what was trending in Japan at that time in both English and Japanese. An abridged JSON response is shown [here](https://github.com/travis-james/JapaneseNewsAPI/blob/main/sample.png).

* **/updatenews** : This is a POST method. It will retrieve headlines and links from Asahi's and NHK's RSS feeds, and what's trending on Twitter in Japan, and translate that information to English using Google's Translate API. If the date already exists in the database, no update will happen.

## Why was this made?
