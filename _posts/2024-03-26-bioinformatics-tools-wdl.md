---
layout: post
title: Building a bioinformatics pipeline using WDL, Python, and Docker
date: 2024-03-26 08:40:16
description: Build a bioinformatics pipeline with popular tools like WDL and Docker
tags: technical data programming
categories: software-engineering bioinformatics code python
giscus_comments: true
---

# Intro

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/pipelines.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

I remember the first time I heard the term "pipeline" in reference to a data pipeline.

I had no idea what it was and even when I had somebody explain it to me the description sounded vague and I still really didn't understand.

And the truth is, the term "pipeline" _is_ really vague. People use it to mean many different things.

Wrote a script in Python that scrapes job postings, cleans them, and saves them in a csv file?

*Pipeline❗️*

Wrote a program that uses Airflow to ingest web server logs in a Postgres database?

*Pipeline❗️*

Add on additional letters like ETL (extract, transform, load), ELT (extract, load, transform) and BLT (bacon, lettuce, tomato) pipelines and it really starts to get confusing.

Over time, I have come to realize an important fact-

> When somebody mentions a "pipeline", context is everything.

Yes, there are many kinds of data pipelines, but the field/industry that person is in will determine what they are thinking of. For instance, if they are a bioinformatician they are almost certain thinking of tools that were made for processing NGS data such as-

- Snakemake
- Nextflow
- Cromwell

Or they were thinking of languages that are commonly used such as Workflow Description Language (WDL) and Common Workflow Language (CWL).

Today we are going to focus on a workflow _language_ called **Workflow Description Language (WDL)** to get an idea of how WDL can help us begin to build pipelines for bioinfomatic analysis.

# Workflow Description Language (WDL)

**WDL** is defined as an open standard for describing data processing workflows with a human-readable and writeable syntax [1](https://github.com/openwdl/wdl).

When working with NGS data, there are many things you need to do to the data before you can actually analyze it.

For example, if you want to conduct an experiment looking at the differential gene expression between patients with breast cancer vs patients without breast cancer and to do this you want run bulk-RNA sequencing on blood samples between both of your groups.

After you sequence your blood samples, you will get a file with ~1M lines of text that look like-

```
@SRR960459.1 HWI-ST330:304:H045HADXX:1:1101:1162:2055 length=100
NAGAACTTGGCGGCGAATGGGCTGACCGCTTCCTCGTGCTTTACGGTATCGCCGCTCCCGATTCGCAGCGCATCGCCTTCTATCGCCTTCTTGACGAGTT
+SRR960459.1 HWI-ST330:304:H045HADXX:1:1101:1162:2055 length=100
#1=DDFFFHHHGHIJJJJIJJJGEGGAFGBHHEHGFBFFDEDECDDA==CB@BDDDDD?;B-<CBDDD>BBBBDDB5<@DDDCDDB@-9ACDDDDB?B<?
```

Not super conducive to a T-test if you ask me.

This file has to be manipulated programmatically to eventually get data which can then be analyzed statistically.

First we have to trim low quality reads.

Then we have to map the reads back to the genome to see what gene they came from.

Then we have to count the reads for each patient.

How do we keep track of all the steps and run them in the correct order?

That's where WDL comes in ✅

To get an idea of how WDL works, we are going to write a simple pipeline that will parse a fastq file and write the base quality scores to a csv file.

It might be helpful to check out this [article](https://jonathjd.github.io/blog/2024/parsing-fastq/) if you aren't familiar with fastq files or base quality scores. We're also going to improve the script we wrote there as well.

# Setting up our environment

To begin, let's set up our development environment using `pip`. The only dependency we will need is `miniwdl` which will allow us to run our WDL workflow.

```bash
$ python -m venv .wdl_env
$ source .wdl_env/bin/activate  
$ pip install miniwdl
$ pip freeze > requirements.txt  
```
Now, we'll set up our Dockerfile to containerize our pipeline.

# Dockerfile

Going into depth on what Docker does is a little outside the scope of this project, but at a high level Docker will ensure our pipeline is reproducable by specifying the exact software, dependencies, libraries, tools, and their version.

```dockerfile
FROM python:3.11

WORKDIR /app

COPY . /app

RUN pip install --no-cache-dir -r requirements.txt
```
Now we need to create the image the container will be built from.

```bash
$ docker build -t jonathjd/mywdl_image1:v1 .
```

Now let's write our script to parse the fastq file with Python!

# Parsing the file with Python

We are going to use the script we previously created [here](https://jonathjd.github.io/blog/2024/parsing-fastq/), but with a couple of adjustments.

- We are going to parse a compressed fastq file.

**Old logic**
```python
def parse_fastq(input_fastq):
    with open(input_fastq, "r") as input_handle:
        for line in input_handle:
            next(input_handle)
            next(input_handle)
            yield next(input_handle).rstrip()
```

**New Logic**
```python
import sys
import csv
import gzip

def parse_fastq(input_fastq):
    if input_fastq.endswith(".gz"):  # Use gzip.open for gzipped files
        with gzip.open(input_fastq, "rt") as input_handle:
            for line in input_handle:
                next(input_handle)
                next(input_handle)
                yield next(input_handle).rstrip()
    else:
        with open(input_fastq, "r") as input_handle:
            for line in input_handle:
                next(input_handle)
                next(input_handle)
                yield next(input_handle).rstrip()
```

- We are going to write out the base quality scores to a **csv file** instead of printing them to the screen.

**Old logic**
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
```

**New Logic**
```python
def write_base_qualities_to_csv(average_quality, csv_output):

    with open(csv_output, mode="w") as output_file:
        writer = csv.writer(output_file)
        writer.writerow(["Base Position", "Average Quality Score"])
        for i, score in enumerate(average_quality, start=1):
            writer.writerow([i, score])
    return


def main(fastq_file, csv_output):
    sum_scores = []
    record_count = 0

    for quality_scores in parse_fastq(fastq_file):
        phred_scores = calculate_phred_scores(quality_scores)

        if len(phred_scores) > len(sum_scores):
            sum_scores += [0] * (len(phred_scores) - len(sum_scores))

        for pos, score in enumerate(phred_scores):
            if pos < len(sum_scores):
                sum_scores[pos] += score

        record_count += 1

    average_scores = [sum_score / record_count for sum_score in sum_scores]

    write_base_qualities_to_csv(
        average_quality=average_scores, csv_output=csv_output
    )
```


- We are going to accept the path to the compressed fastq file and the csv file as input arguments.

**Old logic**
```python
if __name__ == "__main__":
    main()
```

**New Logic**
```python
if __name__ == "__main__":
    fastq_file, csv_output = sys.argv[1], sys.argv[2]
    main(fastq_file=fastq_file, csv_output=csv_output)
```

And that's all! All that is left is to write the logic of our WDL script!

# WDL Script

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/workflow_bi.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

At a very high level, WDL scripts are composed of **tasks** and **workflows**.

I imagine tasks are like _functions_, that execute/do certain things, and workflows are like the _main program logic_ that executes the functions in a specific order.

Let's build the WDL script piece by piece.

First specify the WDL version.

```wdl
version 1.0
```

Next we specify our (only) task.

```wdl
task CalculateBaseQuality {

}
```

This task will have arguments that it takes, along with commands that it passes these arguments into.

```wdl
task CalculateBaseQuality {
    input {
        File fastq_input
        String csv_output
    }

    command {
        python /app/parse_fastq.py ~{fastq_input} ~{csv_output}
    }
}
```

This task also specifies the Docker image that it will execute these tasks in, as well as outputs it expects.

```wdl
task CalculateBaseQuality {
    input {
        File fastq_input
        String csv_output
    }

    command {
        python /app/parse_fastq.py ~{fastq_input} ~{csv_output}
    }

    runtime {
        docker: "jonathjd/mywdl_image1:v1"
    }

    output {
        File output_csv = csv_output
    }
}
```

Finally, we write our workflow which has similar components to the task we just wrote. First we specify the inputs.

```wdl
workflow CalculateBaseQualityWorkflow {
    input {
        File input_fastq
        String output_name = "base_quality.csv"
    }
}
```

Then we call our task that we want to execute (function) and specify the input that will be passed into it.

```wdl
workflow CalculateBaseQualityWorkflow {
    input {
        File input_fastq
        String output_name = "base_quality.csv"
    }

    call CalculateBaseQuality {
        input:
            fastq_input = input_fastq,
            csv_output = output_name
    }
}
```

Finally, we specify the output.

```wdl
workflow CalculateBaseQualityWorkflow {
    input {
        File input_fastq
        String output_name = "base_quality.csv"
    }

    call CalculateBaseQuality {
        input:
            fastq_input = input_fastq,
            csv_output = output_name
    }

    output {
        File base_quality_csv = CalculateBaseQuality.output_csv
    }
}
```

That's it! Now before we run our workflow, we can check our wdl script for syntax errors using-

```bash
$ miniwdl check parse_fastq.wdl
```

If we wanted to see what inputs we need to run our script we can see that with-

```bash
$ miniwdl run parse_fastq.wdl    
```

To run our workflow, we run the run `miniwdl run` command along with the input arguments separated with back slashes.

```bash
$ miniwdl run parse_fastq.wdl \                            
    input_fastq=demo.fastq.gz \
    output_name=base_quality.csv \
    --verbose 
```

Now you should see your base quality scores output to the out folder at this path `./_LAST/out/base_quality_csv/base_quality.csv`.

# Outro

And that's it! We just wrote our first bioinformatics pipeline. Hopefully we can build on this in projects to come.

For the full script head over to my GitHub repo [here](https://github.com/jonathjd/thirty_projects). 

Project 5 ✅

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/vader.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

<script type="text/javascript" src="https://cdnjs.buymeacoffee.com/1.0.0/button.prod.min.js" data-name="bmc-button" data-slug="jdickinson" data-color="#5F7FFF" data-emoji=""  data-font="Lato" data-text="Buy me a coffee" data-outline-color="#000000" data-font-color="#ffffff" data-coffee-color="#FFDD00" ></script>

<hr>
