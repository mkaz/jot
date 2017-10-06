#!/usr/bin/env python

# Jot - a simple to tool to jot notes to a weekly file
#

import argparse
import os
import sys
from datetime import datetime

content = ""
today = datetime.today()
date_format = "%A, %Y-%m-%d %I:%M%p"

parser = argparse.ArgumentParser(description='jot.py - cmdline notes')
parser.add_argument("--filename" )
parser.add_argument("--debug", action="store_true", default=False)
parser.add_argument("--week", action="store_true", default=False)
parser.add_argument("--today", action="store_true", default=False)
parser.add_argument("--date", nargs=1, required=False)
parser.add_argument("--last", nargs='?', required=False)
parser.add_argument("text", nargs='?')
args = parser.parse_args()

if args.text != None:
    content = args.text

### Setup Path & Filename ###############################################################
"""
	1. Check command-line parameter --filename
	2. Check environment variable JOT_FILE
	3. Check environment variable JOT_DIR
	4. Use default ~/Documents/jots/jot-2013-w32.md
"""
if (args.filename):
    filename = args.filename
elif hasattr(os.environ, "JOT_FILE"):
	filename = os.environ["JOT_FILE"]
elif hasattr(os.environ, "JOT_DIR"):
	filename = os.environ["JOT_DIR"] + today.strftime("jot-%Y-w%U.md");
else:
    filename = os.environ["HOME"] + "/Documents/jots/" + today.strftime("jot-%Y-w%U.md");

if args.debug: print("Filename: " , filename)

# open file for appending
if (not os.path.exists(filename)):
    new_file = True
else:
    new_file = False

### End Setup   #########################################################################

def cat_file():
    f = open(filename, 'r')
    data = f.read()
    f.close()
    print(data)

def show_date( day ):
    found_day = False
    num_found = 0
    with open(filename, 'r') as f:
        for line in f:
            line = line.strip()
            if line.startswith('###'):
                datestr = line[4:]
                dt = datetime.strptime(datestr, date_format)
                if dt.date() == day.date():
                    found_day = True
                elif found_day: # date not the same, and already found, so we're done
                    return
            if found_day:
                print(line)


def write_text( content ):
    f = open(filename, 'a')
    if new_file:
        f.write("------------------------------------------------\n");
        f.write("Notes for " + today.strftime("%b %Y -- Week %U") + "\n")
        f.write("------------------------------------------------\n\n");

    # check what the input is coming from args or STDIN
    if len( content ) < 1:
        if os.isatty( sys.stdin.fileno() ):
            print("<< Enter Text to Jot (ctrl-d to end) >>")
        lines = sys.stdin.readlines()
        content = ''.join(lines)

    # clean-parse text
    content.rstrip()
    
    ## append text to file
    if len(content) > 0:          
        f.write("\n\n### {0}\n".format(today.strftime(date_format)));
        f.write(content)

    else:
        print("Nothing to write");
    
    # close up file
    f.close()

if args.week:
    print()
    cat_file()
elif args.today:
    print()
    show_date(today)
elif len(args.date) == 1:
    print()
    dt = datetime.strptime( args.date[0], "%Y-%m-%d")
    show_date(dt)
else:
    write_text(content)
