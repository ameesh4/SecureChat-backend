using System.Threading.Tasks;
using SecureChat.src.db.schema;

namespace SecureChat.src.service.JwtService
{
    public interface IJwtService
    {
        string GenerateToken(string userId);
        Task<User?> ValidateToken(string token);
    }
}