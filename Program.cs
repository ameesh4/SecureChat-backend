using hushline.src.service.UserService;
using hushline.src.db.repository.UserRepository;
using dotenv.net;
using hushline.src.db;

var builder = WebApplication.CreateBuilder(args);

DotEnv.Load();

builder.Services.AddDatabase();
builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.AddScoped<IUserService, UserService>();
builder.Services.AddScoped<IUserRepository, UserRepository>();

if (builder.Environment.IsDevelopment())
{
    builder.Services.AddSwaggerGen(c =>
    {
        c.SwaggerDoc("v1", new() { Title = "Hushline API", Version = "v1" });
    });
}

var app = builder.Build();

app.Run();
