---
layout: post
title: Single-cell RNAseq pt. 2 Processing raw reads
date: 2024-04-04 08:40:16
description: High level overview of how to process raw reads to a counts matrix
tags: technical data biology
categories: data bioinformatics science
giscus_comments: true
---

# Intro

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/scrna_seq_cartoon.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

Today we are continuing our single-cell RNAseq series with an overview of how we go from raw reads contained in a fastq file to a nice data format called a count matrix.

If you need a brief introduction to fastq files go [here](https://jonathjd.github.io/blog/2024/parsing-fastq/) where we went over the fastq file and wrote a script to parse this file type in Python.

If you haven't seen the first article jump [here](https://jonathjd.github.io/blog/2024/intro-to-scrnaseq/) where you can get an overview of 10x Genomics popular microfluidic-droplet scrna-seq assay.

If you want a brief intro to bioinformatic pipelines and even build one that calculates base quality scores with WDL, Python, and Docker check out this article [here](https://jonathjd.github.io/blog/2024/bioinformatics-tools-wdl/).


Let's get started!

<hr>

# Overview of the processing pipeline

We've just ran our samples and got back some fastq files. How do we get to a count matrix?

The general overview of the bioinformatic pipeline is as follows-

```mermaid
graph TD
    A[QC reads] --> B[Map reads to reference]
    B --> C[Assig reads to genes]
    C --> D[Assign reads to cells (cell barcode demultiplexing)]
    C --> D[Counting unique reads (UMI deduplication)]

```


<hr>

# Microfluidic-droplet based methods (10x Genomics)

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/Flow_Focusing_Device.png" class="img-fluid rounded z-depth-1" %}
       <p class="text-center">Image from https://en.wikipedia.org/wiki/Droplet-based_microfluidics#/media/File:Flow_Focusing_Device.png.</p>
   </div>
</div>

- Pool the cells and sequence.

Because each bead contained unique barcodes, we can pool all the cells back together and trace back the read to the cell it originated from. ✅

Let's move on to sequencing protocols.

<hr>

# Sequencing protocols

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/tagged_molecule.jpg" class="img-fluid rounded z-depth-1" %}
        <p class="text-center">Image from https://assets.ctfassets.net/an68im79xiti/1C16trEdzy1Folq5xbOijE/7e6fb1f504e130bd561d898384da99d9/CG000315_ChromiumNextGEMSingleCell3-_GeneExpression_v3.1_DualIndex__RevB.pdf.</p>
    </div>
</div>

<hr>

# Outro

And that's it! In the next article we are going to dive more deeply into the data and considerations you have to take when processing it. 

Then we are going to go into more downstream analysis.

If you want an in-depth explanation and a walkthrough of a common analysis pipeline, this [course](https://www.singlecellcourse.org/introduction-to-single-cell-rna-seq.html) from the Sanger Institute is a great resource.

I hope you found this useful and see you in the next one!

<hr>

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/vader.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

<script type="text/javascript" src="https://cdnjs.buymeacoffee.com/1.0.0/button.prod.min.js" data-name="bmc-button" data-slug="jdickinson" data-color="#5F7FFF" data-emoji=""  data-font="Lato" data-text="Buy me a coffee" data-outline-color="#000000" data-font-color="#ffffff" data-coffee-color="#FFDD00" ></script>

<hr>
