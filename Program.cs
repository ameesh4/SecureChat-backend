using dotenv.net;
using SecureChat.src.service.UserService;
using SecureChat.src.db.repository.UserRepository;
using SecureChat.src.db;

var builder = WebApplication.CreateBuilder(args);

DotEnv.Load();

builder.Services.AddDatabase();
builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.AddScoped<IUserService, UserService>();
builder.Services.AddScoped<IUserRepository, UserRepository>();

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI(c =>
{
    c.SwaggerEndpoint("/swagger/v1/swagger.json", "SecureChat API V1");
    c.RoutePrefix = string.Empty; // Set Swagger UI at the app's root
});

app.MapControllers();

app.Run();
