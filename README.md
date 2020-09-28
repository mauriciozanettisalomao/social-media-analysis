# Social Media Analysis

Work in progress.

## Overview

API to get data from social medias (Twitter is the only implement for now - it needs environment variables).

To run as a container

```
docker build . -t {your-repo}/social-media-analysis

docker run -p 80:80 -e TWITTER_CONSUMER_KEY={your-consumer-key} -e TWITTER_CONSUMER_SECRET={your-consumer-secret} -e TWITTER_ACCESS_TOKEN={your-access-token} -e TWITTER_ACCESS_TOKEN_SECRET={your-token-secret} {your-repo}/social-media-analysis
```

Example to test on Azure:
* [`3 "Brasileirao" portugues mentions`](https://social-media-analysis.azurewebsites.net/socialmedia/twitter/brasileirao/pt?count=3&lang=pt)
* [`5 #dxc english mentions`](https://social-media-analysis.azurewebsites.net/socialmedia/twitter/%23dxc/pt?count=5&lang=en)

## Twitter filters

* [`Standard`](https://developer.twitter.com/en/docs/twitter-api/v1/tweets/search/guides/standard-operators)
* [`Premium|Enterprise`](https://developer.twitter.com/en/docs/twitter-api/v1/tweets/search/guides/premium-operators)

