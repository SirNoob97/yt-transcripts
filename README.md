# yt-transcripts

Project inspired by [youtube-transcript-api](https://github.com/jdepoix/youtube-transcript-api)

```
Usage off: yt-transcripts: [COMMAND] [OPTIONS]

Commands:
  save    Save the transcript to the specified file path.
  list    List available video transcripts.
  fetch   Fetch the transcript.

Options:
  --help, -h      Display command help message.
  --version, -v   Show app version.
```

```
Usage of: ./yt-transcripts list [OPTIONS]

Options:
  -i, --id          Video ID
```

```
Usage of: yt-transcripts save [OPTIONS]

Options:
  -i, --id          Video ID
  -l, --language    Language code in which you want to store the transcript
  -o, --output      Filename in which the data will be stored
```

```
Usage of: ./yt-transcripts fetch [OPTIONS]

Options:
  -i, --id          Video ID
  -l, --language    Language code in which you want to search for the transcript
```
