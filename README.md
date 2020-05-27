# Go unsafe.Pointer vulnerability POCs

This is a series of proof of concept Go programs showing how the use of `unsafe.Pointer` can lead
to different vulnerabilities. There are four examples in total.

These examples accompany the blog post series [Exploitation Exercise with unsafe.Pointer in Go](https://dev.to/jlauinger/exploitation-exercise-with-unsafe-pointer-in-go-information-leak-part-1-1kga).

Each week I will add the code for the next part of the series. So far, there are three parts published:

 1. [Information Leak](https://dev.to/jlauinger/exploitation-exercise-with-unsafe-pointer-in-go-information-leak-part-1-1kga)
 2. [Code Flow Redirection](https://dev.to/jlauinger/exploitation-exercise-with-go-unsafe-pointer-code-flow-redirection-part-2-5hgm)
 3. [ROP and Spawning a Shell](https://dev.to/jlauinger/exploitation-exercise-with-go-unsafe-pointer-rop-and-spawning-a-shell-part-3-4mm7)

This blog posts are written as part of my work on my Master's thesis at the [Software Technology Group](https://www.stg.tu-darmstadt.de/stg/homepage.en.jsp) at TU Darmstadt.

