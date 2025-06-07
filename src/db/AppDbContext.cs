using Microsoft.EntityFrameworkCore;
using SecureChat.db.schema;

namespace SecureChat.db
{
    public class AppDbContext(DbContextOptions<AppDbContext> options) : DbContext(options)
    {
        public DbSet<User> Users { get; set; }
    }
}