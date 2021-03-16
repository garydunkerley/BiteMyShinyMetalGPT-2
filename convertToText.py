import os
import html2text
import codecs

parent = r'/home/gary/go/src/github.com/garydunkerley/screenplayGen'
folder = '/Futurama'
directory = parent + folder

for root, dirs, files in os.walk(directory):
    for filename in files:
        newName = filename.strip('.html') 
        filepath = parent + folder + "/" + filename

        if (filepath.endswith(".html")):
            
            with open(filepath, 'r') as theFile:
                    contents = theFile.read()
            
            newTextFile = open(parent + folder + "/" + newName + ".txt", "w+")
            newTextFile.write(html2text.html2text(contents))
