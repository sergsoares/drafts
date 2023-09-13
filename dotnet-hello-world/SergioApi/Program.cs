var builder = WebApplication.CreateBuilder();
var app = builder.Build();

app.MapGet("/consulta_venda", () => "{\"status\": \"PENDENTE\",\"valor\": 2323,\"tipo\" : \"CARTAO_CREDITO\"}");

app.MapGet("/thales", () => 
  Results.Ok(new Venda("PENDENTE", 250, "CARTAO_CREDITO")) 
);
app.Run();

public record Venda(string status, int valor, string tipo);














