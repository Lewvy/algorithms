# Princeton's Algorithms, Part I - Implementations in Go

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This repository contains my personal implementations and notes for the concepts covered in the Coursera course **[Algorithms, Part I](https://www.coursera.org/learn/algorithms-part1)** by Robert Sedgewick and Kevin Wayne from Princeton University.

---

## My Learning Strategy

My goal with this course is not just to understand the theoretical concepts but to build a deep, practical mastery of them. To achieve this, I've adopted a two-language approach:

* **Go (`.go` files):** All data structures and algorithms presented in the lectures are implemented from scratch in Go. This serves to strengthen my backend development skills and understanding of Go's idioms, memory management, and concurrency patterns.
* **C++:** Associated practice problems, particularly from platforms like LeetCode, are solved using C++. This is my language of choice for competitive programming and for excelling in technical assessments where Go may not be an option.

---

## Repository Structure

The repository is organized by week, mirroring the structure of the course. Each folder contains the Go source code for the data structures and algorithms, along with any relevant notes.

.
├── week1-union-find
│   ├── notes.md
│   └── union_find.go
├── week2-stacks-queues-and-sorts
│   ├── notes.md
│   ├── stack.go
│   ├── queue.go
│   └── elementary_sorts.go
├── week3-mergesort-and-quicksort
│   ├── ...
├── week4-priority-queues-and-heaps
│   ├── ...
├── week5-symbol-tables-and-bsts
│   ├── ...
└── week6-balanced-trees-and-hashtables
├── ...

---

## Topics Covered in Part I

This part of the course provides a comprehensive introduction to the fundamentals of data structures and algorithms.

* **Week 1: Basic Programming Model & Union-Find**
    * Analysis of Algorithms (Asymptotic Notations)
    * Weighted Quick-Union with Path Compression
* **Week 2: Stacks, Queues & Elementary Sorts**
    * Linked Lists, Resizing Arrays
    * Selection Sort, Insertion Sort, Shellsort
* **Week 3: Mergesort & Quicksort**
    * Top-down Mergesort, Bottom-up Mergesort
    * Partitioning, 3-way Quicksort
* **Week 4: Priority Queues & Heaps**
    * Binary Heaps
    * Heapsort
* **Week 5: Symbol Tables & Binary Search Trees**
    * Elementary Symbol Table implementations
    * Core Binary Search Tree operations
* **Week 6: Balanced Search Trees & Hash Tables**
    * Red-Black BSTs
    * Separate Chaining and Linear Probing Hash Tables

---

## Tech Stack

* **Primary Implementation Language:** **Go**
* **Interview Practice Language:** **C++**
* **Testing:** Go's built-in testing package (`go test`)

---

## Acknowledgments

A huge thank you to Professors **Robert Sedgewick** and **Kevin Wayne** for creating and generously sharing this outstanding and foundational course material.

