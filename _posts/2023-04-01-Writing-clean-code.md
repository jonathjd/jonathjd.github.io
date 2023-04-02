---
layout: distill
title:  Writing clean code
date: 2023-04-01 10:40:16
description: Writing clean code in Python
tags: technical software data code
categories: software-engineering code
giscus_comments: true

authors:
  - name: Jonathan Dickinson
    affiliations:
      name: None

toc:
    # if a section has subsections, you can add them as follows:
    # subsections:
    #   - name: Example Child Subsection 1
    #   - name: Example Child Subsection 2
  - name: Overview
  - name: Code Formatting
    subsections:
    - name: Black
    - name: Vertical Spacing
  - name: Variable Names
  - name: Code Smells
    subsections:
    - name: Fake Code Smells
  - name: Pythonic Code == Clean Code?
  - name: Conclusion/Final Thoughts
---

I've been reading the book *"Beyond The Basic Stuff With Python: Best Practices For Writing Clean Code"* by Al Sweigart to try and write better 🐍 Python 🐍 code. I thought it would be a good idea to give my overall opinion on the book as a resource.

<hr>

## Overview

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/als_page.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

I wasn't really planning on getting this book, but my girlfriend and I were at [Powell's](https://www.powells.com/) bookstore in Portland after a conference looking around. I never plan on getting anything but when your at one of the largest bookstores in the world it's tough to make it out without seeing something that catches your eye.

I initially saw Al Sweigarts first book *Automate The Boring Stuff with Python* which has a nice cover and feng shui to it, but after thumbing through a couple pages I thought the content seemed not quite as advanced as I was hoping. When I put it back I saw *"Beyond The Basic Stuff With Python: Best Practices For Writing Clean Code"* and thought it was great. Some of the topics I knew fairly well, some were new and would help me train my engineering brain, so I ponied up the $34.95 and gave it a go.

Like previous thought articles, I'm not going to write a whole bunch and will likely highlight topics I found insightful as I read.

> **_NOTE:_** One final note before we get into it, all of Al's [books](https://inventwithpython.com) are actually free online under a create commons license. I didn't know that upon purchasing the physical copy, but I have no problem supporting such a giving person. Plus the book is pretty nice.

<hr>

## Code Formatting

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/format.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

Formatting is something I'm always paranoid about, because nobody likes code that looks like a mess. Here are some ways Al recommends you can clean up your formatting.

### Black

Using [Black](https://black.readthedocs.io/en/stable/), an "uncompromising code formatter". Named after a pretty sick quote from Henry Ford when he was asked if he would offer other automobile colors for customers in which he replied,

> "You can have any color you want, as long as it's black" - Henry Ford

Essentially, Black formats your code so that it's consistent and (ideally) readable. You can use Black with the following code:

```
$ python3 -m pip install --user black
```

Then you would run it with

```
black yourScript.py
```

It would be worthwhile to read the documentation, but here are a few examples of what you can expect Black to do to your code.

- Change string literals from using single quotes to double quotes

```python
a = 'Pythagorean Theorem'
a = "Pythagorean Theorem"
```

- Remove horizontal white space

```python
array =         [5,   3, 2,2]
array = [5, 3, 2, 2]
```

- Make your lines of Python code 80 characters long

I may try Black out, but I dont know how useful it would be if I'm developing alone. I can see a use case in a team though.

### Vertical Spacing

Vertical spacing is also something I wondered about as well. Here are a few best practices outlined in Al's book (the title is kinda long so I'm not gonna type out the whole thing). I should also mention that although these are mentioned in Al's book they are also a part of the **Python Enhancement Proposal 8**, or [PEP 8](https://peps.python.org/pep-0008/).

- Separate out functions with two blank lines, classes with two blank lines, and methods within a class with one blank line.

Wrong example:
```python
class ThisExample:
    def fakeMethod1():
        pass
    def fakeMethod2():
        pass
def fakeFunction():
    pass
```

Correct example:
```python
class ThisExample:
    def fakeMethod1():
        pass

    def fakeMethod2():
        pass


def fakeFunction():
    pass
```

I'm happy to admit I was not privy to this standard and this was a really helpful tip that drastically improved my code readability.

<hr>

## Variable Names

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/var_names.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

This section was interesting because I frequently see various naming conventions depening on the code base and language I'm working in. After using Python for a little over a year, I want to get the fundamentals down by taking an introduction to computer science class at a community college one summer. The professor had said that all functions and variables needed to be camelCase and that was industry standard. Needless to say that's not entirely true.

> **_NOTE:_** He said a couple other things that were questionable too that we'll probably get into.

- camelCase and snake_case are both acceptable to use, _as long as its consistent_.
- Classes should be in PascalCase formatting.
- Functions, methods, and variable names should be snake_case.
- Constants should be capitalized SNAKE_CASE
- Functions should begin with the verb of their intended use like parse_file() or something similar.
    - I'm 90% sure I read that last one in the book but just a disclaimer I couldn't find the reference to that tip.


<hr>

## Code Smells

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/code_smell.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

Al defines a code smell as a source pattern that signals a potential/future bug 🐛. I'm going to be completely transparent on this one and say that I am guilty of implementing many of the below code smells and I'll post examples of such transgressions with their proper corrections.

- <span>Duplicate code</span>

If you copy and pasted the code from somewhere else in your code base it's likely the source of a code smell. Duplicate code makes updating sections of the code base a nightmare to work with, and honeslty makes it hard to find potential bugs when they arise.

Here is a code snippet of me plotting VO2max over time in Plotly from a DataFrame
```python
def plot_vo2(df):        

    fig = go.Figure()

    fig.add_trace(go.Scatter(
        x=df['Visit Date'],
        y=df['VO2Max (ml/kg/min)'],
        mode='lines+markers',
        name='Your Visit',
        line=dict(shape='linear', color='#FF3F3F'),
        marker=dict(size=14,
        line=dict(width=1.1)
        )))

    fig.update_traces(
        hovertemplate='Visit Date: %{x} <br>VO2Max: %{y} ml/kg/min', 
        hoverlabel = dict(
            bgcolor = '#ffe9e9', 
            bordercolor='black', 
            font=dict(
                color='black',
                family='Arial',
                size=14
            )

    ))
        
    fig.update_layout(
        xaxis=dict(
                showline=True,
                showgrid=False,
                showticklabels=True,
                linewidth=3,
                ticks='outside',
                tickwidth=2,
                tickcolor='black',
                tickfont=dict(
                    family='Arial',
                    size=12,
                    color='rgb(82, 82, 82)',
                ),
            title='Visit Date'
        ),
        yaxis=dict(
            showgrid=True,
            linewidth=3,
            showline=True,
            showticklabels=True,
            ticks='outside',
            ticklen=4,
            tickwidth=2,
            tickcolor='black',
            tickfont=dict(
                family='Arial',
                size=12,
                color='black'
            ),
            title='VO2Max (ml/kg/min)'
        ),
        autosize=False,
        margin=dict(
            autoexpand=False,
            l=100,
            r=20,
            t=110,
        ),
        plot_bgcolor='white'
        )
    st.plotly_chart(fig, use_container_width=True)
    return

```

The only thing that differed from this plot to the ***6*** other Plotly figures were the columns from the DataFrame that were being plotted and the title of the X and Y axes. Imagine just wanting to change one parameter of the plot.

Below would be a much better implementation.

```python
def plot_variable(df, xaxis, yaxis, title):        

    fig = go.Figure()

    fig.add_trace(go.Scatter(
        x=df[xaxis],
        y=df[yaxis],
        mode='lines+markers',
        name='Your Visit',
        line=dict(shape='linear', color='#FF3F3F'),
        marker=dict(size=14,
        line=dict(width=1.1)
        )))

    fig.update_traces(
        hovertemplate='Visit Date: %{x} <br>{xaxis}: %{y}', 
        hoverlabel = dict(
            bgcolor = '#ffe9e9', 
            bordercolor='black', 
            font=dict(
                color='black',
                family='Arial',
                size=14
            )

    ))
        
    fig.update_layout(
        xaxis=dict(
                showline=True,
                showgrid=False,
                showticklabels=True,
                linewidth=3,
                ticks='outside',
                tickwidth=2,
                tickcolor='black',
                tickfont=dict(
                    family='Arial',
                    size=12,
                    color='rgb(82, 82, 82)',
                ),
            title=f'{xaxis}'
        ),
        yaxis=dict(
            showgrid=True,
            linewidth=3,
            showline=True,
            showticklabels=True,
            ticks='outside',
            ticklen=4,
            tickwidth=2,
            tickcolor='black',
            tickfont=dict(
                family='Arial',
                size=12,
                color='black'
            ),
            title=f'{yaxis}'
        ),
        autosize=False,
        margin=dict(
            autoexpand=False,
            l=100,
            r=20,
            t=110,
        ),
        plot_bgcolor='white'
        )
    st.plotly_chart(fig, use_container_width=True)
    return

```

While still not perfect, now I just have to call this function with the correct parameters and do away with the duplicate code. I could even refactor it further to only have to call this function once in a for loop.

- Print debugging

I definitely still print debug. I also use the debugger in the editor as well on occasion but I just find a well scatter assortment of print statements usually gets the job done fairly well. However, what Al is alluding to is leaving these print statements in the code base after the bug as been fixed. To that I can say I have not done suprisingly.

- Hungarian Notation & Variables with Numeric Suffixes

This one was in the variable names chapter but I feel like I should include it in this section because this is definitely something I learned in that intro CS class as well that is a little dated. Hungarian notation is the practice of putting the data type in the variable name. Numeric suffixes is the practice of putting a number before a redundant variable name.

Example:
```c++
   string str1 = "Serendipity Booksellers";
   string str2 = "Main Menu";
   string str3 = "1. Cashier Module";
   string str4 = "2. Inventory Database Module";
   string str5 = "3. Report Module";
   string str6 = "4. Exit";
   string str7 = "Enter Your Choice: ";
```

As you can see this was something that I was led to believe was a best practice. Thankfully, this code snippet is 2 years old at this point and is no longer a problem.

### Fake Code Smells

This section is particularly interesting because I was led to believe that many of these myths were actually true. It would actually be great to get input on these from other engineers in the comments below.

- Myth: Functions should have only 1 return statement

My professor from that intro CS class said this was a problem with people who learned Python as their first language, and that it was a poor practice. I abided by this for **two years**. This paradigm was completely shattered when I asked my friend who was a senior SWE to help me debug a function I had written and he started sprinkling in return statements left and right. I thought he was just messing with me and I asked him if it was ok to put so many return statements in one function and his reply was, *"Yea. Why not?"*. He was certainly right. Why not?

- Myth: Global variables are bad

This was one I was actually really suprised about. This wasn't something that my old school CS prof had told me, but something that I have heard repeatedly. I've even heard global variables referred to as "polluting the global namespace". Inituitively, it also seems like a good practice to at least minimize the number of global variables in a program, which is an idea that Al supports. But the idea that all global variables are bad allegedly is a code smell myth.

- Myth: Comments are unneccesary

I have gone through many iterations of this one. I began (as we probably all did) with not ever commenting my code at all, to commenting every single line (literally), to minimizing my comments again. The idea that the code should be the comments is something that is almost as pervasive as the global variables idea. I've come to put them in when I know for certain I will not know what a part of my codebase does when I come back. But other than that I kind of cross my fingers and hope that my code is readable when I come back (kind of kidding kind of not).

## Pythonic code == Clean Code?

<div style="width:100%;height:0;padding-bottom:52%;position:relative;"><iframe src="https://giphy.com/embed/coxQHKASG60HrHtvkt" width="100%" height="100%" style="position:absolute" frameBorder="0" class="giphy-embed" allowFullScreen></iframe></div><p><a href="https://giphy.com/gifs/coxQHKASG60HrHtvkt">via GIPHY</a></p>

Pythonic code can look sleek and impressive, and you will often see me say that I want to write more "Pythonic" code. This is true for the most part, but I have come to realize that code that is well written and true to the language is "Pythonic".

**For example:**

If you want to write a for loop in c++ that returns the phrase "Fifth iteration!" on the fifth iteration, it would look like so.

```c++

int main()
{
    for (int i = 0; i <= 4; i++) {
        if (i == 4) {
            cout << i << " Fifth iteration!"; 
        }
    }
    return 0;
}

```

Making this more "Pythonic" is really as simple as changing the i variable to "num" and using an f string.

```python
def fifth_iteration():
    
    for num in range(0, 5):
        if num == 4:
            print(f'{i} Fifth iteration!')
    
    return
```

There's definitely an extent to which more Pythonic just means more complicated.

```python

def fifth_iteration():
    
    print([f'{i} Fifth iteration!' for i in range(5) if i == 4])
    
    return
```

The above code is certainly more "Pythonic", but is it more readable? In my opinion, that's a question that should always be asked.

<hr>

## Conclusion/Final Thoughts

This article was by no means comprehensive. There are many more examples of these sections that Al goes over in great detail and with a good amount of clarity. Overall, I am really enjoying this book. It's easy to read and not dry like most technical books tend to be. 

Would I pay $35 for it? Yes. I don't know if I would pay more than $35, especially given that it's free on the internet, but $35 is definitely a fair price for what I have read so far.

I can't give it a conclusive rating yet because I'm not finished with the entire book (just the code formatting section). But I did want to get this out there to hopefully encourage some discussion, and highlight a solid resource I have been using.

<hr>

If you feel like this helped you out, feel free click the link below and support this post. Here's my cat Vader for encouragement.

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/vader.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

<script type="text/javascript" src="https://cdnjs.buymeacoffee.com/1.0.0/button.prod.min.js" data-name="bmc-button" data-slug="jdickinson" data-color="#5F7FFF" data-emoji=""  data-font="Lato" data-text="Buy me a coffee" data-outline-color="#000000" data-font-color="#ffffff" data-coffee-color="#FFDD00" ></script>

