---
layout: post
title: Get weekly rental listings using Python and the RentCastAPI
date: 2024-03-25 08:40:16
description: Write a script that get's weekly rental listings from the RentCastAPI with Python in less than 100 lines of code
tags: technical data programming
categories: software-engineering code python
giscus_comments: true
---

# Intro

My wife and I are planning on moving at the end of this summer.

We want to be able to jump on rental's quickly because we want to get a good place.

So today we're going to write a script using Python that sends me a text message with new listings each week on Wednesday at 10:00am using the **RentCastAPI 🎉**.

Let's get started!

# Script logic

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/rent_matrix.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

The [RentCast](https://www.rentcast.io/api?gclid=Cj0KCQjwwYSwBhDcARIsAOyL0fgg7S2JxQKu96wgawn-sNpOVFdbuEVEZhQO2Tvvh8MIOKOrQ92gz1UaAk9BEALw_wcB) API is pretty well documented and easy to use to get property listings in your area.

To use it, you just have to sign up for an account and generate an API key in your profile.

You have to enter credit card information but don't worry, there are different tiers depending on how many times you call the API and you get 50 *free* calls every 31 days so we'll aim to be below that number.

The logic of our script will be as follows-

1. Call API to get all listings in area code.
2. Save listings to a Pandas DataFrame.
3. Filter listings that have already been seen.
4. Save listings in a csv file.
5. Send a text message with the new listings.

The **rationale** is as follows-

- My area code does not have a large existing number of listings (~24) and that number will very likely not increase drastically, making it more feasible for me to pull the data and comfortably handle it in memory with a single API call.

- Because there aren't too many listings, I can read and write the entire csv file with each execution of the script.

- Along those same lines, I am willing to live with the overhead of using a Pandas DataFrame. It'll make reading, writing, parsing, and sorting data much faster with a decrease in execution speed that is negligible.

If you live in a more populated region with a higher volume of listings, you may want to make the following **adjustments**-

- Filter what columns you save to your csv file to save space.

- Process the csv file line by line using the `yield` keyword to avoid reading the entire file into memory. Similar to what we did in the [parsing](https://jonathjd.github.io/blog/2024/parsing-fastq/) fastq files project.

- Save the data in a different file format like a [parquet](https://www.databricks.com/glossary/what-is-parquet) file.

Let's get started developing our script!

# Imports and the API Call

We'll start with our imports and API call.

```python
import requests
import pandas as pd
import smtplib
import os
from config import API_KEY, EMAIL, PASSWORD, CARRIERS, NUMBER
```

We're importing `requests` to make our API calls.

`pandas` to store the data in a dataframe, and to read and write to a csv file.

`smtplib` to send text messages.

`os` because we will need to verify the existence of our csv file during the first execution of our script.

`config.py` file for our sensitive imports.

### Config.py files

This script will be dealing with some sensitive information including API keys, app passwords, emails, and phone numbers.

> *If you are going to be saving this script to a remote repository like GitHub, it's crucial that you save this information in a config.py file and add the path to this file to a .gitignore file.*

If you'll only be running this locally it still is good practice to save this information in a config.py file but not as essential.

Your config.py file will look like so-

```python
API_KEY = "<API_KEY>"
CARRIERS = {
    "att": "@mms.att.net",
    "tmobile": "@tmomail.net",
    "verizon": "@vtext.com",
    "sprint": "@messaging.sprintpcs.com",
    "mint": "@tmomail.net",
}

EMAIL = "<EMAIL>"
PASSWORD = "<GOOGLE APP PASSWORD>"
NUMBER = "<PHONE NUMBER>"
```
The carriers is not really sensitive information but doesn't hurt to throw into the `config.py` file.

The email is the email you're going to send messages _from_.

The number is the phone number you will send messages _to_.

If you have a gmail account you'll need to create an App password, which can be done [here](https://accounts.google.com/v3/signin/challenge/pk/presend?TL=AEzbmxweoO9vp-_sP5sLF5QJYo1H4rAkIwxmQREhF7GArceVSQHj-n-mYCpgI9GY&cid=1&continue=https://myaccount.google.com/apppasswords&flowName=GlifWebSignIn&ifkv=ARZ0qKL8ZFbnpz6cXbYMvPNVdNQSWUyeUzs0iWeZZbVl_n5AdQzGXMa2E2ytAtBu0ZnirJcl9DOi9A&rart=ANgoxcct182PUy3TXQrNwS1RK4mjjxYFkxwDeQitdRQ0bKSSvsj8niWFX4b0PEqVvcuTOUS3W-rv5bZBQ7zloTodp9fdZz3mzS1panXUW5TWqKP0xO7qLxY&rpbg=1&sarp=1&scc=1&service=accountsettings&theme=mn&hl=en_US).

I found this script from this medium [article](https://medium.com/testingonprod/how-to-send-text-messages-with-python-for-free-a7c92816e1a4) by David Mentgen!

Next, we'll make our API call to grab our data.

```python
def get_rentals() -> pd.DataFrame:

    url = "https://api.rentcast.io/v1/listings/rental/long-term?zipCode=98926&status=Active&limit=100"

    headers = {"accept": "application/json", "X-Api-Key": API_KEY}

    response = requests.get(url, headers=headers)

    rental_listings_df = pd.DataFrame(response.json())

    return rental_listings_df
```

Here we have the URL which contains the arguments outlining what kind of data we are requesting (Active rental listings from 98926)

If we take a look at the RentCastAPI we can see that the data is formatted as such-

```json
[
  {
    "id": "1000-W-26th-St,-Apt-216,-Austin,-TX-78705",
    "formattedAddress": "1000 W 26th St, Apt 216, Austin, TX 78705",
    "addressLine1": "1000 W 26th St",
    "addressLine2": "Apt 216",
    "city": "Austin",
    "state": "TX",
    "zipCode": "78705",
    "county": "Travis",
    "latitude": 30.291138,
    "longitude": -97.747833,
    "propertyType": "Condo",
    "bedrooms": 1,
    "bathrooms": 1,
    "squareFootage": 445,
    "lotSize": 653,
    "yearBuilt": 1973,
    "status": "Active",
    "price": 1275,
    "listedDate": "2022-11-21T00:00:00.000Z",
    "removedDate": null,
    "createdDate": "2019-11-07T15:13:54.790Z",
    "lastSeenDate": "2023-02-25T00:00:00.000Z",
    "daysOnMarket": 96
  },
  {
    "id": "10015-Lake-Creek-Pkwy,-Austin,-TX-78729",
    "formattedAddress": "10015 Lake Creek Pkwy, Austin, TX 78729",
    "addressLine1": "10015 Lake Creek Pkwy",
  }
]
```

And conveniently Pandas can parse this as is and create a DataFrame which saves us a lot of time. 😎


# Filtering rental listings

Next, we want to filter out the listings that we have already seen before from our list so we only have the new listings.

```python
def filter_listings(df) -> pd.DataFrame:

    historic_listings = pd.read_csv("data/rental/cleaned/rental_listings.csv")

    historic_ids = set(historic_listings["id"].tolist())

    new_rental_listings_df = df.query("id not in @historic_ids")

    if not new_rental_listings_df.empty:

        # sort listings
        new_rental_listings_df = new_rental_listings_df.sort_values(by="lastSeenDate")

        df = pd.concat([new_rental_listings_df, historic_listings])

        df.to_csv("data/rental/cleaned/rental_listings.csv", index=False)

        return new_rental_listings_df

    else:
        return new_rental_listings_df
```

To do this we

- Read the listings that we have saved in a csv file.

- Grab the unique id of all the listed rentals we have already seen.

- Filter our data to only include rentals that are not already in our csv file.

- If our DataFrame is not empty, sort the new rentals by listing data so that the most recently posted are first, concatenate these new listings with the old listings, and write to our csv file.

- Return the csv file.

> You might be thinking "Wait, but we haven't made a csv file with our historic listings, so `historic_listings = pd.read_csv("data/rental/cleaned/rental_listings.csv")` isn't reading anything and will throw an error. We'll address this in the `main()` function's logic!

# Send text message

The last function before the `main()` function will send our text message.

```python
def send_text_message(phone_number, carrier, message) -> None:
    recipient = phone_number + CARRIERS[carrier]
    auth = (EMAIL, PASSWORD)

    server = smtplib.SMTP("smtp.gmail.com", 587)
    server.starttls()
    server.login(auth[0], auth[1])

    server.sendmail(auth[0], recipient, message)
    return
```

This function will create a recipient that looks like **1234567892@@vtext.com** which is the address we are sending the text message to.

Then start a smtp session and log into your email with the provided credentials and sends the message.

This function doesn't return anything so we don't really need the `return` keyword, but again it's good practice just to include it (at least in my opinion).

# Main function

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/main_entrypoint_comic.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

Finally we are at the main logic of our script. Which is simple enough!

```python
def main():
    rental_listings_df = get_rentals()

    csv_path = "data/rental/cleaned/rental_listings.csv"

    if os.path.exists(csv_path):

        new_rental_listings_df = filter_listings(df=rental_listings_df)

        if not new_rental_listings_df.empty:
            top_listings_addresses = new_rental_listings_df["formattedAddress"].tolist()
            rental_prices = new_rental_listings_df["price"].tolist()
            days_on_market = new_rental_listings_df["daysOnMarket"].tolist()
            message = (
                f"Your top listings are- {top_listings_addresses} "
                f"and the rental price is {rental_prices} "
                f"they have been on the market for {days_on_market}"
            )
        else:
            message = "No new listings, will try again next Wednesday."

        send_text_message(phone_number=NUMBER, carrier="mint", message=message)

    else:
        rental_listings_df.to_csv(csv_path, index=False)


if __name__ == "__main__":
    main()
```

- First we make our API call and get the data.

- Then we check to see if the path to our csv file containing our data exists, if not we will create it.

- If it does exists, we will filter our listings for anything that is not currently in our csv file.

- If this new DataFrame is not empty (meaning there are new listings) we will send a message with the address, price, and days on market.

- If the new DataFrame is empty, we'll send a message indicating that there were no new listings.

The only thing left to do is set up a **cronjob** that will execute this script every Wednesday at 10am.

# Cronjob

If you aren't familiar with cronjobs, essentially they tell your computer to execute certain tasks at defined intervals of time. To set one up to execute our script, we open the crontab file from our terminal.

```bash
$ crontab -e
```

Enter in the cron syntax outlining how often the task will be executed along with the path to our script.

```
0 10 * * 3 /path/to/your/script.py
```

I use [crontab.guru](https://crontab.guru/#0_10_*_*_3) to figure out the correct intervals.

The above syntax will execute a script at hour 10 (10am) each day on the third day of the week (Wednesday).

Write out the file by pressing esc followed by :wq!

Ensure that the cronjob is correctly set by entering `crontab -l` and you should see your scheduled cronjob! 


# Outro
And that's it! Now we'll be able to stay on top of rental listings and jump on a property when it comes available. 🥷🏽

For the full script head over to my GitHub repo [here](https://github.com/jonathjd/thirty_projects). 

Project 4 ✅

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/vader.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

<script type="text/javascript" src="https://cdnjs.buymeacoffee.com/1.0.0/button.prod.min.js" data-name="bmc-button" data-slug="jdickinson" data-color="#5F7FFF" data-emoji=""  data-font="Lato" data-text="Buy me a coffee" data-outline-color="#000000" data-font-color="#ffffff" data-coffee-color="#FFDD00" ></script>

<hr>
