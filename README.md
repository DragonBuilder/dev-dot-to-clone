# dev.to clone
An attempt at cloning of the tech blogging site [dev.to](https://dev.to/) for educational purposes.

## Introduction

Let's build clone of a blogging app. I'll be attempting to build a clone of [dev.to](https://dev.to/).

## Tech Stack
- Go language
- Firestore

## v1
This will be treated as a first cut, and I will try to reach v1 by implementing each feature and tagging that with a version number.

### Features
- Users can create blogs, known as bloggers.
- User can view blogs created by bloggers.
- Bloggers need to register/login their account before posting a blog.

### Implementation of features

#### Bloggers can create blog (v0.1)

- We will need to create a portal that will be used by bloggers to create and then subsequently publish their blogs.
- We will adhere to KISS (Keep it simple, silly) and not worry about editing or even different kinds for media, like pics or rich text formatting and such. We will defer that to a future version.
- The blogger can only write a Title and Content.
- The content can be written in Markdown format. 
- The blogger can publish their blog by clicking on a Publish button.
- The blog creating page will be hosted at /create endpoint.
- It will have two sections - Title and Content.
- It will also have a Publish button.
- Once the publish button is pressed, the blog will be available at a unique URL which will be dynamically generated.