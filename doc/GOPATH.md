## Introduction

I don't know about you. All I know is about me and how I think. Being here from python background when I first went through a tutorial in go and created first hello world program I have a lot of questions running in my mind. So If you are a developer of my mindset I believe this helps to clarify some of your doubts.

Before we dive into learning Go and its syntax, We need to create a basic understanding of where to put our source code.

These questions may bug you 

1. Why we set some export statements when we install go in our machine.
2. why many opensource git codes are having import statements referring to github. 
3. Where to put our source code when we develop our code in local.
4. What is usual project structure

Before I answer all these questions, If you followed along with me in installation. I need you to type down this in terminal.

	go env

This lists out go environment varialbles necessary for us. Among them there are a few you need to look into

* GOPATH
* GOROOT
* GOARCH



_These will be explained in a bit._


## Go Workspace. 

A Go Workspace is how Go manages our source files, compiled binaries, and cached objects used for faster compilation later. 
A workspace is nothing but a directory in your file system. There will be mainly two directories inside it

* src 		Where our project source code resides.
* bin 		Where all binaries the go tool build will reside

Most Go programmers keep all their Go source code and dependencies in a single workspace. __Note that symbolic links should not be used to link files or directories into your workspace.__ 

### GOPATH

The __GOPATH__ environment variable lists places for Go to look for Go Workspaces. By default GOPATH location will be `/home/<user>/go` We need to change to custom location for our project. This can be acheived by 