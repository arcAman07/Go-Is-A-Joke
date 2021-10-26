# Go-Is-A-Joke
A simple web server which generates jokes based on the user's choices with the help of an external API. The server is made without using any frameworks ( using net/http base package built over the net/tcp package ). Use of templates , io , fmt and log packages have also been used to create a user interface where user can request the type of joke they 
want to listen based on the various categories , blacklistFlags, etc. The Full Joke API documentation link is given below : 

Use of Joke API - https://jokeapi.dev/

Things to be added : JSON data received is yet to be unmarshalled and needs to be converted into a struct to send only the particular fields of data as a reponse to the server.
Compared to node and express, you can see enormous lines of code has to be written and monitored to get the same exact output.( use of an external API )
