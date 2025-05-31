using Microsoft.EntityFrameworkCore;
using SecureChat.src.db.schema;

namespace SecureChat.src.db
{
    public class AppDbContext : DbContext
    {
        public AppDbContext(DbContextOptions<AppDbContext> options)
            : base(options)
        {
        }

        public DbSet<User> Users { get; set; }
    }
}