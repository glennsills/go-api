# go-api
As an exercise I want to recreate the application created by dotnet new webapi

When you run ```dotnet new webapi``` you get a shell app that returns 5 randomly generated weather forecasts records when you call HTTPGET on the the /weatherforecast path.

The first version is working with a couple of small differences

1. This golang app only does http, now https
2. The .NET version supports case-insensitive URLS, the golang version does not
