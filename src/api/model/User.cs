
using SecureChat.src.db.schema;

namespace SecureChat.src.api.model
{
    public class UserLogin
    {
        public string? PhoneNumber { get; set; } = string.Empty;
        public string? Email { get; set; } = string.Empty;
        public string Password { get; set; } = string.Empty;
    }

    public class LoginResponse
    {
        public User? User { get; set; }
        public string? Token { get; set; } = string.Empty;
    }
}