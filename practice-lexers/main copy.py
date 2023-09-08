INTEGER, PLUS, EOF, MINUS = 'INTEGER', 'PLUS', 'EOF', 'MINUS'

class Token(object):
    def __init__(self, type, value):
        self.type = type
        self.value = value

    def __str__(self):
        """String representation of the class instance.

        Examples:
            Token(INTEGER, 3)
            Token(PLUS '+')
        """
        return 'Token({type}, {value})'.format(
            type=self.type,
            value=repr(self.value)
        )

    def __repr__(self):
        return self.__str__()


class Interpreter(object):
    def __init__(self, text):
        self.text = text
        self.pos = 0
        self.current_token = None
        self.lexeme = ''

    def error(self):
        raise Exception('Error parsing input')

    def get_next_token(self):
        while True:
            # set variable locally
            text = self.text
            # text = text.replace(" ", "")
            # text = self.text.replace(" ", "")
            # text = "".join(text.split())
            print(text)
    

            # Identify if we are inside the END of file, if actual position
            # is bigger than size of expression - 1 (because arrays start at 0)
            if self.pos > len(text) - 1:
                # In that case we create a token with type EOF to determine that everything ended.
                return Token(EOF, None)

            # Got actual position inside text and set inside char
            current_char = text[self.pos]

            # Verify if the actual char is a valid digit and can be calculated
            # isdigit() is a Python String function
            if current_char.isdigit():
                # as chart is a digita, we need to create a Token Object with the type and value.
                # token = Token(INTEGER, int(current_char))
                lexeme = lexeme + current_char

                # Like we already processed the char, need to go to next position increment pos
                self.pos += 1

                
                # return token to expr() function
                # return token

            # Verify that chat is plus symbol
            if current_char == '+':
                # Create Token object with actual chat
                token = Token(PLUS, current_char)

                # as token already validated incremented position for next token
                self.pos += 1

                #return to caller()
                return token

            if current_char == '-':
                token = Token(MINUS, current_char)
                self.pos += 1
                return token

            self.error()

    def eat(self, token_type):
        # Verify the type of actual token being processed is equal expected
        # inside expr() in that moment
        if self.current_token.type in token_type:
            self.current_token = self.get_next_token()
        else:
            self.error()

    def expr(self):
        # Use get_next_token to define first token inside the expression
        # (that is composed of 3 tokens that follow the types INTEGER PLUS INTEGER)
        self.current_token = self.get_next_token()

        # Define the current_token as the left side of expressions (that consiste 3 elements)
        left = self.current_token

        # Call eat()
        self.eat(INTEGER)

        # eat(INTEGER) verified that token was INTEGER and already ask for next_token
        # to set in operation variable
        op = self.current_token

        # Valide if seconf token is PLUS type, if yes asked for next token.
        self.eat([MINUS, PLUS])

        # Set the third token as right part of expression.
        right = self.current_token

        # Verify that actual value is expected type and call next token
        # in normal conditions this will trigger EOL inside code
        self.eat(INTEGER)

        # as all 3 elements of expressions was identified the operation can be done.
        if op.type == PLUS:
          result = left.value + right.value

        if op.type == MINUS:
          result = left.value - right.value
        
        # Return result to caller()
        return result



def main():
  while True:
    try:
        text = input('calc> ')
    except EOFError:
        break
    if not text:
        continue

    # Load code inside interpreter
    interpreter = Interpreter(text)

    # Ask for interpreter execute code
    result = interpreter.expr()
    print(result)


if __name__ == '__main__':
  main()