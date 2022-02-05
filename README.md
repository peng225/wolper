# wolper

## What is this

This is a helper program for solving [wordle](https://www.powerlanguage.co.uk/wordle/).

## How to use

1. Get English text data to construct the wolper's dictionary.
2. Assuming the input texts are located in `input` directory, you can build the dictionary file named `dict.txt` by the following command.
   ```
   go run main.go build -i input -o dict.txt
   ```
3. Run the wolper server as follows.
   ```
   go run main.go server -i dict.txt -p [TCP port#]
   ```
4. Run the wolper client as follows.
   ```
   go run main.go query -k [key] -i [included characters] -e [excluded characters] -a [IP address] -p [TCP port#]
   ```
   The sample output is shown below.
   ```
   $ go run main.go query -k fi.er -e powghtl
   address: localhost
   port: 8080
   include: 
   exclude: powghtl
   key: fi.er
   service.ClientQeury called (addrAndPort = localhost:8080).
   Connection succeeded.
   fixer
   finer
   fiber
   ```