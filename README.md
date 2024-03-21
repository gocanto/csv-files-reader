### About it

The given API allows users to send bulk data to your
servers using CSV files. This opens the door ro various integrations
where the information you need to handle in your app is given by
3rd party services.

### Installation

The application has been containerised using [Docker](https://www.docker.com/). Thus, you will need to
have installed in your machine before running this repository code upon tests.

Then you have installed docker in your local machine, you might want to also install
the [Make command line](https://makefiletutorial.com/) to use the handy commands contained within
this repository.

After you have these two (2) installed in your local machine, you can go ahead and clone
this repository like so `git@github.com:gocanto/ohlc-price-data.git` within your
desire directory in your local machine.

Once you have cloned the repository, position yourself in the clone directory from within
you command line a type the following `cp .env.example .env`. This will make sure your
system has the proper env vars needed to run the app locally.

> Notes:
> - Make sure you have git install before cloning this repository.
> - We have shipped a [postman collection](https://github.com/gocanto/ohlc-price-data/tree/main/__fixtures__) for you to have at hand when testing the app.
> - You will need some sort of MySQL client if you would like to inspect the database.

### What endpoints does this application contain?

We only have two (2) endpoints at the moment. Once to upload any required
CSV file (more options in the future) and one to query the data uploaded to the server.

```bash
[POST] http://localhost:8080/upload
[GET] http://localhost:8080/query
```

Please note that the above URLs and ports depend 100% on your local/servers configuration.

### How to use it

You can either use the available PostMan collections within this repository. Otherwose, you
might want to try using your terminal like so:

***Upload data***
```bash
curl --location 'http://localhost:8080/upload' \
--header 'Content-Type: multipart/form-data' \
--form 'file=@"/Users/gus/Sites/ohlc-price-data/__fixtures__/payload.csv"'
```

***Query data***
```bash
curl --location --request GET 'http://127.0.0.1:8080/query?limit=1&offset=3' \
--header 'Content-Type: application/json' \
--data '{
    "symbol": "BTCUSDT",
    "unix": "",
    "foo": 1
}'
```
> Please do note that the above filter is invalid on purpose for you to test
> the filtering ability.

### License

Please see the [license file](https://github.com/gocanto/ohlc-price-data/blob/main/LICENSE) for more information.

## How can I thank you?

There are many ways you would be able to support my open source work. There is not a right one to choose, so the choice is yours.

Nevertheless :grinning:, I would propose the following

- :arrow_up: Follow me on [Twitter](https://twitter.com/gocanto).
- :star: Star the repository.
- :handshake: Open a pull request to fix/improve the codebase.
- :writing_hand: Open a pull request to improve the documentation.

> Thank you for reading this far. :blush:



