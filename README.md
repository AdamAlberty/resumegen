# resumegen

Simple and minimalist CLI tool for generating resumés. It takes resumé data in JSON format as input and outputs PDF.

Best practices for creating resumés were used.

- [Resume Hacking](https://youtu.be/VluPQ6rbb1A?si=10KWqzZnDKir4s6o)
- [Reddit Post by ex-recruiter on how to structure your resumé](https://www.reddit.com/r/jobs/comments/7y8k6p/im_an_exrecruiter_for_some_of_the_top_companies/)

## Installation

Make sure you have **Go compiler** installed on your computer

To build the program, run `go build -o ./resumegen ./cmd/run`

This will build the binary to `./resumegen`

> [!NOTE]  
> For better experience, you can add the program to path. (e.g. moving the program to `/usr/bin` on linux)

## Usage

To find usage of different commands run `resumegen -help`

To generate Skeleton resumé data, run `resumegen generate`

To specify different path to input data, run `resumegen -input /my/custom/path/to/resume_data.json`

To specify different path to output resumé PDF, run `resumegen -output /my/custom/path/to/resume.pdf`
