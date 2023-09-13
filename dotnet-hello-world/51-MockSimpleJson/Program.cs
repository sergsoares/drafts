var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", () => 
  Results.Ok(new Cartao("test"))
);

app.Run();
public record Cartao(string Name);







