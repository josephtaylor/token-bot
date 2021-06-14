# token-bot

This is an ERC-20 Telegram Bot. It has functions to
display the current price information and to display
the current balance for a wallet address.

The telegram bot uses a _polling-based_ implementation,
which eliminates the need to register webhooks.

# Configuration

There is an example configuration file in `./config`
called `config.example.yml`

Rename this file to `config.yml` and add your values to it.

## Dynamic Configuration Values

* `app.etherscan.apiKey` - an etherscan API key.
* `app.pairAddress` - the address of the weth/token pair.
* `app.totalTokens` - the total number of token supply.
* `app.website` - the website hostname
* `app.twitter` - the twitter account with "@"
* `app.telegram.baseUri` - the API url for your telegram bot.

The rest of the values are static and don't need to
be updated.

# Executing the bot

To run the bot with docker:

```bash
docker run \
  -v ./config/config.yml:/config/config.yml \
  josephtaylor/token-bot bot
```

To run the bot outside docker:

```bash
./token bot
```

To enable debug logging use `-d`:

```bash
./token bot -d
```

# Building

## Requirements

* `make`
* `docker`

## Building the Project

run `make` in the root directory.

# Donations

If you use this bot for your project,
consider making a donation!

Send ETH or tokens to the following address:

```
0x5a61F59F41bE917129d12051F19d29B595452535
```