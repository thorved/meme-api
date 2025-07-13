
# Meme API (Because the Internet Needs More Memes)

Welcome to the Meme API, the pinnacle of modern software engineering. Ever wanted to stream memes via HTTP? Of course you have. Who needs productivity when you can have memes?

## Features (If You Can Call Them That)

- `/` — Redirects you to a video. Is it useful? No. Is it fun? Maybe.
- `/health` — Because every API needs to pretend it's healthy.
- `/meme` — Get a random meme. Because your day wasn't already wasted.
- `/meme/{subreddit}` — Specify a subreddit. Because you have very specific meme needs.

## Getting Started (If You Insist)

### Build & Run Locally

```sh
go build -o meme-api ./cmd/meme-api/main.go
./meme-api
```

Runs on port `3000`. Why? Because we said so.

### Docker (For People Who Like Containers)

```sh
docker build -t meme-api .
docker run -p 8080:8080 meme-api
```

Or, if you want to feel fancy:

```sh
docker-compose up --build
```

## License

MIT. Because why not?




## Example Meme (Because You Deserve It)

<div align="center">
  <img src="https://meme.ved.yt/meme" alt="Meme" width="400" />
</div>

## Sarcastic Credit
Special thanks (or blame) to [D3vd/Meme_Api](https://github.com/D3vd/Meme_Api) for making memeapi possible. Without it, this project would just be a sad, meme-less void.

## Disclaimer

This API will not improve your life, but it might make you laugh. Or cry. Or both.
