# wikiwatch

With wikiwatch you can monitor anonymous edits made in the Wikipedia and match the source IP address of the edit with a
known list of IP ranges to assign them to a specific organisation,company,etc. and send a Toot or Tweet.

This tool is meant as a successor to [edsu/anon](https://github.com/edsu/anon), where we switched from using the IRC
channel for getting updates, to the new [EventStreams](https://wikitech.wikimedia.org/wiki/Event_Platform/EventStreams).

We mainly built this tool to migrate our own bots:
  * @bundesedit ([Twitter](https://twitter.com/bundesedit),[Mastodon](https://botsin.space/@bundesedit))
  * @euroedit ([Twitter](https://twitter.com/euroedit),[Mastodon](https://botsin.space/@bundesedit))
  * @landesedit ([Twitter](https://twitter.com/landesedit),[Mastodon](https://botsin.space/@bundesedit))
  * @politikedit ([Twitter](https://twitter.com/politikedit))

But if you also run a bot with this tool, we would love a PR with the name and link added to this list from you :)


## Migration from edsu/anon 
If you previously used the `anon` tool, you can use the *convertRanges.py* tool, to convert the json syntax of the old
ipranges file to the new syntax used by this tool. 

## Get Mastodon credentials
  * Choose a Mastodon instance (There are servers specifically for bots like e.g. botsin.space)
  * Create an Account
  * Go to `<serveraddress>/settings/applications` and create a new application
  * Give it only the `write:statuses` scope
  * Click you application to get client key, client secret and access token
  * Add a `mastodon` section to your config file (check the example config)

As soon as the bot finds a mastodon section in the config, it will start sending toots.

## Get Twitter credentials
  * Sign up for a Developer Account over [here](https://developer.twitter.com/en/apply-for-access)
  * Go to the [Projects&Apps Overview](https://developer.twitter.com/en/portal/projects-and-apps) and create a new
      project
  * At the last step click "New app" and configure the new app inside your project
  * Click the black button labeled "App settings" to go to the settings page of your new app
  * Click "Keys and tokens" in the top navigation
  * Click "Generate" in the box labeled "Access Token and secrets"
  * Add a `twitter` section to your config file (check the example config) and paste the values from the popup (They
      will only be shown once!)

As soon as the bot finds a twitter section in the config, it will start sending tweets.

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
