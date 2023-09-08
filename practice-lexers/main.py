def main(expression):
  tokens = identify_tokens(expression)
  
  result = evaluate(tokens)

  print(result)

def identify_tokens(expression):
  symbols = ['(', ')']

  tokens = []
  lexeme = ''
  for index, char in enumerate(expression):
    if index == len(expression) - 1:
      lexeme += char
      tokens.append({'type': 'int', 'value': int(lexeme) })
      break
      
    if char == '+':
      tokens.append({'type': 'int', 'value': int(lexeme) })
      tokens.append({'type': 'plus', 'value': '+' })
      lexeme = ''
    
    if char == '-':
      tokens.append({'type': 'int', 'value': int(lexeme) })
      tokens.append({'type': 'minus', 'value': '-' })
      lexeme = ''
    
    if char == '*':
      tokens.append({'type': 'int', 'value': int(lexeme) })
      tokens.append({'type': 'multiplication', 'value': '*' })
      lexeme = ''
    
    if char == '/':
      tokens.append({'type': 'int', 'value': int(lexeme) })
      tokens.append({'type': 'division', 'value': '*' })
      lexeme = ''
    
    if char.isdigit():
      lexeme += char
      continue

  return tokens

def evaluate(tokens):
  result = 0
  op = None
  for token in tokens:
    if op == None:
      result = token
    if token['type'] == 'plus':
      op = 'plus'
    
    
    # if token['type'] == 'plus':
    #   result = tokens[0]['value'] + tokens[2]['value']
    # elif token['type'] == 'minus':
    #   result = tokens[0]['value'] - tokens[2]['value']
    # elif token['type'] == 'multiplication':
    #   result = tokens[0]['value'] * tokens[2]['value']
    # elif token['type'] == 'division':
    #   result = tokens[0]['value'] / tokens[2]['value']
    # else:
    #   raise Exception('Invalid operation')
  
  return result

main("3+3")