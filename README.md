# Notes
Notes is a simple command line tool to manage notes. It allows you to add, delete, list and get notes. It also allows you to export and import notes to and from a file.

I've used to use bash_alias for this, but I've decided to create a simple command line tool to manage notes.

## Installation
To install notes, you can use the following command:
```bash
$ make compile 
$ make install
```

## Usage
To use notes, you can use directly from the command line. Since it will be deployed to /usr/local/bin, you can use it from anywhere in the command line.

The notes will be stored in a folder called .notes in your home directory.

the following commands are available:
```bash
$ notes --help
NAME:
   notes - A new cli application

USAGE:
   notes [global options] command [command options]

COMMANDS:
   add      Adds a note to the list
   delete   Deletes a note from the list
   list     Lists all notes
   get      Gets a note from the list
   export   Exports all notes to a file
   import   imports file with notes exported with export command
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```
