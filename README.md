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
