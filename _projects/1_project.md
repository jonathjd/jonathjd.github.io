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

## Tech Stack

- 🐍 Python 🐍
- HTML5
- CSS
- Google Cloud Platform
- Firebase

## Data Science Packages

- 🐼 Pandas 🐼
- NumPy
- Matplotlib
- Seaborn
- Plotly
- Plotly Express

## Notable Features

- Dashboard for clients to track progress over time
- Password protected login
- Update Firebase backend from "Technician" tab
- Calendly calendar to schedule visits


<div class="row justify-content-sm-center">
    <div class="col-sm-8 mt-3 mt-md-0">
        {% include figure.html path="assets/img/cwu-dash.jpg" title="Dashboard" class="img-fluid rounded z-depth-1" %}
    </div>
    <div class="col-sm-4 mt-3 mt-md-0">
        {% include figure.html path="assets/img/cwu-linefig.jpg" title="Line Figure" class="img-fluid rounded z-depth-1" %}
    </div>
</div>


The code is simple.
Just wrap your images with `<div class="col-sm">` and place them inside `<div class="row">` (read more about the <a href="https://getbootstrap.com/docs/4.4/layout/grid/">Bootstrap Grid</a> system).
To make images responsive, add `img-fluid` class to each; for rounded corners and shadows use `rounded` and `z-depth-1` classes.
Here's the code for the last row of images above:

{% raw %}
```html
<div class="row justify-content-sm-center">
    <div class="col-sm-8 mt-3 mt-md-0">
        {% include figure.html path="assets/img/6.jpg" title="example image" class="img-fluid rounded z-depth-1" %}
    </div>
    <div class="col-sm-4 mt-3 mt-md-0">
        {% include figure.html path="assets/img/11.jpg" title="example image" class="img-fluid rounded z-depth-1" %}
    </div>
</div>
```
{% endraw %}
