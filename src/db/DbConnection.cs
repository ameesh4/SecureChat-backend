using dotenv.net;
using Microsoft.EntityFrameworkCore;

namespace hushline.src.db
{
    public static class DbConnection
    {
        public static IServiceCollection AddDatabase(this IServiceCollection services)
        {
            DotEnv.Load();
            var connectionString = Environment.GetEnvironmentVariable("DBURL");
            if (string.IsNullOrEmpty(connectionString))
            {
                throw new InvalidOperationException("Database connection string is not set.");
            }
            services.AddDbContext<AppDbContext>(options =>
                options.UseNpgsql(connectionString));

            return services;
        }
    }
}