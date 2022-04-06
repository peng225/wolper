# wolper

## What is this

This is a helper program for solving [wordle](https://www.powerlanguage.co.uk/wordle/).

## How to build

Just execute the `make` command.

## How to use

### CLI

1. Get English text data to construct the wolper's dictionary.
2. Assuming the input texts are located in `input` directory, you can build the dictionary file named `dict.txt` by the following command.
   ```
   ./wolper build [-i input] [-o dict.txt]
   ```
3. Run the wolper server as follows.
   ```
   ./wolper server [-i dict.txt] [-p TCP port#]
   ```
4. Run the wolper client as follows.
   ```
   ./wolper query [-k key] [-i included characters] [-e excluded characters] [-a IP address] [-p TCP port#] [-u]
   ```
   The sample output is shown below.
   ```
   $ ./wolper query -k fi.er -e powghtl
   address: localhost
   port: 8080
   include: 
   exclude: powghtl
   key: fi.er
   Connection succeeded.
   fiber
   finer
   fixer
   ```

### Web interface

The way to start the wolper server is the same with that of CLI.
Web interface can be used by the following steps.

1. Run the wolper web server as follows.
   ```
   ./wolper web [-a IP address] [-p TCP port#] [--wolper_port TCP port# of wolper server]
   ```
   Note that the IP address here is wolper server's one, not web server's. Also note that the TCP port# is for the web server whcich your are about to start.
2. Access to `http://localhost:"TCP port#"` on your web browser.
