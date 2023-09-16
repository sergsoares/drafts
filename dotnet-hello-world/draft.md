# draft

https://dotnet.microsoft.com/en-us/learn/aspnet/hello-world-tutorial/create

https://www.w3schools.com/cs/cs_for_loop.php

https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/builtin-types/arrays

https://learn.microsoft.com/en-us/aspnet/core/tutorials/first-web-api?view=aspnetcore-6.0&tabs=visual-studio-code

https://learn.microsoft.com/en-us/aspnet/core/tutorials/min-web-api?view=aspnetcore-6.0&tabs=visual-studio

https://learn.microsoft.com/en-us/dotnet/standard/serialization/system-text-json/how-to?pivots=dotnet-8-0

https://blog.jetbrains.com/dotnet/2023/04/25/introduction-to-asp-net-core-minimal-apis/

# Dotnet

dotnet new console -o <NAME>

dotnet new web --no-https -o <NAME>

dotnet new web --no-https -o SergioApi


PROTOCOLO
https://

DOMONIO
stackoverflow.com

PATH
/questions/5002834/saving-a-file-from-stream-to-disk-using-c-sharp


PROTOCOLO
https://

DOMINIO
stackoverflow.com

PATH
/

dotnet new webapi --no-https -o WebApi




dotnet new web --no-https -o 9-WebApp
dotnet new webapi --no-https -o 9-WebapiApp
dotnet new mvc --no-https -o 9-MvcApp
dotnet new webapp --no-https -o 9-WebappApp
dotnet new worker -o 9-WorkerApp
 	

# Tutorial

> dotnet new webapi --no-https -o 9-WebapiApp
> Entra no /swagger pra ver

> Instala esses dois pacotes
dotnet add package Microsoft.AspNetCore.Authentication.JwtBearer --version 6.0.0
dotnet add package System.IdentityModel.Tokens.Jwt --version 6.15.0




====




# Requisição pra listar os usuarios (vai dar erro porque não tem token mas roda ela)
curl --location --request GET 'localhost:4000/users' \
--data-raw ''

Resposta vai ser 
{
    "message": "Unauthorized"
}

# Requisição pra rodar no terminal para gerar o teu token
curl --location --request POST 'localhost:4000/users/authenticate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "thales",
    "password": "thalesthales"
}'

Exemplo da resposta:
```
{
  "id":1,
  "firstName":"Thales",
  "lastName":"Liscano",
  "username":"thales",
  "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJuYmYiOjE2OTQ4MjEyODcsImV4cCI6MTY5NTQyNjA4NywiaWF0IjoxNjk0ODIxMjg3fQ.oaFnlLy0g8HCRoIxRA7hTv-wA7_xvRjLyuPPwdX-DnE"
}
```

# Requisição pra listar os usuarios (nessa precisamos gerar o token, dai copiarmos o valor do do token ali de cima e usamos pra listar)
curl --location --request GET 'localhost:4000/users' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJuYmYiOjE2OTQ4MjEyODcsImV4cCI6MTY5NTQyNjA4NywiaWF0IjoxNjk0ODIxMjg3fQ.oaFnlLy0g8HCRoIxRA7hTv-wA7_xvRjLyuPPwdX-DnE'
* Lembrar de trocar o valor depois do Authorization: pelo valor do token