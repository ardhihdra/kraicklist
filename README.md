# KraickList

## Improvement On Kraicklist (<6 days deadline with currently fulltime jobs 🥲)
see live version -> https://kraicklist-ardhi.herokuapp.com 
1. improve overall UI/UX, (I think it's important, don't cover a book by its cover is hard :smile:), 
2. improve query search prioritize, most match with the title is prioritized. then match with content, then other stuff like thumbnail or any useful information
3. also add some concurrency 👍
4. refactor overall apps structure
5. add some basic css, refactor js, add golang's html/template 
5. add some initial testing 🧪


future works: 
1. work more on search matching, apply more algorithm
2. work on search by image, work on empty result
3. work on peek search result detail before go into the full detail. so user could decide wether to read more or not
4. add last search (some cookies perhaps)
5. deploy on AWS!! (idk why, its just looks cool lol, I mean performance & reliability wont be an issue here, so Heroku is fine~)

## Challange Description

Welcome to Haraj take home challenge!

In this repository you will find simple web app for fictional startup called KraickList. This app will allow users to search ads from given sample data located in `data.gz`.

Currently the app is just rough prototype. The search is case sensitive, limited to exact matches, & the search result is pretty much could be further improved.

You could see the live version of this app [here](https://gentle-forest-97151.herokuapp.com/). Try searching for "iPhone" to see some results.

## Your Mission

Improve the overall app. Think about the problem from user perspective and prioritize your changes according to what you think is most useful.

## Evalution

We will evaluate your submission based on:

1. The approach you are using to identify & solve the problems
2. The quality of your search result
3. The design & testability of your code
4. The method you are using to deploy your app

## Submission

1. Fork this repository and share us the link to your fork after pushing the changes.
2. Host your solution. This project includes Heroku Procfile and in its current state can be deployed easily on free tier. You could also host the app on your own server. Share us the link to your solution.
3. In your submission, share with us what changes you have made and what further changes you would prioritize if you had more time.

## About Sample Data

The data is translated from Arabic to English using Google Translate, so sometimes you will find funny translation on it. 🤣
