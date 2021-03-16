from bs4 import BeautifulSoup
import os

parent = r'/home/gary/go/src/github.com/garydunkerley/screenplayGen'
folder = '/Futurama'
directory = parent + folder

for root, dirs, files in os.walk(directory):
    for filename in files:
        newName = filename.strip('.html') 
        filepath = parent + folder + "/" + filename

        if (filepath.endswith(".html")):
            
            theFile = open(filepath, 'r')
            soup = BeautifulSoup(theFile)
            newTextFile = open(parent + folder + "/" + newName + ".txt", "w+")
            newTextFile.write(soup.get_text())
