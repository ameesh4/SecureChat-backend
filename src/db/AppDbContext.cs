using Microsoft.EntityFrameworkCore;
using hushline.src.db.schema;

namespace hushline.src.db
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