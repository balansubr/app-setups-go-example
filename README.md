Simple Go script demonstrating use of the [Heroku Platform API app-setups resource](https://devcenter.heroku.com/articles/setting-up-apps-using-the-heroku-platform-api)

How to use:

```term
$ git clone git@github.com:balansubr/app-setups-go-example.git
...
$ go build setup.go

./setup -archive https://github.com/balansubr/ruby-rails-sample/tarball/master/
--> Created app fierce-reef-7523
----> App ID:42006200-c4ce-415e-8b76-8ce5ed7d960d
----> Setting up config vars and add-ons......Done.
--> Build 4880aded-9ec9-4f4c-8365-9f121830a276 pending.....................................
----> Build succeeded
.........
--> Postdeploy script completed with exit code 0
--> App setup complete.
```

This will use the API key in your `.netrc` file, created with the standard `heroku login` command.


If you want to specify a key explicitly, use the `-apikey` argument:

    ./setup -apikey <api key> -archive https://github.com/balansubr/ruby-rails-sample/tarball/master/


You can get your API key by running `heroku auth:token` or from the [account page on Dashboard](https://dashboard.heroku.com/account).
