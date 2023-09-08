source = '''
public class Test {

   public static void main(String args[]) {
      int [] numbers = {10, 20, 30, 40, 50};

      for(int x : numbers ) {
         System.out.print( x );
         System.out.print(",");
      }
      System.out.print("\n");
      String [] names = {"James", "Larry", "Tom", "Lacy"};

      for( String name : names ) {
         System.out.print( name );
         System.out.print(",");
      }
   }
}'''


separators = ['\n']
# list_chars = list(source)
keywords = ['public', 'class', 'static', 'String','void', 'main', 'for', 'int']
symbols = ['{','}',')','(', '[', ']', '=', ';', ',','.', ':', "\""]


all_lexemes = []

lexeme = ''
for i,char in enumerate(source):
  if char == " ":
    continue

  if char in symbols and lexeme == ' ':
    all_lexemes.append(char)
    lexeme = ''
    continue
  
  if char in symbols and lexeme != '':
    all_lexemes.append(lexeme)
    all_lexemes.append(char)
    lexeme = ''
    continue

  if char == "\n":
    all_lexemes.append("<newline>")
    continue

  lexeme = lexeme + char
  if lexeme in keywords:
    all_lexemes.append(lexeme)
    lexeme = ''

  if lexeme in symbols :
    all_lexemes.append(lexeme)
    lexeme = ''

for each in all_lexemes:
  print(each)