---
layout: post
title:  Parsing fastq files with python
date: 2024-03-21 10:40:16
description: Parsing fastq files with using python and no libraries
tags: technical software data bioinformatics
categories: bioinformatics code
giscus_comments: true
---

## Intro

I hate to say it but-

My skills are atrophying a bit.

Truth be told nothing terrible. I haven't forgotten how to set up a conda environment, or manipulate a pandas dataframe, but the edges of my skillset have gotten oh so rusty which really makes learning difficult.

Because those edges are where we are going to grow and get better! To fix that I am embarking on 30 days of projects to help round out those edges. I am going to start fairly simply and the projects will gradually get more difficult. I am going to try to keep them to completing in a day, but as the days go on I imagine they will get more difficult and take longer.

The first project is going to be everyone's favorite-

**File Parsing 🎉**

Obviously this isn't revolutionary or very flashy, but it hits on many of the basics and then some due to the large file type we will parse, which is a Fastq file.

## What is a Fastq file?

A fastq file is a text file that contains sequence information that comes from a sequencer (think Illumina, but there are others). In short, when you sequence samples you are essentially trying to read the bases from a strand of DNA. The output is a text file which contains 4 lines-

1. A sequence identifier which gives you information about the sequencing run, such as where the cluster was on the plate.
2. The actual sequence (A, T, C, G, N)
3. A seperator which is a '+' sign.
4. A quality score for each base.

An example of record in a fastq file might look like such-

```
@SRR960459.1 HWI-ST330:304:H045HADXX:1:1101:1162:2055 length=100
NAGAACTTGGCGGCGAATGGGCTGACCGCTTCCTCGTGCTTTACGGTATCGCCGCTCCCGATTCGCAGCGCATCGCCTTCTATCGCCTTCTTGACGAGTT
+SRR960459.1 HWI-ST330:304:H045HADXX:1:1101:1162:2055 length=100
#1=DDFFFHHHGHIJJJJIJJJGEGGAFGBHHEHGFBFFDEDECDDA==CB@BDDDDD?;B-<CBDDD>BBBBDDB5<@DDDCDDB@-9ACDDDDB?B<?
```

Going into each detail of a record is a task for another post. For now we are going to focus on one aspect, the quality score or ***Phred score***.

## Phred Scores
The phred score which is the 4th line in each record contains the likelihood the sequencer made an incorrect base call. This confidence is a probability.

> "But Jon, how is "#" a probability?"

That's a great question! The phred score is encoded using ASCII character's, meaning that the ASCII value that is assigned to each character - 33 is the actual score. Here is the equation-

$$S_{\text{phred}} = \text{ord}(Q_{\text{char}}) - 33$$

Ord is the python function that will return the ASCII character's value which can also be found at this [chart](https://www.ascii-code.com/). The values start at 0, and the first character that can be used in our string is the bang operator which has a value of 33 and goes up from there (" is 34, # is 35 and so on). That's also why we subtract 33 from each character.

Once we get this score we calculate the actual probability using the following equation-

$$P = 10^{\frac{-Q}{10}}$$

That's enough for the background, let's start parsing some files.

## Parsing files

You can download an example fastq file [here](https://zenodo.org/records/3736457).

The first the we are going to do is write a function that will open our file and get the line that contains the quality scores.
```python
def parse_fastq(input_fastq):
    with open(input_fastq, "r") as input_handle:
        for line in input_handle:
            next(input_handle)
            next(input_handle)
            yield next(input_handle).rstrip()
```

The *next()* function will iterate to the next line, so essentially we are saying next(identifier) -> next(sequence) -> next(separater (+)) and end on the quality score.

The yield generator returns a generator object which can be iterated over in a for loop. We do this because there will be millions of sequences which we won't be able to hold in memory.

Next, we create a function that will calculate the quality score from each individual character.
```python
def calculate_phred_scores(quality_score):
    return [ord(char) - 33 for char in quality_score]
```
This function uses list comprehension to calculate the ASCII value for each character in the line and create a list.

Finally, we will define a *main()* function that will contain the bulk of our logic.

```python
def main():
    sum_scores = []
    record_count = 0

    for quality_scores in parse_fastq("data/fastq_files/demo.fastq"):
        phred_scores = calculate_phred_scores(quality_scores)

        if len(phred_scores) > len(sum_scores):
            sum_scores += [0] * (len(phred_scores) - len(sum_scores))

        for pos, score in enumerate(phred_scores):
            if pos < len(sum_scores):
                sum_scores[pos] += score

        record_count += 1

    if record_count > 0:
        average_scores = [sum_score / record_count for sum_score in sum_scores]
        for pos, avg_score in enumerate(average_scores):
            print(f"Position {pos+1}: Average Score = {avg_score:.2f}")

    else:
        print("No record found")


if __name__ == "__main__":
    main()
```

First, we initialize a *sum_scores* variable which will hold the sum of the scores for each base. Remember, we want the average quality score for each individual *base*, not the average score for each individual *sequence*.

Then, we iterate through each quality score and store the result of calculate_phred_scores, which converts the ASCII-encoded quality scores into numerical Phred scores.

We then check to see if the list of phred scores is longer than the sum_scores list and extend it with zeroes if that is the case.

Finally, we divide each score by the number of records to get the average phred score for each base and print it out along with the base position.

> It's crucial to note that this approach assumes that the fastq file is following the aforementioned format. If the file deviates from this format we'll get errors.

## Outro

And that's it! We've just parsed our first fastq file. It's worth noting that this can be improved upon by making class based tools and not a script, plotting the scores, and we didn't do anything with the actual sequences, but it is a start. We will build on this in projects to come.

Project 1 ✅

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/vader.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

<script type="text/javascript" src="https://cdnjs.buymeacoffee.com/1.0.0/button.prod.min.js" data-name="bmc-button" data-slug="jdickinson" data-color="#5F7FFF" data-emoji=""  data-font="Lato" data-text="Buy me a coffee" data-outline-color="#000000" data-font-color="#ffffff" data-coffee-color="#FFDD00" ></script>

<hr>
