var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

var seuNome = "Sergio";

app.MapGet("/", () => $"Meu nome é {seuNome}!");

app.Run();
