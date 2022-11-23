# wikiwatch

With wikiwatch you can monitor anonymous edits made in the Wikipedia and match the source IP address of the edit with a
known list of IP ranges to assign them to a specific organisation,company,etc. and send a Toot or Tweet.

This tool is meant as a successor to [edsu/anon](https://github.com/edsu/anon).

We mainly built this tool to migrate our own bots:
  * @bundesedit
  * @euroedit
  * @landesedit
  * @politikedit

If you also run a bot with this tool, we would love a PR with the name and link added to this list from you :)


## Migration from edsu/anon 
If you previously used the `anon` tool, you can use the *convertRanges.py* tool, to convert the json syntax of the old
ipranges file to the new syntax used by this tool. 

## Get Mastodon credentials
Todo...

## Get Twitter credentials
Todo...

## Usage
Take the *config-example.json* and copy it to *config.json* and addapt to your needs. If you don't plan on using Twitter
or Mastodon, you can completly remove the whole block from the config. Depending on the number of organisation you want
to monitor, you can either put the IP ranges directly into the config file under *organisations* or put them in a
seperate json (this way you can publish them publicly) and provide the path to the json in *organisations_file*.

Then either take the compiled binary from the [Github releases](https://github.com/codemonauts/wikiwatch/releases) or
download the Docker container:
### Standalone
```bash
./wikiwatch -config ./config.json -loglevel INFO
```

or take the example systemd unitfile from this repo and deploy it.


### Docker
```bash
docker run --rm -v ${PWD}/config.json:/config.json ghcr.io/codemonauts/wikiwatch
```



With ❤ by [codemonauts](https://codemonauts.com)
