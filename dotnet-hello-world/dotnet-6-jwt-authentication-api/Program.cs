using WebApi.Helpers;
using WebApi.Services;

var builder = WebApplication.CreateBuilder(args);

var services = builder.Services;
services.AddCors();
services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

services.Configure<AppSettings>(builder.Configuration.GetSection("AppSettings"));
services.AddScoped<IUserService, UserService>();

var app = builder.Build();
app.UseCors(x => x
    .AllowAnyOrigin()
    .AllowAnyMethod()
    .AllowAnyHeader());

// Isso aqui configura o JWT (que o mecanismo pra fazer token de autenticacao)
app.UseMiddleware<JwtMiddleware>();
app.UseSwagger();
app.UseSwaggerUI();
app.MapControllers();

app.Run("http://localhost:4000");