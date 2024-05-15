---
layout: post
title: Using extracellular vesicles to detect cancer
date: 2026-04-15 08:40:16
description: Taking a look at Mercy Halo, an exciting new assay
tags: technical data biology
categories: data biology science
giscus_comments: true
---

# Intro

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/scrna_pt2_banner.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

Over these last few months I have been exploring new technologies analysis approaches that are emerging in biotech and today I wanted to highlight a newer technology in the cancer diagnostic space-

**Mercy Halo 🎉**

The Mercy Halo is the flagship product from a biotech startup based in Massachusetts called **Mercy BioAnalytics**.

Through Mercy Halo, this quickly growing startup is leveraging the biological insights from _extracellular vesicles_, or *"EV's"* to detect early stage cancer.

Today, I want to outline this innovative approach to a long standing problem and highlight the cool technology that's going into this platform.

<hr>

# Quick review: Extracellular vesicles (EVs)

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/scrna_diagram.png" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

The first question we should address is-

> What are extracellular vesicles, or EV's?

The textbook definition of EV's is <span>that they are cell-derived membrane-surrounded vesicles that carry bioactive molecules and deliver them to recipient cells</span>(1).

I usually envision EV's like small cells that are meant to transport molecules through the extracellular space. Similar to somatic cells that have a lipid bilayer that's meant to regulate the cell like the osmotic pressure and to keep out larger molecules that should not be passing in and out. But unlike somatic cells these EVs don't have any organelles, they don't replicate, and they don't generate energy. They are meant to transport various molecules to surrounding cells.

The molecules contained within these EVs could be proteins, RNA, DNA, histones, and even mitochondrial DNA. All of which can be delivered to a neighboring cell who would recieve such molecules.

This could lead us to another question worth asking-

> How are EV's made?

Another hallmark of EVs is that they are usually created _from_ somatic cells. There are a few ways an EV can be created but one way is through a process known as **exocytosis**.

## Exocytosis

Let's say you have a cell that wants to get rid of some denatured proteins or RNA fragments in the cytosol of the cell.

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/exo_schem_1.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

The cell can do a few things with these free floating molecules. It can break them down further and recycle the constituent parts (amino acids, nucleotides). But it can also begin to form an **endosome** through a process called **endocytosis**.

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/exo_schem_2.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

This endosome can move to envelope these misfolded proteins and RNA fragments and create **intraluminal vesicles (ILVs)** that are contained within the endosome, which is now referred to as a **multivesicular endosome (MVEs)**

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/exo_schem_3.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

This MVE can now secrete the ILVs containing the misfolded proteins and RNA fragments into the surrounding extracellular matrix.

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/exo_schem_4.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

Now this is just one way to generate one kind of EV called an exosome, but there are many other kinds of EVs like:

- Autophagic EVs
- Stressed EVs
- Matrix vesicles

All of which may contain slightly different cargo and and may be secreted by slightly different methods but all fall under the umbrella as an EV.

Apart from the cargo that is actually contained within the EV, there are cell-surface receptors and other membrane bound proteins such as tetraspanins that exist on the membrane. These function to bind to various parts of the extracellular matrix as well as communicate with surrounding cells. Both of which would be very useful properties for a cancer cell to propogate.

<div class="row mt-3">
   <div class="col-sm mt-3 mt-md-0">
       {% include figure.html path="assets/img/exo_schem_5.jpg" class="img-fluid rounded z-depth-1" %}
   </div>
</div>

<hr>

# Cancer cell-derived EVs

There are a few disconcerting trends that have appeared in the literature concerning EVs

1. They often carry biomolecules that are pathophysiological in nature (1).
2. Cancer cells often secrete EVs at a higher rate than normal cells (2).
3. Cancer cells secrete EVs at a higher rate in late stage cancer cells than early stage (2).

This makes sense on a couple of different fronts. In particular, when cells degrade aberrhant proteins or other molecules there normally is a predefined pathway to carry this out like in ubqiunation, cell autophagy, and mitophagy. When cells begin to dump large amounts molecules into the cytosol or extracellular matrix en masse this can often indicate that something is wrong, and the cell is about to "call-it-quits".

EVs that are secreted from cancer cells often have a few different properties that make them particularly nefarious.

- Promote tumor angiogenesis, extravasation, and intravasation.
- Promote the differentiation of functions of immunosuppressive cells.
- Promote the differentiation of different cell types into pro-tumorigenic, immunosuppressive, anti-inflammatory, and chemoresistant phenotypes
- Induce apoptosis.
- Disable natural killer cells.

All of which serve to foster an microenvironment that is killing or converting healthy cells to cancer cells. Or serve to foster a microenvironment that is resistant to the natural immune response and chemotherapy.

<hr>

# Early cancer detection

As one might be able to see, the environment that cancer cells foster often have a certain biological signature and this means they can be detected. This is the aim of the detection platform **Mercy Halo**.

This assay's workflow has a few different steps that leverage the specific phenotype of cancer EVs, the first of which is **Size Exclusion Chromatography (SEC)**.

## SEC
## 
##
##
##

# Step 1: QC reads


<hr>

# Step 2: Map reads to reference


<hr>

# Step 3: Map reads to cells


<hr>

# Step 4: Count unique reads

<hr>

# Outro

<hr>

<div class="row mt-3">
    <div class="col-sm mt-3 mt-md-0">
        {% include figure.html path="assets/img/vader.jpg" class="img-fluid rounded z-depth-1" %}
    </div>
</div>

<script type="text/javascript" src="https://cdnjs.buymeacoffee.com/1.0.0/button.prod.min.js" data-name="bmc-button" data-slug="jdickinson" data-color="#5F7FFF" data-emoji=""  data-font="Lato" data-text="Buy me a coffee" data-outline-color="#000000" data-font-color="#ffffff" data-coffee-color="#FFDD00" ></script>

<hr>
