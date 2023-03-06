---
layout: page
title: CWU Fitlab
description: CRUD web application using Streamlit
img: assets/img/fitlab.jpg
importance: 1
category: work
---

## Overview

CWU Fitlab is a full-stack web application created for an upper level practicum for seniors at Central Washington University. CWU FitLab is a comprehensive fitness assessment experience provided by trained CWU students pursuing a degree in Exercise Science. This application makes it easier to manage client scheduling, track client progress over time, and provide access to client assessment data.

[Application](https://cwu-fitlab.streamlit.app/)

<hr>

## Tech Stack

- 🐍 Python 🐍
- HTML5
- CSS
- Google Cloud Platform
- Firebase

<hr>

## Data Science Packages

- 🐼 Pandas 🐼
- NumPy
- Matplotlib
- Seaborn
- Plotly
- Plotly Express

<hr>

## Notable Features

- Dashboard for clients to track progress over time
- Password protected login
- Update Firebase backend from "Technician" tab
- Calendly calendar to schedule visits

<hr>

<div class="row justify-content-sm-center">
    <div class="col-sm-8 mt-3 mt-md-0">
        {% include figure.html path="assets/img/cwu-dash.jpg" title="Dashboard" class="img-fluid rounded z-depth-1" %}
    </div>
    <div class="col-sm-4 mt-3 mt-md-0">
        {% include figure.html path="assets/img/cwu-linefig.jpg" title="Line Figure" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

<hr>

## Code sample

Code sample of a function that fetches client data from the backend and creates a pandas dataframe to display on the technican page.

{% highlight python linenos %}

def fetch_agg_data():
    collection_ref = db.collection("clients")
    subcollections = collection_ref.list_documents()
    client_list = []
    # Create pandas df to export as csv
    df = pd.DataFrame()

    # fetch client ID's and append to dict
    for subcollection in subcollections:
        client_list.append(subcollection.id)

    for c in client_list:
        data = fetch_client_data(c)
        data["Client"] = c
        df = pd.concat([df, data], axis=0)

    df.index = range(0,len(df))

    return df   

{% endhighlight %}
