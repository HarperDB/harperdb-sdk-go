/*
This is a client for the HarperDB database.
It mirrors the HTTP API as of version 2.2.0 and makes it
very easy to get up and running with HarperDB and your Go application.

For more information see: https://docs.harperdb.io/

Basics

Instantiate a new client:

	client := harperdb.NewClient("http://localhost:9925", "username", "password")

*/
package harperdb
