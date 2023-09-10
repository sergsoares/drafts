# Write an interpreter as described in this article off the top of your head, 
# without peeking into the code from the article. 
# Write some tests for your interpreter, and make sure they pass.

def main():
  expression = "2 + 7 * 4"

  # tokens = lexer()
  # Lexer
  # Ignore spaces
  # operators
  # precende of operator

  tokens = []
  result = interpreter(tokens)

  if result == 30:
    print("> OK")
  else:
    print("> FAIL")




def interpreter(tokens):
  tokens = [
    ("INTEGER",3),
    ("PLUS", '+'),
    ("INTEGER",5),
    ("PLUS", '+'),
    ("INTEGER",2),
  ]

  result = tokens[0][1]
  for each in tokens:
    if each[0][0] == 'MINUS':
      result -= each[0][1]
    elif each[0][0] == 'PLUS':
      result += each[0][1]


    print(result)
    print(each)





















main()