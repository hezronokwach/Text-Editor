# Text Transformation Program
  

This program reads text from a file, applies various transformations specified within the text, and writes the transformed text to another file.
  

### Directory Structure

  

The directory structure of this project is organized as follows:

  

-  `main_function`: Contains the main Go file `main.go`, which is responsible for executing the text transformation process.

- Go files implementing helper functions for text transformation. These files are imported as packages in `main.go`.

- Test files with the extension `_test.go` for each function in the project. These test files ensure the correctness of the implemented functions.

  

## Features

  

-  **Transformation Patterns**: The program recognizes specific patterns within the text to perform transformations. These patterns include:

-  `(up)`: Convert the previous word to uppercase.

-  `(low)`: Convert the previous word to lowercase.

-  `(cap)`: Capitalize the previous word.

-  `(hex)`: Convert the previous word from hexadecimal to integer.

-  `(bin)`: Convert the previous word to binary.

-  `(up,`, `(low,`, `(cap,`: Perform transformations with a following number. For example, `(up,3)` would convert the previous 3 words to uppercase.

-  **Vowel Addition**: After applying transformations, the program adds vowels to the words to ensure grammatical correctness.

-  **Punctuation**: Finally, the program punctuates the words based on certain rules.


### How to Run
  

To run the program, execute the following command in the terminal:
  

go run . sample.txt result.txt

` sample.txt` is the input file containing the text to be transformed and `result.txt` is the output file containing the text transformed
  

### Example

Suppose the input file `sample.txt` contains the following text:
  
~~~
I am exactly how they describe me am me (cap, 1) : ' awesome '
~~~

After running the program, the text will be transformed as follows:
~~~  

I am exactly how they describe me am Me: 'awesome'
~~~

And the transformed text will be saved to the `result.txt` file.
  

### Dependencies
  

The program uses the `reloaded` package for various text transformations. Ensure that the `reloaded` package is available in your Go environment.