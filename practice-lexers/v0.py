#phrase = "deploy falhou hoje"
phrase = "the beautiful white moon"
separator = ' '
list_chars = list(phrase)

lexeme = ''
for i,v in enumerate(list_chars):
  lexeme = lexeme + list_chars[i]

  if list_chars[i] == separator:
    print("Lexeme found: ", lexeme)
    lexeme = ''
    
  if i + 1 >= len(list_chars):
    print("Lexeme found: ", lexeme)
    break