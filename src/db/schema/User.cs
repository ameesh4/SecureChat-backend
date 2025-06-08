using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace SecureChat.db.schema
{
    [Table("users")]
    public class User
    {
        [Key]
        public int Id { get; set; }
        [StringLength(50)]
        public required string Name { get; set; }
        [StringLength(15)]
        public string? PhoneNumber { get; set; }
        [StringLength(72)]
        public string? Email { get; set; }
        [StringLength(1000)]
        public string? Password { get; set; }
        [StringLength(1000)]
        public string? RefreshToken { get; set; }
        public DateTime? CreatedAt { get; set; }
        public DateTime? UpdatedAt { get; set; }
    }
}