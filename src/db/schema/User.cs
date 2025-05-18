using System.ComponentModel.DataAnnotations;

namespace hushline.src.db.schema
{
    public class User
    {
        [Key]
        public int Id { get; set; }
        public required string Name { get; set; }
        public string? PhoneNumber { get; set; }
        public string? Email { get; set; }
        public required string Password { get; set; }
        public DateTime CreatedAt { get; set; }
        public DateTime UpdatedAt { get; set; }
    }
}