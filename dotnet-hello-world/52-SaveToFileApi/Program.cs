var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", () => { 
  var outline = "Hello World!";
  using (System.IO.StreamWriter outfile = new System.IO.StreamWriter("yourFile.txt")) {
    outfile.Write(outline);
  }
  Results.Ok("OK");
});

app.Run();
