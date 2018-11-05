# A simple demo on some web frameworks

## Install dependencies
~~~~
> go get ./...
~~~~

## Build and run
#### Echo
~~~~
> cd echo
> go build
> ./echo
~~~~

#### banjo
~~~~
> cd banjo
> go build
> ./banjo
~~~~
banjo doesn't support middleware and path/url query param, so only 1 API implemented

## Test with cURL
~~~~
> curl http://127.0.0.1:4321/user/radaiming/followers
[
  {
    "avatar_url": "http://xxx...",
    "events_url": "",
    "followers_url": "",
    "following_url": "",
    "gists_url": "",
    "gravatar_id": "",
    "html_url": "",
    "id": 0,
    "login": "",
    "node_id": "",
    "organizations_url": "",
    "received_events_url": "",
    "repos_url": "",
    "site_admin": false,
    "starred_url": "",
    "subscriptions_url": "",
    "type": "",
    "url": ""
  }
]
~~~~

~~~~
> curl -H 'Content-Type: application/json' --data-binary '{"message": "release v1.0.0"}' http://127.0.0.1:4321/repos/radaiming/web-frameworks-demo/git/tags
{
  "message": "release v1.0.0",
  "node_id": "",
  "object": {
    "sha": "",
    "type": "",
    "url": ""
  },
  "sha": "",
  "tag": "",
  "tagger": {
    "date": "",
    "email": "",
    "name": ""
  },
  "url": "",
  "verification": {
    "payload": null,
    "reason": "",
    "signature": null,
    "verified": false
  }
}
~~~~